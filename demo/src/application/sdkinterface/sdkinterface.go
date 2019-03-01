package sdkinterface

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scryinfo/iscap/demo/src/sdk"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainevents"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainoperations"
	"github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
	"github.com/scryinfo/iscap/demo/src/sdk/scryclient"
	cif "github.com/scryinfo/iscap/demo/src/sdk/scryclient/chaininterfacewrapper"
	"io/ioutil"
	"math/big"
)

var (
	protocolContractAddr                        = "0xbb7bae05bdbc0ed9e514ce18122fc6b4cbcca346"
	tokenContractAddr                           = "0xc67d1847fb1b00173dcdbc00c7cbe32651537daa"
	keyPassword                                 = "12345"
)

func init() {
	err := sdk.Init("http://127.0.0.1:7545/", "192.168.1.6:48080", getContracts(), 0, "/ip4/127.0.0.1/tcp/5001", common.HexToAddress(protocolContractAddr), common.HexToAddress(tokenContractAddr))
	if err != nil {
		fmt.Println("failed to initialize sdk, error:", err)
		return
	}
}

func SellerPublishData(pubData PubData) (string, error) {
	se, _ := scryclient.NewScryClient(pubData.Seller)	// change to global variable.
	se.SubscribeEvent("DataPublish", onPublish)

	var pd [][]byte = make([][]byte, len(pubData.ProofData))
	for i := 0; i < len(pubData.ProofData); i++ {
		pd[i] = []byte(pubData.ProofData[i])
	}
	var s = common.BytesToAddress([]byte(pubData.Seller))
	var p *big.Int
	var ok bool
	if p, ok = new(big.Int).SetString(pubData.Price, 10);!ok {
		return "Set price failed.", nil
	}

	txParam := chainoperations.TransactParams{s, keyPassword, big.NewInt(0), false}
	result, err := cif.Publish(&txParam, p, []byte(pubData.MetaData), pd, len(pubData.ProofData), []byte(pubData.DespData), false)
	return result, err
}

func getContracts() []chainevents.ContractInfo {
	protocolEvents := []string{"DataPublish", "TransactionCreate", "RegisterVerifier", "VerifiersChosen", "Vote", "Buy", "ReadyForDownload", "TransactionClose"}
	tokenEvents := []string{"Approval"}

	contracts := []chainevents.ContractInfo{
		{protocolContractAddr, getAbiText("github.com/scryinfo/iscap/demo/src/testconsole/ScryProtocol.abi"), protocolEvents},
		{tokenContractAddr, getAbiText("github.com/scryinfo/iscap/demo/src/testconsole/ScryToken.abi"), tokenEvents},
	}

	return contracts
}

func getAbiText(fileName string) string {
	abi, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("failed to read abi text", err)
		return ""
	}

	return string(abi)
}

func onPublish(event events.Event) bool {
	return true
}
