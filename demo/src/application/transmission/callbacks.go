package transmission

import (
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/scryinfo/iscap/demo/src/application/definition"
	"github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
	"github.com/scryinfo/iscap/demo/src/sdk/util/storage/ipfsaccess"
	rlog "github.com/sirupsen/logrus"
)

const IPFSOutDir = "D:/desktop"

func onPublish(event events.Event) bool {
	go func() {
		//var details *definition.DataDetails = getDetails(event.Data.Get("despDataId").(string))
		//details.Price = event.Data.Get("price").(float64)
		// supportVerify is not implement.
		if err := bootstrap.SendMessage(window, "onPublish", event); err != nil {
			rlog.Error("failed to send onPublish event, error:", err)
		}
	}()
	return true
}

func getDetails(ipfsID string) *definition.DataDetails {
	//ipfsGet -> modify file name -> read file -> json.unmarshal -> delete file
	var err error
	if err = ipfsaccess.GetIAInstance().GetFromIPFS(ipfsID, IPFSOutDir); err != nil {
		rlog.Error("Node: ipfs get failed. ", err.Error())
		return nil
	}
	return nil
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
		// event.Data.metaDataIdEncWithSeller → ...EncWithBuyer
		if err := bootstrap.SendMessage(window, "onPurchase", event); err != nil {
			rlog.Error("failed to send onPurchase event, error:", err)
		}
	}()
	return true
}

func onReadyForDownload(event events.Event) bool {
	go func() {
		rlog.Debug("Node: ready.for.download.callback. ", event)
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
