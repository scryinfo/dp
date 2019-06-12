package sdkinterface

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryinfo/dp/demo/src/application/definition"
	"github.com/scryinfo/dp/demo/src/application/sdkinterface/settings"
	"github.com/scryinfo/dp/demo/src/sdk"
	"github.com/scryinfo/dp/demo/src/sdk/core/chainevents"
	"github.com/scryinfo/dp/demo/src/sdk/core/chainoperations"
	"github.com/scryinfo/dp/demo/src/sdk/scryclient"
	cif "github.com/scryinfo/dp/demo/src/sdk/scryclient/chaininterfacewrapper"
	"github.com/scryinfo/dp/demo/src/sdk/util/accounts"
	"github.com/scryinfo/dp/demo/src/sdk/util/storage/ipfsaccess"
	"github.com/scryinfo/dp/dots/service"
	"math/big"
	"os"
)

const IPFSOutDir = "D:/desktop"

var (
	curUser  *scryclient.ScryClient
	deployer *scryclient.ScryClient
	scryinfo *settings.scryinfo
	err      error
)

func SetScryInfo(si *settings.scryinfo) {
	scryinfo = si
}

func SetFromBlock(fromBlock uint64) {
	sdk.StartScan(fromBlock)
}

func CreateUserWithLogin(password string) (string, error) {
	var client *scryclient.ScryClient
	if client, err = scryclient.CreateScryClient(password); err != nil {
		return "", errors.Wrap(err, "Create new user failed. ")
	}

	curUser = client

	return curUser.Account.Address, nil
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
	if address, err = service.GetAMIns().ImportAccount([]byte(keyJson), oldPassword, newPassword); err != nil {
		return nil, errors.Wrap(err, "Import interface failed. ")
	}

	client := scryclient.NewScryClient(address)
	return client, nil
}

func TransferTokenFromDeployer(token *big.Int) error {
	if deployer == nil {
		deployer, err = importAccount(scryinfo.Chain.Contracts.DeployerKeyJson,
			scryinfo.Chain.Contracts.DeployerPassword,
			scryinfo.Chain.Contracts.DeployerPassword)
		if err != nil {
			return errors.Wrap(err, "Deployer init failed. ")
		}
	}

	if curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(deployer.Account.Address),
		Password: scryinfo.Chain.Contracts.DeployerPassword, Value: big.NewInt(0), Pending: false}
	if err = cif.TransferTokens(&txParam, common.HexToAddress(curUser.Account.Address), token); err != nil {
		return errors.Wrap(err, "Transfer token failed. ")
	}

	return nil
}

func SubscribeEvents(eventName []string, cb ...chainevents.EventCallback) error {
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

func UnsubscribeEvents(eventName []string) error {
	for i := 0; i < len(eventName); i++ {
		if err = curUser.UnSubscribeEvent(eventName[i]); err != nil {
			return errors.Wrap(err, "unsubscribe events "+eventName[i]+" failed. ")
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
		[]byte(data.IDs.MetaDataID),
		data.IDs.ProofDataIDs,
		len(data.IDs.ProofDataIDs),
		data.IDs.DetailsID,
		data.SupportVerify)
}

func ApproveTransferForRegisterAsVerifier(password string) error {
	return approveTransfer(password, common.HexToAddress(scryinfo.Chain.Contracts.ProtocolAddr), big.NewInt(10000))
}

func ApproveTransferForBuying(password string) error {
	return approveTransfer(password, common.HexToAddress(scryinfo.Chain.Contracts.ProtocolAddr), big.NewInt(1600))
}

func approveTransfer(password string, protocolContractAddr common.Address, token *big.Int) error {
	if curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = cif.ApproveTransfer(&txParam, protocolContractAddr, token); err != nil {
		return errors.Wrap(err, "Contract transfer token from buyer failed. ")
	}

	return nil
}

func CreateTransaction(publishId string, password string, startVerify bool) error {
	if curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = cif.PrepareToBuy(&txParam, publishId, startVerify); err != nil {
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
	if metaDataIdEncWithBuyer, err = service.GetAMIns().ReEncrypt(metaDataIDEncSeller, seller, buyer, password); err != nil {
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

func CancelTransaction(txId, password string) error {
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = cif.CancelTransaction(&txParam, tID); err != nil {
		return errors.Wrap(err, "Cancel transaction failed. ")
	}
	return nil
}

func DecryptAndGetMetaDataFromIPFS(password string, metaDataIdEncWithBuyer []byte, buyer, extension string) (string, error) {
	var metaDataIDByte []byte
	if metaDataIDByte, err = service.GetAMIns().Decrypt(metaDataIdEncWithBuyer, buyer, password); err != nil {
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

func ConfirmDataTruth(txId string, password string, truth bool) error {
	if curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	if err = cif.ConfirmDataTruth(&txParam, tID, truth); err != nil {
		return errors.Wrap(err, "Confirm data truth failed. ")
	}
	return nil
}

func RegisterAsVerifier(password string) error {
	if curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = cif.RegisterAsVerifier(&txParam); err != nil {
		return errors.Wrap(err, "Register as verifier failed. ")
	}
	return nil
}

func Vote(password, txId string, judge bool, comment string) error {
	if curUser == nil {
		return errors.New("Current user is nil. ")
	}
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = cif.Vote(&txParam, tID, judge, comment); err != nil {
		return errors.Wrap(err, "Vote failed. ")
	}
	return nil
}

func CreditToVerifiers(creditData *definition.CreditData) error {
	if curUser == nil {
		return errors.New("Current user is nil. ")
	}
	tID, ok := new(big.Int).SetString(creditData.SelectedTx.TransactionID, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	txParam := chainoperations.TransactParams{From: common.HexToAddress(curUser.Account.Address),
		Password: creditData.Password, Value: big.NewInt(0), Pending: false}

	if creditData.Credit.Verifier1Revert {
		credit := uint8(creditData.Credit.Verifier1Credit)
		if err = cif.CreditsToVerifier(&txParam, tID, 0, credit); err != nil {
			return errors.Wrap(err, "Credit failed. ")
		}
	}
	if creditData.Credit.Verifier2Revert {
		credit := uint8(creditData.Credit.Verifier2Credit)
		if err = cif.CreditsToVerifier(&txParam, tID, 1, credit); err != nil {
			return errors.Wrap(err, "Credit failed. ")
		}
	}

	return nil
}
