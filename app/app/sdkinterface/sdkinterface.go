package sdkinterface

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryInfo/dp/app/app"
	definition2 "github.com/scryInfo/dp/app/app/definition"
	settings2 "github.com/scryInfo/dp/app/app/settings"
	sdk2 "github.com/scryInfo/dp/dots/binary/sdk"
	chainevents2 "github.com/scryInfo/dp/dots/binary/sdk/core/chainevents"
	chainoperations2 "github.com/scryInfo/dp/dots/binary/sdk/core/chainoperations"
	scry "github.com/scryInfo/dp/dots/binary/sdk/scry"
	accounts2 "github.com/scryInfo/dp/dots/binary/sdk/util/accounts"
	ipfsaccess2 "github.com/scryInfo/dp/dots/binary/sdk/util/storage/ipfsaccess"
	"math/big"
	"os"
)

const IPFSOutDir = "D:/desktop" //todo config


func SetScryInfo(si *settings2.ScryInfo) {
	app.GetGapp().ScryInfo = si
}

func SetFromBlock(fromBlock uint64) {
	sdk2.StartScan(fromBlock)
}

func CreateUserWithLogin(password string) (string, error) {
	client, err := scry.CreateScryClient(password)
	if err != nil {
		return "", errors.Wrap(err, "Create new user failed. ")
	}

	app.GetGapp().CurUser = client

	return app.GetGapp().CurUser.Account().Address, nil
}

func UserLogin(address string, password string) (bool, error) {
	var err error
	var client scry.Client
	if client = scry.NewScryClient(address); client == nil {
		return false, errors.New("Call NewScryClient failed. ")
	}

	var succ bool
	if succ, err = client.Authenticate(password); err != nil {
		return false, errors.Wrap(err, "Authenticate user infomation failed. ")
	}

	if succ {
		app.GetGapp().CurUser = client
	}

	return succ, nil
}

func importAccount(keyJson string, oldPassword string, newPassword string) (scry.Client, error) {
	var address string
	var err error
	if address, err = accounts2.GetAMInstance().ImportAccount([]byte(keyJson), oldPassword, newPassword); err != nil {
		return nil, errors.Wrap(err, "Import account failed. ")
	}

	client := scry.NewScryClient(address)
	return client, nil
}

func TransferTokenFromDeployer(token *big.Int) error {
	var err error
	if app.GetGapp().Deployer == nil {
		app.GetGapp().Deployer, err = importAccount(app.GetGapp().ScryInfo.Chain.Contracts.DeployerKeyJson,
			app.GetGapp().ScryInfo.Chain.Contracts.DeployerPassword,
			app.GetGapp().ScryInfo.Chain.Contracts.DeployerPassword)
		if err != nil {
			return errors.Wrap(err, "Deployer init failed. ")
		}
	}

	if app.GetGapp().CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{From: common.HexToAddress(app.GetGapp().Deployer.Account().Address),
		Password: app.GetGapp().ScryInfo.Chain.Contracts.DeployerPassword, Value: big.NewInt(0), Pending: false}
	if err = app.GetGapp().ChainWrapper.TransferTokens(&txParam, common.HexToAddress(app.GetGapp().CurUser.Account().Address), token); err != nil {
		return errors.Wrap(err, "Transfer token failed. ")
	}

	return nil
}

func SubscribeEvents(eventName []string, cb ...chainevents2.EventCallback) error {
	var err error
	if cb == nil || len(cb) != len(eventName) {
		return errors.New("Quantity of callback functions is wrong. ")
	}

	for i := 0; i < len(eventName); i++ {
		if err = app.GetGapp().CurUser.SubscribeEvent(eventName[i], cb[i]); err != nil {
			return errors.Wrap(err, "Subscribe event failed. ")
		}
	}
	return nil
}

func UnsubscribeEvents(eventName []string) error {
	var err error
	for i := 0; i < len(eventName); i++ {
		if err = app.GetGapp().CurUser.UnSubscribeEvent(eventName[i]); err != nil {
			return errors.Wrap(err, "unsubscribe events "+eventName[i]+" failed. ")
		}
	}
	return nil
}

