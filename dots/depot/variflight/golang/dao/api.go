// Scry Info.  All rights reserved.
// license that can be found in the license file.

package dao

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/db/gorms"
)

const (
	ApiTypeId = "e19f09c0-ac20-4d1e-a810-130900f94657"
)

var defaultApiConfig = ApiConfig{
	commonPartOfURL:  "http://open-al.variflight.com/api/flight?",
	flightDateLayout: "2006-01-02",
	flightTimeLayout: "2006-01-02 15:04:05",

	dataAge: 30 * 60, // 30 minutes

	TestMode: false,
}

type ApiConfig struct {
	// required parameters for data source VariFlight API validation
	Appid       string `json:"appid"`
	Appsecurity string `json:"appsecurity"`

	// other shared parts for API calling
	commonPartOfURL  string `json:"-"`
	flightDateLayout string `json:"-"`
	flightTimeLayout string `json:"-"`

	// time control parameters (measured in second) for avoiding unnecessary extra API calling
	dataAge time.Duration `json:"data_age, omitempty"`

	// TestMode only for test purpose with fake datum, when in production it must be false
	TestMode bool `json:"test_mode, omitempty"`
}

// Api coordinate the overall logic of calling API from https://www.variflight.com/, caching and storing the responded flight data.
type Api struct {
	config *ApiConfig

	Gorms *gorms.Gorms `dot:""`

	storer *storer
	cacher *cacher
}

func ApiTypeLives() []*dot.TypeLives {
	return []*dot.TypeLives{
		{
			Meta: dot.Metadata{
				TypeId: ApiTypeId,
				NewDoter: func(conf []byte) (dot.Dot, error) {
					_conf := defaultApiConfig
					if err := dot.UnMarshalConfig(conf, &_conf); err != nil {
						dot.Logger().Debugln("UnMarshalConfig(ApiConfig) failed", zap.Error(err))
						os.Exit(1)
					}
					dot.Logger().Debug(func() string {
						return fmt.Sprintf("ApiConfig: %v", _conf)
					})
					return &Api{config: &_conf}, nil
				},
			},
			//Lives: []dot.Live{
			//	{
			//		TypeId: ApiTypeId,
			//		RelyLives: map[string]dot.LiveId{
			//			"Gorms": gorms.TypeId,
			//		},
			//	},
			//},
		},
		gorms.TypeLives()[0],
	}
}

func (a *Api) AfterAllInject(l dot.Line) {
	db := a.Gorms.Db.DB()
	driverName := a.Gorms.Db.Dialect().GetName()
	a.storer = newStorer(db, driverName)
	a.cacher = newCacher()
}

// Call coordinates the overall logic of API calling, caching and storing the responded flight data.
func (a *Api) Call(paramsFunc APIParamsConfFunc) ([]VariFlightData, error) {
	apiParams := paramsFunc()
	dot.Logger().Info(func() string {
		return fmt.Sprintf("Api Call. parameters: %v", apiParams)
	})

	if a.config.TestMode {
		return a.callInTestMode(apiParams)
	}

	apiParams.Opts["appid"] = a.config.Appid

	token := token(apiParams.Opts, a.config.Appsecurity)
	reqURL := url(a.config.commonPartOfURL, apiParams.Opts, a.config.Appsecurity)

	// read local data and return if valid
	localData, err := a.localData(token)
	if err != nil {
		return nil, err
	}
	if localData != nil && a.isAcceptable(localData) {
		return localData.value, nil
	}

	// call Api for remote data
	remoteData, remoteDataBytes, err := a.call(apiParams.Method, reqURL)
	if err != nil {
		return nil, err
	}

	// create new record of local data if not exists
	if localData == nil {
		if err := a.create(&data{
			token:            token,
			digest:           fmt.Sprint(md5.Sum(remoteDataBytes)),
			updatedAtTime:    time.Now(),
			flightDateLayout: a.config.flightDateLayout,
			value:            remoteData,
			valueJSONString:  string(remoteDataBytes),
		}); err != nil {
			return nil, err
		}
		return remoteData, nil
	}

	// otherwise update local data
	if err := a.update(localData, remoteDataBytes, remoteData); err != nil {
		return nil, err
	}
	return remoteData, nil
}

func (a *Api) localData(token string) (*data, error) {
	d := a.cacher.read(token)
	if d != nil {
		return d, nil
	}

	d, err := a.storer.read(token)
	if err != nil {
		return nil, err
	}

	if d == nil {
		return nil, nil
	}

	d.valueJSONString = ""
	a.cacher.create(d)
	return d, nil
}

// Note dynamic flight data may vary over time. If data has been just updated within during of minDurAgainstExtraRequest,
// then avoid extra calling Api,
// but return the cached data directly, because in this case we can roughly think that the flight data should
// has low chance to change, so it's unnecessary to call the Api again for the same _token.
// Otherwise, we should repeat calling Api for latest flight data.
func (a *Api) isAcceptable(d *data) bool {
	return d.isUpdatedWithin((a.config.dataAge) * time.Second)
}

