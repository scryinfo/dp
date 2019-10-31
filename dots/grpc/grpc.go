package grpc

import "github.com/scryinfo/dp/dots/binary/scry"

type BinaryGrpcServer interface {
	SetChainWrapper(w scry.ChainWrapper)
}
