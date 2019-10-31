package cec

import (
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/app/server"
	database "github.com/scryinfo/dp/dots/app/storage"
	"github.com/scryinfo/dp/dots/app/storage/definition"
	scry2 "github.com/scryinfo/dp/dots/binary/scry"
	"github.com/scryinfo/dp/dots/eth/event"
	"github.com/scryinfo/dp/dots/storage/ipfs"
	"go.uber.org/zap"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

type Callbacks struct {
	CurUser      scry2.Client
	EventNames   []string
	EventHandler []event.Callback
	FlagChan     chan bool // flag for scanned approval event.
	config       cbsConfig
	WS           *server.WSServer `dot:""`
	Storage      *storage.Ipfs    `dot:""`
	DB           *database.SQLite `dot:""`
}

type cbsConfig struct {
	ProofsOutDir string `json:"proofsOutDir"`
}

const (
	CBsTypeId = "36b2b9b7-1559-4d57-a388-f8224072a5d1"
)

func (c *Callbacks) Create(l dot.Line) error {
	c.FlagChan = make(chan bool, 10)

	c.EventNames = []string{
		"Approval",
		"Publish",
		"VerifiersChosen",
		"AdvancePurchase",
		"ConfirmPurchase",
		"ReEncrypt",
		"TransactionClose",
		"RegisterVerifier",
		"VoteResult",
		"VerifierDisable",
		"ArbitrationBegin",
		"ArbitrationResult",
	}

	c.EventHandler = []event.Callback{
		c.onApprove,
		c.onPublish,
		c.onVerifiersChosen,
		c.onAdvancePurchase,
		c.onConfirmPurchase,
		c.onReEncrypt,
		c.onTransactionClose,
		c.onRegisterAsVerifier,
		c.onVoteResult,
		c.onVerifierDisable,
		c.onArbitrationBegin,
		c.onArbitrationResult,
	}

	if len(c.EventNames) != len(c.EventHandler) {
		return errors.New("Quantities of name and function are not matched. (callbacks) ")
	}

	return nil
}

//construct dot
func newCBsDot(conf interface{}) (dot.Dot, error) {
	var err error
	var bs []byte
	if bt, ok := conf.([]byte); ok {
		bs = bt
	} else {
		return nil, dot.SError.Parameter
	}

	dConf := &cbsConfig{}
	err = dot.UnMarshalConfig(bs, dConf)
	if err != nil {
		return nil, err
	}

	d := &Callbacks{config: *dConf}

	return d, err
}

func CBsTypeLive() []*dot.TypeLives {
	return []*dot.TypeLives{
		{
			Meta: dot.Metadata{
				TypeId: CBsTypeId,
				NewDoter: func(conf interface{}) (dot.Dot, error) {
					return newCBsDot(conf)
				},
			},
		},
		server.WebSocketTypeLive(),
		storage.IpfsTypeLive(),
		database.SQLiteTypeLive(),
	}
}

type Details struct {
	Title               string   `json:"Title,omitempty"`
	Keys                string   `json:"Keys,omitempty"`
	Description         string   `json:"Description,omitempty"`
	Seller              string   `json:"Seller,omitempty"`
	MetaDataExtension   string   `json:"MetaDataExtension,omitempty"`
	ProofDataExtensions []string `json:"ProofDataExtensions,omitempty"`
}

func (c *Callbacks) onPublish(event event.Event) bool {
	var dl definition.DataList
	{
		dl.CreatedTime = time.Now().Unix()
		dl.Price = event.Data.Get("price").(*big.Int).String()
		dl.PublishId = event.Data.Get("publishId").(string)
		dl.SupportVerify = event.Data.Get("supportVerify").(bool)

		var (
			details Details
			err     error
		)
		if details, err = c.getPubDataDetails(event.Data.Get("despDataId").(string)); err != nil {
			dot.Logger().Errorln("", zap.NamedError("onPublish: get publish data details failed. ", err))
		}

		dl.Title = details.Title
		dl.Keys = details.Keys
		dl.Description = details.Description
		dl.Seller = strings.ToLower(details.Seller)
		dl.MetaDataExtension = details.MetaDataExtension

		if dl.ProofDataExtensions, err = json.Marshal(details.ProofDataExtensions); err != nil {
			dot.Logger().Errorln("json marshal failed", zap.NamedError("[]string to []byte", err))
			return false
		}
	}

	if num, err := c.DB.Read(&definition.DataList{}, "", "publish_id = ?", dl.PublishId); num == 0 { // "error": "record not found"
		if num, err := c.DB.Insert(&dl); num != 1 || err != nil {
			dot.Logger().Errorln("db insert failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
			return false
		}
	} else if num != 1 {
		dot.Logger().Errorln("db read failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
		return false
	}

	if !c.updateFromBlock(event.BlockNumber) {
		return false
	}

	if err := c.WS.SendMessage("onPublish", dl); err != nil {
		dot.Logger().Errorln(server.EventSendFailed, zap.NamedError("onPublish", err))
	}

	return true
}

func (c *Callbacks) getPubDataDetails(ipfsId string) (detailsData Details, err error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("onPublish.callback: get publish data details failed. ", er))
		}
	}()

	var fileName string
	{
		outDir := c.config.ProofsOutDir
		if err = c.Storage.Get(ipfsId, outDir); err != nil {
			return
		}

		oldFileName := outDir + "/" + ipfsId
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

func (c *Callbacks) onApprove(_ event.Event) bool {
	c.FlagChan <- true
	return true
}

func (c *Callbacks) onVerifiersChosen(event event.Event) bool {
	var tx definition.Transaction
	{
		tx.PublishId = event.Data.Get("publishId").(string)
		tx.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
		tx.State = int(event.Data.Get("state").(uint8))

		if err := c.getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), c.getProofFileExtensions(tx.PublishId)); err != nil {
			dot.Logger().Errorln("", zap.NamedError("Node - onVC.callback: get and rename proof files failed. ", err))
			return false
		}
	}

	if !c.makeTxWithDataDetails(&tx, tx.PublishId) {
		return false
	}

	// update acc info
	{
		var acc definition.Account

		if num, err := c.DB.Read(&acc, "", "address = ?", c.CurUser.Account().Addr); num != 1 || err != nil {
			dot.Logger().Errorln("id not in slice", zap.Int64("affect rows number", num), zap.NamedError("error", err))
			return false
		}

		if bs, err := appendSlice(acc.Verify, tx.TransactionId); err == nil {
			acc.Verify = bs
		} else {
			dot.Logger().Errorln("append verify failed. ", zap.NamedError("error", err))
			return false
		}

		if num, err := c.DB.Update(&acc, map[string]interface{}{
			"verify":     acc.Verify,
			"from_block": event.BlockNumber,
		}, "address = ?", c.CurUser.Account().Addr); num != 1 || err != nil {
			dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
			return false
		}
	}

	if err := c.WS.SendMessage("onVerifiersChosen", tx); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onVerifiersChosen"+server.EventSendFailed, err))
	}

	return true
}

