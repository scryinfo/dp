package transmission

import (
	"encoding/json"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/btcsuite/btcutil/base58"
	"github.com/scryinfo/iscap/demo/src/application/definition"
	"github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
	"github.com/scryinfo/iscap/demo/src/sdk/util/storage/ipfsaccess"
	rlog "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"os"
)

const IPFSOutDir = "D:/desktop"

func onPublish(event events.Event) bool {
	go func() {
		var (
			details definition.DataDetails
			err error
		)
		if details, err = getPubDataDetails(event.Data.Get("despDataId").(string)); err != nil {
			rlog.Error("Node: get publish data details failed. ", err)
		}

		bigint := event.Data.Get("price").(*big.Int)
		details.Price = bigint.String()
		details.PublishID = event.Data.Get("publishId").(string)
		// details.SupportVerify is not implement.

		if err := bootstrap.SendMessage(window, "onPublish", details); err != nil {
			rlog.Error("failed to send onPublish event, error:", err)
		}
	}()
	return true
}

// Get publish data details from details' ipfsID.
// ipfsGet -> modify file name -> read file -> json.unmarshal -> delete file
func getPubDataDetails(ipfsID string) (definition.DataDetails, error) {
	var err error
	if err = ipfsaccess.GetIAInstance().GetFromIPFS(ipfsID, IPFSOutDir); err != nil {
		rlog.Error("Node - onPublish.callback: ipfs get failed. ", err.Error())
		return definition.DataDetails{}, err
	}

	oldFileName := IPFSOutDir + "/" + ipfsID
	newFileName := oldFileName + ".txt"
	if err = os.Rename(oldFileName, newFileName); err != nil {
		rlog.Error("Node - onPublish.callback: rename details file failed. ", err)
		return definition.DataDetails{}, err
	}

	var details []byte
	if details, err = ioutil.ReadFile(newFileName); err != nil {
		rlog.Error("Node - onPublish.callback: read details file failed. ", err)
		return definition.DataDetails{}, err
	}
	var detailsData definition.DataDetails = definition.DataDetails{}
	if err = json.Unmarshal(details, &detailsData); err != nil {
		rlog.Error("Node - onPublish.callback: json unmarshal details failed. ", err)
		return definition.DataDetails{}, err
	}

	if err = os.Remove(newFileName); err != nil {
		rlog.Debug("Node - onPublish.callback: delete details file failed. ", err)
	}

	return detailsData, nil
}

func onApprove(event events.Event) bool {
	go func() {
		// -
		if err := bootstrap.SendMessage(window, "onApprove", event); err != nil {
			rlog.Error("failed to send onApprove event, error:", err)
		}
	}()
	return true
}

func onTransactionCreate(event events.Event) bool {
	go func() {
		rlog.Info("Node: in onTransactionCreate.callback. ")
		var transaction definition.TransactionDetails = definition.TransactionDetails{}
		transaction.ProofFileNames = getProofFiles(event.Data.Get("proofIds").([]string))
		transaction.PublishID = event.Data.Get("publishId").(string)
		transaction.TransactionID = event.Data.Get("transactionId").(string)
		transaction.TxState = event.Data.Get("state").(string)

		rlog.Info("Node: before send message to js. ")
		if err := bootstrap.SendMessage(window, "onTransactionCreate", transaction); err != nil {
			rlog.Error("failed to send onTransactionCreate event, error:", err)
		}
		rlog.Info("Node: after send message to js. ")
	}()
	return true
}

// Get proof files from proofIDs.
// ipfsGet -> modify file name, how can I get extension?
func getProofFiles(ipfsIDs []string) []string {
	var (
		err    error
		proofs []string
	)
	var proofIDsByte32 [][32]byte = make([][32]byte, len(ipfsIDs))
	for i := 0; i< len(ipfsIDs); i++ {
		if len(ipfsIDs[i]) != 32 {
			rlog.Error("proofIDs' length is not 32. ")
			return nil
		}
		copy(proofIDsByte32[i][:], ipfsIDs[i])
	}
	for i := 0; i < len(ipfsIDs); i++ {
		var ipfsID string = ipfsBytes32ToHash(proofIDsByte32[i])
		// optimize: user can set IPFS out dir himself.
		if err = ipfsaccess.GetIAInstance().GetFromIPFS(ipfsID, IPFSOutDir); err != nil {
			rlog.Error("Node - onTransactionCreate.callback: ipfs get failed. ", err.Error())
			return nil
		}
		// add extension here.
	}

	return proofs
}
func ipfsBytes32ToHash(ipfsb [32]byte) string {
	var byte34 []byte = make([]byte, 34)
	// if ipfs modify encrypt algorithm, byte will change together.
	copy(byte34[:], []byte{byte(18), byte(32)})
	copy(byte34[2:], ipfsb[:])

	return base58.Encode(byte34)
}

func onPurchase(event events.Event) bool {
	go func() {
		var pd definition.PurchaseDetails = definition.PurchaseDetails{
			TransactionID: event.Data.Get("transactionId").(string),
			MetaDataIdEncWithSeller: event.Data.Get("metaDataIdEncSeller").(string),
			TxState: event.Data.Get("state").(string),
		}
		if err := bootstrap.SendMessage(window, "onPurchase", pd); err != nil {
			rlog.Error("failed to send onPurchase event, error:", err)
		}
	}()
	return true
}

func onReadyForDownload(event events.Event) bool {
	go func() {
		var rd definition.ReEncryptDetails = definition.ReEncryptDetails{
			TransactionID: event.Data.Get("transactionId").(string),
			MetaDataIdEncWithBuyer: event.Data.Get("metaDataIdEncBuyer").(string),
			TxState: event.Data.Get("state").(string),
		}
		rlog.Debug("Node: ready.for.download.callback. ", rd)
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
