// Scry Info.  All rights reserved.
// license that can be found in the license file.

package core

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/scryinfo/dp/demo/src/sdk/core/chainevents"
	"github.com/scryinfo/dp/demo/src/sdk/util/accounts"
	"github.com/scryinfo/dp/demo/src/sdk/util/storage/ipfsaccess"
	rlog "github.com/sirupsen/logrus"
)

// Connector
type Connector struct {
	ctx  context.Context
	conn *ethclient.Client
}

// StartEngine start
func StartEngine(ethNodeAddr string,
	asServiceAddr string,
	contracts []chainevents.ContractInfo,
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
		rlog.Error("failed to initialize interface service, error:", err)
		return nil, err
	}

	chainevents.StartEventProcessing(connector.conn, contracts)

	return connector.conn, nil
}

// StartScan
func StartScan(fromBlock uint64) {
	chainevents.SetFromBlock(fromBlock)
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
