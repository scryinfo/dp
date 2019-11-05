package grpc

import "github.com/scryinfo/dp/dots/binary/scry"

// BinaryGrpcServer binary grpc server
type BinaryGrpcServer interface {
	SetChainWrapper(w scry.ChainWrapper)
}
