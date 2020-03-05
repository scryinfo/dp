// Scry Info.  All rights reserved.
// license that can be found in the license file.

package VariFlight

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/scryinfo/dot/dots/db/gorms"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.uber.org/zap"

	"github.com/scryinfo/dot/dot"
)

const (
	VariFlightCallerTypeId = "e19f09c0-ac20-4d1e-a810-130900f94657"
)

var defaultVariFlightCallerConfig = VariFlightCallerConfig{
	commonPartOfURL:  "http://open-al.variflight.com/api/flight?",
	flightDateLayout: "2006-01-02 15:04:05",
	flightTimeLayout: "2006-01-02",

	MinDurAgainstExtraRequest: 1800e9,          // 30 minutes
	MaxDurAgainstExtraRequest: 5 * 24 * 3600e9, // 5 days
}

type VariFlightCallerConfig struct {
	// required parameters for data source VariFlight API validation
	Appid       string `json:"appid"`
	Appsecurity string `json:"appsecurity"`

	// other shared parts for API calling
	commonPartOfURL  string `json:"-"`
	flightDateLayout string `json:"-"`
	flightTimeLayout string `json:"-"`

	// time control parameters for avoiding unnecessary extra API calling
	MinDurAgainstExtraRequest uint64 `json:"minDurAgainstExtraRequest, omitempty"`
	MaxDurAgainstExtraRequest uint64 `json:"maxDurAgainstExtraRequest, omitempty"`
}

// VariFlightCaller coordinate the overall logic of calling API from https://www.variflight.com/, caching and storing the responded flight data.
type VariFlightCaller struct {
	config *VariFlightCallerConfig

	Gorms *gorms.Gorms `dot:""`

	storer *storer
	cacher *cacher
}

func VariFlightCallerTypeLives() []*dot.TypeLives {
	return []*dot.TypeLives{
		{
			Meta: dot.Metadata{
				TypeId: VariFlightCallerTypeId,
				NewDoter: func(conf []byte) (dot.Dot, error) {
					_conf := defaultVariFlightCallerConfig
					if err := dot.UnMarshalConfig(conf, &_conf); err != nil {
						dot.Logger().Debugln("UnMarshalConfig(VariFlightCallerConfig) failed", zap.Error(err))
						os.Exit(1)
					}
					dot.Logger().Debug(func() string {
						return spew.Sprintf("VariFlightCallerConfig: %#+v", _conf)
					})
					return &VariFlightCaller{config: &_conf}, nil
				},
			},
			//Lives: []dot.Live{
			//	{
			//		TypeId: VariFlightCallerTypeId,
			//		RelyLives: map[string]dot.LiveId{
			//			"Gorms": gorms.TypeId,
			//		},
			//	},
			//},
		},
		gorms.TypeLives()[0],
	}
}

func (a *VariFlightCaller) AfterAllInject(l dot.Line) {
	db := a.Gorms.Db.DB()
	driverName := a.Gorms.Db.Dialect().GetName()
	a.storer = newStorer(db, driverName)
	a.cacher = newCacher()
}

// Call coordinates the overall logic of API calling, caching and storing the responded flight data.
func (a *VariFlightCaller) Call(paramsFunc APIParamsConfFunc) ([]VariFlightData, error) {
	dot.Logger().Info(func() string {
		return spew.Sprintf("VariFlightCaller Call. parameters: %#+v", paramsFunc())
	})
	apiParams := paramsFunc()
	apiParams.Opts["appid"] = a.config.Appid

	token := token(apiParams.Opts, a.config.Appsecurity)
	reqURL := url(a.config.commonPartOfURL, apiParams.Opts, a.config.Appsecurity)

	// read data from cache
	data := a.cacher.read(token)
	// data not exists
	if data == nil {
		// call VariFlightCaller for data
		apiData, apiDataBytes, err := a.call(apiParams.Method, reqURL)
		if err != nil {
			return nil, err
		}
		data := newData(token, fmt.Sprint(md5.Sum(apiDataBytes)), time.Now(), a.config.flightDateLayout, apiData, string(apiDataBytes))

		// storer data
		if err = a.storer.create(data); err != nil {
			return nil, err
		}

		// cache data
		data.valueJSONString = ""
		a.cacher.create(data)

		return apiData, nil
	}

	// data exists
	// Note dynamic flight data may vary over time. If data has been just updated within during of minDurAgainstExtraRequest,
	// or if it's departure plan date is far beyond the during of maxDurAgainstExtraRequest, then avoid extra calling VariFlightCaller,
	// but return the cached data directly, because in this case we can roughly think that the flight data should
	// has low chance to change, so it's unnecessary to call the VariFlightCaller again for the same _token.
	if data.isUpdatedWithin(time.Duration(a.config.MinDurAgainstExtraRequest)) || data.isDepPlanDateBeyond(time.Duration(a.config.MaxDurAgainstExtraRequest)) {
		return data.value, nil
	}

	// Otherwise, we should repeat calling VariFlightCaller for latest flight data.
	apiData, apiDataBytes, err := a.call(apiParams.Method, reqURL)
	if err != nil {
		return nil, err
	}
	// if data hasn't varied in fact, only the updatedAtTime property needs to be updated.
	newDigest := fmt.Sprint(md5.Sum(apiDataBytes))
	if data.isSameDigest(newDigest) {
		if err := a.storer.updateUpdateAtTime(token, time.Now()); err != nil {
			return nil, err
		}
		a.cacher.updateUpdateAtTime(token, time.Now())
		return data.value, nil
	}
	// but if data has varied over time, more data property need to be updated as follow.
	if err := a.storer.update(token, newDigest, time.Now(), string(apiDataBytes)); err != nil {
		return nil, err
	}
	a.cacher.update(token, newDigest, time.Now(), apiData)
	return apiData, nil
}

// call calls API and returns the decoded data and the raw body bytes.
func (a *VariFlightCaller) call(method APIMethod, url string) ([]VariFlightData, []byte, error) {
	dot.Logger().Info(func() string {
		return spew.Sprintf("VariFlightCaller call. method: %v, url: %v\n", method, url)
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
