package websocket

import (
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryInfo/dp/app/app"
	settings2 "github.com/scryInfo/dp/app/app/settings"
	events2 "github.com/scryInfo/dp/dots/binary/sdk/core/ethereum/events"
	ipfsaccess2 "github.com/scryInfo/dp/dots/binary/sdk/util/storage/ipfsaccess"
	rlog "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
)

func onPublish(event events2.Event) bool {
	var err error
	go func() {
		var op settings2.OnPublish
		if op, err = getPubDataDetails(event.Data.Get("despDataId").(string)); err != nil {
			rlog.Error(errors.Wrap(err, "onPublish: get publish data details failed. "))
		}
		op.Block = event.BlockNumber
		op.Price = event.Data.Get("price").(*big.Int).String()
		op.PublishID = event.Data.Get("publishId").(string)
		op.SupportVerify = event.Data.Get("supportVerify").(bool)

		if err = sendMessage("onPublish", op); err != nil {
			rlog.Error(errors.Wrap(err, "onPublish"+EventSendFailed))
		}
	}()
	return true
}

func getPubDataDetails(ipfsID string) (detailsData settings2.OnPublish, err error) {
	defer func() {
		if er := recover(); er != nil {
			rlog.Error(errors.Wrap(er.(error), "onPublish.callback: get publish data details failed. "))
		}
	}()

	for {
		var fileName string
		{
			outDir := app.GetGapp().ScryInfo.Config.IPFSOutDir
			if err = ipfsaccess2.GetIAInstance().GetFromIPFS(ipfsID, outDir); err != nil {
				break
			}

			oldFileName := outDir + "/" + ipfsID
			fileName = oldFileName + ".txt"

			if err = os.Rename(oldFileName, fileName); err != nil {
				break
			}
		}

		var details []byte
		if details, err = ioutil.ReadFile(fileName); err != nil {
			break
		}
		if err = json.Unmarshal(details, &detailsData); err != nil {
			break
		}

		if err = os.Remove(fileName); err != nil {
			err = errors.Wrap(err, "onPublish.callback: delete details file failed. ")
			rlog.Debug(err)
		}

		break
	}

	return
}

func onApprove(_ events2.Event) bool {
	go func() {
		// -
	}()
	return true
}

func onVerifiersChosen(event events2.Event) bool {
	var err error
	go func() {
		var (
			ovc        settings2.OnVerifiersChosen
			extensions []string
		)
		ovc.PublishID = event.Data.Get("publishId").(string)
		if err = sendMessage("onProofFilesExtensions", ovc.PublishID); err != nil {
			rlog.Error(errors.Wrap(err, "onProofFilesExtensions"+EventSendFailed))
		}

		ovc.Block = event.BlockNumber
		ovc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		ovc.TxState = setTxState(event.Data.Get("state").(uint8))

		extensions = <-channel
		if ovc.ProofFileNames, err = getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
			rlog.Error(errors.Wrap(err, "Node - onVC.callback: get and rename proof files failed. "))
		}

		if err = sendMessage("onVerifiersChosen", ovc); err != nil {
			rlog.Error(errors.Wrap(err, "onVerifiersChosen"+EventSendFailed))
		}
	}()

	return true
}

func onTransactionCreate(event events2.Event) bool {
	go func() {
		var (
			otc        settings2.OnTransactionCreate
			extensions []string
			err error
		)
		otc.PublishID = event.Data.Get("publishId").(string)
		if err = sendMessage("onProofFilesExtensions", otc.PublishID); err != nil {
			rlog.Error(errors.Wrap(err, "onProofFilesExtensions"+EventSendFailed))
		}

		otc.Block = event.BlockNumber
		otc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		otc.Buyer = event.Data.Get("users").([]common.Address)[0].String()
		otc.StartVerify = event.Data.Get("needVerify").(bool)
		otc.TxState = setTxState(event.Data.Get("state").(uint8))

		extensions = <-channel
		if otc.ProofFileNames, err = getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
			rlog.Error(errors.Wrap(err, "Node - onTC.callback: get and rename proof files failed. "))
		}
		if err = sendMessage("onTransactionCreate", otc); err != nil {
			rlog.Error(errors.Wrap(err, "onTransactionCreate"+EventSendFailed))
		}
	}()

	return true
}

