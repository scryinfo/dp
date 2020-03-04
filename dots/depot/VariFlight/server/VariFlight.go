// Scry Info.  All rights reserved.
// license that can be found in the license file.

package server

import (
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/grpc/gserver"

	"github.com/scryinfo/dp/dots/depot/VariFlight/server/_proto"
	"github.com/scryinfo/dp/dots/depot/VariFlight/server/grpcServer"
)

const (
	VariFlightWebSocketGrpcServerTypeId = "70c1a120-03af-4e55-a6fe-feb44b5d2761"

	defaultServicePath = "/variflight"
)

// VariFlightWebSocketGrpcServer component serves VariFlight flight data against gRPC-WebSocket request,
// initially routed by the HTTP GET method and the given path.
//
// If ServicePath is not configured, then use defaultServicePath instead.
//
// todo: rename concisely
type VariFlightWebSocketGrpcServer struct {
	WebSocket *gserver.WebSocket `dot:""`

	grpcServer              *grpc.Server
	variFlightServiceServer *grpcServer.VariFlightGrpcServer

	config *VariFlightWebSocketGrpcServerConfig
}

type VariFlightWebSocketGrpcServerConfig struct {
	// required parameters for data source VariFlight API validation
	Appid       string `json:"appid"`
	Appsecurity string `json:"appsecurity"`

	// time control parameters for avoiding unnecessary extra API calling
	MinDurAgainstExtraRequest time.Duration `json:"minDurAgainstExtraRequest"`
	MaxDurAgainstExtraRequest time.Duration `json:"maxDurAgainstExtraRequest"`

	// arguments for connecting a database
	DriverName     string `json:"driver_name"`
	DataSourceName string `json:"data_source_name"`

	// path to route grpc service
	ServicePath string `json:"service_path"`
}

func VariFlightWebSocketGrpcServerConfigTypeLive() *dot.ConfigTypeLives {
	return &dot.ConfigTypeLives{
		TypeIdConfig: VariFlightWebSocketGrpcServerTypeId,
		ConfigInfo:   &VariFlightWebSocketGrpcServerConfig{},
	}
}

func VariFlightWebSocketGrpcServerTypeLives() []*dot.TypeLives {
	variFlightWebSocketGrpcServerTypeLive := func() *dot.TypeLives {
		return &dot.TypeLives{
			Meta: dot.Metadata{
				TypeId: VariFlightWebSocketGrpcServerTypeId,
				NewDoter: func(conf []byte) (dot.Dot, error) {
					var toConf = VariFlightWebSocketGrpcServerConfig{}
					if err := dot.UnMarshalConfig(conf, &toConf); err != nil {
						return nil, err
					}
					if toConf.ServicePath == "" {
						toConf.ServicePath = defaultServicePath
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
	allTypeLives := []*dot.TypeLives{variFlightWebSocketGrpcServerTypeLive()}
	allTypeLives = append(allTypeLives, gserver.WebSocketTypeLives()...)
	return allTypeLives
}

func (s *VariFlightWebSocketGrpcServer) AfterAllInject(l dot.Line) {
	s.grpcServer = grpc.NewServer()
	s.variFlightServiceServer = grpcServer.New(s.config.Appid, s.config.Appsecurity, s.config.MinDurAgainstExtraRequest, s.config.MaxDurAgainstExtraRequest, s.config.DriverName, s.config.DataSourceName)
	_proto.RegisterVariFlightDataServiceServer(s.grpcServer, s.variFlightServiceServer)

	s.WebSocket.GET(s.config.ServicePath, s.grpcServer)
}

// todo: code to close the db, otherwise use a common db component (e.g. gorms).
func (s *VariFlightWebSocketGrpcServer) Stop(ignore bool) error {
	if s.grpcServer != nil {
		s.grpcServer = nil
	}
	if s.variFlightServiceServer != nil {
		s.variFlightServiceServer = nil
	}
	return nil
}

// For test purpose only.
func NewVariFlightWebSocketGrpcServerTest() *VariFlightWebSocketGrpcServer {
	config := VariFlightWebSocketGrpcServerConfig {
		Appid: "VariFlight_test_appid",
		Appsecurity: "VariFlight_test_appsecurity",

		MinDurAgainstExtraRequest: time.Minute * 30,
		MaxDurAgainstExtraRequest: time.Hour * 24,

		DriverName: "postgres",
		DataSourceName: "",

		ServicePath: "",
	}

	bytes, err := json.Marshal(&config);
	if err != nil {
		dot.Logger().Fatal(func() string {
			return fmt.Sprintf("failed to marshal config, error: %v", err)
		})
	}

	newer := func(conf []byte) (dot.Dot, error) {
		var toConf = VariFlightWebSocketGrpcServerConfig{}
		if err := dot.UnMarshalConfig(conf, &toConf); err != nil {
			return nil, err
		}
		if toConf.ServicePath == "" {
			toConf.ServicePath = defaultServicePath
		}
		return &VariFlightWebSocketGrpcServer{config: &toConf}, nil
	}

	d, err := newer(bytes)
	if err != nil {
		dot.Logger().Fatal(func() string {
			return fmt.Sprintf("failed to get new dot, error: %v", err)
		})
	}

	cp, ok := d.(*VariFlightWebSocketGrpcServer)
	if !ok {
		dot.Logger().Fatal(func() string {
			return fmt.Sprint("obtained component is not *VariFlightWebSocketGrpcServer")
		})
	}

	cp.AfterAllInject(nil)

	return cp
}