// Scry Info.  All rights reserved.
// license that can be found in the license file.

package msg_handler

import (
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	app2 "github.com/scryinfo/dp/dots/app"
	"github.com/scryinfo/dp/dots/app/connection"
	"github.com/scryinfo/dp/dots/app/settings"
	"github.com/scryinfo/dp/dots/eth/event"
	"github.com/scryinfo/dp/dots/storage"
	"go.uber.org/zap"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
)

func onPublish(event event.Event) bool {
	var op settings.OnPublish
	{
		var err error
		if op, err = getPubDataDetails(event.Data.Get("despDataId").(string)); err != nil {
			dot.Logger().Errorln("", zap.NamedError("onPublish: get publish data details failed. ", err))
		}
		op.Block = event.BlockNumber
		op.Price = event.Data.Get("price").(*big.Int).String()
		op.PublishID = event.Data.Get("publishId").(string)
		op.SupportVerify = event.Data.Get("supportVerify").(bool)
	}

	if err := app2.GetGapp().Connection.SendMessage("onPublish", op); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onPublish"+ connection.EventSendFailed, err))
	}

	return true
}

func getPubDataDetails(ipfsID string) (detailsData settings.OnPublish, err error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("onPublish.callback: get publish data details failed. ", er))
		}
	}()

	var fileName string
	{
		outDir := storage.GetIPFSConfig().OutDir
		if err = storage.GetIPFSIns().Get(ipfsID, outDir); err != nil {
			return
		}

		oldFileName := outDir + "/" + ipfsID
		fileName = oldFileName + ".txt"

		if err = os.Rename(oldFileName, fileName); err != nil {
			return
		}
	}

	{
		var details []byte
		if details, err = ioutil.ReadFile(fileName); err != nil {
			return
		}
		if err = json.Unmarshal(details, &detailsData); err != nil {
			return
		}
	}

	if err = os.Remove(fileName); err != nil {
		dot.Logger().Debugln("", zap.NamedError("onPublish.callback: delete details file failed. ", err))
	}

	return
}

func onApprove(_ event.Event) bool {
	return true
}

func onVerifiersChosen(event event.Event) bool {
	var ovc settings.OnVerifiersChosen
	{
		ovc.PublishID = event.Data.Get("publishId").(string)
		if err := app2.GetGapp().Connection.SendMessage("onProofFilesExtensions", ovc.PublishID); err != nil {
			dot.Logger().Errorln("", zap.NamedError("onProofFilesExtensions"+connection.EventSendFailed, err))
		}

		ovc.Block = event.BlockNumber
		ovc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		ovc.TxState = setTxState(event.Data.Get("state").(uint8))

		extensions := <-extChan
		var err error
		if ovc.ProofFileNames, err = getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
			dot.Logger().Errorln("", zap.NamedError("Node - onVC.callback: get and rename proof files failed. ", err))
		}
	}

	if err := app2.GetGapp().Connection.SendMessage("onVerifiersChosen", ovc); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onVerifiersChosen"+connection.EventSendFailed, err))
	}

	return true
}

func onTransactionCreate(event event.Event) bool {
	var otc settings.OnTransactionCreate
	{
		otc.PublishID = event.Data.Get("publishId").(string)
		if err := app2.GetGapp().Connection.SendMessage("onProofFilesExtensions", otc.PublishID); err != nil {
			dot.Logger().Errorln("", zap.NamedError("onProofFilesExtensions"+connection.EventSendFailed, err))
		}

		otc.Block = event.BlockNumber
		otc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		otc.Buyer = event.Data.Get("users").([]common.Address)[0].String()
		otc.StartVerify = event.Data.Get("needVerify").(bool)
		otc.TxState = setTxState(event.Data.Get("state").(uint8))

		extensions := <-extChan
		var err error
		if otc.ProofFileNames, err = getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
			dot.Logger().Errorln("", zap.NamedError("Node - onTC.callback: get and rename proof files failed. ", err))
		}
	}

	if err := app2.GetGapp().Connection.SendMessage("onTransactionCreate", otc); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onTransactionCreate"+connection.EventSendFailed, err))
	}

	return true
}

func getAndRenameProofFiles(ipfsIDs [][32]byte, extensions []string) ([]string, error) {
	if len(ipfsIDs) != len(extensions) {
		return nil, errors.New("Quantity of IPFS IDs or extensions is wrong. ")
	}

	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("in callback: get and rename proof files failed. ", er))
		}
	}()

	var proofs = make([]string, len(ipfsIDs))

	outDir := storage.GetIPFSConfig().OutDir
	for i := 0; i < len(ipfsIDs); i++ {
		ipfsID := ipfsBytes32ToHash(ipfsIDs[i])
		if err := storage.GetIPFSIns().Get(ipfsID, outDir); err != nil {
			err = errors.Wrap(err, "Node - callback: IPFS get failed. ")
			break
		}
		oldFileName := outDir + "/" + ipfsID
		newFileName := oldFileName + extensions[i]
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
	// if ipfs change encrypt algorithm, byte 18 and 32 will change together.
	copy(byte34[:2], []byte{byte(18), byte(32)})
	copy(byte34[2:], ipfsb[:])

	return base58.Encode(byte34)
}

