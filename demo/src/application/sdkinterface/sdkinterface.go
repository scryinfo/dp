package sdkinterface

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryinfo/iscap/demo/src/application/definition"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface/settings"
	"github.com/scryinfo/iscap/demo/src/sdk"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainevents"
	"github.com/scryinfo/iscap/demo/src/sdk/core/chainoperations"
	"github.com/scryinfo/iscap/demo/src/sdk/scryclient"
	cif "github.com/scryinfo/iscap/demo/src/sdk/scryclient/chaininterfacewrapper"
	"github.com/scryinfo/iscap/demo/src/sdk/util/accounts"
	"github.com/scryinfo/iscap/demo/src/sdk/util/storage/ipfsaccess"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

const (
	sep             = "|"
	IPFSOutDir      = "D:/desktop"
)

var (
	curUser    *scryclient.ScryClient
	deployer   *scryclient.ScryClient
	scryInfo   *settings.ScryInfo
	err        error
)

func Initialize(fromBlock uint64) error {
	// load definition
	if scryInfo, err = settings.LoadSettings(); err != nil {
		return errors.Wrap(err, "Load chain settings failed. ")
	}

	// initialization
	var contracts []chainevents.ContractInfo
	if contracts, err = getContracts(scryInfo.Chain.Contracts.ProtocolAddr,
		scryInfo.Chain.Contracts.TokenAddr,
		scryInfo.Chain.Contracts.ProtocolAbiPath,
		scryInfo.Chain.Contracts.TokenAbiPath,
		scryInfo.Chain.Contracts.ProtocolEvents,
		scryInfo.Chain.Contracts.TokenEvents); err != nil {
		return errors.Wrap(err, "Contracts init failed. ")
	}

	if err = sdk.Init(scryInfo.Chain.Ethereum.EthNode, contracts, fromBlock); err != nil {
		return errors.Wrap(err, "SDK init failed. ")
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

func CreateUserWithLogin(password string) (string, error) {
	var client *scryclient.ScryClient
	if client, err = scryclient.CreateScryClient(password); err != nil {
		return "", errors.Wrap(err, "Create new user failed. ")
	}

	curUser = client

	return client.Account.Address, nil
}

func UserLogin(address string, password string) (bool, error) {
	var client *scryclient.ScryClient
	if client = scryclient.NewScryClient(address); client == nil {
		return false, errors.New("Call NewScryClient failed. ")
	}

	var succ bool
	if succ, err = client.Authenticate(password); err != nil {
		return false, errors.Wrap(err, "Authenticate user infomation failed. ")
	}

	if succ {
		curUser = client
	}

	return succ, nil
}

func importAccount(keyJson string, oldPassword string, newPassword string) (*scryclient.ScryClient, error) {
	var address string
	if address, err = accounts.GetAMInstance().ImportAccount([]byte(keyJson), oldPassword, newPassword); err != nil {
		return nil, errors.Wrap(err, "Import account failed. ")
	}

	client := scryclient.NewScryClient(address)
	return client, nil
}

func TransferTokenFromDeployer(token *big.Int) error {
	var err error
	if deployer == nil {
		deployer, err = importAccount(scryInfo.Chain.Contracts.DeployerKeyJson,
			scryInfo.Chain.Contracts.DeployerPassword,
			scryInfo.Chain.Contracts.DeployerPassword)
		if err != nil {
			return errors.Wrap(err, "Deployer init failed. ")
		}
	}

	if curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(deployer.Account.Address),
		Password: scryInfo.Chain.Contracts.DeployerPassword, Value:    big.NewInt(0), Pending:  false}
	if err = cif.TransferTokens(&txParam, common.HexToAddress(curUser.Account.Address), token); err != nil {
		return errors.Wrap(err, "Transfer token failed. ")
	}

	return nil
}

func SubScribeEvents(eventName []string, cb ...chainevents.EventCallback) error {
	if cb == nil || len(cb) != len(eventName) {
		return errors.New("Quantity of callback functions is wrong. ")
	}

	for i := 0; i < len(eventName); i++ {
		if err = curUser.SubscribeEvent(eventName[i], cb[i]); err != nil {
			return errors.Wrap(err, "Subscribe event failed. ")
		}
	}
	return nil
}

func PublishData(data *definition.PublishData) (string, error) {
	if curUser == nil {
		return "", errors.New("Current user is nil. ")
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
	return approveTransfer(password, common.HexToAddress(scryInfo.Chain.Contracts.ProtocolAddr))
}

func approveTransfer(password string, protocolContractAddr common.Address) error {
	if curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = cif.ApproveTransfer(&txParam, protocolContractAddr, big.NewInt(1600)); err != nil {
		return errors.Wrap(err, "Contract transfer token from buyer failed. ")
	}

	return nil
}

func CreateTransaction(publishId string, password string) error {
	if curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = cif.PrepareToBuy(&txParam, publishId); err != nil {
		return errors.Wrap(err, "Transaction create failed. ")
	}

	return nil
}

func Buy(txId string, password string) error {
	if curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	if err = cif.BuyData(&txParam, tID); err != nil {
		return errors.Wrap(err, "Buy data failed. ")
	}

	return nil
}

func SubmitMetaDataIdEncWithBuyer(txId string, password, seller, buyer string, metaDataIDEncSeller []byte) error {
	var metaDataIdEncWithBuyer []byte
	if metaDataIdEncWithBuyer, err = accounts.GetAMInstance().ReEncrypt([]byte(metaDataIDEncSeller), seller, buyer, password);err != nil {
		return errors.Wrap(err, "Re-encrypt meta data ID failed. ")
	}
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = cif.SubmitMetaDataIdEncWithBuyer(&txParam, tID, metaDataIdEncWithBuyer); err != nil {
		return errors.Wrap(err, "Submit encrypted ID with buyer failed. ")
	}
	return nil
}

func BuyerDecryptAndGetMetaDataFromIPFS(password string, metaDataIdEncWithBuyer []byte, buyer, extension string) (string, error) {
	var metaDataIDByte []byte
	if metaDataIDByte, err = accounts.GetAMInstance().Decrypt(metaDataIdEncWithBuyer, buyer, password); err != nil {
		return "", errors.Wrap(err, "Decrypt meta data ID encrypted with buyer failed. ")
	}
	if err = ipfsaccess.GetIAInstance().GetFromIPFS(string(metaDataIDByte)); err != nil {
		return "", errors.Wrap(err, "Get meta data from IPFS failed. ")
	}
	oldFileName := IPFSOutDir + "/" + string(metaDataIDByte)
	newFileName := oldFileName + extension
	if err = os.Rename(oldFileName, newFileName); err != nil {
		return "", errors.Wrap(err, "Add extension to meta data failed. ")
	}
	return newFileName, nil
}

func ConfirmDataTruth(txId string, password string, arbitrate bool) error {
	if curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	if err = cif.ConfirmDataTruth(&txParam, tID, arbitrate); err != nil {
		return errors.Wrap(err, "Confirm data truth failed. ")
	}
	return nil
}

func getContracts(protocolContractAddr string,
	tokenContractAddr string,
	protocolAbiPath string,
	tokenAbiPath string,
	protocolEvents string,
	tokenEvents string) ([]chainevents.ContractInfo, error) {
	pe := strings.Split(protocolEvents, sep)
	te := strings.Split(tokenEvents, sep)

	var (
		protocolAbi string
		tokenAbi string
	)
	if protocolAbi, err = getAbiText(protocolAbiPath); err != nil {
		return nil, errors.Wrap(err, "Read protocol abi file failed. ")
	}
	if tokenAbi, err = getAbiText(tokenAbiPath); err != nil {
		return nil, errors.Wrap(err, "Read token abi file failed. ")
	}

	contracts := []chainevents.ContractInfo{
		{protocolContractAddr, protocolAbi, pe},
		{tokenContractAddr, tokenAbi, te},
	}

	return contracts, nil
}

func getAbiText(fileName string) (string, error) {
	abi, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", errors.Wrap(err, "Read abi file failed. ")
	}

	return string(abi), nil
}
