package transmission

import (
	"encoding/json"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryinfo/iscap/demo/src/application/definition"
	"github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
	"github.com/scryinfo/iscap/demo/src/sdk/util/storage/ipfsaccess"
	rlog "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"os"
)

const (
	IPFSOutDir = "D:/desktop"
	EventSendFailed = " event send failed. "
)
var err error

func onPublish(event events.Event) bool {
	go func() {
		var (
			op *definition.OnPublish
		)
		if op, err = getPubDataDetails(event.Data.Get("despDataId").(string)); err != nil {
			rlog.Error(errors.Wrap(err, "onPublish: get publish data details failed. "))
		}
		op.Block = event.BlockNumber
		op.Price = event.Data.Get("price").(*big.Int).String()
		op.PublishID = event.Data.Get("publishId").(string)
		op.SupportVerify = event.Data.Get("supportVerify").(bool)

		if err = bootstrap.SendMessage(window, "onPublish", &op); err != nil {
			rlog.Error(errors.Wrap(err, "onPublish"+EventSendFailed))
		}
	}()
	return true
}

// Get publish data details from details' ipfsID.
// ipfsGet -> modify file name -> read file -> json.unmarshal -> delete file
func getPubDataDetails(ipfsID string) (*definition.OnPublish, error) {
	defer func() {
		if er := recover(); er != nil {
			rlog.Error(errors.Wrap(er.(error), "onPublish.callback: get publish data details failed. "))
		}
	}()
	if err = ipfsaccess.GetIAInstance().GetFromIPFS(ipfsID); err != nil {
		return nil, errors.Wrap(err, "Node - onPublish.callback: ipfs get failed. ")
	}

	oldFileName := IPFSOutDir + "/" + ipfsID
	newFileName := oldFileName + ".txt"
	if err = os.Rename(oldFileName, newFileName); err != nil {
		return nil, errors.Wrap(err, "Node - onPublish.callback: rename details file failed. ")
	}

	var details []byte
	if details, err = ioutil.ReadFile(newFileName); err != nil {
		return nil, errors.Wrap(err, "Node - onPublish.callback: read details file failed. ")
	}
	var detailsData definition.OnPublish
	if err = json.Unmarshal(details, &detailsData); err != nil {
		return nil, errors.Wrap(err, "Node - onPublish.callback: json unmarshal details failed. ")
	}

	if err = os.Remove(newFileName); err != nil {
		rlog.Debug("Node - onPublish.callback: delete details file failed. ", err)
	}

	return &detailsData, nil
}

func onApprove(event events.Event) bool {
	go func() {
		var oa definition.OnApprove
		oa.Block = event.BlockNumber

		if err = bootstrap.SendMessage(window, "onApprove", oa); err != nil {
			rlog.Error(errors.Wrap(err, "onApprove"+EventSendFailed))
		}
	}()
	return true
}

func onTransactionCreate(event events.Event) bool {
	go func() {
		var (
			otc definition.OnTransactionCreate
			extensions []string
		)
		otc.PublishID = event.Data.Get("publishId").(string)
		if err = bootstrap.SendMessage(window, "onProofFilesExtensions", otc.PublishID); err != nil {
			rlog.Error(errors.Wrap(err, "onProofFilesExtensions"+EventSendFailed))
		}

		otc.Block = event.BlockNumber
		otc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		otc.Buyer = event.Data.Get("users").([]common.Address)[0].String()
		otc.StartVerify = event.Data.Get("startVerify").(bool)
		otc.Verifier1 = event.Data.Get("verifiers").([]common.Address)[0].String()
		otc.Verifier2 = event.Data.Get("verifiers").([]common.Address)[1].String()
		otc.TxState = setTxState(event.Data.Get("state").(uint8))

		extensions = <- channel
		if otc.ProofFileNames, err = getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
			rlog.Error(errors.Wrap(err, "Node - onTC.callback: get and rename proof files failed. "))
		}
		if err = bootstrap.SendMessage(window, "onTransactionCreate", otc); err != nil {
			rlog.Error(errors.Wrap(err, "onTransactionCreate"+EventSendFailed))
		}
	}()
	return true
}

