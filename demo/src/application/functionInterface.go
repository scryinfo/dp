package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scryinfo/iscap/demo/src/sdk"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainevents"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainoperations"
	"github.com/scryinfo/iscap/demo/src/sdk/scryclient"
	cif "github.com/scryinfo/iscap/demo/src/sdk/scryclient/chaininterfacewrapper"
	"github.com/scryinfo/iscap/demo/src/sdk/util/accounts"
	"io/ioutil"
	"math/big"
)

var (
	protocolContractAddr = "0xbb7bae05bdbc0ed9e514ce18122fc6b4cbcca346"
	tokenContractAddr    = "0xc67d1847fb1b00173dcdbc00c7cbe32651537daa"
	deployerKeyJson      = `{"version":3,"id":"8db8b2a0-ec6e-40ea-9808-631117870070","address":"61ad28110ce3911a9aafabba551cdc932a02bd52","crypto":{"ciphertext":"b4835e7a3ea3a132b172f1609ed310b7345323c552791b36017d761e6fe748f0","cipherparams":{"iv":"880c3c504350c97d6b5469d9333c3feb"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"2ae6c42c17a67f271e15de48b743b9a9c400b1413f8d9ccbf8389be86e84b938","n":262144,"r":8,"p":1},"mac":"5c26d0cf4925208e40f1a791d713a2a71b7287b7c09fdf192c21ad8fe158b388"}}`
	keyPassword          = "12345"
)

func init() {
	err := sdk.Init("http://127.0.0.1:7545/", "192.168.1.6:48080", getContracts(), 0, "/ip4/127.0.0.1/tcp/5001", common.HexToAddress(protocolContractAddr), common.HexToAddress(tokenContractAddr))
	if err != nil {
		fmt.Println("failed to initialize sdk, error:", err)
		return
	}
}

//func testAccounts() {
//	fmt.Println("Start testing accounts...")
//
//	ac, err := accounts.GetAMInstance().CreateAccount("12345")
//	if err != nil {
//		fmt.Println("failed to create account, error:", err)
//	}
//
//	rv, err := accounts.GetAMInstance().AuthAccount(ac.Address, "12345")
//	if err != nil {
//		fmt.Println("failed to authenticate account, error:", err)
//	}
//
//	if rv {
//		fmt.Println("Account authentication passed")
//	} else {
//		fmt.Println("Account authentication not passed")
//	}
//
//	fmt.Println("Test end")
//}

func getClient(keyJson string) *scryclient.ScryClient {
	client, err := scryclient.NewScryClient(getPublicAddress(keyJson))
	if err != nil {
		fmt.Println("failed to create contract client. error:", err)
		return nil
	}

	return client
}

func ImportAccount(keyJson string, oldPassword string, newPassword string) (*scryclient.ScryClient, error) {
	address, err := accounts.GetAMInstance().ImportAccount([]byte(keyJson), oldPassword, newPassword)
	if err != nil {
		fmt.Println("failed to import account. error:", err)
		return nil, err
	}

	client, err := scryclient.NewScryClient(address)
	if err != nil {
		fmt.Println("failed to create contract client. error:", err)
		return nil, err
	}

	return client, nil
}

func SellerPublishData(pubData PubData) (string, bool) {
	var pd [][]byte = make([][]byte, len(pubData.ProofData))
	for i := 0; i < len(pubData.ProofData); i++ {
		pd[i] = []byte(pubData.ProofData[i])
	}
	var seller = common.BytesToAddress([]byte(pubData.Seller))

	txParam := chainoperations.TransactParams{seller, keyPassword, big.NewInt(0), false}
	result, err := cif.Publish(&txParam, pubData.Price, []byte(pubData.MetaData), pd, len(pubData.ProofData), []byte(pubData.DespData), false)
	ok := true
	if err != nil {
		ok = false
		result = err.Error()
	}
	return result, ok
}

func getContracts() []chainevents.ContractInfo {
	protocolEvents := []string{"DataPublish", "TransactionCreate", "RegisterVerifier", "VerifiersChosen", "Vote", "Buy", "ReadyForDownload", "TransactionClose"}
	tokenEvents := []string{"Approval"}

	contracts := []chainevents.ContractInfo{
		{protocolContractAddr, getAbiText("./ScryProtocol.abi"), protocolEvents},
		{tokenContractAddr, getAbiText("./ScryToken.abi"), tokenEvents},
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

func getPublicAddress(keyJson string) string {
	publicAddress := chainoperations.DecodeKeystoreAddress([]byte(keyJson))
	return publicAddress
}
