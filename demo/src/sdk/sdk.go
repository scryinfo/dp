// Scry Info.  All rights reserved.
// license that can be found in the license file.

package sdk

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryinfo/dp/demo/src/sdk/core"
	ce "github.com/scryinfo/dp/demo/src/sdk/core/chainevents"
	"github.com/scryinfo/dp/demo/src/sdk/scryclient/chaininterfacewrapper"
	"github.com/scryinfo/dp/demo/src/sdk/settings"
	rlog "github.com/sirupsen/logrus"
	"os"
)

const (
	startEngineFailed         = "failed to start engine"
	initContractWrapperFailed = "failed to initialize contract interface"
	loadPathFailed            = "failed to load log path"
	initSdkFailed             = "failed to initialize sdk"

	protocolAbi = `[
    {
      "inputs": [
        {
          "name": "_token",
          "type": "address"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "seqNo",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "users",
          "type": "address[]"
        }
      ],
      "name": "RegisterVerifier",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "seqNo",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "publishId",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "price",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "despDataId",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "supportVerify",
          "type": "bool"
        },
        {
          "indexed": false,
          "name": "users",
          "type": "address[]"
        }
      ],
      "name": "DataPublish",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "seqNo",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "transactionId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "publishId",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "proofIds",
          "type": "bytes32[]"
        },
        {
          "indexed": false,
          "name": "state",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "users",
          "type": "address[]"
        }
      ],
      "name": "VerifiersChosen",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "seqNo",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "transactionId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "publishId",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "proofIds",
          "type": "bytes32[]"
        },
        {
          "indexed": false,
          "name": "needVerify",
          "type": "bool"
        },
        {
          "indexed": false,
          "name": "state",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "users",
          "type": "address[]"
        }
      ],
      "name": "TransactionCreate",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "seqNo",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "transactionId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "judge",
          "type": "bool"
        },
        {
          "indexed": false,
          "name": "comments",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "state",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "index",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "users",
          "type": "address[]"
        }
      ],
      "name": "Vote",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "seqNo",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "transactionId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "publishId",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "metaDataIdEncSeller",
          "type": "bytes"
        },
        {
          "indexed": false,
          "name": "state",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "buyer",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "index",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "users",
          "type": "address[]"
        }
      ],
      "name": "Buy",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "seqNo",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "transactionId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "metaDataIdEncBuyer",
          "type": "bytes"
        },
        {
          "indexed": false,
          "name": "state",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "index",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "users",
          "type": "address[]"
        }
      ],
      "name": "ReadyForDownload",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "seqNo",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "transactionId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "state",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "index",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "users",
          "type": "address[]"
        }
      ],
      "name": "TransactionClose",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "seqNo",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "verifier",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "users",
          "type": "address[]"
        }
      ],
      "name": "VerifierDisable",
      "type": "event"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "seqNo",
          "type": "string"
        }
      ],
      "name": "registerAsVerifier",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "seqNo",
          "type": "string"
        },
        {
          "name": "publishId",
          "type": "string"
        },
        {
          "name": "price",
          "type": "uint256"
        },
        {
          "name": "metaDataIdEncSeller",
          "type": "bytes"
        },
        {
          "name": "proofDataIds",
          "type": "bytes32[]"
        },
        {
          "name": "despDataId",
          "type": "string"
        },
        {
          "name": "supportVerify",
          "type": "bool"
        }
      ],
      "name": "publishDataInfo",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "seqNo",
          "type": "string"
        },
        {
          "name": "publishId",
          "type": "string"
        },
        {
          "name": "startVerify",
          "type": "bool"
        }
      ],
      "name": "createTransaction",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "seqNo",
          "type": "string"
        },
        {
          "name": "txId",
          "type": "uint256"
        },
        {
          "name": "judge",
          "type": "bool"
        },
        {
          "name": "comments",
          "type": "string"
        }
      ],
      "name": "vote",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "seqNo",
          "type": "string"
        },
        {
          "name": "txId",
          "type": "uint256"
        }
      ],
      "name": "buyData",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "seqNo",
          "type": "string"
        },
        {
          "name": "txId",
          "type": "uint256"
        }
      ],
      "name": "cancelTransaction",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "seqNo",
          "type": "string"
        },
        {
          "name": "txId",
          "type": "uint256"
        },
        {
          "name": "encryptedMetaDataId",
          "type": "bytes"
        }
      ],
      "name": "submitMetaDataIdEncWithBuyer",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "seqNo",
          "type": "string"
        },
        {
          "name": "txId",
          "type": "uint256"
        },
        {
          "name": "truth",
          "type": "bool"
        }
      ],
      "name": "confirmDataTruth",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "deposit",
          "type": "uint256"
        }
      ],
      "name": "setVerifierDepositToken",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "num",
          "type": "uint8"
        }
      ],
      "name": "setVerifierNum",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "bonus",
          "type": "uint256"
        }
      ],
      "name": "setVerifierBonus",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "seqNo",
          "type": "string"
        },
        {
          "name": "txId",
          "type": "uint256"
        },
        {
          "name": "verifierIndex",
          "type": "uint8"
        },
        {
          "name": "credit",
          "type": "uint8"
        }
      ],
      "name": "creditsToVerifier",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`
	tokenAbi = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"INITIAL_SUPPLY","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_subtractedValue","type":"uint256"}],"name":"decreaseApproval","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_addedValue","type":"uint256"}],"name":"increaseApproval","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`
)

var err error

// Init init user service, contracts and SDK instance
func Init(
	ethNodeAddr string,
	keyServiceAddr string,
	protocolAddr string,
	tokenAddr string,
	ipfsNodeAddr string,
	logPath string,
	appId string,
) error {
	settings.SetAppId(appId)
	settings.SetLogPath(logPath)

	err = initLog()
	if err != nil {
		fmt.Println(initSdkFailed, err)
		return err
	}

	contracts := getContracts(protocolAddr, tokenAddr)
	conn, err := core.StartEngine(
		ethNodeAddr,
		keyServiceAddr,
		contracts,
		ipfsNodeAddr)
	if err != nil {
		rlog.Error(startEngineFailed, err)
		return errors.New(startEngineFailed)
	}

	err = chaininterfacewrapper.Initialize(
		common.HexToAddress(contracts[0].Address),
		common.HexToAddress(contracts[1].Address),
		conn)
	if err != nil {
		rlog.Error(initContractWrapperFailed, err)
		return errors.New(initContractWrapperFailed)
	}

	return nil
}

func getContracts(
	protocolAddr string,
	tokenAddr string) []ce.ContractInfo {
	protocolEvents := []string{
		"DataPublish",
		"TransactionCreate",
		"RegisterVerifier",
		"VerifiersChosen",
		"Vote",
		"Buy",
		"ReadyForDownload",
		"TransactionClose",
		"VerifierDisable",
	}
	tokenEvents := []string{"Approval"}

	contracts := []ce.ContractInfo{
		{
			protocolAddr,
			protocolAbi,
			protocolEvents,
		}, {
			tokenAddr,
			tokenAbi,
			tokenEvents,
		},
	}

	return contracts
}

func initLog() error {
	filePath := settings.GetLogPath()
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(loadPathFailed, err)
		return err
	}

	rlog.SetFormatter(&rlog.TextFormatter{})
	rlog.SetOutput(f)
	rlog.SetLevel(rlog.DebugLevel)

	return nil
}

// StartScan start scan
func StartScan(fromBlock uint64) {
	core.StartScan(fromBlock)
}
