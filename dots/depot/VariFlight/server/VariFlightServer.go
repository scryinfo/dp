package server

import (
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"github.com/davecgh/go-spew/fmt"

	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/grpc/gserver"
	"github.com/scryinfo/dp/dots/depot/VariFlight/server/_proto"
	VariFlight "github.com/scryinfo/dp/dots/depot/VariFlight/server/VariFlightApiCaller"
)

const (
	VariFlightServerTypeId = "ca39e667-ddaa-47cb-989b-d888ef4b2585"
)

var defaultVariFlightServerConfig = VariFlightServerConfig{
	ServicePath: "/scry/variflight",
}

type VariFlightServerConfig struct {
	ServicePath string `json:"service_path, omitempty"`
}

type VariFlightServer struct {
	config *VariFlightServerConfig

	VariFlightApiCaller *VariFlight.VariFlightApiCaller `dot:""`
	WebSocket           *gserver.WebSocket              `dot:""`

	serviceServer *variflightServiceServer
	grpcServer    *grpc.Server
}

func VariFlightServerTypeLives() []*dot.TypeLives {
	typeLives := []*dot.TypeLives{
		{
			Meta: dot.Metadata{
				TypeId: VariFlightServerTypeId,
				NewDoter: func(conf []byte) (dot.Dot, error) {
					_conf := defaultVariFlightServerConfig
					if err := dot.UnMarshalConfig(conf, &_conf); err != nil {
						dot.Logger().Debugln("UnMarshalConfig(VariFlightServerConfig) failed.", zap.Error(err))
						os.Exit(1)
					}
					dot.Logger().Debug(func() string {
						return fmt.Sprintf("VariFlightServerConfig: %v", _conf)
					})
					return &VariFlightServer{config: &_conf}, nil
				},
			},
			//Lives: []dot.Live{
			//	{
			//		TypeId: VariFlightServerTypeId,
			//		RelyLives: map[string]dot.LiveId{
			//			"VariFlightApiCaller": VariFlight.VariFlightApiCallerTypeId,
			//			"WebSocket":        gserver.WebSocketTypeId,
			//		},
			//	},
			//},
		},
	}

	typeLives = append(typeLives, VariFlight.VariFlightCallerTypeLives()...)
	typeLives = append(typeLives, gserver.WebSocketTypeLives()...)
	return typeLives
}

func (s *VariFlightServer) AfterAllInject(l dot.Line) {
	s.grpcServer = grpc.NewServer()
	s.serviceServer = newVariFlightServiceServer(s.VariFlightApiCaller)
	_proto.RegisterVariFlightDataServiceServer(s.grpcServer, s.serviceServer)

	s.WebSocket.GET(s.config.ServicePath, s.grpcServer)
}
