// Scry Info.  All rights reserved.
// license that can be found in the license file.

package msg_handler

import (
	"encoding/json"
	"github.com/scryinfo/dot/dot"
	app2 "github.com/scryinfo/dp/dots/app"
	"github.com/scryinfo/dp/dots/app/sdkinterface"
	"github.com/scryinfo/dp/dots/app/settings"
	"go.uber.org/zap"
	"io/ioutil"
	"math/big"
	"time"
)

const (
	verifierNum            = 2
	verifierBonus          = 300
	registerAsVerifierCost = 10000

	arbitratorNum = 1
	arbitratorBonus = 500

	sep = "|"
)

var (
	extChan   = make(chan []string, 4)
	eventName = []string{
		"Approval",
		"DataPublish",
		"VerifiersChosen",
		"TransactionCreate",
		"Buy",
		"ReadyForDownload",
		"TransactionClose",
		"RegisterVerifier",
		"Vote",
		"VerifierDisable",
		"ArbitrationBegin",
		"ArbitrationResult",
	}
)

func MessageHandlerInit() {
	app2.GetGapp().Connection.AddCallbackFunc("login.verify", loginVerify)
	app2.GetGapp().Connection.AddCallbackFunc("create.new.account", createNewAccount)
	app2.GetGapp().Connection.AddCallbackFunc("block.set", blockSet)
	app2.GetGapp().Connection.AddCallbackFunc("logout", logout)
	app2.GetGapp().Connection.AddCallbackFunc("publish", publish)
	app2.GetGapp().Connection.AddCallbackFunc("buy", buy)
	app2.GetGapp().Connection.AddCallbackFunc("extensions", extensions)
	app2.GetGapp().Connection.AddCallbackFunc("purchase", purchase)
	app2.GetGapp().Connection.AddCallbackFunc("reEncrypt", reEncrypt)
	app2.GetGapp().Connection.AddCallbackFunc("cancel", cancel)
	app2.GetGapp().Connection.AddCallbackFunc("decrypt", decrypt)
	app2.GetGapp().Connection.AddCallbackFunc("confirm", confirm)
	app2.GetGapp().Connection.AddCallbackFunc("register", register)
	app2.GetGapp().Connection.AddCallbackFunc("verify", verify)
	app2.GetGapp().Connection.AddCallbackFunc("credit", credit)
	app2.GetGapp().Connection.AddCallbackFunc("arbitrate", arbitrate)
	app2.GetGapp().Connection.AddCallbackFunc("get.eth.balance", getEthBalance)
	app2.GetGapp().Connection.AddCallbackFunc("get.token.balance", getTokenBalance)
	app2.GetGapp().Connection.AddCallbackFunc("acc.backup", backup)
	app2.GetGapp().Connection.AddCallbackFunc("acc.restore", restore)
}

func loginVerify(mi *settings.MessageIn) (payload interface{}, err error) {
	var ai settings.AccInfo
	if err = json.Unmarshal(mi.Payload, &ai); err != nil {
		return
	}
	if payload, err = app2.GetGapp().CurUser.UserLogin(ai.Account, ai.Password); !(payload.(bool)) {
		return
	}

	return
}

func createNewAccount(mi *settings.MessageIn) (payload interface{}, err error) {
	var pwd settings.AccInfo
	if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
		return
	}
	if payload, err = app2.GetGapp().CurUser.CreateUserWithLogin(pwd.Password); err != nil {
		return
	}

	return
}

func blockSet(mi *settings.MessageIn) (payload interface{}, err error) {
	var sid settings.SDKInitData
	if err = json.Unmarshal(mi.Payload, &sid); err != nil {
		return
	}
	if err = app2.GetGapp().CurUser.SubscribeEvents(eventName, onApprove, onPublish, onVerifiersChosen, onTransactionCreate, onPurchase, onReadyForDownload,
		onClose, onRegisterAsVerifier, onVote, onVerifierDisable, onArbitrationBegin, onArbitrationResult); err != nil {
		return
	}
	if err = sdkinterface.SetFromBlock(uint64(sid.FromBlock)); err != nil {
		dot.Logger().Errorln("set fromBlock failed. ", zap.NamedError("", err))
		return
	}
	// when an user login success, he will get 10,000,000 eth and tokens for test. in 'block.set' case.
	if err = app2.GetGapp().CurUser.TransferEthFromDeployer(big.NewInt(10000000)); err != nil { // for test
		return
	}
	if err = app2.GetGapp().CurUser.TransferTokenFromDeployer(big.NewInt(10000000)); err != nil { // for test
		return
	}
	payload = true

	return
}

func logout(_ *settings.MessageIn) (payload interface{}, err error) {
	if err = app2.GetGapp().CurUser.UnsubscribeEvents(eventName); err != nil {
		return
	}
	payload = true

	return
}

func publish(mi *settings.MessageIn) (payload interface{}, err error) {
	var pd settings.PublishData
	if err = json.Unmarshal(mi.Payload, &pd); err != nil {
		return
	}
	if payload, err = app2.GetGapp().CurUser.PublishData(&pd); err != nil {
		return
	}

	return
}

func buy(mi *settings.MessageIn) (payload interface{}, err error) {
	var bd settings.BuyData
	if err = json.Unmarshal(mi.Payload, &bd); err != nil {
		return
	}

	fee := int64(bd.SelectedData.Price)
	if bd.StartVerify {
		fee += int64(verifierNum * verifierBonus) + int64(arbitratorNum * arbitratorBonus)
	}
	if err = app2.GetGapp().CurUser.ApproveTransferToken(bd.Password, big.NewInt(fee)); err != nil {
		return
	}

	if err = app2.GetGapp().CurUser.CreateTransaction(bd.SelectedData.PublishID, bd.Password, bd.StartVerify); err != nil {
		return
	}
	payload = true

	return
}

