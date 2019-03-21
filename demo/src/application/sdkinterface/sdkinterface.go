package sdkinterface

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scryinfo/iscap/demo/src/application/definition"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface/settings"
	"github.com/scryinfo/iscap/demo/src/sdk"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainevents"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainoperations"
	"github.com/scryinfo/iscap/demo/src/sdk/scryclient"
	cif "github.com/scryinfo/iscap/demo/src/sdk/scryclient/chaininterfacewrapper"
	"github.com/scryinfo/iscap/demo/src/sdk/util/accounts"
	"github.com/scryinfo/iscap/demo/src/sdk/util/storage/ipfsaccess"
	rlog "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

const (
	failedToInitSDK = "failed to initialize sdk."
	sep             = "|"
	IPFSOutDir      = "D:/desktop"
)

var (
	curUser  *scryclient.ScryClient = nil
	deployer *scryclient.ScryClient = nil
	scryInfo *settings.ScryInfo     = nil
	err      error                  = nil
)

func Initialize() error {
	// load definition
	scryInfo, err = settings.LoadSettings()
	if err != nil {
		rlog.Error(failedToInitSDK, err)
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
		rlog.Error(failedToInitSDK, err)
		return err
	}

	return nil
}

//new user
//func CreateUser(password string) (*scryclient.ScryClient, error) {
//	client, err := scryclient.CreateScryClient(password)
//	if err != nil {
//		return nil, err
//	}
//
//	return client, nil
//}

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
		rlog.Error("failed to import account. error:", err)
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
			rlog.Error("failed to transfer token, error:", err)
			return err
		}
	}

	if curUser == nil {
		rlog.Error("failed to transfer token, null current user")
		return errors.New("failed to transfer token, null current user")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(deployer.Account.Address),
		Password: scryInfo.Chain.Contracts.DeployerPassword,
		Value:    big.NewInt(0),
		Pending:  false}
	err = cif.TransferTokens(&txParam, common.HexToAddress(curUser.Account.Address), token)
	if err != nil {
		rlog.Error("failed to transfer token, error:", err)
		return err
	}

	return nil
}

func SubScribeEvents(eventName []string, cb ...chainevents.EventCallback) error {
	if cb == nil {
		rlog.Error("null callback function")
		return errors.New("failed to subscribe event, callback function is null")
	}
	if len(cb) != len(eventName) {
		rlog.Error("invalid callback function numbers")
		return errors.New("failed to subscribe event, callback function's quantity invalid")
	}

	for i := 0; i < len(eventName); i++ {
		if err = curUser.SubscribeEvent(eventName[i], cb[i]); err != nil {
			rlog.Error("subscribe ", eventName[i], " event failed. ")
			return errors.New("failed to subscribe " + eventName[i] + " event" + err.Error())
		}
	}
	return nil
}

func PublishData(data *definition.PubDataIDs) (string, error) {
	if curUser == nil {
		rlog.Error("no current user")
		return "", errors.New("failed to publish data, current user is null")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: data.Password, Value: big.NewInt(0), Pending: false}

	return cif.Publish(&txParam,
		big.NewInt(int64(data.Price)),
		[]byte(data.MetaDataID),
		data.ProofDataIDs,
		len(data.ProofDataIDs),
		data.DetailsID,
		data.SupportVerify)
}

func ApproveTransferForBuying(password string) error {
	return approveTransfer(
		password,
		common.HexToAddress(scryInfo.Chain.Contracts.ProtocolAddr))
}

func approveTransfer(password string, protocolContractAddr common.Address) error {
	if curUser == nil {
		rlog.Error("no current user")
		return errors.New("failed to approve transfer, current user is null")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	err := cif.ApproveTransfer(&txParam,
		protocolContractAddr,
		big.NewInt(1600))
	if err != nil {
		rlog.Error("ApproveTransfer:", err)
		return errors.New("failed to approve transfer, error:" + err.Error())
	}

	return nil
}

func CreateTransaction(publishId string, password string) error {
	if curUser == nil {
		rlog.Error("no current user")
		return errors.New("failed to CreateTransaction, current user is null")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	err := cif.PrepareToBuy(&txParam, publishId)
	if err != nil {
		rlog.Error("failed to CreateTransaction, error:", err)
		return errors.New("failed to CreateTransaction, error:" + err.Error())
	}

	return nil
}

func Buy(txId string, password string) error {
	if curUser == nil {
		rlog.Error("no current user")
		return errors.New("failed to buy, current user is null")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		rlog.Error("failed to set txID to *big.Int. ", txId)
		return errors.New("failed to set txID to *big.Int")
	}
	err := cif.BuyData(&txParam, tID)
	if err != nil {
		rlog.Error("failed to buyData, error:", err)
		return errors.New("failed to buyData, error:" + err.Error())
	}

	return nil
}

func SubmitMetaDataIdEncWithBuyer(txId string, password, buyer, seller string, metaDataIDEncSeller []byte) error {
	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}

	var metaDataIdEncWithBuyer []byte
	var err error
	metaDataIdEncWithBuyer, err = accounts.GetAMInstance().ReEncrypt([]byte(metaDataIDEncSeller), seller, buyer, password)
	if err != nil {
		rlog.Error("failed to reEncrypt meta data ID, error:", err)
		return errors.New("failed to SubmitMetaDataIdEncWithBuyer, error:" + err.Error())
	}
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		rlog.Error("failed to set txID to *big.Int.")
		return errors.New("failed to set txID to *big.Int")
	}
	err = cif.SubmitMetaDataIdEncWithBuyer(&txParam, tID, metaDataIdEncWithBuyer)
	if err != nil {
		rlog.Error("failed to SubmitMetaDataIdEncWithBuyer, error:", err)
		return errors.New("failed to SubmitMetaDataIdEncWithBuyer, error:" + err.Error())
	}
	return nil
}

func BuyerDecryptAndGetMetaDataFromIPFS(password string, metaDataIdEncWithBuyer []byte, buyer, extension string) (string, error) {
	metaDataIDByte, err := accounts.GetAMInstance().Decrypt(metaDataIdEncWithBuyer, buyer, password)
	if err != nil {
		return "", err
	}
	metaDataID := string(metaDataIDByte)
	if err := ipfsaccess.GetIAInstance().GetFromIPFS(metaDataID, IPFSOutDir); err != nil {
		return "", err
	}
	oldFileName := IPFSOutDir + "/" + metaDataID
	newFileName := oldFileName + "." + extension
	if err = os.Rename(oldFileName, newFileName); err != nil {
		rlog.Error("Node: rename meta data failed. ", err)
		return "", err
	}
	return newFileName, nil
}

func ConfirmDataTruth(txId string, password string, arbitrate bool) error {
	if curUser == nil {
		rlog.Error("no current user")
		return errors.New("failed to buy, current user is null")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		rlog.Error("failed to set txID to *big.Int.")
		return errors.New("failed to set txID to *big.Int")
	}
	err := cif.ConfirmDataTruth(&txParam, tID, arbitrate)
	if err != nil {
		rlog.Error("failed to ConfirmDataTruth, error:", err)
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
		rlog.Error("failed to read abi text", err)
		return ""
	}

	return string(abi)
}
