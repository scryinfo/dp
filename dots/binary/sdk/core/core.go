// Scry Info.  All rights reserved.
// license that can be found in the license file.

package core

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	chainevents2 "github.com/scryinfo/dp/dots/binary/sdk/core/chainevents"
	"go.uber.org/zap"
)

type Connector struct {
	ctx  context.Context
	conn *ethclient.Client
}

//start
func StartEngine(ethNodeAddr string,
	contracts []chainevents2.ContractInfo,
) (*ethclient.Client, error) {
	logger := dot.Logger()

	defer func() {
		if er := recover(); er != nil {
			logger.Errorln("", zap.Any("failed to initialize start engine, error:", er))
		}
	}()

	connector, err := newConnector(ethNodeAddr)
	if err != nil {
		logger.Errorln("", zap.NamedError("failed to initialize connector. error: ", err))
		return nil, err
	}

	chainevents2.StartEventProcessing(connector.conn, contracts)

	return connector.conn, nil
}

func StartScan(fromBlock uint64) {
	chainevents2.SetFromBlock(fromBlock)
}

func newConnector(ethNodeAddr string) (*Connector, error) {
	cn, err := ethclient.Dial(ethNodeAddr)
	if err != nil {
		return nil, errors.Wrap(err, "Connect to node: "+ethNodeAddr+" failed. ")
	}

	return &Connector{
		ctx:  context.Background(),
		conn: cn,
	}, nil
}
