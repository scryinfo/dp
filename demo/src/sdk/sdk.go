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

var (
	START_ENGINE_FAILED                    = "failed to start engine"
	INIT_CONTRACT_INTERFACE_WRAPPER_FAILED = "failed to initialize contract interface"
	LOAD_PATH_FAILED                       = "failed to load log path"
	INIT_SDK_FAILED                        = "failed to initialize sdk"
)

func Init(ethNodeAddr string,
            asServiceAddr string,
            contracts []chainevents.ContractInfo,
            fromBlock uint64,
            ipfsNodeAddr string) error {

	err := initLog()
	if err != nil {
		fmt.Println(INIT_SDK_FAILED, err)
		return err
	}

	conn, err := core.StartEngine(ethNodeAddr, asServiceAddr, contracts, fromBlock, ipfsNodeAddr)
	if err != nil {
		rlog.Error(START_ENGINE_FAILED, err)
		return errors.New(START_ENGINE_FAILED)
	}

	err = chaininterfacewrapper.Initialize(common.HexToAddress(contracts[0].Address),
                                            common.HexToAddress(contracts[1].Address),
                                            conn)
	if err != nil {
		rlog.Error(INIT_CONTRACT_INTERFACE_WRAPPER_FAILED, err)
		return errors.New(INIT_CONTRACT_INTERFACE_WRAPPER_FAILED)
	}

	return nil
}

func initLog() error {
	ph, err := settings.LoadLogPath()
	if err != nil {
		fmt.Println(LOAD_PATH_FAILED, err)
		return err
	}

	filePath := ph.Dir + "/" + ph.File
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(LOAD_PATH_FAILED, err)
		return err
	}

	rlog.SetFormatter(&rlog.TextFormatter{})
	rlog.SetOutput(f)
	rlog.SetLevel(rlog.DebugLevel)

	return nil
}
