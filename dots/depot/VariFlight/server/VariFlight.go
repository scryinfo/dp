// Scry Info.  All rights reserved.
// license that can be found in the license file.

package server

import (
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/depot/VariFlight/server/_proto"
	"google.golang.org/grpc"
	"time"

	"github.com/scryinfo/dot/dots/grpc/gserver"

	"./grpcServer"
)

const (
	VariFlightWebSocketGrpcServerTypeId = ""
)

type VariFlightWebSocketGrpcServerConfig struct {
	// required parameters for data source VariFlight API validation
	Appid       string `json:"appid"`
	Appsecurity string `json:"appsecurity"`

	// todo: rename
	// time control parameters for avoiding unnecessary extra API calling
	MinDurAgainstExtraRequest time.Duration `json:"minDurAgainstExtraRequest"`
	MaxDurAgainstExtraRequest time.Duration `json:"maxDurAgainstExtraRequest"`

	// arguments for connecting a database
	DriverName     string
	DataSourceName string
}

type VariFlightWebSocketGrpcServer struct {
	config *VariFlightWebSocketGrpcServerConfig

	//variFlightHttpHandler *httpHandler.VariFlightHttpHandler

	WebSocket *gserver.WebSocket `dot:""`

	grpcServer   *grpc.Server
	domainServer *grpcServer.VariFlightGrpcServer
}

func VariFlightWebSocketGrpcServerConfigTypeLive() *dot.ConfigTypeLives {
	return &dot.ConfigTypeLives{
		TypeIdConfig: VariFlightWebSocketGrpcServerTypeId,
		ConfigInfo:   &VariFlightWebSocketGrpcServerConfig{},
	}
}

func VariFlightWebSocketGrpcServerTypeLive() *dot.TypeLives {
	return &dot.TypeLives{
		Meta: dot.Metadata{
			TypeId: VariFlightWebSocketGrpcServerTypeId,
			NewDoter: func(conf []byte) (dot dot.Dot, err error) {
				var toConf = VariFlightWebSocketGrpcServerConfig{}
				if err := dot.UnMarshalConfig(conf, &toConf); err != nil {
					return nil, err
				}
				return &VariFlightWebSocketGrpcServer{config: &toConf}, nil
			},
		},
		Lives: []dot.Live{
			{
				TypeId: VariFlightWebSocketGrpcServerTypeId,
				RelyLives: map[string]dot.LiveId{
					"WebSocket": gserver.WebSocketTypeId,
				},
			},
		},
	}
}

func VariFlightWebSocketGrpcServerAndRelyTypeLives() []*dot.TypeLives {
	allTypeLives := []*dot.TypeLives{VariFlightWebSocketGrpcServerTypeLive()}
	allTypeLives = append(allTypeLives, gserver.WebSocketAndRelyTypeLives()...)
	return allTypeLives
}

func (s *VariFlightWebSocketGrpcServer) AfterAllInject(l dot.Line) {
	s.grpcServer = s.WebSocket.ServerNobl.Server()
	s.domainServer = grpcServer.New(s.config.Appid, s.config.Appsecurity, s.config.MinDurAgainstExtraRequest, s.config.MaxDurAgainstExtraRequest, s.config.DriverName, s.config.DataSourceName)
	_proto.RegisterVariFlightDataServiceServer(s.grpcServer, s.domainServer)
}

// todo: to close the db, otherwise use gorms component.
func (s *VariFlightWebSocketGrpcServer) Stop(ignore bool) error {

}
