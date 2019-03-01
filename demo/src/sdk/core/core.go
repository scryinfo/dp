package core

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainevents"
	"github.com/scryinfo/iscap/demo/src/sdk/util/accounts"
	"github.com/scryinfo/iscap/demo/src/sdk/util/storage/ipfsaccess"
	rlog "github.com/sirupsen/logrus"
)

type Connector struct {
	ctx  context.Context
	conn *ethclient.Client
}

//start
func StartEngine(ethNodeAddr string,
                    asServiceAddr string,
                    contracts []chainevents.ContractInfo,
                    fromBlock uint64,
                    ipfsNodeAddr string) (*ethclient.Client, error) {

	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to initialize start engine, error:", err)
		}
	}()

	err := ipfsaccess.GetIAInstance().Initialize(ipfsNodeAddr)
	if err != nil {
		rlog.Error("failed to initialize ipfs. error: ", err)
		return nil, err
	}

	connector, err := newConnector(ethNodeAddr)
	if err != nil {
		rlog.Error("failed to initialize connector. error: ", err)
		return nil, err
	}

	err = accounts.GetAMInstance().Initialize(asServiceAddr)
	if err != nil {
		rlog.Error("failed to initialize account service, error:", err)
		return nil, err
	}

	chainevents.StartEventProcessing(connector.conn, contracts, fromBlock)

	return connector.conn, nil
}

func newConnector(ethNodeAddr string) (*Connector, error) {
	cn, err := ethclient.Dial(ethNodeAddr)
	if err != nil {
		rlog.Error("failed to connect to node:" + ethNodeAddr)
		return nil, err
	}

	return &Connector{
		ctx:  context.Background(),
		conn: cn,
	}, nil
}