func (c *Callbacks) onAdvancePurchase(event event.Event) bool {
	var tx definition.Transaction
	{
		tx.CreatedTime = time.Now().Unix()
		tx.PublishId = event.Data.Get("publishId").(string)
		tx.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
		tx.Buyer = strings.ToLower(event.Data.Get("users").([]common.Address)[1].String())
		tx.StartVerify = event.Data.Get("needVerify").(bool)
		tx.State = int(event.Data.Get("state").(uint8))

		if err := c.getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), c.getProofFileExtensions(tx.PublishId)); err != nil {
			dot.Logger().Errorln("", zap.NamedError("Node - onAP.callback: get and rename proof files failed. ", err))
		}
	}

	if !c.makeTxWithDataDetails(&tx, tx.PublishId) {
		return false
	}

	tx.Identify = 1
	if tx.Buyer == strings.ToLower(c.CurUser.Account().Addr) {
		tx.Identify = 2
		if num, err := c.DB.Read(&definition.Transaction{}, "", "transaction_id = ?", tx.TransactionId); num == 0 { // "error": "record not found"
			if num, err := c.DB.Insert(&tx); num != 1 || err != nil {
				dot.Logger().Errorln("db insert failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
				return false
			}
		} else if num != 1 {
			dot.Logger().Errorln("db read failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
			return false
		}
	}

	if !c.updateFromBlock(event.BlockNumber) {
		return false
	}

	if err := c.WS.SendMessage("onAdvancePurchase", tx); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onTransactionCreate"+server.EventSendFailed, err))
	}

	return true
}