func extensions(mi *settings.MessageIn) (payload interface{}, err error) {
	var p settings.Prepared
	if err = json.Unmarshal(mi.Payload, &p); err != nil {
		return
	}
	extChan <- p.Extensions
	payload = true

	return
}

func purchase(mi *settings.MessageIn) (payload interface{}, err error) {
	var pd settings.PurchaseData
	if err = json.Unmarshal(mi.Payload, &pd); err != nil {
		return
	}
	if err = app2.GetGapp().CurUser.Buy(pd.SelectedTx.TransactionID, pd.Password); err != nil {
		return
	}
	payload = true

	return
}

func reEncrypt(mi *settings.MessageIn) (payload interface{}, err error) {
	var re settings.ReEncryptData
	if err = json.Unmarshal(mi.Payload, &re); err != nil {
		return
	}
	if err = app2.GetGapp().CurUser.ReEncryptMetaDataIdBySeller(re.SelectedTx.TransactionID, re.Password, re.SelectedTx.Seller,
		re.SelectedTx.MetaDataIDEncWithSeller); err != nil {
		return
	}
	payload = true

	return
}

func cancel(mi *settings.MessageIn) (payload interface{}, err error) {
	var pd settings.PurchaseData
	if err = json.Unmarshal(mi.Payload, &pd); err != nil {
		return
	}
	if err = app2.GetGapp().CurUser.CancelTransaction(pd.SelectedTx.TransactionID, pd.Password); err != nil {
		return
	}
	payload = true

	return
}

func decrypt(mi *settings.MessageIn) (payload interface{}, err error) {
	var dd settings.DecryptData
	if err = json.Unmarshal(mi.Payload, &dd); err != nil {
		return
	}
	if payload, err = app2.GetGapp().CurUser.DecryptAndGetMetaDataFromIPFS(dd.Password, dd.SelectedTx.MetaDataIDEncrypt,
		dd.SelectedTx.User, dd.SelectedTx.MetaDataExtension); err != nil {
		return
	}

	return
}

func confirm(mi *settings.MessageIn) (payload interface{}, err error) {
	var cd settings.ConfirmData
	if err = json.Unmarshal(mi.Payload, &cd); err != nil {
		return
	}
	if err = app2.GetGapp().CurUser.ConfirmDataTruth(cd.SelectedTx.TransactionID, cd.Password, cd.Truth); err != nil {
		return
	}
	payload = true

	return
}

func register(mi *settings.MessageIn) (payload interface{}, err error) {
	var rvd settings.RegisterVerifierData
	if err = json.Unmarshal(mi.Payload, &rvd); err != nil {
		return
	}
	if err = app2.GetGapp().CurUser.ApproveTransferToken(rvd.Password, big.NewInt(registerAsVerifierCost)); err != nil {
		return
	}
	if err = app2.GetGapp().CurUser.RegisterAsVerifier(rvd.Password); err != nil {
		return
	}
	payload = true

	return
}

func verify(mi *settings.MessageIn) (payload interface{}, err error) {
	var vd settings.VerifyData
	if err = json.Unmarshal(mi.Payload, &vd); err != nil {
		return
	}
	if err = app2.GetGapp().CurUser.Vote(vd.Password, vd.TransactionID, vd.Verify.Suggestion, vd.Verify.Comment); err != nil {
		return
	}
	payload = true

	return
}

func credit(mi *settings.MessageIn) (payload interface{}, err error) {
	var cd settings.CreditData
	if err = json.Unmarshal(mi.Payload, &cd); err != nil {
		return
	}
	if err = app2.GetGapp().CurUser.CreditToVerifiers(&cd); err != nil {
		return
	}
	payload = true

	return
}

func arbitrate(mi *settings.MessageIn) (payload interface{}, err error) {
	var ad settings.ArbitrateData
	if err = json.Unmarshal(mi.Payload, &ad); err != nil {
		return
	}
	if err = app2.GetGapp().CurUser.Arbitrate(ad.Password, ad.SelectedTx.TransactionId, ad.ArbitrateResult); err != nil {
		return
	}
	payload = true

	return
}

func getEthBalance(mi *settings.MessageIn) (payload interface{}, err error) {
	var pwd settings.AccInfo
	if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
		return
	}

	var balance string
	if balance, err = app2.GetGapp().CurUser.GetEthBalance(pwd.Password); err != nil {
		return
	}

	payload = balance + sep + time.Now().String()

	return
}

func getTokenBalance(mi *settings.MessageIn) (payload interface{}, err error) {
	var pwd settings.AccInfo
	if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
		return
	}

	var balance string
	if balance, err = app2.GetGapp().CurUser.GetTokenBalance(pwd.Password); err != nil {
		return
	}

	payload = balance + sep + time.Now().String()

	return
}

func backup(mi *settings.MessageIn) (interface{}, error) {
	return true, ioutil.WriteFile(app2.GetGapp().ScryInfo.AccsBackupFile, mi.Payload, 0777)
}

func restore(_ *settings.MessageIn) (interface{}, error) {
	return ioutil.ReadFile(app2.GetGapp().ScryInfo.AccsBackupFile)
}
