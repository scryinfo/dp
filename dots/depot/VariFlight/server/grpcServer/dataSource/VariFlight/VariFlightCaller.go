package VariFlight

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// VariFlightCaller coordinate the overall logic of calling API from https://www.variflight.com/, caching and storing the responded flight data.
type VariFlightCaller struct {
	// required parameters for API validation
	Appid            string
	RegistrationCode string

	// other shared parts for API calling
	commonPartOfURL  string
	flightDateLayout string
	flightTimeLayout string

	// time control parameters for avoiding unnecessary extra API calling
	MinDurAgainstExtraRequest time.Duration
	MaxDurAgainstExtraRequest time.Duration

	// arguments for connecting a database
	DriverName     string
	DataSourceName string

	storer *storer
	cacher *cacher
}

func New(appid, registrationCode string, minDurAgainstExtraRequest, maxDurAgainstExtraRequest time.Duration, driverName, dataSourceName string) *VariFlightCaller {
	return &VariFlightCaller{
		Appid:            appid,
		RegistrationCode: registrationCode,

		commonPartOfURL:  defaultCommonPartURL,
		flightDateLayout: defaultFlightDateLayout,
		flightTimeLayout: defaultFlightTimeLayout,

		MinDurAgainstExtraRequest: minDurAgainstExtraRequest,
		MaxDurAgainstExtraRequest: maxDurAgainstExtraRequest,

		DriverName:     driverName,
		DataSourceName: dataSourceName,

		cacher: newCacher(),
		storer: newStorer(driverName, dataSourceName),
	}
}

