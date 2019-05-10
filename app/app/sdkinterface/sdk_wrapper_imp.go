package sdkinterface

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryInfo/dp/app/app"
	settings2 "github.com/scryInfo/dp/app/app/settings"
	sdk2 "github.com/scryInfo/dp/dots/binary/sdk"
	chainevents2 "github.com/scryInfo/dp/dots/binary/sdk/core/chainevents"
	chainoperations2 "github.com/scryInfo/dp/dots/binary/sdk/core/chainoperations"
	"github.com/scryInfo/dp/dots/binary/sdk/scry"
	accounts2 "github.com/scryInfo/dp/dots/binary/sdk/util/accounts"
	ipfsaccess2 "github.com/scryInfo/dp/dots/binary/sdk/util/storage/ipfsaccess"
	"math/big"
	"os"
)

type sdkWrapperImp struct {
	CurUser scry.Client
}

func NewSDKWrapperImp() SDKWrapper {
	return &sdkWrapperImp{}
}

func SetFromBlock(fromBlock uint64) {
	sdk2.StartScan(fromBlock)
}

func (swi *sdkWrapperImp) CreateUserWithLogin(password string) (string, error) {
	client, err := scry.CreateScryClient(password, app.GetGapp().ChainWrapper)
	if err != nil {
		return "", errors.Wrap(err, "Create new user failed. ")
	}

	swi.CurUser = client

	return client.Account().Address, nil
}

func (swi *sdkWrapperImp) UserLogin(address string, password string) (bool, error) {
	var client scry.Client
	if client = scry.NewScryClient(address, app.GetGapp().ChainWrapper); client == nil {
		return false, errors.New("Call NewScryClient failed. ")
	}

	login, err := client.Authenticate(password)
	if err != nil {
		return false, errors.Wrap(err, "Authenticate user information failed. ")
	}
	if login {
		swi.CurUser = client
	}

	return true, nil
}

func (swi *sdkWrapperImp) TransferTokenFromDeployer(token *big.Int) error {
	var err error
	if app.GetGapp().Deployer == nil {
		app.GetGapp().Deployer, err = importAccount(app.GetGapp().ScryInfo.Chain.Contracts.DeployerKeyJson,
			app.GetGapp().ScryInfo.Chain.Contracts.DeployerPassword,
			app.GetGapp().ScryInfo.Chain.Contracts.DeployerPassword)
		if err != nil {
			return errors.Wrap(err, "Deployer init failed. ")
		}
	}

	if swi.CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{
		From: common.HexToAddress(app.GetGapp().Deployer.Account().Address),
		Password: app.GetGapp().ScryInfo.Chain.Contracts.DeployerPassword,
		Value: big.NewInt(0),
		Pending: false}
	if err = app.GetGapp().ChainWrapper.TransferTokens(&txParam, common.HexToAddress(swi.CurUser.Account().Address), token); err != nil {
		return errors.Wrap(err, "Transfer token from deployer failed. ")
	}

	return nil
}
func importAccount(keyJson string, oldPassword string, newPassword string) (scry.Client, error) {
	address, err := accounts2.GetAMInstance().ImportAccount([]byte(keyJson), oldPassword, newPassword)
	if err != nil {
		return nil, errors.Wrap(err, "Import account failed. ")
	}

	return scry.NewScryClient(address, app.GetGapp().ChainWrapper), nil
}

func (swi *sdkWrapperImp) SubscribeEvents(eventName []string, cb ...chainevents2.EventCallback) error {
	if swi.CurUser == nil {
		return errors.New("Current user is nil. ")
	}
	if cb == nil || len(cb) != len(eventName) {
		return errors.New("Quantity of event names or callback functions is wrong. ")
	}

	for i := 0; i < len(eventName); i++ {
		if err := swi.CurUser.SubscribeEvent(eventName[i], cb[i]); err != nil {
			return errors.Wrap(err, "Subscribe event failed. ")
		}
	}

	return nil
}

func (swi *sdkWrapperImp) UnsubscribeEvents(eventName []string) error {
	if swi.CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	for i := 0; i < len(eventName); i++ {
		if err := swi.CurUser.UnSubscribeEvent(eventName[i]); err != nil {
			return errors.Wrap(err, "Unsubscribe failed, event:  "+eventName[i]+" . ")
		}
	}

	return nil
}

