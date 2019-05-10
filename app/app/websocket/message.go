package websocket

import (
	"encoding/json"
	sdkinterface2 "github.com/scryInfo/dp/app/app/sdkinterface"
	settings2 "github.com/scryInfo/dp/app/app/settings"
	"math/big"
)

const (
	verifierNum            = 2
	verifierBonus          = 300
	registerAsVerifierCost = 10000
)

var (
	channel   = make(chan []string, 3) // todo: think how to reduce this global variable with no influence on function.
	curUser   sdkinterface2.SDKWrapper
	eventName = []string{"DataPublish", "Approval", "VerifiersChosen", "TransactionCreate", "Buy", "ReadyForDownload", "TransactionClose",
		"RegisterVerifier", "Vote", "VerifierDisable"}
)

func SetCurUser(cu sdkinterface2.SDKWrapper) {
	curUser = cu
}

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

func loginVerify(mi *settings2.MessageIn) (payload interface{}, err error) {
	var ai settings2.AccInfo
	if err = json.Unmarshal(mi.Payload, &ai); err != nil {
		return
	}
	if payload, err = curUser.UserLogin(ai.Account, ai.Password); !(payload.(bool)) {
		return
	}

	return
}

func createNewAccount(mi *settings2.MessageIn) (payload interface{}, err error) {
	var pwd settings2.AccInfo
	if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
		return
	}
	if payload, err = curUser.CreateUserWithLogin(pwd.Password); err != nil {
		return
	}

	return
}

func blockSet(mi *settings2.MessageIn) (payload interface{}, err error) {
	var sid settings2.SDKInitData
	if err = json.Unmarshal(mi.Payload, &sid); err != nil {
		return
	}
	if err = curUser.SubscribeEvents(eventName, onPublish, onApprove, onVerifiersChosen, onTransactionCreate, onPurchase, onReadyForDownload,
		onClose, onRegisterAsVerifier, onVote, onVerifierDisable); err != nil {
		return
	}
	sdkinterface2.SetFromBlock(uint64(sid.FromBlock))
	// when an user login success, he will get 1,000,000 tokens for test. in 'block.set' case.
	if err = curUser.TransferTokenFromDeployer(big.NewInt(1000000)); err != nil { // for test
		return
	}
	payload = true

	return
}

func logout(_ *settings2.MessageIn) (payload interface{}, err error) {
	if err = curUser.UnsubscribeEvents(eventName); err != nil {
		return
	}
	payload = true

	return
}

func publish(mi *settings2.MessageIn) (payload interface{}, err error) {
	var pd settings2.PublishData
	if err = json.Unmarshal(mi.Payload, &pd); err != nil {
		return
	}
	if payload, err = curUser.PublishData(&pd); err != nil {
		return
	}

	return
}

func buy(mi *settings2.MessageIn) (payload interface{}, err error) {
	var bd settings2.BuyData
	if err = json.Unmarshal(mi.Payload, &bd); err != nil {
		return
	}

	fee := int64(bd.SelectedData.Price)
	if bd.StartVerify {
		fee += int64(verifierNum * verifierBonus)
	}
	if err = curUser.ApproveTransferToken(bd.Password, big.NewInt(fee)); err != nil {
		return
	}

	if err = curUser.CreateTransaction(bd.SelectedData.PublishID, bd.Password, bd.StartVerify); err != nil {
		return
	}
	payload = true

	return
}

func extensions(mi *settings2.MessageIn) (payload interface{}, err error) {
	var p settings2.Prepared
	if err = json.Unmarshal(mi.Payload, &p); err != nil {
		return
	}
	channel <- p.Extensions
	payload = true

	return
}

func purchase(mi *settings2.MessageIn) (payload interface{}, err error) {
	var pd settings2.PurchaseData
	if err = json.Unmarshal(mi.Payload, &pd); err != nil {
		return
	}
	if err = curUser.Buy(pd.SelectedTx.TransactionID, pd.Password); err != nil {
		return
	}
	payload = true

	return
}

func reEncrypt(mi *settings2.MessageIn) (payload interface{}, err error) {
	var re settings2.ReEncryptData
	if err = json.Unmarshal(mi.Payload, &re); err != nil {
		return
	}
	if err = curUser.SubmitMetaDataIdEncWithBuyer(re.SelectedTx.TransactionID, re.Password, re.SelectedTx.Seller,
		re.SelectedTx.Buyer, re.SelectedTx.MetaDataIDEncWithSeller); err != nil {
		return
	}
	payload = true

	return
}

func cancel(mi *settings2.MessageIn) (payload interface{}, err error) {
	var pd settings2.PurchaseData
	if err = json.Unmarshal(mi.Payload, &pd); err != nil {
		return
	}
	if err = curUser.CancelTransaction(pd.SelectedTx.TransactionID, pd.Password); err != nil {
		return
	}
	payload = true

	return
}

func decrypt(mi *settings2.MessageIn) (payload interface{}, err error) {
	var dd settings2.DecryptData
	if err = json.Unmarshal(mi.Payload, &dd); err != nil {
		return
	}
	if payload, err = curUser.DecryptAndGetMetaDataFromIPFS(dd.Password, dd.SelectedTx.MetaDataIDEncrypt,
		dd.SelectedTx.User, dd.SelectedTx.MetaDataExtension); err != nil {
		return
	}

	return
}

func confirm(mi *settings2.MessageIn) (payload interface{}, err error) {
	var cd settings2.ConfirmData
	if err = json.Unmarshal(mi.Payload, &cd); err != nil {
		return
	}
	if err = curUser.ConfirmDataTruth(cd.SelectedTx.TransactionID, cd.Password, cd.Truth); err != nil {
		return
	}
	payload = true

	return
}

func register(mi *settings2.MessageIn) (payload interface{}, err error) {
	var rvd settings2.RegisterVerifierData
	if err = json.Unmarshal(mi.Payload, &rvd); err != nil {
		return
	}
	if err = curUser.ApproveTransferToken(rvd.Password, big.NewInt(registerAsVerifierCost)); err != nil {
		return
	}
	if err = curUser.RegisterAsVerifier(rvd.Password); err != nil {
		return
	}
	payload = true

	return
}

func verify(mi *settings2.MessageIn) (payload interface{}, err error) {
	var vd settings2.VerifyData
	if err = json.Unmarshal(mi.Payload, &vd); err != nil {
		return
	}
	if err = curUser.Vote(vd.Password, vd.TransactionID, vd.Verify.Suggestion, vd.Verify.Comment); err != nil {
		return
	}
	payload = true

	return
}

func credit(mi *settings2.MessageIn) (payload interface{}, err error) {
	var cd settings2.CreditData
	if err = json.Unmarshal(mi.Payload, &cd); err != nil {
		return
	}
	if err = curUser.CreditToVerifiers(&cd); err != nil {
		return
	}
	payload = true

	return
}
