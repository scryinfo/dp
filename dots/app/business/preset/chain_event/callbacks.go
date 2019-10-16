package cec

import (
    "encoding/json"
    "github.com/btcsuite/btcutil/base58"
    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/app"
    "github.com/scryinfo/dp/dots/app/business/definition"
    "github.com/scryinfo/dp/dots/app/server"
    "github.com/scryinfo/dp/dots/eth/event"
    "github.com/scryinfo/dp/dots/storage"
    "go.uber.org/zap"
    "io/ioutil"
    "math/big"
    "os"
    "strconv"
)

type Callbacks struct {
    EventNames   []string
    EventHandler []event.Callback
    ExtChan      chan []string
    FlagChan chan bool // flag for scanned approval event.
    config       cbsConfig
    WS           *app.WSServer  `dot:""`
    Storage      *storage.Ipfs `dot:""`
}

type cbsConfig struct {
    ProofsOutDir string `json:"proofsOutDir"`
}

const (
    CBsTypeId = "36b2b9b7-1559-4d57-a388-f8224072a5d1"
)

func (c *Callbacks) Create(l dot.Line) error {
    c.ExtChan = make(chan []string, 10)
    c.FlagChan = make(chan bool, 10)

    c.EventNames = []string{
        "Approval",
        "Publish",
        "VerifiersChosen",
        "AdvancePurchase",
        "ConfirmPurchase",
        "ReEncrypt",
        "FinishPurchase",
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
        c.onFinishPurchase,
        c.onRegisterAsVerifier,
        c.onVoteResult,
        c.onVerifierDisable,
        c.onArbitrationBegin,
        c.onArbitrationResult,
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
        app.WebSocketTypeLive(),
        storage.IpfsTypeLive(),
    }
}

func (c *Callbacks) onPublish(event event.Event) bool {
    var op definition.Callbacks
    {
        var err error
        if op, err = c.getPubDataDetails(event.Data.Get("despDataId").(string)); err != nil {
            dot.Logger().Errorln("", zap.NamedError("onPublish: get publish data details failed. ", err))
        }
        op.Block = event.BlockNumber
        op.Price = event.Data.Get("price").(*big.Int).String()
        op.PublishId = event.Data.Get("publishId").(string)
        op.SupportVerify = event.Data.Get("supportVerify").(bool)
    }

    if err := c.WS.SendMessage("onPublish", op); err != nil {
        dot.Logger().Errorln("", zap.NamedError("onPublish"+server.EventSendFailed, err))
    }

    return true
}

func (c *Callbacks) getPubDataDetails(ipfsId string) (detailsData definition.Callbacks, err error) {
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
    var ovc definition.Callbacks
    {
        ovc.PublishId = event.Data.Get("publishId").(string)
        if err := c.WS.SendMessage("onProofFilesExtensions", ovc.PublishId); err != nil {
            dot.Logger().Errorln("", zap.NamedError("onProofFilesExtensions"+server.EventSendFailed, err))
        }

        ovc.Block = event.BlockNumber
        ovc.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
        ovc.TransactionState = setTxState(event.Data.Get("state").(uint8))

        extensions := <- c.ExtChan
        var err error
        if err = c.getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
            dot.Logger().Errorln("", zap.NamedError("Node - onVC.callback: get and rename proof files failed. ", err))
        }
    }

    if err := c.WS.SendMessage("onVerifiersChosen", ovc); err != nil {
        dot.Logger().Errorln("", zap.NamedError("onVerifiersChosen"+server.EventSendFailed, err))
    }

    return true
}

func (c *Callbacks) onAdvancePurchase(event event.Event) bool {
    var otc definition.Callbacks
    {
        otc.PublishId = event.Data.Get("publishId").(string)
        if err := c.WS.SendMessage("onProofFilesExtensions", otc.PublishId); err != nil {
            dot.Logger().Errorln("", zap.NamedError("onProofFilesExtensions"+server.EventSendFailed, err))
        }

        otc.Block = event.BlockNumber
        otc.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
        otc.Buyer = event.Data.Get("users").([]common.Address)[1].String()
        otc.StartVerify = event.Data.Get("needVerify").(bool)
        otc.TransactionState = setTxState(event.Data.Get("state").(uint8))

        extensions := <- c.ExtChan
        var err error
        if err = c.getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
            dot.Logger().Errorln("", zap.Strings("show extensions: ", extensions),
                zap.NamedError("Node - onAP.callback: get and rename proof files failed. ", err))
        }
    }

    if err := c.WS.SendMessage("onAdvancePurchase", otc); err != nil {
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
    var op definition.Callbacks
    {
        op.Block = event.BlockNumber
        op.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
        op.EncryptedId.EncryptedId = event.Data.Get("metaDataIdEncSeller").([]byte)
        op.PublishId = event.Data.Get("publishId").(string)
        op.TransactionState = setTxState(event.Data.Get("state").(uint8))
    }

    if err := c.WS.SendMessage("onConfirmPurchase", op); err != nil {
        dot.Logger().Errorln("", zap.NamedError("onPurchase"+server.EventSendFailed, err))
    }

    return true
}