// Get proof files from proofIDs.
// ipfsGet -> modify file name, how can I get extension?
func getAndRenameProofFiles(ipfsIDs [][32]byte, extensions []string) ([]string, error) {
	defer func() {
		if er := recover(); er != nil {
			rlog.Error(errors.Wrap(er.(error), "onTransactionCreate.callback: get and rename proof files failed. "))
		}
	}()
	var (
		proofs = make([]string, len(ipfsIDs))
		oldFileName, newFileName string
	)
	if len(ipfsIDs) != len(extensions) {
		return nil, errors.New("Invalid IPFS IDs or extensions. ")
	}
	for i := 0; i < len(ipfsIDs); i++ {
		ipfsID := ipfsBytes32ToHash(ipfsIDs[i])
		if err = ipfsaccess.GetIAInstance().GetFromIPFS(ipfsID); err != nil {
			return nil, errors.Wrap(err, "Node - onTransactionCreate.callback: IPFS get failed. ")
		}
		oldFileName = IPFSOutDir + "/" + ipfsID
		newFileName = oldFileName + extensions[i]
		if err = os.Rename(oldFileName, newFileName); err != nil {
			return nil, errors.Wrap(err, "Node - onTransactionCreate.callback: rename proof file failed. ")
		}
		proofs[i] = newFileName
	}

	return proofs, nil
}
func ipfsBytes32ToHash(ipfsb [32]byte) string {
	byte34 := make([]byte, 34)
	// if ipfs modify encrypt algorithm, byte will change together.
	copy(byte34[:2], []byte{byte(18), byte(32)})
	copy(byte34[2:], ipfsb[:])

	return base58.Encode(byte34)
}

func onPurchase(event events.Event) bool {
	go func() {
		var op definition.OnPurchase
		op.Block = event.BlockNumber
		op.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		op.MetaDataIdEncWithSeller = event.Data.Get("metaDataIdEncSeller").([]byte)
		op.TxState = setTxState(event.Data.Get("state").(uint8))

		if err = bootstrap.SendMessage(window, "onPurchase", op); err != nil {
			rlog.Error(errors.Wrap(err, "onPurchase"+EventSendFailed))
		}
	}()
	return true
}

func onReadyForDownload(event events.Event) bool {
	go func() {
		var orfd definition.OnReadyForDownload
		orfd.Block = event.BlockNumber
		orfd.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		orfd.MetaDataIdEncWithBuyer = event.Data.Get("metaDataIdEncBuyer").([]byte)
		orfd.TxState = setTxState(event.Data.Get("state").(uint8))

		if err = bootstrap.SendMessage(window, "onReadyForDownload", orfd); err != nil {
			rlog.Error(errors.Wrap(err, "onReadyForDownload"+EventSendFailed))
		}
	}()
	return true
}

func onClose(event events.Event) bool {
	go func() {
		var oc definition.OnClose
		oc.Block = event.BlockNumber
		oc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		oc.TxState = setTxState(event.Data.Get("state").(uint8))

		if err = bootstrap.SendMessage(window, "onClose", oc); err != nil {
			rlog.Error(errors.Wrap(err, "onClose"+EventSendFailed))
		}
	}()

	return true
}

func onRegisterAsVerifier(event events.Event) bool {
	go func() {
		var orav definition.OnRegisterAsVerifier
		orav.Block = event.BlockNumber

		if err = bootstrap.SendMessage(window, "onRegisterVerifier", orav); err != nil {
			rlog.Error(errors.Wrap(err, "onRegisterVerifier"+EventSendFailed))
		}
	}()

	return true
}

func onVote(event events.Event) bool {
	go func() {
		var (
			ov definition.OnVote
			judge bool
			comment string
		)
		ov.Block = event.BlockNumber
		ov.VerifierIndex = event.Data.Get("index").(*big.Int).String()
		judge = event.Data.Get("judge").(bool)
		comment = event.Data.Get("comments").(string)
		ov.VerifierResponse = setJudge(judge) + ", " + comment
		ov.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		ov.TxState = setTxState(event.Data.Get("state").(uint8))

		if err = bootstrap.SendMessage(window, "onVote", ov); err != nil {
			rlog.Error(errors.Wrap(err, "onVote"+EventSendFailed))
		}
	}()

	return true
}

func onVerifierDisable(event events.Event) bool {
	go func() {
		var ovd definition.OnVerifierDisable
		ovd.Block = event.BlockNumber

		if err = bootstrap.SendMessage(window, "onVerifierDisable", ovd); err != nil {
			rlog.Error(errors.Wrap(err, "onVerifierDisable"+EventSendFailed))
		}
	}()

	return true
}

func setTxState(state byte) string {
	var str string
	switch state {
	case 1:
		str = "Created"
	case 2:
		str = "Voted"
	case 3:
		str = "Buying"
	case 4:
		str = "ReadyForDownload"
	case 5:
		str = "Arbitrating"
	case 6:
		str = "Payed"
	case 7:
		str = "Closed"
	}
	return str
}

func setJudge(judge bool) string {
	var str string = "Not suggest to buy"
	if judge {
		str = "Suggest to buy"
	}
	return str
}