func onPurchase(event event.Event) bool {
	var op settings.OnPurchase
	{
		op.Block = event.BlockNumber
		op.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		op.MetaDataIdEncWithSeller = event.Data.Get("metaDataIdEncSeller").([]byte)
		op.PublishID = event.Data.Get("publishId").(string)
		op.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
		op.TxState = setTxState(event.Data.Get("state").(uint8))
	}

	if err := app2.GetGapp().Connection.SendMessage("onPurchase", op); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onPurchase"+connection.EventSendFailed, err))
	}

	return true
}

func onReadyForDownload(event event.Event) bool {
	var orfd settings.OnReadyForDownload
	{
		orfd.Block = event.BlockNumber
		orfd.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		orfd.MetaDataIdEncWithBuyer = event.Data.Get("metaDataIdEncBuyer").([]byte)
		orfd.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
		orfd.TxState = setTxState(event.Data.Get("state").(uint8))
	}

	if err := app2.GetGapp().Connection.SendMessage("onReadyForDownload", orfd); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onReadyForDownload"+connection.EventSendFailed, err))
	}

	return true
}

func onClose(event event.Event) bool {
	var oc settings.OnClose
	{
		oc.Block = event.BlockNumber
		oc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		oc.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
		oc.TxState = setTxState(event.Data.Get("state").(uint8))
	}

	if err := app2.GetGapp().Connection.SendMessage("onClose", oc); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onClose"+connection.EventSendFailed, err))
	}

	return true
}

func onRegisterAsVerifier(event event.Event) bool {
	var orav settings.OnRegisterAsVerifier
	{
		orav.Block = event.BlockNumber
	}

	if err := app2.GetGapp().Connection.SendMessage("onRegisterVerifier", orav); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onRegisterVerifier"+connection.EventSendFailed, err))
	}

	return true
}

func onVote(event event.Event) bool {
	var ov settings.OnVote
	{
		ov.Block = event.BlockNumber
		ov.VerifierIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
		ov.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
		ov.TxState = setTxState(event.Data.Get("state").(uint8))

		judge := event.Data.Get("judge").(bool)
		comment := event.Data.Get("comments").(string)
		ov.VerifierResponse = setJudge(judge) + ", " + comment
	}

	if err := app2.GetGapp().Connection.SendMessage("onVote", ov); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onVote"+connection.EventSendFailed, err))
	}

	return true
}

func onVerifierDisable(event event.Event) bool {
	var ovd settings.OnVerifierDisable
	{
		ovd.Block = event.BlockNumber
	}

	if err := app2.GetGapp().Connection.SendMessage("onVerifierDisable", ovd); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onVerifierDisable"+connection.EventSendFailed, err))
	}

	return true
}

func onArbitrationBegin(event event.Event) bool {
	var oab settings.OnArbitrationBegin
	{
		oab.PublishId = event.Data.Get("publishId").(string)
		if err := app2.GetGapp().Connection.SendMessage("onProofFilesExtensions", oab.PublishId); err != nil {
			dot.Logger().Errorln("", zap.NamedError("onProofFilesExtensions"+connection.EventSendFailed, err))
		}

		oab.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
		oab.MetaDataIdEncArbitrator = event.Data.Get("metaDataIdEncArbitrator").([]byte)
		oab.Block = event.BlockNumber

		extensions := <-extChan
		var err error
		if oab.ProofFileNames, err = getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
			dot.Logger().Errorln("", zap.NamedError("Node - onVC.callback: get and rename proof files failed. ", err))
		}
	}

	if err := app2.GetGapp().Connection.SendMessage("onArbitrationBegin", oab); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onArbitrationBegin"+connection.EventSendFailed, err))
	}

	return true
}

func onArbitrationResult(event event.Event) bool {
	var oar settings.OnArbitrationResult
	{
		oar.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
		oar.ArbitrateResult = setArbitrateResult(event.Data.Get("judge").(bool))
		oar.User = strconv.Itoa(int(event.Data.Get("identify").(uint8)))
		oar.Block = event.BlockNumber
	}

	if err := app2.GetGapp().Connection.SendMessage("onArbitrationResult", oar); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onArbitrationResult"+connection.EventSendFailed, err))
	}

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

func setArbitrateResult(result bool) string {
	str := "Arbitrate Result: "
	if result {
		str += "true. "
	} else {
		str += "false. "
	}

	return str
}