func (c *Callbacks) getAndRenameProofFiles(ipfsIds [][32]byte, extensions []string) error {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("in callback: get and rename proof files failed. ", er))
		}
	}()

	if len(ipfsIds) != len(extensions) {
		return errors.New("Quantity of IPFS Ids or extensions is wrong. " + strconv.Itoa(len(ipfsIds)) + ", " + strconv.Itoa(len(extensions)))
	}

	outDir := c.config.ProofsOutDir
	for i := 0; i < len(ipfsIds); i++ {
		ipfsId := ipfsBytes32ToHash(ipfsIds[i])
		if err := c.Storage.Get(ipfsId, outDir); err != nil {
			err = errors.Wrap(err, "Node - callback: IPFS get failed. ")
			break
		}
		oldFileName := outDir + "/" + ipfsId
		newFileName := oldFileName + extensions[i]
		if err := os.Rename(oldFileName, newFileName); err != nil {
			err = errors.Wrap(err, "Node - callback: rename proof file failed. ")
			break
		}
	}

	return nil
}
func ipfsBytes32ToHash(ipfsb [32]byte) string {
	byte34 := make([]byte, 34)
	// if ipfs change encrypt algorithm, byte 18 and 32 will change together.
	copy(byte34[:2], []byte{byte(18), byte(32)})
	copy(byte34[2:], ipfsb[:])

	return base58.Encode(byte34)
}

func (c *Callbacks) onConfirmPurchase(event event.Event) bool {
	var tx definition.Transaction
	if num, err := c.DB.Update(&tx, map[string]interface{}{
		"meta_data_id_enc_with_seller": string(event.Data.Get("metaDataIdEncSeller").([]byte)),
		"state":                        int(event.Data.Get("state").(uint8)),
	}, "transaction_id = ?", event.Data.Get("transactionId").(*big.Int).String()); num != 1 || err != nil {
		dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
		return false
	}

	switch strings.ToLower(c.CurUser.Account().Addr) {
	case strings.ToLower(tx.Seller):
		tx.Identify = 1
	case strings.ToLower(tx.Buyer):
		tx.Identify = 2
	default:
		tx.Identify = 3
	}

	// update acc info
	{
		var acc definition.Account

		if num, err := c.DB.Read(&acc, "", "address = ?", c.CurUser.Account().Addr); num != 1 || err != nil {
			dot.Logger().Errorln("db read failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
			return false
		}

		if bs, err := DeleteSlice(acc.Verify, tx.TransactionId); err == nil {
			acc.Verify = bs
		} else {
			dot.Logger().Errorln("delete verify failed. ", zap.NamedError("error", err))
			return false
		}

		if num, err := c.DB.Update(&acc, map[string]interface{}{
			"verify":     acc.Verify,
			"from_block": event.BlockNumber,
		}, "address = ?", c.CurUser.Account().Addr); num != 1 || err != nil {
			dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
			return false
		}
	}

	if err := c.WS.SendMessage("onConfirmPurchase", tx); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onPurchase"+server.EventSendFailed, err))
	}

	return true
}

func (c *Callbacks) onReEncrypt(event event.Event) bool {
	var tx definition.Transaction
	if num, err := c.DB.Update(&tx, map[string]interface{}{
		"meta_data_id_enc_with_buyer": string(event.Data.Get("metaDataIdEncBuyer").([]byte)),
		"state":                       int(event.Data.Get("state").(uint8)),
	}, "transaction_id = ?", event.Data.Get("transactionId").(*big.Int).String()); num != 1 || err != nil {
		dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
		return false
	}

	switch strings.ToLower(c.CurUser.Account().Addr) {
	case strings.ToLower(tx.Seller):
		tx.Identify = 1
	case strings.ToLower(tx.Buyer):
		tx.Identify = 2
	default:
		dot.Logger().Errorln("Invalid identify! (on re-encrypt)")
	}

	if !c.updateFromBlock(event.BlockNumber) {
		return false
	}

	if err := c.WS.SendMessage("onReEncrypt", tx); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onReadyForDownload"+server.EventSendFailed, err))
	}

	return true
}

