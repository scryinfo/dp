package grpcServer

import (
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"../_proto"
	"./dataSource/VariFlight"
)

var _ _proto.VariFlightDataServiceServer = (*VariFlightGrpcServer)(nil)

var variFligtAPICodeToGrpcCode = map[VariFlight.VariFlightStatusCode]codes.Code{
	VariFlight.UserNotExists:          codes.Unauthenticated,
	VariFlight.CheckPending:           codes.Unavailable,
	VariFlight.IPIsNotInWhiteList:     codes.Unauthenticated,
	VariFlight.ParamValidationFailure: codes.InvalidArgument,
	VariFlight.NoDataPermission:       codes.PermissionDenied,
	VariFlight.NoData:                 codes.NotFound,
	VariFlight.UnknownError:           codes.Unknown,
}

//  VariFlightGrpcServer serve flight data provided by https://www.variflight.com/.
type VariFlightGrpcServer struct {
	// required parameters for data source VariFlight API validation
	Appid            string
	RegistrationCode string

	// time control parameters for avoiding unnecessary extra API calling
	MinDurAgainstExtraRequest time.Duration
	MaxDurAgainstExtraRequest time.Duration

	// arguments for connecting a database
	DriverName     string
	DataSourceName string

	variFlightCaller *VariFlight.VariFlightCaller
}

func New(appid, registrationCode string, minDurAgainstExtraRequest, maxDurAgainstExtraRequest time.Duration, driverName, dataSourceName string) *VariFlightGrpcServer {
	return &VariFlightGrpcServer{
		Appid:            appid,
		RegistrationCode: registrationCode,

		MinDurAgainstExtraRequest: minDurAgainstExtraRequest,
		MaxDurAgainstExtraRequest: maxDurAgainstExtraRequest,

		DriverName:     driverName,
		DataSourceName: dataSourceName,

		variFlightCaller: VariFlight.New(appid, registrationCode, minDurAgainstExtraRequest, maxDurAgainstExtraRequest, driverName, dataSourceName),
	}
}

func (s *VariFlightGrpcServer) GetFlightDataByFlightNumber(req *_proto.GetFlightDataByFlightNumberRequest, srv _proto.VariFlightDataService_GetFlightDataByFlightNumberServer) error {
	params := VariFlight.GetFlightDataByFlightNumber(req.FlightNumber, req.Date)
	variFlightDatas, err := s.variFlightCaller.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func (s *VariFlightGrpcServer) GetFlightDataBetweenTwoAirports(req *_proto.GetFlightDataBetweenTwoAirportsRequest, srv _proto.VariFlightDataService_GetFlightDataBetweenTwoAirportsServer) error {
	params := VariFlight.GetFlightDataBetweenTwoAirports(req.DepartureAirport, req.ArrivalAirport, req.Date)
	variFlightDatas, err := s.variFlightCaller.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func (s *VariFlightGrpcServer) GetFlightDataBetweenTwoCities(req *_proto.GetFlightDataBetweenTwoCitiesRequest, srv _proto.VariFlightDataService_GetFlightDataBetweenTwoCitiesServer) error {
	params := VariFlight.GetFlightDataBetweenTwoCities(req.DepartureCity, req.ArrivalCity, req.Date)
	variFlightDatas, err := s.variFlightCaller.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func (s *VariFlightGrpcServer) GetFlightDataByDepartureAndArrivalStatus(req *_proto.GetFlightDataAtOneAirportByStatusRequest, srv _proto.VariFlightDataService_GetFlightDataByDepatureAndArrivalStatusServer) error {
	params := VariFlight.GetFlightDataByDepartureAndArrivalStatus(req.Airport, req.Status, "", "", req.Date)
	variFlightDatas, err := s.variFlightCaller.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func protoVariFlightData(data *VariFlight.VariFlightData) *_proto.VariFlightData {
	return &_proto.VariFlightData{
		Fcategory:             data.Fcategory,
		FlightNo:              data.FlightNo,
		FlightCompany:         data.FlightCompany,
		FlightDepcode:         data.FlightDepcode,
		FlightArrcode:         data.FlightArrcode,
		FlightDeptimePlanDate: data.FlightDeptimePlanDate,
		FlightArrtimePlanDate: data.FlightArrtimePlanDate,
		FlightDeptimeDate:     data.FlightDeptimeDate,
		FlightArrtimeDate:     data.FlightArrtimeDate,
		FlightState:           data.FlightState,
		FlightHTerminal:       data.FlightHTerminal,
		FlightTerminal:        data.FlightTerminal,
		OrgTimezone:           data.OrgTimezone,
		DstTimezone:           data.DstTimezone,
		ShareFlightNo:         data.ShareFlightNo,
		StopFlag:              data.StopFlag,
		ShareFlag:             data.ShareFlag,
		VirtualFlag:           data.VirtualFlag,
		LegFlag:               data.LegFlag,
		FlightDep:             data.FlightDep,
		FlightArr:             data.FlightArr,
		FlightDepAirport:      data.FlightDepAirport,
		FlightArrAirport:      data.FlightArrAirport,
	}
}

func protoError(err interface{}) error {
	switch err := err.(type) {
	case VariFlight.GetUrlError:
		return status.Errorf(codes.FailedPrecondition, "GetUrlError: ", err)
	case VariFlight.VariFlightDataQueryError:
		return status.Errorf(variFligtAPICodeToGrpcCode[err.StatusCode], "VariFlightDataQueryError: %v ", err)
	case VariFlight.DecodeJsonError:
		return status.Errorf(codes.FailedPrecondition, "DecodeJsonError: %v", err)
	case VariFlight.DBAccessError:
		return status.Errorf(codes.FailedPrecondition, "DBAccessError: %v", err)
	default:
		return status.Errorf(codes.Unknown, "unknown error: %v", err)
	}
}