func (swi *sdkWrapperImp) PublishData(data *settings2.PublishData) (string, error) {
	if swi.CurUser == nil {
		return "", errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{
		From: common.HexToAddress(swi.CurUser.Account().Address),
		Password: data.Password,
		Value: big.NewInt(0),
		Pending: false}

	return app.GetGapp().ChainWrapper.Publish(&txParam,
		big.NewInt(int64(data.Price)),
		[]byte(data.IDs.MetaDataID),
		data.IDs.ProofDataIDs,
		len(data.IDs.ProofDataIDs),
		data.IDs.DetailsID,
		data.SupportVerify)
}

func (swi *sdkWrapperImp) ApproveTransferToken(password string, quantity *big.Int) error {
	protocolAddr := common.HexToAddress(app.GetGapp().ScryInfo.Chain.Contracts.ProtocolAddr)
	return swi.approveTransfer(password, protocolAddr, quantity)
}

func (swi *sdkWrapperImp) approveTransfer(password string, protocolContractAddr common.Address, token *big.Int) error {
	if swi.CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{
		From: common.HexToAddress(swi.CurUser.Account().Address),
		Password: password,
		Value: big.NewInt(0),
		Pending: false}
	if err := app.GetGapp().ChainWrapper.ApproveTransfer(&txParam, protocolContractAddr, token); err != nil {
		return errors.Wrap(err, "Contract transfer token from buyer failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) CreateTransaction(publishId string, password string, startVerify bool) error {
	if swi.CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{
		From: common.HexToAddress(swi.CurUser.Account().Address),
		Password: password,
		Value: big.NewInt(0),
		Pending: false}
	if err := app.GetGapp().ChainWrapper.PrepareToBuy(&txParam, publishId, startVerify); err != nil {
		return errors.Wrap(err, "Transaction create failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) Buy(txId string, password string) error {
	if swi.CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := chainoperations2.TransactParams{
		From: common.HexToAddress(swi.CurUser.Account().Address),
		Password: password,
		Value: big.NewInt(0),
		Pending: false}
	if err := app.GetGapp().ChainWrapper.BuyData(&txParam, tID); err != nil {
		return errors.Wrap(err, "Buy data failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) SubmitMetaDataIdEncWithBuyer(txId string, password, seller, buyer string, metaDataIDEncSeller []byte) error {
	metaDataIdEncWithBuyer, err := accounts2.GetAMInstance().ReEncrypt(metaDataIDEncSeller, seller, buyer, password)
	if err != nil {
		return errors.Wrap(err, "Re-encrypt meta data ID failed. ")
	}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := chainoperations2.TransactParams{
		From: common.HexToAddress(swi.CurUser.Account().Address),
		Password: password,
		Value: big.NewInt(0),
		Pending: false}
	if err := app.GetGapp().ChainWrapper.SubmitMetaDataIdEncWithBuyer(&txParam, tID, metaDataIdEncWithBuyer); err != nil {
		return errors.Wrap(err, "Submit encrypted ID with buyer failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) CancelTransaction(txId, password string) error {
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := chainoperations2.TransactParams{
		From: common.HexToAddress(swi.CurUser.Account().Address),
		Password: password,
		Value: big.NewInt(0),
		Pending: false}
	if err := app.GetGapp().ChainWrapper.CancelTransaction(&txParam, tID); err != nil {
		return errors.Wrap(err, "Cancel transaction failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) DecryptAndGetMetaDataFromIPFS(password string, metaDataIdEncWithBuyer []byte, buyer, extension string) (string, error) {
	var oldFileName string
	{
		metaDataIDByte, err := accounts2.GetAMInstance().Decrypt(metaDataIdEncWithBuyer, buyer, password)
		if err != nil {
			return "", errors.Wrap(err, "Decrypt meta data ID encrypted with buyer failed. ")
		}
		outDir := app.GetGapp().ScryInfo.Config.IPFSOutDir
		if err := ipfsaccess2.GetIAInstance().GetFromIPFS(string(metaDataIDByte), outDir); err != nil {
			return "", errors.Wrap(err, "Get meta data from IPFS failed. ")
		}
		oldFileName = outDir + "/" + string(metaDataIDByte)
	}

	newFileName := oldFileName + extension
	if err := os.Rename(oldFileName, newFileName); err != nil {
		return "", errors.Wrap(err, "Add extension to meta data failed. ")
	}

	return newFileName, nil
}

func (swi *sdkWrapperImp) ConfirmDataTruth(txId string, password string, truth bool) error {
	if swi.CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := chainoperations2.TransactParams{
		From: common.HexToAddress(swi.CurUser.Account().Address),
		Password: password,
		Value: big.NewInt(0),
		Pending: false}
	if err := app.GetGapp().ChainWrapper.ConfirmDataTruth(&txParam, tID, truth); err != nil {
		return errors.Wrap(err, "Confirm data truth failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) RegisterAsVerifier(password string) error {
	if swi.CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := chainoperations2.TransactParams{
		From: common.HexToAddress(swi.CurUser.Account().Address),
		Password: password,
		Value: big.NewInt(0),
		Pending: false}
	if err := app.GetGapp().ChainWrapper.RegisterAsVerifier(&txParam); err != nil {
		return errors.Wrap(err, "Register as verifier failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) Vote(password, txId string, judge bool, comment string) error {
	if swi.CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := chainoperations2.TransactParams{
		From: common.HexToAddress(swi.CurUser.Account().Address),
		Password: password,
		Value: big.NewInt(0),
		Pending: false}
	if err := app.GetGapp().ChainWrapper.Vote(&txParam, tID, judge, comment); err != nil {
		return errors.Wrap(err, "Vote failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) CreditToVerifiers(creditData *settings2.CreditData) error {
	if swi.CurUser == nil {
		return errors.New("Current user is nil. ")
	}

	tID, ok := new(big.Int).SetString(creditData.SelectedTx.TransactionID, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := chainoperations2.TransactParams{
		From: common.HexToAddress(swi.CurUser.Account().Address),
		Password: creditData.Password,
		Value: big.NewInt(0),
		Pending: false}

	if creditData.Credit.Verifier1Revert {
		credit := uint8(creditData.Credit.Verifier1Credit)
		if err := app.GetGapp().ChainWrapper.CreditsToVerifier(&txParam, tID, 0, credit); err != nil {
			return errors.Wrap(err, "Credit verifier1 failed. ")
		}
	}
	if creditData.Credit.Verifier2Revert {
		credit := uint8(creditData.Credit.Verifier2Credit)
		if err := app.GetGapp().ChainWrapper.CreditsToVerifier(&txParam, tID, 1, credit); err != nil {
			return errors.Wrap(err, "Credit verifier2 failed. ")
		}
	}

	return nil
}
