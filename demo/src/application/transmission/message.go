package transmission

import (
	"encoding/json"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/scryinfo/iscap/demo/src/application/definition"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface"
	"github.com/scryinfo/iscap/demo/src/sdk/scryclient"
	"math/big"
)

var window *astilectron.Window    = nil

func SetWindow(w *astilectron.Window) {
	window = w
}

func HandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (interface{}, error) {
	var (
		payload interface{} = nil
		err     error       = nil
	)

	switch m.Name {
	// when an user jump from login page to home page, he will get 1000,0000 tokens for test.
	case "login.verify":
		var ai definition.AccInfo = definition.AccInfo{}
		if err = json.Unmarshal(m.Payload, &ai); err != nil {
			break
		}
		var ok bool
		if ok, err = sdkinterface.UserLogin(ai.Account, ai.Password); !ok {
			break
		}
		err = sdkinterface.TransferTokenFromDeployer(big.NewInt(10000000)) // test
		payload = ok
		return payload, nil
	case "create.new.account":
		var pwd definition.AccInfo = definition.AccInfo{}
		if err = json.Unmarshal(m.Payload, &pwd); err != nil {
			break
		}
		var user *scryclient.ScryClient = nil
		if user, err = sdkinterface.CreateUserWithLogin(pwd.Password); err != nil {
			break
		}
		payload = user.Account.Address
		return payload, nil
	case "save.keystore":
		if err = sdkinterface.TransferTokenFromDeployer(big.NewInt(10000000)); err != nil {   // test
			break
		}
		payload = true
		return payload, nil
	case "publish":
		var pd definition.PubDataIDs = definition.PubDataIDs{}
		if err = json.Unmarshal(m.Payload, &pd); err != nil {
			break
		}
		payload, err = sdkinterface.PublishData(&pd, onPublish)
		if err != nil {
			break
		}
		return payload, nil
	case "buy":
		var bd definition.BuyData = definition.BuyData{}
		if err = json.Unmarshal(m.Payload, &bd); err != nil {
			break
		}
		if err = sdkinterface.ApproveTransferForBuying(bd.Password, onApprove); err != nil {
			break
		}
		if err = sdkinterface.CreateTransaction(bd.PublishID, bd.Password, onTransactionCreate); err != nil {
			break
		}
		payload = true
		return payload, nil
	case "purchase":
		var pd definition.PurchaseData = definition.PurchaseData{}
		if err = json.Unmarshal(m.Payload, &pd); err != nil {
			break
		}
		if err = sdkinterface.Buy(pd.TransactionID, pd.Password, onPurchase); err != nil {
			break
		}
		payload = true
		return payload, nil
	case "reEncrypt":
		var re definition.ReEncryptData = definition.ReEncryptData{}
		if err = json.Unmarshal(m.Payload, &re); err != nil {
			break
		}
		if err = sdkinterface.SubmitMetaDataIdEncWithBuyer(re.SelectedTx.TransactionID, re.Password,
			re.SelectedTx.MetaDataIDEncWithSeller, re.SelectedTx.Buyer, re.SelectedTx.Seller, onReadyForDownload); err != nil {
			break
		}
		payload = true
		return payload, nil
	case "decrypt":
		var dd definition.DecryptData = definition.DecryptData{}
		if err = json.Unmarshal(m.Payload, &dd); err != nil {
			break
		}
		if payload, err = sdkinterface.BuyerDecryptAndGetMetaDataFromIPFS(dd.Password, []byte(dd.SelectedTx.MetaDataIDEncWithBuyer),
			dd.SelectedTx.MetaDataExtension, dd.SelectedTx.Buyer); err != nil {
				break
		}
		return payload, nil
	case "confirm":
		var cd definition.ConfirmData = definition.ConfirmData{}
		if err = json.Unmarshal(m.Payload, &cd); err != nil {
			break
		}
		if err = sdkinterface.ConfirmDataTruth(cd.TransactionID, cd.Password,cd.Arbitrate, onClose); err != nil {
			break
		}
		payload = true
		return payload, nil
	}

	payload = err.Error()
	return payload, err
}