func (c *Callbacks) onReEncrypt(event event.Event) bool {
    var orfd definition.Callbacks
    {
        orfd.Block = event.BlockNumber
        orfd.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
        orfd.EncryptedId.EncryptedId = event.Data.Get("metaDataIdEncBuyer").([]byte)
        orfd.TransactionState = setTxState(event.Data.Get("state").(uint8))
    }

    if err := c.WS.SendMessage("onReEncrypt", orfd); err != nil {
        dot.Logger().Errorln("", zap.NamedError("onReadyForDownload"+server.EventSendFailed, err))
    }

    return true
}

func (c *Callbacks) onFinishPurchase(event event.Event) bool {
    var oc definition.Callbacks
    {
        oc.Block = event.BlockNumber
        oc.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
        oc.TransactionState = setTxState(event.Data.Get("state").(uint8))
    }

    if err := c.WS.SendMessage("onFinishPurchase", oc); err != nil {
        dot.Logger().Errorln("", zap.NamedError("onClose"+server.EventSendFailed, err))
    }

    return true
}

func (c *Callbacks) onRegisterAsVerifier(event event.Event) bool {
    var orav definition.Callbacks
    {
        orav.Block = event.BlockNumber
    }

    if err := c.WS.SendMessage("onRegisterVerifier", orav); err != nil {
        dot.Logger().Errorln("", zap.NamedError("onRegisterVerifier"+server.EventSendFailed, err))
    }

    return true
}

func (c *Callbacks) onVoteResult(event event.Event) bool {
    var ov definition.Callbacks
    {
        ov.Block = event.BlockNumber
        ov.VerifyResult.VerifierIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
        ov.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
        ov.TransactionState = setTxState(event.Data.Get("state").(uint8))

        judge := event.Data.Get("judge").(bool)
        comment := event.Data.Get("comments").(string)
        ov.VerifyResult.VerifierResponse = setJudge(judge) + ", " + comment
    }

    if err := c.WS.SendMessage("onVoteResult", ov); err != nil {
        dot.Logger().Errorln("", zap.NamedError("onVote"+server.EventSendFailed, err))
    }

    return true
}

func (c *Callbacks) onVerifierDisable(event event.Event) bool {
    var ovd definition.Callbacks
    {
        ovd.VerifierDisabled.VerifierAddress = event.Data.Get("verifier").(common.Address).String()
        ovd.Block = event.BlockNumber
    }

    if err := c.WS.SendMessage("onVerifierDisable", ovd); err != nil {
        dot.Logger().Errorln("", zap.NamedError("onVerifierDisable"+server.EventSendFailed, err))
    }

    return true
}

func (c *Callbacks) onArbitrationBegin(event event.Event) bool {
    var oab definition.Callbacks
    {
        oab.PublishId = event.Data.Get("publishId").(string)
        if err := c.WS.SendMessage("onProofFilesExtensions", oab.PublishId); err != nil {
            dot.Logger().Errorln("", zap.NamedError("onProofFilesExtensions"+server.EventSendFailed, err))
        }

        oab.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
        oab.EncryptedId.EncryptedId = event.Data.Get("metaDataIdEncArbitrator").([]byte)
        oab.Block = event.BlockNumber

        extensions := <- c.ExtChan
        var err error
        if err = c.getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
            dot.Logger().Errorln("", zap.NamedError("Node - onVC.callback: get and rename proof files failed. ", err))
        }
    }

    if err := c.WS.SendMessage("onArbitrationBegin", oab); err != nil {
        dot.Logger().Errorln("", zap.NamedError("onArbitrationBegin"+server.EventSendFailed, err))
    }

    return true
}

func (c *Callbacks) onArbitrationResult(event event.Event) bool {
    var oar definition.Callbacks
    {
        oar.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
        oar.Arbitrate.ArbitrateResult = event.Data.Get("judge").(bool)
        oar.Block = event.BlockNumber
    }

    if err := c.WS.SendMessage("onArbitrationResult", oar); err != nil {
        dot.Logger().Errorln("", zap.NamedError("onArbitrationResult"+server.EventSendFailed, err))
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
