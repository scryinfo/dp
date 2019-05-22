// Scry Info.  All rights reserved.
// license that can be found in the license file.

package websocket

import (
	"encoding/json"
	app2 "github.com/scryinfo/dp/dots/app"
	"github.com/scryinfo/dp/dots/app/sdkinterface"
	"github.com/scryinfo/dp/dots/app/settings"
	"math/big"
)

const (
	verifierNum            = 2
	verifierBonus          = 300
	registerAsVerifierCost = 10000
)

var (
	extChan   = make(chan []string, 3)
	eventName = []string{"DataPublish", "Approval", "VerifiersChosen", "TransactionCreate", "Buy", "ReadyForDownload", "TransactionClose",
		"RegisterVerifier", "Vote", "VerifierDisable"}
)

func MessageHandlerInit() {
	addCallbackFunc("login.verify", loginVerify)
	addCallbackFunc("create.new.account", createNewAccount)
	addCallbackFunc("block.set", blockSet)
	addCallbackFunc("logout", logout)
	addCallbackFunc("publish", publish)
	addCallbackFunc("buy", buy)
	addCallbackFunc("extensions", extensions)
	addCallbackFunc("purchase", purchase)
	addCallbackFunc("reEncrypt", reEncrypt)
	addCallbackFunc("cancel", cancel)
	addCallbackFunc("decrypt", decrypt)
	addCallbackFunc("confirm", confirm)
	addCallbackFunc("register", register)
	addCallbackFunc("verify", verify)
	addCallbackFunc("credit", credit)
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
	if err = app2.GetGapp().CurUser.SubscribeEvents(eventName, onPublish, onApprove, onVerifiersChosen, onTransactionCreate, onPurchase, onReadyForDownload,
		onClose, onRegisterAsVerifier, onVote, onVerifierDisable); err != nil {
		return
	}
	sdkinterface.SetFromBlock(uint64(sid.FromBlock))
	// when an user login success, he will get 1,000,000 tokens for test. in 'block.set' case.
	if err = app2.GetGapp().CurUser.TransferTokenFromDeployer(big.NewInt(1000000)); err != nil { // for test
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
		fee += int64(verifierNum * verifierBonus)
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
	if err = app2.GetGapp().CurUser.SubmitMetaDataIdEncWithBuyer(re.SelectedTx.TransactionID, re.Password, re.SelectedTx.Seller,
		re.SelectedTx.Buyer, re.SelectedTx.MetaDataIDEncWithSeller); err != nil {
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
