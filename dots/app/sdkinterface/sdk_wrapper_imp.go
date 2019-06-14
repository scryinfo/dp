// Scry Info.  All rights reserved.
// license that can be found in the license file.

package sdkinterface

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/app/settings"
	"github.com/scryinfo/dp/dots/auth"
	"github.com/scryinfo/dp/dots/binary"
	"github.com/scryinfo/dp/dots/binary/scry"
	"github.com/scryinfo/dp/dots/eth/event"
	"github.com/scryinfo/dp/dots/eth/event/listen"
	"github.com/scryinfo/dp/dots/eth/transaction"
	"github.com/scryinfo/dp/dots/storage"
	"go.uber.org/zap"
	"math/big"
	"os"
)

type sdkWrapperImp struct {
	curUser scry.Client
	dp      *settings.AccInfo
	cw      scry.ChainWrapper
}

var _ SDKWrapper = (*sdkWrapperImp)(nil)

func CreateSDKWrapperImp(cw scry.ChainWrapper) SDKWrapper {
	return &sdkWrapperImp{
		cw:   cw,
		dp: &settings.AccInfo{
			Account: "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8",
			Password: "111111",
		},
	}
}

func SetFromBlock(fromBlock uint64) error {
	l := dot.GetDefaultLine()
	if l == nil {
		return errors.New("the line do not create, do not call it")
	}
	d, err := l.ToInjecter().GetByLiveId(listen.ListenerTypeId)
	if err != nil {
		dot.Logger().Errorln("get listen dot failed. ", zap.NamedError("", err))
	}
	if g, ok := d.(*listen.Listener); ok {
		g.SetFromBlock(fromBlock)
		return nil
	}

	return errors.New("do not get Listener dot")
}

func (swi *sdkWrapperImp) CreateUserWithLogin(password string) (string, error) {
	client, err := scry.CreateScryClient(password, swi.cw)
	if err != nil {
		return "", errors.Wrap(err, "Create new user failed. ")
	}

	swi.curUser = client

	return client.Account().Addr, nil
}

func (swi *sdkWrapperImp) UserLogin(address string, password string) (bool, error) {
	var client scry.Client
	if client = scry.NewScryClient(address, swi.cw); client == nil {
		return false, errors.New("Call NewScryClient failed. ")
	}

	login, err := client.Authenticate(password)
	if err != nil {
		return false, errors.Wrap(err, "Authenticate user information failed. ")
	}
	if login {
		swi.curUser = client
	} else {
		return false, errors.New("Login verify failed. ")
	}

	return true, nil
}