func (a *Api) update(d *data, newBytes []byte, newValue []VariFlightData) error {
	// if data hasn't varied in fact, only the updatedAtTime property needs to be updated.
	newDigest := fmt.Sprint(md5.Sum(newBytes))
	if d.isSameDigest(newDigest) {
		if err := a.storer.updateUpdateAtTime(d.token, time.Now()); err != nil {
			return nil
		}
		a.cacher.updateUpdateAtTime(d.token, time.Now())
		return nil
	}
	// but if data has varied over time, more data property need to be updated as follow.
	if err := a.storer.update(d.token, newDigest, time.Now(), string(newBytes)); err != nil {
		return nil
	}
	a.cacher.update(d.token, newDigest, time.Now(), newValue)
	return nil
}

func (a *Api) create(d *data) error {
	if err := a.storer.create(d); err != nil {
		return err
	}
	d.valueJSONString = ""
	a.cacher.create(d)
	return nil
}

// call calls API and returns the decoded data and the raw body bytes.
func (a *Api) call(method APIMethod, url string) ([]VariFlightData, []byte, error) {
	dot.Logger().Info(func() string {
		return fmt.Sprintf("Api call. method: %v, url: %v\n", method, url)
	})
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, newGetUrlError(method, url, err)
	}
	defer resp.Body.Close()

	var buffer bytes.Buffer
	bodyReader := io.TeeReader(resp.Body, &buffer)

	decoder := json.NewDecoder(bodyReader)

	if resp.StatusCode != http.StatusOK {
		var errInfo VariFlightErrorInfo
		if err := decoder.Decode(&errInfo); err != nil {
			return nil, nil, newDecodeJsonError(method, url, resp.StatusCode, resp.Status, err)
		}
		return nil, nil, newDataQueryError(method, url, VariFlightStatusCode(resp.StatusCode), VariFlightStatus(resp.Status), errInfo.ErrorCode, errInfo.Error)
	}

	var apiData []VariFlightData
	if err := decoder.Decode(&apiData); err != nil {
		return nil, nil, newDecodeJsonError(method, url, resp.StatusCode, resp.Status, err)
	}

	return apiData, buffer.Bytes(), nil
}

// callInTestMode only serves for test purpose with fake datum, without complex logic on accessing real data source API.
// This is not for unit test, but used as a quick way to test and check browser whether it's http requests can finally reach this code.
func (a *Api) callInTestMode(params *APIParams) ([]VariFlightData, error) {
	datum := []VariFlightData{
		{
			FlightNo:              "CA4506",
			FlightDeptimePlanDate: "2020-03-19",
			FlightDep:             "南京",
			FlightArr:             "成都",
			FlightDepAirport:      "南京禄口",
			FlightArrAirport:      "成都双流",
		},
		{
			FlightNo:              "CA4505",
			FlightDeptimePlanDate: "2020-03-19",
			FlightDep:             "南京",
			FlightArr:             "成都",
			FlightDepAirport:      "南京禄口",
			FlightArrAirport:      "成都双流",
		},
		{
			FlightNo:              "CA4504",
			FlightDeptimePlanDate: "2020-03-19",
			FlightDep:             "南京",
			FlightArr:             "成都",
			FlightDepAirport:      "南京禄口",
			FlightArrAirport:      "成都双流",
		},
		{
			FlightNo:              "CA4503",
			FlightDeptimePlanDate: "2020-03-20",
			FlightDep:             "南京",
			FlightArr:             "成都",
			FlightDepAirport:      "南京禄口",
			FlightArrAirport:      "成都双流",
		},
		{
			FlightNo:              "CA4502",
			FlightDeptimePlanDate: "2020-03-20",
			FlightDep:             "南京",
			FlightArr:             "成都",
			FlightDepAirport:      "南京禄口",
			FlightArrAirport:      "成都双流",
		},
		{
			FlightNo:              "CA4501",
			FlightDeptimePlanDate: "2020-03-20",
			FlightDep:             "南京",
			FlightArr:             "成都",
			FlightDepAirport:      "南京禄口",
			FlightArrAirport:      "成都双流",
		},
	}
	results := []VariFlightData{}

	switch params.Method {
	case APIMethodByFlightNumber:
		flightNo, date := params.Opts["fnum"], params.Opts["date"]
		for _, data := range datum {
			if data.FlightNo == flightNo && data.FlightDeptimePlanDate == date {
				results = append(results, data)
			}
		}
	case APIMethodByAirports:
		dep, arr, date := params.Opts["dep"], params.Opts["arr"], params.Opts["date"]
		for _, data := range datum {
			if data.FlightDepAirport == dep && data.FlightArrAirport == arr && data.FlightDeptimePlanDate == date {
				results = append(results, data)
			}
		}
	case APIMethodByCities:
		depCity, arrCity, date := params.Opts["dep"], params.Opts["arr"], params.Opts["date"]
		for _, data := range datum {
			if data.FlightDep == depCity && data.FlightArr == arrCity && data.FlightDeptimePlanDate == date {
				results = append(results, data)
			}
		}
	case APIMethodByStatus:
		return nil, errors.New("Service unimplemented")
	default:
		return nil, errors.New("Forbidden request")
	}
	return results, nil
}
