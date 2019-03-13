package sdkinterface

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scryinfo/iscap/demo/src/application/definition"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface/settings"
	"github.com/scryinfo/iscap/demo/src/sdk"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainevents"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainoperations"
	"github.com/scryinfo/iscap/demo/src/sdk/scryclient"
	cif "github.com/scryinfo/iscap/demo/src/sdk/scryclient/chaininterfacewrapper"
	"github.com/scryinfo/iscap/demo/src/sdk/util/accounts"
	rlog "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"strings"
)

const (
	failedToInitSDK = "failed to initialize sdk."
	sep             = "|"
)

var (
	curUser  *scryclient.ScryClient = nil
	deployer *scryclient.ScryClient = nil
	scryInfo *settings.ScryInfo     = nil
)

func Initialize() error {
	// load definition
	var err error
	scryInfo, err = settings.LoadSettings()
	if err != nil {
		fmt.Println(failedToInitSDK, err)
		return err
	}

	// initialization
	contracts := getContracts(scryInfo.Chain.Contracts.ProtocolAddr,
		scryInfo.Chain.Contracts.TokenAddr,
		scryInfo.Chain.Contracts.ProtocolAbiPath,
		scryInfo.Chain.Contracts.TokenAbiPath,
		scryInfo.Chain.Contracts.ProtocolEvents,
		scryInfo.Chain.Contracts.TokenEvents)

	err = sdk.Init(scryInfo.Chain.Ethereum.EthNode,
		scryInfo.Services.Keystore,
		contracts,
		0,
		scryInfo.Services.Ipfs)
	if err != nil {
		fmt.Println(failedToInitSDK, err)
		return err
	}
	rlog.Info("Node: sdk init finished. ")

	return nil
}

//new user
func CreateUser(password string) (*scryclient.ScryClient, error) {
	client, err := scryclient.CreateScryClient(password)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateUserWithLogin(password string) (*scryclient.ScryClient, error) {
	client, err := scryclient.CreateScryClient(password)
	if err != nil {
		return nil, err
	}

	curUser = client

	return client, nil
}

func UserLogin(address string, password string) (bool, error) {
	client := scryclient.NewScryClient(address)
	if client == nil {
		return false, errors.New("failed to call NewScryClient")
	}

	succ, err := client.Authenticate(password)
	if err != nil {
		return false, errors.New("failed to authenticate user")
	}

	if succ {
		curUser = client
	}

	return succ, nil
}

func ImportAccount(keyJson string, oldPassword string, newPassword string) (*scryclient.ScryClient, error) {
	address, err := accounts.GetAMInstance().ImportAccount([]byte(keyJson), oldPassword, newPassword)
	if err != nil {
		fmt.Println("failed to import account. error:", err)
		return nil, err
	}

	client := scryclient.NewScryClient(address)
	return client, nil
}

func TransferTokenFromDeployer(token *big.Int) error {
	var err error
	if deployer == nil {
		deployer, err = ImportAccount(scryInfo.Chain.Contracts.DeployerKeyJson,
			scryInfo.Chain.Contracts.DeployerPassword,
			scryInfo.Chain.Contracts.DeployerPassword)
		if err != nil {
			fmt.Println("failed to transfer token, error:", err)
			return err
		}
	}

	if curUser == nil {
		fmt.Println("failed to transfer token, null current user")
		return errors.New("failed to transfer token, null current user")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(deployer.Account.Address),
		Password: scryInfo.Chain.Contracts.DeployerPassword,
		Value:    big.NewInt(0),
		Pending:  false}
	err = cif.TransferTokens(&txParam, common.HexToAddress(curUser.Account.Address), token)
	if err != nil {
		fmt.Println("failed to transfer token, error:", err)
		return err
	}

	return nil
}

func PublishData(data *definition.PubDataIDs, cb chainevents.EventCallback) (string, error) {
	if curUser == nil {
		fmt.Println("no current user")
		return "", errors.New("failed to publish data, current user is null")
	}

    if cb == nil {
        fmt.Println("null callback function")
        return "", errors.New("failed to publish data, callback function is null")
    }

	curUser.SubscribeEvent("DataPublish", cb)
	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: data.Password,
		Value:    big.NewInt(0),
		Pending:  false}

	return cif.Publish(&txParam,
		big.NewInt(int64(data.Price)),
		[]byte(data.MetaDataID),
		data.ProofDataIDs,
		len(data.ProofDataIDs),
		data.DetailsID,
		data.SupportVerify)
}

