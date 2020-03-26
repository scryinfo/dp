// Scry Info.  All rights reserved.
// license that can be found in the license file.

package golang

import (
	"github.com/scryinfo/dot/dots/grpc/gserver"
	"google.golang.org/grpc"

	"github.com/scryinfo/dot/dot"

	"github.com/scryinfo/dp/dots/depot/variflight/golang/protobuf"
)

const (
	GrpcWebSocketServerTypeId = "ca39e667-ddaa-47cb-989b-d888ef4b2585"
)

type GrpcWebSocketServer struct {
	ServiceServer    *GrpcServiceServer `dot:""`
	WebSocketWrapper *gserver.WebSocket `dot:""`
}

func GrpcWebSocketServerTypeLives() []*dot.TypeLives {
	typeLives := []*dot.TypeLives{
		{
			Meta: dot.Metadata{
				TypeId: GrpcWebSocketServerTypeId,
				NewDoter: func(conf []byte) (dot.Dot, error) {
					return &GrpcWebSocketServer{}, nil
				},
			},
			//Lives: []dot.Live{
			//	{
			//		TypeId: GrpcWebSocketServerTypeId,
			//		RelyLives: map[string]dot.LiveId{
			//			"Api": data.ApiTypeId,
			//			"WebSocketWrapper":        gserver.WebSocketTypeId,
			//		},
			//	},
			//},
		},
	}

	typeLives = append(typeLives, GrpcServiceServerTypeLives()...)
	typeLives = append(typeLives, gserver.WebSocketTypeLives()...)
	return typeLives
}

func (s *GrpcWebSocketServer) AfterAllInject(l dot.Line) {
	grpcServer := grpc.NewServer()
	protobuf.RegisterVariFlightDataServiceServer(grpcServer, s.ServiceServer)

	s.WebSocketWrapper.Wrap(grpcServer)
}

