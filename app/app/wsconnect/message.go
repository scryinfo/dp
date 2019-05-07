package wsconnect

import (
	"encoding/json"
	definition2 "github.com/scryInfo/dp/app/app/definition"
	sdkinterface2 "github.com/scryInfo/dp/app/app/sdkinterface"
	"math/big"
)

var (
	channel   = make(chan []string, 3)
	eventName = []string{"DataPublish", "Approval", "VerifiersChosen", "TransactionCreate", "Buy", "ReadyForDownload", "TransactionClose",
		"RegisterVerifier", "Vote", "VerifierDisable"}
)

func init() {
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

func loginVerify(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var ai definition2.AccInfo
		if err = json.Unmarshal(mi.Payload, &ai); err != nil {
			break
		}
		if payload, err = sdkinterface2.UserLogin(ai.Account, ai.Password); !(payload.(bool)) {
			break
		}
		break
	}
	return payload, nil
}

func createNewAccount(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var pwd definition2.AccInfo
		if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
			break
		}
		if payload, err = sdkinterface2.CreateUserWithLogin(pwd.Password); err != nil {
			break
		}
		break
	}
	return payload, err
}

func blockSet(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var sid definition2.SDKInitData
		if err = json.Unmarshal(mi.Payload, &sid); err != nil {
			break
		}
		if err = sdkinterface2.SubscribeEvents(eventName, onPublish, onApprove, onVerifiersChosen, onTransactionCreate, onPurchase, onReadyForDownload,
			onClose, onRegisterAsVerifier, onVote, onVerifierDisable); err != nil {
			break
		}
		sdkinterface2.SetFromBlock(uint64(sid.FromBlock))
		// when an user login success, he will get 1,000,000 tokens for test. in 'block.set' case.
		if err = sdkinterface2.TransferTokenFromDeployer(big.NewInt(1000000)); err != nil { // for test
			break
		}
		payload = true
		break
	}
	return payload, err
}

func logout(_ *definition2.MessageIn) (payload interface{}, err error) {
	for {
		if err = sdkinterface2.UnsubscribeEvents(eventName); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func publish(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var pd definition2.PublishData
		if err = json.Unmarshal(mi.Payload, &pd); err != nil {
			break
		}
		payload, err = sdkinterface2.PublishData(&pd)
		if err != nil {
			break
		}
		break
	}
	return payload, err
}

func buy(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var bd definition2.BuyData
		if err = json.Unmarshal(mi.Payload, &bd); err != nil {
			break
		}
		if err = sdkinterface2.ApproveTransferForBuying(bd.Password); err != nil {
			break
		}
		if err = sdkinterface2.CreateTransaction(bd.SelectedData.PublishID, bd.Password, bd.StartVerify); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func extensions(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var p definition2.Prepared
		if err = json.Unmarshal(mi.Payload, &p); err != nil {
			break
		}
		channel <- p.Extensions
		payload = true
		break
	}
	return payload, err
}

func purchase(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var pd definition2.PurchaseData
		if err = json.Unmarshal(mi.Payload, &pd); err != nil {
			break
		}
		if err = sdkinterface2.Buy(pd.SelectedTx.TransactionID, pd.Password); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func reEncrypt(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var re definition2.ReEncryptData
		if err = json.Unmarshal(mi.Payload, &re); err != nil {
			break
		}
		if err = sdkinterface2.SubmitMetaDataIdEncWithBuyer(re.SelectedTx.TransactionID, re.Password, re.SelectedTx.Seller,
			re.SelectedTx.Buyer, re.SelectedTx.MetaDataIDEncWithSeller); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func cancel(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var pd definition2.PurchaseData
		if err = json.Unmarshal(mi.Payload, &pd); err != nil {
			break
		}
		if err = sdkinterface2.CancelTransaction(pd.SelectedTx.TransactionID, pd.Password); err != nil { // sdkinterface not implement.
			break
		}
		payload = true
		break
	}
	return payload, err
}

func decrypt(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var dd definition2.DecryptData
		if err = json.Unmarshal(mi.Payload, &dd); err != nil {
			break
		}
		if payload, err = sdkinterface2.DecryptAndGetMetaDataFromIPFS(dd.Password, dd.SelectedTx.MetaDataIDEncrypt,
			dd.SelectedTx.User, dd.SelectedTx.MetaDataExtension); err != nil {
			break
		}
		break
	}
	return payload, err
}

func confirm(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var cd definition2.ConfirmData
		if err = json.Unmarshal(mi.Payload, &cd); err != nil {
			break
		}
		if err = sdkinterface2.ConfirmDataTruth(cd.SelectedTx.TransactionID, cd.Password, cd.Truth); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func register(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var rvd definition2.RegisterVerifierData
		if err = json.Unmarshal(mi.Payload, &rvd); err != nil {
			break
		}
		if err = sdkinterface2.ApproveTransferForRegisterAsVerifier(rvd.Password); err != nil {
			break
		}
		if err = sdkinterface2.RegisterAsVerifier(rvd.Password); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func verify(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var vd definition2.VerifyData
		if err = json.Unmarshal(mi.Payload, &vd); err != nil {
			break
		}
		if err = sdkinterface2.Vote(vd.Password, vd.TransactionID, vd.Verify.Suggestion, vd.Verify.Comment); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func credit(mi *definition2.MessageIn) (payload interface{}, err error) {
	for {
		var cd definition2.CreditData
		if err = json.Unmarshal(mi.Payload, &cd); err != nil {
			break
		}
		if err = sdkinterface2.CreditToVerifiers(&cd); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}
