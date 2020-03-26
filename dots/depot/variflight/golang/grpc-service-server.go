// Scry Info.  All rights reserved.
// license that can be found in the license file.

package golang

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"log"

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

//  GrpcServiceServer serve flight data provided by https://www.variflight.com/.
type GrpcServiceServer struct {
	DataSource *dao.Api `dot:""`
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

func (s *GrpcServiceServer) GetFlightDataByFlightNumber(srv protobuf.VariFlightDataService_GetFlightDataByFlightNumberServer) error {
	srv.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	srv.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))

	log.Println("started stream\n\n")
	defer log.Println("returened stream\n\n")
	for {
		req, err := srv.Recv()
		if err ==io.EOF {
			log.Println("returened stream. Encountered EOF.\n\n")
			return nil
		}
		if err != nil {
			log.Printf("failed to Recv(): %#+v\n\n", err)
			return err
		}
		log.Printf("received req: %v\n\n", req)

		params := dao.GetFlightDataByFlightNumber(req.FlightNumber, req.Date)
		log.Printf("params: %v\n\n", params)
		log.Printf("GrpcServiceServer DataSource: %#+v\n\n", s.DataSource)
		variFlightDatas, err := s.DataSource.Call(params)
		if err != nil {
			return protoError(err)
		}
		log.Printf("params called.")

		for _, variFlightData := range variFlightDatas {
			if err := srv.Send(protoVariFlightData(&variFlightData)); err != nil {
				log.Printf("failed to sent variFlightData: %#+v\n\n", variFlightData)
				return err
			}
			log.Printf("sent variFlightData: %+#v\n\n", variFlightData)
		}
	}
	return nil
}

func (s *GrpcServiceServer) GetFlightDataBetweenTwoAirports(srv protobuf.VariFlightDataService_GetFlightDataBetweenTwoAirportsServer) error  {
	srv.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	srv.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))

	log.Println("started stream\n\n")
	defer log.Println("returened stream\n\n")
	for {
		req, err := srv.Recv()
		if err ==io.EOF {
			log.Println("returened stream. Encountered EOF.\n\n")
			return nil
		}
		if err != nil {
			log.Printf("failed to Recv(): %#+v\n\n", err)
			return err
		}

		params := dao.GetFlightDataBetweenTwoAirports(req.DepartureAirport, req.ArrivalAirport, req.Date)
		variFlightDatas, err := s.DataSource.Call(params)
		if err != nil {
			return protoError(err)
		}

		for _, variFlightData := range variFlightDatas {
			if err := srv.Send(protoVariFlightData(&variFlightData)); err != nil {
				log.Printf("failed to sent variFlightData: %#+v\n\n", variFlightData)
				return err
			}
			log.Printf("sent variFlightData: %+#v\n\n", variFlightData)
		}
	}
	return nil
}

func (s *GrpcServiceServer) GetFlightDataBetweenTwoCities(srv protobuf.VariFlightDataService_GetFlightDataBetweenTwoCitiesServer) error {
	srv.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	srv.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))

	log.Println("started stream\n\n")
	defer log.Println("returened stream\n\n")
	for {
		req, err := srv.Recv()
		if err ==io.EOF {
			log.Println("returened stream. Encountered EOF.\n\n")
			return nil
		}
		if err != nil {
			log.Printf("failed to Recv(): %#+v\n\n", err)
			return err
		}

		params := dao.GetFlightDataBetweenTwoCities(req.DepartureCity, req.ArrivalCity, req.Date)
		variFlightDatas, err := s.DataSource.Call(params)
		if err != nil {
			return protoError(err)
		}

		for _, variFlightData := range variFlightDatas {
			if err := srv.Send(protoVariFlightData(&variFlightData)); err != nil {
				log.Printf("failed to sent variFlightData: %#+v\n\n", variFlightData)
				return err
			}
			log.Printf("sent variFlightData: %+#v\n\n", variFlightData)
		}
	}
	return nil
}

func (s *GrpcServiceServer) GetFlightDataByDepartureAndArrivalStatus(srv protobuf.VariFlightDataService_GetFlightDataByDepartureAndArrivalStatusServer) error  {
	srv.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	srv.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))

	log.Println("started stream\n\n")
	defer log.Println("returened stream\n\n")
	for {
		req, err := srv.Recv()
		if err ==io.EOF {
			log.Println("returened stream. Encountered EOF.\n\n")
			return nil
		}
		if err != nil {
			log.Printf("failed to Recv(): %#+v\n\n", err)
			return err
		}

		params := dao.GetFlightDataByDepartureAndArrivalStatus(req.Airport, req.Status, "", "", req.Date)
		variFlightDatas, err := s.DataSource.Call(params)
		if err != nil {
			return protoError(err)
		}

		for _, variFlightData := range variFlightDatas {
			if err := srv.Send(protoVariFlightData(&variFlightData)); err != nil {
				log.Printf("failed to sent variFlightData: %#+v\n\n", variFlightData)
				return err
			}
			log.Printf("sent variFlightData: %+#v\n\n", variFlightData)
		}
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
