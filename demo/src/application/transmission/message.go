package transmission

import (
	"encoding/json"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/scryinfo/iscap/demo/src/application/definition"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface"
	rlog "github.com/sirupsen/logrus"
	"math/big"
)

var (
	window    *astilectron.Window
	channel   = make(chan []string, 3)
	eventName = []string{"DataPublish", "Approval", "TransactionCreate", "Buy", "ReadyForDownload", "TransactionClose",
		"RegisterVerifier", "Vote", "VerifierDisable"}
)

func SetWindow(w *astilectron.Window) {
	window = w
}

func HandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (interface{}, error) {
	var (
		payload interface{}
		err     error
	)

	switch m.Name {
	// when an user jump from login page to home page, he will get 1000,0000 tokens for test.
	case "login.verify":
		var ai definition.AccInfo
		if err = json.Unmarshal(m.Payload, &ai); err != nil {
			break
		}
		if payload, err = sdkinterface.UserLogin(ai.Account, ai.Password); !(payload.(bool)) {
			break
		}
		return payload, nil
	case "create.new.account":
		var pwd definition.AccInfo
		if err = json.Unmarshal(m.Payload, &pwd); err != nil {
			break
		}
		if payload, err = sdkinterface.CreateUserWithLogin(pwd.Password); err != nil {
			break
		}
		return payload, nil
	case "sdk.init":
		var sid definition.SDKInitData
		if err = json.Unmarshal(m.Payload, &sid); err != nil {
			break
		}
		if err = sdkinterface.SubscribeEvents(eventName, onPublish, onApprove, onTransactionCreate, onPurchase, onReadyForDownload,
			onClose, onRegisterAsVerifier, onVote, onVerifierDisable); err != nil {
			break
		}
		sdkinterface.SetFromBlock(uint64(sid.FromBlock))
		if err = sdkinterface.TransferTokenFromDeployer(big.NewInt(1000000)); err != nil { // for test
			break
		}
		payload = true
		return payload, nil
	case "logout":
		if err = sdkinterface.UnsubscribeEvents(eventName); err != nil {
			break
		}
		payload = true
		return payload, nil
	case "publish":
		var pd definition.PublishData
		if err = json.Unmarshal(m.Payload, &pd); err != nil {
			break
		}
		payload, err = sdkinterface.PublishData(&pd)
		if err != nil {
			break
		}
		return payload, nil
	case "buy":
		var bd definition.BuyData
		if err = json.Unmarshal(m.Payload, &bd); err != nil {
			break
		}
		// optimize: here need to give out the summary of token buyer approve contract transfer, now it is 1600.
		if err = sdkinterface.ApproveTransferForBuying(bd.Password); err != nil {
			break
		}
		if err = sdkinterface.CreateTransaction(bd.SelectedData.PublishID, bd.Password, bd.StartVerify); err != nil {
			break
		}
		payload = true
		return payload, nil
	case "extensions":
		var p definition.Prepared
		if err = json.Unmarshal(m.Payload, &p); err != nil {
			break
		}
		channel <- p.Extensions
		payload = true
		return payload, nil
	case "purchase":
		var pd definition.PurchaseData
		if err = json.Unmarshal(m.Payload, &pd); err != nil {
			break
		}
		if err = sdkinterface.Buy(pd.SelectedTx.TransactionID, pd.Password); err != nil {
			break
		}
		payload = true
		return payload, nil
	case "reEncrypt":
		var re definition.ReEncryptData
		if err = json.Unmarshal(m.Payload, &re); err != nil {
			break
		}
		if err = sdkinterface.SubmitMetaDataIdEncWithBuyer(re.SelectedTx.TransactionID, re.Password, re.SelectedTx.Seller,
			re.SelectedTx.Buyer, re.SelectedTx.MetaDataIDEncWithSeller); err != nil {
			break
		}
		payload = true
		return payload, nil
	case "cancel":
		var pd definition.PurchaseData
		if err = json.Unmarshal(m.Payload, &pd); err != nil {
			break
		}
		if err = sdkinterface.CancelTransaction(pd.SelectedTx.TransactionID, pd.Password); err != nil { // sdkinterface not implement.
			break
		}
		payload = true
		return payload, nil
	case "decrypt":
		var dd definition.DecryptData
		if err = json.Unmarshal(m.Payload, &dd); err != nil {
			break
		}
		if payload, err = sdkinterface.DecryptAndGetMetaDataFromIPFS(dd.Password, dd.SelectedTx.MetaDataIDEncrypt,
			dd.SelectedTx.User, dd.SelectedTx.MetaDataExtension); err != nil {
			break
		}
		return payload, nil
	case "confirm":
		var cd definition.ConfirmData
		if err = json.Unmarshal(m.Payload, &cd); err != nil {
			break
		}
		if err = sdkinterface.ConfirmDataTruth(cd.SelectedTx.TransactionID, cd.Password, cd.Truth); err != nil {
			break
		}
		payload = true
		return payload, nil
	case "register":
		var rvd definition.RegisterVerifierData
		if err = json.Unmarshal(m.Payload, &rvd); err != nil {
			break
		}
		if err = sdkinterface.ApproveTransferForRegisterAsVerifier(rvd.Password); err != nil {
			break
		}
		if err = sdkinterface.RegisterAsVerifier(rvd.Password); err != nil {
			break
		}
		payload = true
		return payload, nil
	case "verify":
		var vd definition.VerifyData
		if err = json.Unmarshal(m.Payload, &vd); err != nil {
			break
		}
		if err = sdkinterface.Vote(vd.Password, vd.TransactionID, vd.Verify.Suggestion, vd.Verify.Comment); err != nil {
			break
		}
		payload = true
		return payload, nil
	case "credit":
		var cd definition.CreditData
		if err = json.Unmarshal(m.Payload, &cd); err != nil {
			break
		}
		rlog.Info("Node: in credit msg from js. ", string(m.Payload), " ", cd)
		if err = sdkinterface.CreditToVerifiers(&cd); err != nil {
			break
		}
		payload = true
		return payload, nil
	}

	rlog.Error("Handle message: ", m.Name, " failed. ", err)
	payload = err.Error()
	return payload, err
}
