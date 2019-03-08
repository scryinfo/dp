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
	// when an user jump from login page to home page, he will get 10000 tokens for test.
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
		err = sdkinterface.TransferTokenFromDeployer(big.NewInt(10000)) // test
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
		err = sdkinterface.TransferTokenFromDeployer(big.NewInt(10000)) // test
		if err != nil {
			break
		}
		payload = true
		return payload, nil
	case "get.datalist":
		payload = []definition.Datalist{
			{"Qm461", "title1", 1, "test tags461", "test description461", false},
			{"Qm462", "title2", 2, "test tags462", "test description462", false},
			{"Qm463", "title3", 3, "test tags463", "test description463", false},
			{"Qm464", "title4", 4, "test tags464", "test description464", false},
			{"Qm465", "title5", 5, "test tags465", "test description465", false},
			{"Qm466", "title6", 6, "test tags466", "test description466", false},
			{"Qm467", "title7", 7, "test tags467", "test description467", false},
			{"Qm468", "title8", 8, "test tags468", "test description468", false},
			{"Qm469", "title9", 9, "test tags469", "test description469", false},
			{"Qm4610", "title10", 10, "test tags4610", "test description4610", false},
			{"Qm4611", "title11", 11, "test tags4611", "test description4611", false},
			{"Qm4612", "title12", 12, "test tags4612", "test description4612", false},
			{"Qm4613", "title13", 13, "test tags4613", "test description4613", false},
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
		rlog.Debug("Node: show buy details from js. ", string(m.Payload))
		if err = json.Unmarshal(m.Payload, &bd); err != nil {
			break
		}
		rlog.Debug("Node: show buy details. ", bd)
		// optimize: approve contract transfer how much money (1600 now) ? data price + rewards may be a solution.
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
		payload = true
		return payload, nil
	case "publish":
		rlog.Debug("Node: show publish data from js. ", string(m.Payload))
		var pd definition.PubDataIDs = definition.PubDataIDs{}
		if err = json.Unmarshal(m.Payload, &pd); err != nil {
			break
		}
		rlog.Debug("Node: show publish data. ", pd)
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
		rlog.Debug("Node: publish.callback. ", event)
		if err := bootstrap.SendMessage(window, "onPublish", event); err != nil {
			rlog.Error("failed to send onPublish event, error:", err)
		}
	}()
	return true
}

func onApprove(event events.Event) bool {
	go func() {
		rlog.Debug("Node: approve.callback. ", event)
		if err := bootstrap.SendMessage(window, "onApprove", event); err != nil {
			rlog.Error("failed to send onApprove event, error:", err)
		}
	}()
	return true
}

func onTransactionCreat(event events.Event) bool {
	//txId = event.Data.Get("transactionId").(*big.Int)
	go func() {
		rlog.Debug("Node: transaction.create.callback. ", event)
		if err := bootstrap.SendMessage(window, "onTransactionCreat", event); err != nil {
			rlog.Error("failed to send onTransactionCreat event, error:", err)
		}
	}()
	return true
}