func ApproveTransferForBuying(password string, cb chainevents.EventCallback) error {
    return approveTransfer(cb,
        password,
        common.HexToAddress(scryInfo.Chain.Contracts.ProtocolAddr))
}

func approveTransfer(cb chainevents.EventCallback, password string, protocolContractAddr common.Address) error {
    if curUser == nil {
        fmt.Println("no current user")
        return errors.New("failed to approve transfer, current user is null")
    }

    curUser.SubscribeEvent("Approval", cb)

    txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
        Password: password,
        Value: big.NewInt(0),
        Pending: false}
    err := cif.ApproveTransfer(&txParam,
        protocolContractAddr,
        big.NewInt(1600))
    if err != nil {
        fmt.Println("ApproveTransfer:", err)
        return errors.New("failed to approve transfer, error:" + err.Error())
    }

    return nil
}

func CreateTransaction(publishId string, password string, cb chainevents.EventCallback) error {
    if curUser == nil {
        fmt.Println("no current user")
        return errors.New("failed to prepare to buy, current user is null")
    }

    curUser.SubscribeEvent("TransactionCreate", cb)

    txParam := chainoperations.TransactParams{
        From: common.HexToAddress(curUser.Account.Address),
        Password: password,
        Value: big.NewInt(0),
        Pending: false}
    err := cif.PrepareToBuy(&txParam, publishId)
    if err != nil {
        fmt.Println("failed to prepareToBuy, error:", err)
        return errors.New("failed to prepareToBuy, error:" + err.Error())
    }

    return nil
}

func Buy(txId float64, password string, cb chainevents.EventCallback) error {
	if curUser == nil {
		fmt.Println("no current user")
		return errors.New("failed to buy, current user is null")
	}

	curUser.SubscribeEvent("Buy", cb)

	txParam := chainoperations.TransactParams{
		From: common.HexToAddress(curUser.Account.Address),
		Password: password,
		Value: big.NewInt(0),
		Pending: false}

	err := cif.BuyData(&txParam, big.NewInt(int64(txId)))
	if err != nil {
		fmt.Println("failed to buyData, error:", err)
		return errors.New("failed to buyData, error:" + err.Error())
	}

	return nil
}

func SubmitMetaDataIdEncWithBuyer(txId float64, password string, cb chainevents.EventCallback) error {
	curUser.SubscribeEvent("ReadyForDownload", cb)

	txParam := chainoperations.TransactParams{
		From: common.HexToAddress(curUser.Account.Address),
		Password: password,
		Value: big.NewInt(0),
		Pending: false}

	metaDataIdEncWithBuyer := []byte("qzfCOkBRLj50jBos+eIk2J4Dl5D2caxeyQQTVzxgcGe6qfL7qNXBT9LgGYGrl98andDM3oS" +
		"AE1dXDHceS1yiyqnGw/f23pN3lBNcTVuRPB9JZ699mErK4J3ryokdlgLJ3lzUU/RMXFw4nU8894Jsv/nlG+db3fq0fyvl6vTZrw==")
	err := cif.SubmitMetaDataIdEncWithBuyer(&txParam, big.NewInt(int64(txId)), metaDataIdEncWithBuyer)
	if err != nil {
		fmt.Println("failed to SubmitMetaDataIdEncWithBuyer, error:", err)
		return errors.New("failed to SubmitMetaDataIdEncWithBuyer, error:" + err.Error())
	}
	return nil
}

func ConfirmDataTruth(txId float64, password string,arbitrate bool, cb chainevents.EventCallback) error {
	curUser.SubscribeEvent("TransactionClose", cb)

	txParam := chainoperations.TransactParams{
		From: common.HexToAddress(curUser.Account.Address),
		Password: password,
		Value: big.NewInt(0),
		Pending: false}
	err := cif.ConfirmDataTruth(&txParam, big.NewInt(int64(txId)), arbitrate)
	if err != nil {
		fmt.Println("failed to ConfirmDataTruth, error:", err)
		return errors.New("failed to ConfirmDataTruth, error:" + err.Error())
	}
	return nil
}

func getContracts(protocolContractAddr string,
	tokenContractAddr string,
	protocolAbiPath string,
	tokenAbiPath string,
	protocolEvents string,
	tokenEvents string) []chainevents.ContractInfo {
	pe := strings.Split(protocolEvents, sep)
	te := strings.Split(tokenEvents, sep)

	contracts := []chainevents.ContractInfo{
		{protocolContractAddr, getAbiText(protocolAbiPath), pe},
		{tokenContractAddr, getAbiText(tokenAbiPath), te},
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
