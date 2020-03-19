// Scry Info.  All rights reserved.
// license that can be found in the license file.

package golang

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/depot/variflight/golang/dao"
	"github.com/scryinfo/dp/dots/depot/variflight/golang/protobuf"
)

const (
	GrpcServiceServerTypeId = "a0b3f061-8e33-42b8-943a-f0542ad2533f"
)

var _ protobuf.VariFlightDataServiceServer = (*GrpcServiceServer)(nil)

var variFligtAPICodeToGrpcCode = map[dao.VariFlightStatusCode]codes.Code{
	dao.UserNotExists:          codes.Unauthenticated,
	dao.CheckPending:           codes.Unavailable,
	dao.IPIsNotInWhiteList:     codes.Unauthenticated,
	dao.ParamValidationFailure: codes.InvalidArgument,
	dao.NoDataPermission:       codes.PermissionDenied,
	dao.NoData:                 codes.NotFound,
	dao.UnknownError:           codes.Unknown,
}

type DataSource interface {
	Call(paramsFunc dao.APIParamsConfFunc) ([]dao.VariFlightData, error)
}

//  GrpcServiceServer serve flight data provided by https://www.variflight.com/.
type GrpcServiceServer struct {
	DataSource DataSource `dot:""`
}

func GrpcServiceServerTypeLives() []*dot.TypeLives {
	typeLives := []*dot.TypeLives{
		{
			Meta: dot.Metadata{
				TypeId: GrpcServiceServerTypeId,
				NewDoter: func(conf []byte) (dot.Dot, error) {
					return &GrpcServiceServer{}, nil
				},
			},
		},
	}
	typeLives = append(typeLives, dao.ApiTypeLives()...)
	return typeLives
}

func (s *GrpcServiceServer) GetFlightDataByFlightNumber(req *protobuf.GetFlightDataByFlightNumberRequest, srv protobuf.VariFlightDataService_GetFlightDataByFlightNumberServer) error {
	srv.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	srv.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))

	params := dao.GetFlightDataByFlightNumber(req.FlightNumber, req.Date)
	variFlightDatas, err := s.DataSource.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func (s *GrpcServiceServer) GetFlightDataBetweenTwoAirports(req *protobuf.GetFlightDataBetweenTwoAirportsRequest, srv protobuf.VariFlightDataService_GetFlightDataBetweenTwoAirportsServer) error {
	srv.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	srv.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))

	params := dao.GetFlightDataBetweenTwoAirports(req.DepartureAirport, req.ArrivalAirport, req.Date)
	variFlightDatas, err := s.DataSource.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func (s *GrpcServiceServer) GetFlightDataBetweenTwoCities(req *protobuf.GetFlightDataBetweenTwoCitiesRequest, srv protobuf.VariFlightDataService_GetFlightDataBetweenTwoCitiesServer) error {
	srv.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	srv.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))

	params := dao.GetFlightDataBetweenTwoCities(req.DepartureCity, req.ArrivalCity, req.Date)
	variFlightDatas, err := s.DataSource.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func (s *GrpcServiceServer) GetFlightDataByDepartureAndArrivalStatus(req *protobuf.GetFlightDataAtOneAirportByStatusRequest, srv protobuf.VariFlightDataService_GetFlightDataByDepartureAndArrivalStatusServer) error {
	srv.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	srv.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))

	params := dao.GetFlightDataByDepartureAndArrivalStatus(req.Airport, req.Status, "", "", req.Date)
	variFlightDatas, err := s.DataSource.Call(params)
	if err != nil {
		return protoError(err)
	}

	for _, variFlightData := range variFlightDatas {
		srv.Send(protoVariFlightData(&variFlightData))
	}

	return nil
}

func protoVariFlightData(data *dao.VariFlightData) *protobuf.VariFlightData {
	return &protobuf.VariFlightData{
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
	case dao.GetUrlError:
		return status.Errorf(codes.FailedPrecondition, "GetUrlError: ", err)
	case dao.VariFlightDataQueryError:
		return status.Errorf(variFligtAPICodeToGrpcCode[err.StatusCode], "VariFlightDataQueryError: %v ", err)
	case dao.DecodeJsonError:
		return status.Errorf(codes.FailedPrecondition, "DecodeJsonError: %v", err)
	case dao.DBAccessError:
		return status.Errorf(codes.FailedPrecondition, "DBAccessError: %v", err)
	default:
		return status.Errorf(codes.Unknown, "unknown error: %v", err)
	}
}