func getAndRenameProofFiles(ipfsIDs [][32]byte, extensions []string) ([]string, error) {
	if len(ipfsIDs) != len(extensions) {
		return nil, errors.New("Quantity of IPFS IDs or extensions is wrong. ")
	}

	defer func() {
		if er := recover(); er != nil {
			rlog.Error(errors.Wrap(er.(error), "in callback: get and rename proof files failed. "))
		}
	}()

	var (
		proofs                   = make([]string, len(ipfsIDs))
		oldFileName, newFileName string
	)

	outDir := app.GetGapp().ScryInfo.Config.IPFSOutDir
	for i := 0; i < len(ipfsIDs); i++ {
		ipfsID := ipfsBytes32ToHash(ipfsIDs[i])
		if err := ipfsaccess2.GetIAInstance().GetFromIPFS(ipfsID, outDir); err != nil {
			err = errors.Wrap(err, "Node - callback: IPFS get failed. ")
			break
		}
		oldFileName = outDir + "/" + ipfsID
		newFileName = oldFileName + extensions[i]
		if err := os.Rename(oldFileName, newFileName); err != nil {
			err = errors.Wrap(err, "Node - callback: rename proof file failed. ")
			break
		}
		proofs[i] = newFileName
	}

	return proofs, nil
}
func ipfsBytes32ToHash(ipfsb [32]byte) string {
	byte34 := make([]byte, 34)
	// if ipfs change encrypt algorithm, byte will change together.
	copy(byte34[:2], []byte{byte(18), byte(32)})
	copy(byte34[2:], ipfsb[:])

	return base58.Encode(byte34)
}

func onPurchase(event events2.Event) bool {
	go func() {
		var op settings2.OnPurchase

		op.Block = event.BlockNumber
		op.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		op.MetaDataIdEncWithSeller = event.Data.Get("metaDataIdEncSeller").([]byte)
		op.PublishID = event.Data.Get("publishId").(string)
		op.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
		op.TxState = setTxState(event.Data.Get("state").(uint8))

		// temp
		op.Buyer = event.Data.Get("buyer").(common.Address).String()

		if err := sendMessage("onPurchase", op); err != nil {
			rlog.Error(errors.Wrap(err, "onPurchase"+EventSendFailed))
		}
	}()

	return true
}

func onReadyForDownload(event events2.Event) bool {
	go func() {
		var orfd settings2.OnReadyForDownload

		orfd.Block = event.BlockNumber
		orfd.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		orfd.MetaDataIdEncWithBuyer = event.Data.Get("metaDataIdEncBuyer").([]byte)
		orfd.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
		orfd.TxState = setTxState(event.Data.Get("state").(uint8))

		if err := sendMessage("onReadyForDownload", orfd); err != nil {
			rlog.Error(errors.Wrap(err, "onReadyForDownload"+EventSendFailed))
		}
	}()

	return true
}

func onClose(event events2.Event) bool {
	go func() {
		var oc settings2.OnClose

		oc.Block = event.BlockNumber
		oc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		oc.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
		oc.TxState = setTxState(event.Data.Get("state").(uint8))

		if err := sendMessage("onClose", oc); err != nil {
			rlog.Error(errors.Wrap(err, "onClose"+EventSendFailed))
		}
	}()

	return true
}

func onRegisterAsVerifier(event events2.Event) bool {
	go func() {
		var orav settings2.OnRegisterAsVerifier

		orav.Block = event.BlockNumber

		if err := sendMessage("onRegisterVerifier", orav); err != nil {
			rlog.Error(errors.Wrap(err, "onRegisterVerifier"+EventSendFailed))
		}
	}()

	return true
}

func onVote(event events2.Event) bool {
	go func() {
		var ov      settings2.OnVote

		ov.Block = event.BlockNumber
		ov.VerifierIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
		ov.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		ov.TxState = setTxState(event.Data.Get("state").(uint8))

		judge := event.Data.Get("judge").(bool)
		comment := event.Data.Get("comments").(string)
		ov.VerifierResponse = setJudge(judge) + ", " + comment

		if err := sendMessage("onVote", ov); err != nil {
			rlog.Error(errors.Wrap(err, "onVote"+EventSendFailed))
		}
	}()

	return true
}

func onVerifierDisable(event events2.Event) bool {
	go func() {
		var ovd settings2.OnVerifierDisable

		ovd.Block = event.BlockNumber

		if err := sendMessage("onVerifierDisable", ovd); err != nil {
			rlog.Error(errors.Wrap(err, "onVerifierDisable"+EventSendFailed))
		}
	}()

	return true
}

func setTxState(state byte) (str string) {
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
		str = "Closed"
	default:
		str = "Unknown TxState!"
	}

	return
}

func setJudge(judge bool) (str string) {
	if judge {
		str = "Suggest to buy"
	} else {
		str = "Not suggest to buy"
	}

	return
}
