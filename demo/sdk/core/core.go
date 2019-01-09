package core

import (
	"../util/storage/ipfsaccess"
	"./chainevents"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Connector struct {
	ctx context.Context
	conn *ethclient.Client
}

//start
func StartEngine(ethNodeAddr string,
	            protocolContractAddr string,
				protocolContractABI string,
				ipfsNodeAddr string) (*ethclient.Client, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("failed to initialize start engine, error:", err)
		}
	}()

	err := ipfsaccess.GetInstance().Initialize(ipfsNodeAddr)
	if err != nil {
		fmt.Println("failed to initialize ipfs. error: " + err.Error())
		return nil, err
	}

	connector, err := newConnector(ethNodeAddr, protocolContractAddr)
	if err != nil {
		fmt.Println("failed to initialize connector. error: " + err.Error())
		return nil, err
	}

	chainevents.StartEventProcessing(connector.conn, protocolContractAddr, protocolContractABI)

	return connector.conn, nil
}


func newConnector(ethNodeAddr string,
	protocolContractAddr string) (*Connector, error) {
	cn, err := ethclient.Dial(ethNodeAddr)
	if err != nil {
		fmt.Println("failed to connect to node:" + ethNodeAddr)
		return nil, err
	}

	return &Connector{
		ctx: context.Background(),
		conn: cn,
	}, nil
}
