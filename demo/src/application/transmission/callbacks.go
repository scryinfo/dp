package transmission

import (
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"

	rlog "github.com/sirupsen/logrus"
)

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
