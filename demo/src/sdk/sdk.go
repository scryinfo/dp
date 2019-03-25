package sdk

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/scryinfo/iscap/demo/src/sdk/core"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainevents"
	"github.com/scryinfo/iscap/demo/src/sdk/scryclient/chaininterfacewrapper"
	"github.com/scryinfo/iscap/demo/src/sdk/settings"
	rlog "github.com/sirupsen/logrus"
	"os"
)

const (
	LoadPathFailed                     = "failed to load log path"
)

var err error

func Init(ethNodeAddr string, contracts []chainevents.ContractInfo, fromBlock uint64) error {
	var conn *ethclient.Client
	if conn, err = core.StartEngine(ethNodeAddr, contracts, fromBlock); err != nil {
		return errors.Wrap(err, "SDK init failed. ")
	}


	if err = chaininterfacewrapper.Initialize(common.HexToAddress(contracts[0].Address),
		common.HexToAddress(contracts[1].Address), conn); err != nil {
		return errors.Wrap(err, "Contract interface init failed. ")
	}

	return nil
}

func InitLog() error {
	ph, err := settings.LoadLogPath()
	if err != nil {
		fmt.Println(LoadPathFailed, err)
		return err
	}

	filePath := ph.Dir + "/" + ph.File
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(LoadPathFailed, err)
		return err
	}

	rlog.SetFormatter(&rlog.TextFormatter{})
	rlog.SetOutput(f)
	rlog.SetLevel(rlog.DebugLevel)

	return nil
}
