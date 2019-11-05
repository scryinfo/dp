package grpc

import "github.com/scryinfo/dp/dots/binary/scry"

// BinaryGrpcServer
type BinaryGrpcServer interface {
	SetChainWrapper(w scry.ChainWrapper)
}
