package _go

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"github.com/davecgh/go-spew/fmt"

	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/grpc/gserver"
	"github.com/scryinfo/dp/dots/depot/VariFlight/go/_proto"
	VariFlight "github.com/scryinfo/dp/dots/depot/VariFlight/go/VariFlightApiCaller"
)

const (
	VariFlightServerTypeId = "ca39e667-ddaa-47cb-989b-d888ef4b2585"
)

type VariFlightServer struct {
	VariFlightApiCaller *VariFlight.VariFlightApiCaller `dot:""`
	WebSocket           *gserver.WebSocket              `dot:""`
}

func VariFlightServerTypeLives() []*dot.TypeLives {
	typeLives := []*dot.TypeLives{
		{
			Meta: dot.Metadata{
				TypeId: VariFlightServerTypeId,
				NewDoter: func(conf []byte) (dot.Dot, error) {
					return &VariFlightServer{}, nil
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

func (s *VariFlightServer) AfterAllInject(l dot.Line) {s.grpcServer = grpc.NewServer()
	_proto.RegisterVariFlightDataServiceServer(grpc.NewServer(), newVariFlightServiceServer(s.VariFlightApiCaller))

	s.WebSocket.Wrap(s.grpcServer)
}
