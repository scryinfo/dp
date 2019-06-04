// Scry Info.  All rights reserved.
// license that can be found in the license file.

package sdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryinfo/dp/dots/binary/core"
	"github.com/scryinfo/dp/dots/binary/core/chainevents"
	"github.com/scryinfo/dp/dots/binary/scry"
	"github.com/scryinfo/dp/dots/binary/sdk/settings"
)

const (
	startEngineFailed         = "failed to start engine"
	initContractWrapperFailed = "failed to initialize contract interface"

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

func Init(
	ethNodeAddr string,
	protocolAddr string,
	tokenAddr string,
	keyServiceAddr string,
	ipfsNodeAddr string,
	appId string,
) (scry.ChainWrapper, error) {
	settings.SetAppId(appId)

	contracts := getContracts(protocolAddr, tokenAddr)
	conn, err := core.StartEngine(
		ethNodeAddr,
		keyServiceAddr,
		contracts,
		ipfsNodeAddr)
	if err != nil {
		return nil, errors.New(startEngineFailed)
	}

	//todo
	chain, err := scry.NewChainWrapper(
		common.HexToAddress(contracts[0].Address),
		common.HexToAddress(contracts[1].Address),
		conn)
	if err != nil {
		return nil, errors.New(initContractWrapperFailed)
	}

	return chain, err
}

func getContracts(
	protocolAddr string,
	tokenAddr string) []chainevents.ContractInfo {
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

	contracts := []chainevents.ContractInfo{
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

func StartScan(fromBlock uint64) {
	core.StartScan(fromBlock)
}
