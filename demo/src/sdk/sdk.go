package sdk

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryinfo/iscap/demo/src/sdk/core"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainevents"
	"github.com/scryinfo/iscap/demo/src/sdk/scryclient/chaininterfacewrapper"
	"github.com/scryinfo/iscap/demo/src/sdk/settings"
	rlog "github.com/sirupsen/logrus"
	"os"
)

const (
	StartEngineFailed                  = "failed to start engine"
	InitConTractInterfaceWrapperFailed = "failed to initialize contract interface"
	LoadPathFailed                     = "failed to load log path"
)

func Init(ethNodeAddr string,
	contracts []chainevents.ContractInfo,
	fromBlock uint64) error {

	conn, err := core.StartEngine(ethNodeAddr, contracts, fromBlock)
	if err != nil {
		rlog.Error(StartEngineFailed, err)
		return errors.New(StartEngineFailed)
	}

	err = chaininterfacewrapper.Initialize(common.HexToAddress(contracts[0].Address),
		common.HexToAddress(contracts[1].Address),
		conn)
	if err != nil {
		rlog.Error(InitConTractInterfaceWrapperFailed, err)
		return errors.New(InitConTractInterfaceWrapperFailed)
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
