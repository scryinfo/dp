package core

import (
	"./chainevents"
	"github.com/ethereum/go-ethereum/ethclient"
)

//start
func StartEngine(conn *ethclient.Client,
				protocolContractAddr string,
				protocolContractABI string) {
	chainevents.StartEventProcessing(conn, protocolContractAddr, protocolContractABI)
}



