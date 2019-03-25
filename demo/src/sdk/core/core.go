package core

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainevents"
	rlog "github.com/sirupsen/logrus"
)

type Connector struct {
	ctx  context.Context
	conn *ethclient.Client
}

//start
func StartEngine(ethNodeAddr string,
	contracts []chainevents.ContractInfo,
	fromBlock uint64) (*ethclient.Client, error) {

	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to initialize start engine, error:", err)
		}
	}()

	connector, err := newConnector(ethNodeAddr)
	if err != nil {
		return nil, errors.Wrap(err, "Init connector failed. ")
	}

	chainevents.StartEventProcessing(connector.conn, contracts, fromBlock)

	return connector.conn, nil
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
