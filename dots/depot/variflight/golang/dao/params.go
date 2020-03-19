// Scry Info.  All rights reserved.
// license that can be found in the license file.

package dao

const (
	defaultPage    = "1"
	defaultPerpage = "20"

	APIMethodByFlightNumber APIMethod = "GetFlightDataByFlightNumber"
	APIMethodByAirports     APIMethod = "GetFlightDataBetweenTwoAirports"
	APIMethodByCities       APIMethod = "GetFlightDataBetweenTwoCities"
	APIMethodByStatus       APIMethod = "GetFlightDataByDepartureAndArrivalStatus"
)

// APIMethod is method of calling Api from https://www.variflight.com/
type APIMethod string

// APIParams contains APIMethod and its required parameters, with appid, appsecurity and _token excluded.
type APIParams struct {
	Method APIMethod
	Opts   map[string]string
}

type APIParamsConfFunc func() *APIParams

func GetFlightDataByFlightNumber(flightNumber, date string) APIParamsConfFunc {
	return func() *APIParams {
		return &APIParams{
			Method: APIMethodByFlightNumber,
			Opts: map[string]string{
				"fnum": flightNumber,
				"date": date,
			},
		}
	}
}

func GetFlightDataBetweenTwoAirports(departureAirport, arrivalAirport, date string) APIParamsConfFunc {
	return func() *APIParams {
		return &APIParams{
			Method: APIMethodByAirports,
			Opts: map[string]string{
				"dep":  departureAirport,
				"arr":  arrivalAirport,
				"date": date,
			},
		}
	}
}

func GetFlightDataBetweenTwoCities(departureCity, arrivalCity, date string) APIParamsConfFunc {
	return func() *APIParams {
		return &APIParams{
			Method: APIMethodByCities,
			Opts: map[string]string{
				"depcity": departureCity,
				"arrcity": arrivalCity,
				"date":    date,
			},
		}
	}
}

func GetFlightDataByDepartureAndArrivalStatus(airport, status, page, dataItemsNumberPerPage, date string) APIParamsConfFunc {
	return func() *APIParams {
		if page == "" {
			page = defaultPage
		}
		if dataItemsNumberPerPage == "" {
			dataItemsNumberPerPage = defaultPerpage
		}
		return &APIParams{
			Method: APIMethodByStatus,
			Opts: map[string]string{
				"airport": airport,
				"status":  status,
				"page":    page,
				"perpage": dataItemsNumberPerPage,
				"date":    date,
			},
		}
	}
}