func (c *Callbacks) onTransactionClose(event event.Event) bool {
	var tx definition.Transaction
	if num, err := c.DB.Update(&tx, map[string]interface{}{
		"state": int(event.Data.Get("state").(uint8)),
	}, "transaction_id = ?", event.Data.Get("transactionId").(*big.Int).String()); num != 1 || err != nil {
		dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
		return false
	}

	switch strings.ToLower(c.CurUser.Account().Addr) {
	case strings.ToLower(tx.Seller):
		tx.Identify = 1
	case strings.ToLower(tx.Buyer):
		tx.Identify = 2
	default:
		tx.Identify = 3
	}

	if !c.updateFromBlock(event.BlockNumber) {
		return false
	}

	if err := c.WS.SendMessage("onTransactionClose", tx); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onClose"+server.EventSendFailed, err))
	}

	return true
}

func (c *Callbacks) onRegisterAsVerifier(event event.Event) bool {
	if num, err := c.DB.Update(&definition.Account{}, map[string]interface{}{
		"from_block":  int64(event.BlockNumber + 1),
		"is_verifier": true,
	}, "address = ?", c.CurUser.Account().Addr); num != 1 || err != nil {
		dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
		return false
	}

	if err := c.WS.SendMessage("onRegisterVerifier", ""); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onRegisterVerifier"+server.EventSendFailed, err))
	}

	return true
}

func (c *Callbacks) onVoteResult(event event.Event) bool {
	var tx definition.Transaction

	switch int(event.Data.Get("index").(uint8)) {
	case 0:
		tx.Identify = 2
	case 1:
		tx.Identify = 3
		if num, err := c.DB.Update(&tx, map[string]interface{}{
			"verifier1_response": setJudge(event.Data.Get("judge").(bool)) + ". " + event.Data.Get("comments").(string),
			"state":              int(event.Data.Get("state").(uint8)),
		}, "transaction_id = ?", event.Data.Get("transactionId").(*big.Int).String()); num != 1 || err != nil {
			dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
			return false
		}
	case 2:
		tx.Identify = 3
		if num, err := c.DB.Update(&tx, map[string]interface{}{
			"verifier2_response": setJudge(event.Data.Get("judge").(bool)) + ". " + event.Data.Get("comments").(string),
			"state":              int(event.Data.Get("state").(uint8)),
		}, "transaction_id = ?", event.Data.Get("transactionId").(*big.Int).String()); num != 1 || err != nil {
			dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
			return false
		}
	}

	if !c.updateFromBlock(event.BlockNumber) {
		return false
	}

	if err := c.WS.SendMessage("onVoteResult", tx); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onVote"+server.EventSendFailed, err))
	}

	return true
}

func (c *Callbacks) onVerifierDisable(event event.Event) bool {
	var acc definition.Account
	{
		acc.Address = event.Data.Get("verifier").(common.Address).String()
	}

	if !c.updateFromBlock(event.BlockNumber) {
		return false
	}

	if err := c.WS.SendMessage("onVerifierDisable", acc); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onVerifierDisable"+server.EventSendFailed, err))
	}

	return true
}

func (c *Callbacks) onArbitrationBegin(event event.Event) bool {
	var tx definition.Transaction

	if err := c.getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), c.getProofFileExtensions(event.Data.Get("publishId").(string))); err != nil {
		dot.Logger().Errorln("", zap.NamedError("Node - onVC.callback: get and rename proof files failed. ", err))
	}

	if num, err := c.DB.Update(&tx, map[string]interface{}{
		"meta_data_id_enc_with_arbitrator": string(event.Data.Get("metaDataIdEncArbitrator").([]byte)),
	}, "transaction_id = ?", event.Data.Get("transactionId").(*big.Int).String()); num != 1 || err != nil {
		dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
		return false
	}

	// update acc info
	{
		var acc definition.Account

		if num, err := c.DB.Read(&acc, "", "address = ?", c.CurUser.Account().Addr); num != 1 || err != nil {
			dot.Logger().Errorln("db read failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
			return false
		}

		if bs, err := appendSlice(acc.Arbitrate, tx.TransactionId); err == nil {
			acc.Arbitrate = bs
		} else {
			dot.Logger().Errorln("append arbitrate failed. ", zap.NamedError("error", err))
			return false
		}

		if num, err := c.DB.Update(&acc, map[string]interface{}{
			"arbitrate":  acc.Arbitrate,
			"from_block": event.BlockNumber,
		}, "address = ?", c.CurUser.Account().Addr); num != 1 || err != nil {
			dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
			return false
		}
	}

	if err := c.WS.SendMessage("onArbitrationBegin", tx); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onArbitrationBegin"+server.EventSendFailed, err))
	}

	return true
}

