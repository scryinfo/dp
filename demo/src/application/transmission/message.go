package transmission

import (
	"encoding/json"
	"errors"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/scryinfo/iscap/demo/src/application/definition"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface"
	"github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
	"github.com/scryinfo/iscap/demo/src/sdk/scryclient"
	rlog "github.com/sirupsen/logrus"
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
	case "get.datalist":
		payload = []definition.Datalist{
			{"Qm461", "title1", 1, "test tags461", "test description461", false},
		}
		return payload, nil
	case "get.transaction":
		payload = []definition.Transaction{
			{"title1", 1, 1, "0x1234567890123456789012345678901234567890", "0x1524783212578655202365479511235413256752", definition.Created,
				"1,v1r", "2,v2r", "3,v3r", true},
		}
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
		if err = sdkinterface.SubmitMetaDataIdEncWithBuyer(pd.TransactionID, pd.Password, onReadyForDownload); err != nil {
			break
		}
		if err = sdkinterface.ConfirmDataTruth(pd.TransactionID, pd.Password, onClose); err != nil {
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
	}

	payload = err.Error()
	return payload, err
}

func onPublish(event events.Event) bool {
	go func() {
		if err := bootstrap.SendMessage(window, "onPublish", event); err != nil {
			rlog.Error("failed to send onPublish event, error:", err)
		}
	}()
	return true
}

func onApprove(event events.Event) bool {
	go func() {
		if err := bootstrap.SendMessage(window, "onApprove", event); err != nil {
			rlog.Error("failed to send onApprove event, error:", err)
		}
	}()
	return true
}

func onTransactionCreat(event events.Event) bool {
	//txId = event.Data.Get("transactionId").(*big.Int)
	go func() {
		if err := bootstrap.SendMessage(window, "onTransactionCreat", event); err != nil {
			rlog.Error("failed to send onTransactionCreat event, error:", err)
		}
	}()
	return true
}

func onPurchase(event events.Event) bool {
	go func() {
		rlog.Debug("Node: purchase.callback. ", event)
		// event.data.metaDataIdEncWithSeller → ...EncWithBuyer
		if err := bootstrap.SendMessage(window, "onPurchase", event); err != nil {
			rlog.Error("failed to send onPurchase event, error:", err)
		}
	}()
	return true
}

func onReadyForDownload(event events.Event) bool {
	go func() {
		rlog.Debug("Node: ready.for.download.callback. ", event)
		//metaDataIdEncWithBuyer = event.Data.Get("metaDataIdEncBuyer").([]byte)
		if err := bootstrap.SendMessage(window, "onReadyForDownload", event); err != nil {
			rlog.Error("failed to send onReadyForDownload event, error:", err)
		}
	}()
	return true
}

func onClose(event events.Event) bool {
	go func() {
		rlog.Debug("Node: confirm.data.truth.callback. ", event)
		if err := bootstrap.SendMessage(window, "onClose", event); err != nil {
			rlog.Error("failed to send onClose event, error:", err)
		}
	}()

	return true
}