func PublishData(data *definition2.PublishData) (string, error) {
	if app.GetGapp().CurUser == nil {
		return "", errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{From: common.HexToAddress(app.GetGapp().CurUser.Account().Address),
		Password: data.Password, Value: big.NewInt(0), Pending: false}

	return app.GetGapp().ChainWrapper.Publish(&txParam,
		big.NewInt(int64(data.Price)),
		[]byte(data.IDs.MetaDataID),
		data.IDs.ProofDataIDs,
		len(data.IDs.ProofDataIDs),
		data.IDs.DetailsID,
		data.SupportVerify)
}

func ApproveTransferForRegisterAsVerifier(password string) error {
	return approveTransfer(password, common.HexToAddress(app.GetGapp().ScryInfo.Chain.Contracts.ProtocolAddr), big.NewInt(10000))
}

func ApproveTransferForBuying(password string) error {
	return approveTransfer(password, common.HexToAddress(app.GetGapp().ScryInfo.Chain.Contracts.ProtocolAddr), big.NewInt(1600))
}

func approveTransfer(password string, protocolContractAddr common.Address, token *big.Int) error {
	var err error
	if app.GetGapp().CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{From: common.HexToAddress(app.GetGapp().CurUser.Account().Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = app.GetGapp().ChainWrapper.ApproveTransfer(&txParam, protocolContractAddr, token); err != nil {
		return errors.Wrap(err, "Contract transfer token from buyer failed. ")
	}

	return nil
}

func CreateTransaction(publishId string, password string, startVerify bool) error {
	var err error
	if app.GetGapp().CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{From: common.HexToAddress(app.GetGapp().CurUser.Account().Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = app.GetGapp().ChainWrapper.PrepareToBuy(&txParam, publishId, startVerify); err != nil {
		return errors.Wrap(err, "Transaction create failed. ")
	}

	return nil
}

func Buy(txId string, password string) error {
	var err error
	if app.GetGapp().CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{From: common.HexToAddress(app.GetGapp().CurUser.Account().Address),
		Password: password, Value: big.NewInt(0), Pending: false}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	if err = app.GetGapp().ChainWrapper.BuyData(&txParam, tID); err != nil {
		return errors.Wrap(err, "Buy data failed. ")
	}

	return nil
}

func SubmitMetaDataIdEncWithBuyer(txId string, password, seller, buyer string, metaDataIDEncSeller []byte) error {
	var err error
	var metaDataIdEncWithBuyer []byte
	if metaDataIdEncWithBuyer, err = accounts2.GetAMInstance().ReEncrypt(metaDataIDEncSeller, seller, buyer, password); err != nil {
		return errors.Wrap(err, "Re-encrypt meta data ID failed. ")
	}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	txParam := chainoperations2.TransactParams{From: common.HexToAddress(app.GetGapp().CurUser.Account().Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = app.GetGapp().ChainWrapper.SubmitMetaDataIdEncWithBuyer(&txParam, tID, metaDataIdEncWithBuyer); err != nil {
		return errors.Wrap(err, "Submit encrypted ID with buyer failed. ")
	}
	return nil
}

func CancelTransaction(txId, password string) error {
	var err error
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	txParam := chainoperations2.TransactParams{From: common.HexToAddress(app.GetGapp().CurUser.Account().Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = app.GetGapp().ChainWrapper.CancelTransaction(&txParam, tID); err != nil {
		return errors.Wrap(err, "Cancel transaction failed. ")
	}
	return nil
}

func DecryptAndGetMetaDataFromIPFS(password string, metaDataIdEncWithBuyer []byte, buyer, extension string) (string, error) {
	var err error
	var metaDataIDByte []byte
	if metaDataIDByte, err = accounts2.GetAMInstance().Decrypt(metaDataIdEncWithBuyer, buyer, password); err != nil {
		return "", errors.Wrap(err, "Decrypt meta data ID encrypted with buyer failed. ")
	}
	if err = ipfsaccess2.GetIAInstance().GetFromIPFS(string(metaDataIDByte)); err != nil {
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
	var err error
	if app.GetGapp().CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{From: common.HexToAddress(app.GetGapp().CurUser.Account().Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	if err = app.GetGapp().ChainWrapper.ConfirmDataTruth(&txParam, tID, truth); err != nil {
		return errors.Wrap(err, "Confirm data truth failed. ")
	}
	return nil
}

func RegisterAsVerifier(password string) error {
	var err error
	if app.GetGapp().CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{From: common.HexToAddress(app.GetGapp().CurUser.Account().Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = app.GetGapp().ChainWrapper.RegisterAsVerifier(&txParam); err != nil {
		return errors.Wrap(err, "Register as verifier failed. ")
	}
	return nil
}

func Vote(password, txId string, judge bool, comment string) error {
	var err error
	if app.GetGapp().CurUser == nil {
		return errors.New("Current user is nil. ")
	}
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	txParam := chainoperations2.TransactParams{From: common.HexToAddress(app.GetGapp().CurUser.Account().Address),
		Password: password, Value: big.NewInt(0), Pending: false}
	if err = app.GetGapp().ChainWrapper.Vote(&txParam, tID, judge, comment); err != nil {
		return errors.Wrap(err, "Vote failed. ")
	}
	return nil
}

func CreditToVerifiers(creditData *definition2.CreditData) error {
	var err error
	if app.GetGapp().CurUser == nil {
		return errors.New("Current user is nil. ")
	}
	tID, ok := new(big.Int).SetString(creditData.SelectedTx.TransactionID, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}
	txParam := chainoperations2.TransactParams{From: common.HexToAddress(app.GetGapp().CurUser.Account().Address),
		Password: creditData.Password, Value: big.NewInt(0), Pending: false}

	if creditData.Credit.Verifier1Revert {
		credit := uint8(creditData.Credit.Verifier1Credit)
		if err = app.GetGapp().ChainWrapper.CreditsToVerifier(&txParam, tID, 0, credit); err != nil {
			return errors.Wrap(err, "Credit failed. ")
		}
	}
	if creditData.Credit.Verifier2Revert {
		credit := uint8(creditData.Credit.Verifier2Credit)
		if err = app.GetGapp().ChainWrapper.CreditsToVerifier(&txParam, tID, 1, credit); err != nil {
			return errors.Wrap(err, "Credit failed. ")
		}
	}

	return nil
}
