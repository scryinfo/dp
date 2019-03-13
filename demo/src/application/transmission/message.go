package transmission

import (
	"encoding/json"
	"errors"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/scryinfo/iscap/demo/src/application/definition"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface"
	"github.com/scryinfo/iscap/demo/src/sdk/scryclient"
	"math/big"
)

var (
	user   *scryclient.ScryClient = nil
	window *astilectron.Window    = nil
)

func SetWindow(w *astilectron.Window) {
	window = w
}

func HandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (interface{}, error) {
	var (
		payload interface{} = nil
		err     error       = errors.New("")
	)

	switch m.Name {
	// when an user jump from login page to home page, he will get 1000,0000 tokens for test.
	case "login.verify":
		var ai definition.AccInfo = definition.AccInfo{}
		err = json.Unmarshal(m.Payload, &ai)
		if err != nil {
			break
		}
		var ok bool
		ok, err = sdkinterface.UserLogin(ai.Account, ai.Password)
		if !ok {
			break
		}
		err = sdkinterface.TransferTokenFromDeployer(big.NewInt(10000000)) // test
		payload = ok
		return payload, nil
	case "create.new.account":
		var pwd definition.AccInfo = definition.AccInfo{}
		err = json.Unmarshal(m.Payload, &pwd)
		if err != nil {
			break
		}
		user, err = sdkinterface.CreateUserWithLogin(pwd.Password)
		if err != nil {
			break
		}
		payload = user.Account.Address
		return payload, nil
	case "save.keystore":
		err = sdkinterface.TransferTokenFromDeployer(big.NewInt(10000000)) // test
		if err != nil {
			break
		}
		payload = true
		return payload, nil
	case "buy":
		var bd definition.BuyData = definition.BuyData{}
		if err = json.Unmarshal(m.Payload, &bd); err != nil {
			break
		}
		// optimize: approve transfer how much money (1600 now)? data price + rewards + gas estimated may be a solution.
		if err = sdkinterface.ApproveTransferForBuying(bd.Password, onApprove); err != nil {
			break
		}
		// optimize: not support buy a group of data one time, 'ids'([]string) adjust to 'id'(string).
		if err = sdkinterface.CreateTransaction(bd.IDs, bd.Password, onTransactionCreat); err != nil {
			break
		}
		payload = true
		return payload, nil
	case "purchase":
		var pd definition.PurchaseData = definition.PurchaseData{}
		// optimize: not support buy a group of data one time, 'ids'([]string) adjust to 'id'(string).
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
		if err = sdkinterface.SubmitMetaDataIdEncWithBuyer(re.SelectedTx.TransactionID, re.Password, onReadyForDownload);
		err != nil {
			break
		}
		payload = true
		return payload, nil
	case "confirm":
		var cd definition.ConfirmData = definition.ConfirmData{}
		if err = json.Unmarshal(m.Payload, &cd); err != nil {
			break
		}
		if err = sdkinterface.ConfirmDataTruth(cd.TransactionID, cd.Password,cd.Arbitrate, onClose); err != nil {
			break
		}
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
	}

	payload = err.Error()
	return payload, err
}