func (c *Callbacks) onArbitrationResult(event event.Event) bool {
	var tx definition.Transaction

	if num, err := c.DB.Update(&tx, map[string]interface{}{
		"arbitrate_result": event.Data.Get("judge").(bool),
	}, "transaction_id = ?", event.Data.Get("transactionId").(*big.Int).String()); num != 1 || err != nil {
		dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
		return false
	}

	switch strings.ToLower(c.CurUser.Account().Addr) {
	case strings.ToLower(tx.Seller):
		tx.Identify = 1
	case strings.ToLower(tx.Buyer):
		tx.Identify = 2
	}

	if !c.updateFromBlock(event.BlockNumber) {
		return false
	}

	if err := c.WS.SendMessage("onArbitrationResult", tx); err != nil {
		dot.Logger().Errorln("", zap.NamedError("onArbitrationResult"+server.EventSendFailed, err))
	}

	return true
}

func (c *Callbacks) updateFromBlock(blockNumber uint64) bool {
	if num, err := c.DB.Update(&definition.Account{}, map[string]interface{}{"from_block": int64(blockNumber + 1)}, "address = ?", c.CurUser.Account().Addr); num != 1 || err != nil {
		dot.Logger().Errorln("db update failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
		return false
	}

	return true
}

func (c *Callbacks) getProofFileExtensions(pubId string) (extensions []string) {
	var dl definition.DataList
	if num, err := c.DB.Read(&dl, "", "publish_id = ?", pubId); num != 1 || err != nil {
		dot.Logger().Errorln("db read failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
		return
	}

	if err := json.Unmarshal(dl.ProofDataExtensions, &extensions); err != nil {
		dot.Logger().Errorln("json unmarshal failed", zap.NamedError("in getProofFileExtensions", err))
		return
	}

	return
}

func (c *Callbacks) makeTxWithDataDetails(tx *definition.Transaction, pubId string) bool {
	var dl definition.DataList
	if num, err := c.DB.Read(&dl, "", "publish_id = ?", pubId); num != 1 || err != nil {
		dot.Logger().Errorln("db read failed", zap.Int64("affect rows number", num), zap.NamedError("error", err))
		return false
	}

	tx.PublishId = dl.PublishId
	tx.Title = dl.Title
	tx.Price = dl.Price
	tx.Keys = dl.Keys
	tx.Description = dl.Description
	tx.Seller = dl.Seller
	tx.SupportVerify = dl.SupportVerify
	tx.MetaDataExtension = dl.MetaDataExtension
	tx.ProofDataExtensions = dl.ProofDataExtensions

	return true
}

func setJudge(judge bool) (str string) {
	if judge {
		str = "Suggest to buy"
	} else {
		str = "Not suggest to buy"
	}

	return
}

func appendSlice(bs []byte, str string) (result []byte, err error) {
	var ss []string
	if bs != nil {
		if err = json.Unmarshal(bs, &ss); err != nil {
			err = errors.Wrap(err, "json unmarshal failed. ")
			return
		}
	}

	if index := getIndex(ss, str); index == -1 {
		ss = append(ss, str)
	}
	if result, err = json.Marshal(ss); err != nil {
		err = errors.Wrap(err, "json marshal failed. ")
		return
	}

	return
}

func DeleteSlice(bs []byte, str string) (result []byte, err error) {
	var ss []string
	if bs != nil {
		if err = json.Unmarshal(bs, &ss); err != nil {
			err = errors.Wrap(err, "json unmarshal failed. ")
			return
		}
	}

	if index := getIndex(ss, str); index != -1 {
		ss = append(ss[:index], ss[index+1:]...)
	}

	if result, err = json.Marshal(ss); err != nil {
		err = errors.Wrap(err, "json marshal failed. ")
		return
	}

	return
}

func getIndex(ss []string, str string) int {
	for i := range ss {
		if str == ss[i] {
			return i
		}
	}

	return -1
}