func (swi *sdkWrapperImp) TransferEthFromDeployer(eth *big.Int) error {
	if swi.curUser == nil {
		return errors.New("Current user is nil. ")
	}

	if err := swi.curUser.TransferEthFrom(common.HexToAddress(swi.dp.Account), swi.dp.Password, eth, swi.cw.Conn()); err != nil {
		return errors.Wrap(err, "Transfer eth from deployer failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) TransferTokenFromDeployer(token *big.Int) error {
	if swi.curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(swi.dp.Account),
		Password: swi.dp.Password,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	if err := swi.cw.TransferTokens(&txParam, common.HexToAddress(swi.curUser.Account().Addr), token); err != nil {
		return errors.Wrap(err, "Transfer token from deployer failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) SubscribeEvents(eventName []string, cb ...event.Callback) error {
	if swi.curUser == nil {
		return errors.New("Current user is nil. ")
	}
	if cb == nil || len(cb) != len(eventName) {
		return errors.New("Quantity of event names or callback functions is wrong. ")
	}

	for i := 0; i < len(eventName); i++ {
		if err := swi.curUser.SubscribeEvent(eventName[i], cb[i]); err != nil {
			return errors.Wrap(err, "Subscribe event failed. ")
		}
	}

	return nil
}

func (swi *sdkWrapperImp) UnsubscribeEvents(eventName []string) error {
	if swi.curUser == nil {
		return errors.New("Current user is nil. ")
	}

	for i := 0; i < len(eventName); i++ {
		if err := swi.curUser.UnSubscribeEvent(eventName[i]); err != nil {
			return errors.Wrap(err, "Unsubscribe failed, event:  "+eventName[i]+" . ")
		}
	}

	return nil
}

func (swi *sdkWrapperImp) PublishData(data *settings.PublishData) (string, error) {
	if swi.curUser == nil {
		return "", errors.New("Current user is nil. ")
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(swi.curUser.Account().Addr),
		Password: data.Password,
		Value:    big.NewInt(0),
		Pending:  false,
	}

	return swi.cw.Publish(&txParam,
		big.NewInt(int64(data.Price)),
		[]byte(data.IDs.MetaDataID),
		data.IDs.ProofDataIDs,
		len(data.IDs.ProofDataIDs),
		data.IDs.DetailsID,
		data.SupportVerify)
}

func (swi *sdkWrapperImp) ApproveTransferToken(password string, quantity *big.Int) error {
	logger := dot.Logger()

	l := dot.GetDefaultLine()
	if l == nil {
		logger.Errorln("the line do not create, do not call it")
		return nil
	}
	d, err := l.ToInjecter().GetByLiveId(binary.BinLiveId)
	if err != nil {
		logger.Errorln(err.Error())
		return nil
	}
	g, ok := d.(*binary.Binary)
	if !ok {
		logger.Errorln("do not get the IPFS dot")
		return nil
	}
	protocolAddress := g.Config().ProtocolContractAddr
	protocolAddr := common.HexToAddress(protocolAddress)
	return swi.approveTransfer(password, protocolAddr, quantity)
}

func (swi *sdkWrapperImp) approveTransfer(password string, protocolContractAddr common.Address, token *big.Int) error {
	if swi.curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(swi.curUser.Account().Addr),
		Password: password,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	if err := swi.cw.ApproveTransfer(&txParam, protocolContractAddr, token); err != nil {
		return errors.Wrap(err, "Contract transfer token from buyer failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) CreateTransaction(publishId string, password string, startVerify bool) error {
	if swi.curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(swi.curUser.Account().Addr),
		Password: password,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	if err := swi.cw.PrepareToBuy(&txParam, publishId, startVerify); err != nil {
		return errors.Wrap(err, "Transaction create failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) Buy(txId string, password string) error {
	if swi.curUser == nil {
		return errors.New("Current user is nil. ")
	}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(swi.curUser.Account().Addr),
		Password: password,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	if err := swi.cw.BuyData(&txParam, tID); err != nil {
		return errors.Wrap(err, "Buy data failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) SubmitMetaDataIdEncWithBuyer(txId string, password, seller, buyer string, metaDataIDEncSeller []byte) error {
	metaDataIdEncWithBuyer, err := auth.GetAccIns().ReEncrypt(metaDataIDEncSeller, seller, buyer, password)
	if err != nil {
		return errors.Wrap(err, "Re-encrypt meta data ID failed. ")
	}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(swi.curUser.Account().Addr),
		Password: password,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	if err := swi.cw.SubmitMetaDataIdEncWithBuyer(&txParam, tID, metaDataIdEncWithBuyer); err != nil {
		return errors.Wrap(err, "Submit encrypted ID with buyer failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) CancelTransaction(txId, password string) error {
	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(swi.curUser.Account().Addr),
		Password: password,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	if err := swi.cw.CancelTransaction(&txParam, tID); err != nil {
		return errors.Wrap(err, "Cancel transaction failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) DecryptAndGetMetaDataFromIPFS(password string, metaDataIdEncWithBuyer []byte, buyer, extension string) (string, error) {
	var oldFileName string
	{
		metaDataIDByte, err := auth.GetAccIns().Decrypt(metaDataIdEncWithBuyer, buyer, password)
		if err != nil {
			return "", errors.Wrap(err, "Decrypt meta data ID encrypted with buyer failed. ")
		}
		outDir := storage.GetIPFSConfig().OutDir
		if err := storage.GetIPFSIns().Get(string(metaDataIDByte), outDir); err != nil {
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
	if swi.curUser == nil {
		return errors.New("Current user is nil. ")
	}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(swi.curUser.Account().Addr),
		Password: password,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	if err := swi.cw.ConfirmDataTruth(&txParam, tID, truth); err != nil {
		return errors.Wrap(err, "Confirm data truth failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) RegisterAsVerifier(password string) error {
	if swi.curUser == nil {
		return errors.New("Current user is nil. ")
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(swi.curUser.Account().Addr),
		Password: password,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	if err := swi.cw.RegisterAsVerifier(&txParam); err != nil {
		return errors.Wrap(err, "Register as verifier failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) Vote(password, txId string, judge bool, comment string) error {
	if swi.curUser == nil {
		return errors.New("Current user is nil. ")
	}

	tID, ok := new(big.Int).SetString(txId, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(swi.curUser.Account().Addr),
		Password: password,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	if err := swi.cw.Vote(&txParam, tID, judge, comment); err != nil {
		return errors.Wrap(err, "Vote failed. ")
	}

	return nil
}

func (swi *sdkWrapperImp) CreditToVerifiers(creditData *settings.CreditData) error {
	if swi.curUser == nil {
		return errors.New("Current user is nil. ")
	}

	tID, ok := new(big.Int).SetString(creditData.SelectedTx.TransactionID, 10)
	if !ok {
		return errors.New("Set to *big.Int failed. ")
	}

	txParam := transaction.TxParams{
		From:     common.HexToAddress(swi.curUser.Account().Addr),
		Password: creditData.Password,
		Value:    big.NewInt(0),
		Pending:  false,
	}

	if creditData.Credit.Verifier1Revert {
		credit := uint8(creditData.Credit.Verifier1Credit)
		if err := swi.cw.CreditsToVerifier(&txParam, tID, 0, credit); err != nil {
			return errors.Wrap(err, "Credit verifier1 failed. ")
		}
	}
	if creditData.Credit.Verifier2Revert {
		credit := uint8(creditData.Credit.Verifier2Credit)
		if err := swi.cw.CreditsToVerifier(&txParam, tID, 1, credit); err != nil {
			return errors.Wrap(err, "Credit verifier2 failed. ")
		}
	}

	return nil
}

func (swi *sdkWrapperImp) GetEthBalance(password string) (string, error) {
	if swi.curUser == nil {
		return "", errors.New("Current user is nil. ")
	}

	balance, err := swi.curUser.GetEth(common.HexToAddress(swi.curUser.Account().Addr), swi.cw.Conn())
	if err != nil {
		return "", errors.Wrap(err, "Get eth balance failed. ")
	}

	return balance.String(), nil
}

func (swi *sdkWrapperImp) GetTokenBalance(password string) (string, error) {
	if swi.curUser == nil {
		return "", errors.New("Current user is nil. ")
	}

	address := common.HexToAddress(swi.curUser.Account().Addr)
	txParam := transaction.TxParams{
		From:     address,
		Password: password,
		Value:    big.NewInt(0),
		Pending:  false,
	}
	balance, err := swi.cw.GetTokenBalance(&txParam, address)
	if err != nil {
		return "", errors.Wrap(err, "Get token balance failed. ")
	}

	return balance.String(), nil
}
