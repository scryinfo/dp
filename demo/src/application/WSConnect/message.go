package WSConnect

import (
	"encoding/json"
	"github.com/scryinfo/dp/demo/src/application/definition"
	"github.com/scryinfo/dp/demo/src/application/sdkinterface"
	"math/big"
)

var (
	channel   = make(chan []string, 3)
	eventName = []string{"DataPublish", "Approval", "VerifiersChosen", "TransactionCreate", "Buy", "ReadyForDownload", "TransactionClose",
		"RegisterVerifier", "Vote", "VerifierDisable"}
)

func init() {
	addCallbackFunc("login.verify", loginVerify)
	addCallbackFunc("create.new.interface", createNewAccount)
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

func loginVerify(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var ai definition.AccInfo
		if err = json.Unmarshal(mi.Payload, &ai); err != nil {
			break
		}
		if payload, err = sdkinterface.UserLogin(ai.Account, ai.Password); !(payload.(bool)) {
			break
		}
		break
	}
	return payload, nil
}

func createNewAccount(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var pwd definition.AccInfo
		if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
			break
		}
		if payload, err = sdkinterface.CreateUserWithLogin(pwd.Password); err != nil {
			break
		}
		break
	}
	return payload, err
}

func blockSet(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var sid definition.SDKInitData
		if err = json.Unmarshal(mi.Payload, &sid); err != nil {
			break
		}
		if err = sdkinterface.SubscribeEvents(eventName, onPublish, onApprove, onVerifiersChosen, onTransactionCreate, onPurchase, onReadyForDownload,
			onClose, onRegisterAsVerifier, onVote, onVerifierDisable); err != nil {
			break
		}
		sdkinterface.SetFromBlock(uint64(sid.FromBlock))
		// when an user login success, he will get 1,000,000 tokens for test. in 'block.set' case.
		if err = sdkinterface.TransferTokenFromDeployer(big.NewInt(1000000)); err != nil { // for test
			break
		}
		payload = true
		break
	}
	return payload, err
}

func logout(_ *definition.MessageIn) (payload interface{}, err error) {
	for {
		if err = sdkinterface.UnsubscribeEvents(eventName); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func publish(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var pd definition.PublishData
		if err = json.Unmarshal(mi.Payload, &pd); err != nil {
			break
		}
		payload, err = sdkinterface.PublishData(&pd)
		if err != nil {
			break
		}
		break
	}
	return payload, err
}

func buy(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var bd definition.BuyData
		if err = json.Unmarshal(mi.Payload, &bd); err != nil {
			break
		}
		if err = sdkinterface.ApproveTransferForBuying(bd.Password); err != nil {
			break
		}
		if err = sdkinterface.CreateTransaction(bd.SelectedData.PublishID, bd.Password, bd.StartVerify); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func extensions(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var p definition.Prepared
		if err = json.Unmarshal(mi.Payload, &p); err != nil {
			break
		}
		channel <- p.Extensions
		payload = true
		break
	}
	return payload, err
}

func purchase(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var pd definition.PurchaseData
		if err = json.Unmarshal(mi.Payload, &pd); err != nil {
			break
		}
		if err = sdkinterface.Buy(pd.SelectedTx.TransactionID, pd.Password); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func reEncrypt(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var re definition.ReEncryptData
		if err = json.Unmarshal(mi.Payload, &re); err != nil {
			break
		}
		if err = sdkinterface.SubmitMetaDataIdEncWithBuyer(re.SelectedTx.TransactionID, re.Password, re.SelectedTx.Seller,
			re.SelectedTx.Buyer, re.SelectedTx.MetaDataIDEncWithSeller); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func cancel(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var pd definition.PurchaseData
		if err = json.Unmarshal(mi.Payload, &pd); err != nil {
			break
		}
		if err = sdkinterface.CancelTransaction(pd.SelectedTx.TransactionID, pd.Password); err != nil { // sdkinterface not implement.
			break
		}
		payload = true
		break
	}
	return payload, err
}

func decrypt(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var dd definition.DecryptData
		if err = json.Unmarshal(mi.Payload, &dd); err != nil {
			break
		}
		if payload, err = sdkinterface.DecryptAndGetMetaDataFromIPFS(dd.Password, dd.SelectedTx.MetaDataIDEncrypt,
			dd.SelectedTx.User, dd.SelectedTx.MetaDataExtension); err != nil {
			break
		}
		break
	}
	return payload, err
}

func confirm(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var cd definition.ConfirmData
		if err = json.Unmarshal(mi.Payload, &cd); err != nil {
			break
		}
		if err = sdkinterface.ConfirmDataTruth(cd.SelectedTx.TransactionID, cd.Password, cd.Truth); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func register(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var rvd definition.RegisterVerifierData
		if err = json.Unmarshal(mi.Payload, &rvd); err != nil {
			break
		}
		if err = sdkinterface.ApproveTransferForRegisterAsVerifier(rvd.Password); err != nil {
			break
		}
		if err = sdkinterface.RegisterAsVerifier(rvd.Password); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func verify(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var vd definition.VerifyData
		if err = json.Unmarshal(mi.Payload, &vd); err != nil {
			break
		}
		if err = sdkinterface.Vote(vd.Password, vd.TransactionID, vd.Verify.Suggestion, vd.Verify.Comment); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}

func credit(mi *definition.MessageIn) (payload interface{}, err error) {
	for {
		var cd definition.CreditData
		if err = json.Unmarshal(mi.Payload, &cd); err != nil {
			break
		}
		if err = sdkinterface.CreditToVerifiers(&cd); err != nil {
			break
		}
		payload = true
		break
	}
	return payload, err
}
