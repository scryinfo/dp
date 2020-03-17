// Scry Info.  All rights reserved.
// license that can be found in the license file.

package _go

import (
	"github.com/scryinfo/dp/dots/depot/VariFlight/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	VariFlight "github.com/scryinfo/dp/dots/depot/VariFlight/go/VariFlightApiCaller"
)

var _ proto.VariFlightDataServiceServer = (*variflightServiceServer)(nil)

var variFligtAPICodeToGrpcCode = map[VariFlight.VariFlightStatusCode]codes.Code{
	VariFlight.UserNotExists:          codes.Unauthenticated,
	VariFlight.CheckPending:           codes.Unavailable,
	VariFlight.IPIsNotInWhiteList:     codes.Unauthenticated,
	VariFlight.ParamValidationFailure: codes.InvalidArgument,
	VariFlight.NoDataPermission:       codes.PermissionDenied,
	VariFlight.NoData:                 codes.NotFound,
	VariFlight.UnknownError:           codes.Unknown,
}

//  variflightServiceServer serve flight data provided by https://www.variflight.com/.
type variflightServiceServer struct {
	variFlightApiCaller *VariFlight.VariFlightApiCaller
}

func newVariFlightServiceServer(vfCaller *VariFlight.VariFlightApiCaller) *variflightServiceServer {
	return &variflightServiceServer{vfCaller}
}

func (s *variflightServiceServer) GetFlightDataByFlightNumber(req *proto.GetFlightDataByFlightNumberRequest, srv proto.VariFlightDataService_GetFlightDataByFlightNumberServer) error {
	params := VariFlight.GetFlightDataByFlightNumber(req.FlightNumber, req.Date)
	variFlightDatas, err := s.variFlightApiCaller.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func (s *variflightServiceServer) GetFlightDataBetweenTwoAirports(req *proto.GetFlightDataBetweenTwoAirportsRequest, srv proto.VariFlightDataService_GetFlightDataBetweenTwoAirportsServer) error {
	params := VariFlight.GetFlightDataBetweenTwoAirports(req.DepartureAirport, req.ArrivalAirport, req.Date)
	variFlightDatas, err := s.variFlightApiCaller.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func (s *variflightServiceServer) GetFlightDataBetweenTwoCities(req *proto.GetFlightDataBetweenTwoCitiesRequest, srv proto.VariFlightDataService_GetFlightDataBetweenTwoCitiesServer) error {
	params := VariFlight.GetFlightDataBetweenTwoCities(req.DepartureCity, req.ArrivalCity, req.Date)
	variFlightDatas, err := s.variFlightApiCaller.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func (s *variflightServiceServer) GetFlightDataByDepartureAndArrivalStatus(req *proto.GetFlightDataAtOneAirportByStatusRequest, srv proto.VariFlightDataService_GetFlightDataByDepatureAndArrivalStatusServer) error {
	params := VariFlight.GetFlightDataByDepartureAndArrivalStatus(req.Airport, req.Status, "", "", req.Date)
	variFlightDatas, err := s.variFlightApiCaller.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func protoVariFlightData(data *VariFlight.VariFlightData) *proto.VariFlightData {
	return &proto.VariFlightData{
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