// Call coordinates the overall logic of API calling, caching and storing the responded flight data.
func (a *VariFlightCaller) Call(paramsFunc APIParamsConfFunc) ([]VariFlightData, error) {
	apiParams := paramsFunc()
	apiParams.Opts["appid"] = a.Appid

	token := token(apiParams.Opts, a.RegistrationCode)
	reqURL := url(a.commonPartOfURL, apiParams.Opts, a.RegistrationCode)

	// read data from cache
	data := a.cacher.read(token)
	// data not exists
	if data == nil {
		// call VariFlightCaller for data
		apiData, apiDataBytes, err := a.call(apiParams.Method, reqURL)
		if err != nil {
			return nil, err
		}
		data := newData(token, fmt.Sprint(md5.Sum(apiDataBytes)), time.Now(), a.flightDateLayout, apiData, string(apiDataBytes))

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
	// Note dynamic flight data may vary over time. If data has been just updated within during of MinDurAgainstExtraRequest,
	// or if it's departure plan date is far beyond the during of MaxDurAgainstExtraRequest, then avoid extra calling VariFlightCaller,
	// but return the cached data directly, because in this case we can roughly think that the flight data should
	// has low chance to change, so it's unnecessary to call the VariFlightCaller again for the same _token.
	if data.isUpdatedWithin(a.MinDurAgainstExtraRequest) || data.isDepPlanDateBeyond(a.MaxDurAgainstExtraRequest) {
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

//// ......
//func (a *VariFlightCaller) GetFlightDataByFlightNumber(flightNumber, date string) ([]VariFlightData, error) {
//	var queryParams = map[string]string{}
//	queryParams["fnum"] = flightNumber
//	queryParams["date"] = date
//	queryParams["appid"] = a.Appid
//
//	_token := utils.token(queryParams, a.RegistrationCode)
//	reqURL := utils.url(a.commonPartOfURL, queryParams, a.RegistrationCode)
//
//	// read data from cache
//	data := a.cacher.read(_token)
//	// data not exists
//	if data == nil {
//		// call VariFlightCaller for data
//		var apiData []VariFlightData
//		var apiDataBytes []byte
//		var err error
//		if apiData ,apiDataBytes, err = a.call(APIMethodByFlightNumber, reqURL); err != nil {
//			return nil, err
//		}
//		data := newData(_token, fmt.Sprint(md5.Sum(apiDataBytes)), time.Now(), a.flightDateLayout, apiData, string(apiDataBytes))
//
//		// storer data
//		if err = a.storer.Create(data); err != nil {
//			return nil, err
//		}
//
//		// cache data
//		data.valueJSONString = ""
//		a.cacher.create(data)
//
//		return apiData, nil
//	}
//
//	// data exists
//	// Note dynamic flight data may vary over time. If data has been just updated within during of MinDurAgainstExtraRequest,
//	// or if it's departure plan date is far beyond the during of MaxDurAgainstExtraRequest, then avoid extra calling VariFlightCaller,
//	// but return the cached data directly, because in this case we can roughly think that the flight data should
//	// has low chance to change, so it's unnecessary to call the VariFlightCaller again for the same _token.
//	if data.isUpdatedWithin(a.MinDurAgainstExtraRequest) || data.isDepPlanDateBeyond(a.MaxDurAgainstExtraRequest) {
//		return data.value, nil
//	}
//
//    // Otherwise, we should repeat calling VariFlightCaller for latest flight data.
//	apiData ,apiDataBytes, err := a.call(APIMethodByFlightNumber, reqURL)
//	if err != nil {
//		return nil, err
//	}
//	// if data hasn't varied in fact, only the updatedAtTime property needs to be updated.
//	newDigest := fmt.Sprint(md5.Sum(apiDataBytes))
//	if data.isSameDigest(newDigest) {
//		if err := a.storer.updateUpdateAtTime(_token, time.Now()); err != nil {
//			return nil, err
//		}
//		a.cacher.updateUpdateAtTime(_token, time.Now())
//		return data.value, nil
//	}
//	// but if data has varied over time, more data property need to be updated as follow.
//	if err := a.storer.update(_token, newDigest, time.Now(), string(apiDataBytes)); err != nil {
//		return nil, err
//	}
//	a.cacher.update(_token, newDigest, time.Now(), apiData)
//	return apiData, nil
//}

//func (a *VariFlightCaller) GetFlightDataBetweenTwoAirports(departureAirport, arrivalAirport, date string) ([]VariFlightData, error) {
//	var queryParams = map[string]string{}
//	queryParams["dep"] = departureAirport
//	queryParams["arr"] = arrivalAirport
//	queryParams["date"] = date
//	queryParams["Appid"] = a.Appid
//
//	urlWithToken := defaultCommonPartURL + utils.queryWithToken(queryParams, a.RegistrationCode)
//
//	var flightDataResp []VariFlightData
//
//	if err := a.call(APIMethodByAirports, urlWithToken, &flightDataResp); err != nil {
//		return nil, err
//	}
//	return flightDataResp, nil
//}
//
//func (a *VariFlightCaller) GetFlightDataBetweenTwoCities(departureCity, arrivalCity, date string) ([]VariFlightData, error) {
//	var queryParams = map[string]string{}
//	queryParams["depcity"] = departureCity
//	queryParams["arrcity"] = arrivalCity
//	queryParams["date"] = date
//	queryParams["Appid"] = a.Appid
//
//	urlWithToken := defaultCommonPartURL + utils.queryWithToken(queryParams, a.RegistrationCode)
//
//	var flightDataResp []VariFlightData
//
//	if err := a.call(APIMethodByCities, urlWithToken, &flightDataResp); err != nil {
//		return nil, err
//	}
//	return flightDataResp, nil
//}
//
//func (a *VariFlightCaller) GetFlightDataByDepartureAndArrivalStatus(airport, status, page, dataItemsNumberPerPage, date string) ([]VariFlightData, error) {
//	var queryParams = map[string]string{}
//	queryParams["airport"] = airport
//	queryParams["status"] = status
//	queryParams["page"] = page
//	queryParams["perpage"] = dataItemsNumberPerPage
//	queryParams["date"] = date
//	queryParams["Appid"] = a.Appid
//
//	urlWithToken := defaultCommonPartURL + utils.queryWithToken(queryParams, a.RegistrationCode)
//
//	var flightDataResp []VariFlightData
//
//	if err := a.call(APIMethodByStatus, urlWithToken, &flightDataResp); err != nil {
//		return nil, err
//	}
//	return flightDataResp, nil
//}
