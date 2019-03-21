package transmission

import (
	"encoding/json"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
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
		details.Block = event.BlockNumber
		details.Price = event.Data.Get("price").(*big.Int).String()
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
		var ad definition.ApproveDetails = definition.ApproveDetails{}
		ad.Block = event.BlockNumber
		
		if err := bootstrap.SendMessage(window, "onApprove", ad); err != nil {
			rlog.Error("failed to send onApprove event, error:", err)
		}
	}()
	return true
}

func onTransactionCreate(event events.Event) bool {
	go func() {
		var td definition.TransactionDetails = definition.TransactionDetails{}
		td.Block = event.BlockNumber
		td.ProofFileNames = getProofFiles(event.Data.Get("proofIds").([][32]uint8))
		td.PublishID = event.Data.Get("publishId").(string)
		td.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		td.Buyer = event.Data.Get("users").([]common.Address)[0].String()
		td.TxState = setTxState(event.Data.Get("state").(uint8))

		if err := bootstrap.SendMessage(window, "onTransactionCreate", td); err != nil {
			rlog.Error("failed to send onTransactionCreate event, error:", err)
		}
	}()
	return true
}

// Get proof files from proofIDs.
// ipfsGet -> modify file name, how can I get extension?
func getProofFiles(ipfsIDs [][32]byte) []string {
	var (
		err    error
		proofs []string = make([]string, len(ipfsIDs))
	)
	for i := 0; i < len(ipfsIDs); i++ {
		var ipfsID string = ipfsBytes32ToHash(ipfsIDs[i])
		// optimize: user can set IPFS out dir himself.
		if err = ipfsaccess.GetIAInstance().GetFromIPFS(ipfsID, IPFSOutDir); err != nil {
			rlog.Error("Node - onTransactionCreate.callback: ipfs get failed. ", err.Error())
			return nil
		}
		proofs[i] = ipfsID
		// add extension here.
	}

	return proofs
}
func ipfsBytes32ToHash(ipfsb [32]byte) string {
	var byte34 []byte = make([]byte, 34)
	// if ipfs modify encrypt algorithm, byte will change together.
	copy(byte34[:2], []byte{byte(18), byte(32)})
	copy(byte34[2:], ipfsb[:])

	return base58.Encode(byte34)
}

func onPurchase(event events.Event) bool {
	go func() {
		var pd definition.PurchaseDetails = definition.PurchaseDetails{}
		pd.Block = event.BlockNumber
		pd.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		pd.MetaDataIdEncWithSeller = event.Data.Get("metaDataIdEncSeller").([]byte)
		pd.TxState = setTxState(event.Data.Get("state").(uint8))

		if err := bootstrap.SendMessage(window, "onPurchase", pd); err != nil {
			rlog.Error("failed to send onPurchase event, error:", err)
		}
	}()
	return true
}

func onReadyForDownload(event events.Event) bool {
	go func() {
		var rd definition.ReEncryptDetails = definition.ReEncryptDetails{}
		rd.Block = event.BlockNumber
		rd.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		rd.MetaDataIdEncWithBuyer = event.Data.Get("metaDataIdEncBuyer").([]byte)
		rd.TxState = setTxState(event.Data.Get("state").(uint8))

		if err := bootstrap.SendMessage(window, "onReadyForDownload", rd); err != nil {
			rlog.Error("failed to send onReadyForDownload event, error:", err)
		}
	}()
	return true
}

func onClose(event events.Event) bool {
	go func() {
		var cd definition.CloseDetails = definition.CloseDetails{}
		cd.Block = event.BlockNumber
		cd.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		cd.TxState = setTxState(event.Data.Get("state").(uint8))

		if err := bootstrap.SendMessage(window, "onClose", cd); err != nil {
			rlog.Error("failed to send onClose event, error:", err)
		}
	}()

	return true
}

func setTxState(state byte) string {
	var str string
	switch state {
	case 1: str = "Created"
	case 2: str = "Voted"
	case 3: str = "Buying"
	case 4: str = "ReadyForDownload"
	case 5: str = "Arbitrating"
	case 6: str = "Payed"
	case 7: str = "Closed"
	}
	return str
}