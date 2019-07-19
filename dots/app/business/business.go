package business

import (
    "encoding/json"
    "github.com/btcsuite/btcutil/base58"
    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/binary"
    scry2 "github.com/scryinfo/dp/dots/binary/scry"
    "github.com/scryinfo/dp/dots/app"
    "github.com/scryinfo/dp/dots/app/business/definition"
    ci "github.com/scryinfo/dp/dots/app/server"
    "github.com/scryinfo/dp/dots/eth/event"
    "github.com/scryinfo/dp/dots/eth/transaction"
    "go.uber.org/zap"
    "io/ioutil"
    "math/big"
    "os"
    "strconv"
    "time"
)

type Business struct {
    PresetMsgNames    []string
    PresetMsgHandlers []ci.PresetFunc
    CurUser           scry2.Client
    EventNames        []string
    EventHandler      []event.Callback
    deployer          *definition.AccInfo
    extChan           chan []string
    config            businessConfig
    Bin               *binary.Binary `dot:"binary"`
    WS                *app.WSServer  `dot:""`
}

type businessConfig struct {
    IPFSOutDir     string `json:"IPFSOutDir"`
    AccsBackupFile string `json:"accsBackupFile"`
}

const (
    BusTypeId = "64a3ff50-50de-447c-b0b9-401fff8c4fa4"
    BusLiveId = "64a3ff50-50de-447c-b0b9-401fff8c4fa4"

    verifierNum            = 2
    verifierBonus          = 300
    registerAsVerifierCost = 10000

    arbitratorNum   = 1
    arbitratorBonus = 500

    sep = "|"
)

func (b *Business) Create(l dot.Line) error {
    b.extChan = make(chan []string, 4)
    
    b.deployer = &definition.AccInfo{
        Account:  "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8",
        Password: "111111",
    }

    b.PresetMsgNames = []string{
        "login.verify",
        "create.new.account",
        "block.set",
        "logout",
        "publish",
        "buy",
        "extensions",
        "purchase",
        "reEncrypt",
        "cancel",
        "decrypt",
        "confirm",
        "register",
        "verify",
        "credit",
        "arbitrate",
        "get.eth.balance",
        "get.token.balance",
        "acc.backup",
        "acc.restore",
    }

    b.PresetMsgHandlers = []ci.PresetFunc{
        b.LoginVerify,
        b.CreateNewAccount,
        b.BlockSet,
        b.Logout,
        b.Publish,
        b.Buy,
        b.Extensions,
        b.Purchase,
        b.ReEncrypt,
        b.Cancel,
        b.Decrypt,
        b.Confirm,
        b.Register,
        b.Verify,
        b.Credit,
        b.Arbitrate,
        b.GetEthBalance,
        b.GetTokenBalance,
        b.Backup,
        b.Restore,
    }

    b.EventNames = []string{
        "Approval",
        "DataPublish",
        "VerifiersChosen",
        "TransactionCreate",
        "Buy",
        "ReadyForDownload",
        "TransactionClose",
        "RegisterVerifier",
        "Vote",
        "VerifierDisable",
        "ArbitrationBegin",
        "ArbitrationResult",
    }
    
    b.EventHandler = []event.Callback{
        b.onApprove, 
        b.onPublish, 
        b.onVerifiersChosen, 
        b.onTransactionCreate, 
        b.onPurchase, 
        b.onReadyForDownload,
        b.onClose, 
        b.onRegisterAsVerifier, 
        b.onVote, 
        b.onVerifierDisable, 
        b.onArbitrationBegin, 
        b.onArbitrationResult,
    }

    return nil
}

func (b *Business) Start(ignore bool) error {
    if err := b.WS.PresetMsgHandleFuncs(b.PresetMsgNames, b.PresetMsgHandlers); err != nil {
        return err
    }

    if err := b.WS.ListenAndServe(); err != nil {
        dot.Logger().Errorln("Start http web server failed. ", zap.NamedError("error", err))
        return errors.New("Start http web server failed. ")
    }

    return nil
}

//construct dot
func newBusDot(conf interface{}) (dot.Dot, error) {
    var err error
    var bs []byte
    if bt, ok := conf.([]byte); ok {
        bs = bt
    } else {
        return nil, dot.SError.Parameter
    }

    dConf := &businessConfig{}
    err = dot.UnMarshalConfig(bs, dConf)
    if err != nil {
        return nil, err
    }

    d := &Business{config: *dConf}

    return d, err
}

func BusTypeLive() []*dot.TypeLives {
    t := []*dot.TypeLives{
        {
            Meta: dot.Metadata{
                TypeId: BusTypeId,
                NewDoter: func(conf interface{}) (dot.Dot, error) {
                    return newBusDot(conf)
                },
            },
            Lives: []dot.Live{
                {
                    LiveId:    BusLiveId,
                    RelyLives: map[string]dot.LiveId{"binary": binary.BinLiveId},
                },
            },
        },
        app.WebSocketTypeLive(),
    }

    t = append(t, binary.BinTypeLive()...)

    return t
}

func (b *Business) LoginVerify(mi *ci.MessageIn) (payload interface{}, err error) {
    var ai definition.AccInfo
    if err = json.Unmarshal(mi.Payload, &ai); err != nil {
        return
    }

    var client scry2.Client
    if client = scry2.NewScryClient(ai.Account, b.Bin.ChainWrapper()); client == nil {
        err = errors.New("Call NewScryClient failed. ")
        return
    }

    var login bool
    if login, err = client.Authenticate(ai.Password); err != nil {
        err = errors.Wrap(err, "Authenticate user information failed. ")
        return
    }
    if login {
        b.CurUser = client
    } else {
        err = errors.New("Login verify failed. ")
        return
    }

    payload = true

    return
}

func (b *Business) CreateNewAccount(mi *ci.MessageIn) (payload interface{}, err error) {
    var pwd definition.AccInfo
    if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
        return
    }

    var client scry2.Client
    if client, err = scry2.CreateScryClient(pwd.Password, b.Bin.ChainWrapper()); err != nil {
        err = errors.Wrap(err, "Create new user failed. ")
        return
    }

    b.CurUser = client

    payload = client.Account().Addr

    return
}

func (b *Business) BlockSet(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }
    
    var sid definition.SDKInitData
    if err = json.Unmarshal(mi.Payload, &sid); err != nil {
        return
    }

    if len(b.EventNames) != len(b.EventHandler) {
        err = errors.New("Quantities of name and function are not matched. ")
        return
    }

    for i := range b.EventNames {
        if err = b.CurUser.SubscribeEvent(b.EventNames[i], b.EventHandler[i]); err != nil {
            err = errors.Wrap(err, "Subscribe event failed. ")
            return
        }
    }
    
    b.Bin.Listener.SetFromBlock(uint64(sid.FromBlock))
    
    // when an user login success, he will get 10,000,000 eth and tokens for test. in 'block.set' case.
    if err = b.CurUser.TransferEthFrom(common.HexToAddress(b.deployer.Account),
        b.deployer.Password,
        big.NewInt(10000000), 
        b.Bin.ChainWrapper().Conn(), 
    ); err != nil {
        err = errors.Wrap(err, "Transfer eth from deployer failed. ")
        return 
    }

    txParam := transaction.TxParams{
        From:     common.HexToAddress(b.deployer.Account),
        Password: b.deployer.Password,
        Value:    big.NewInt(0),
        Pending:  false,
    }
    if err = b.Bin.ChainWrapper().TransferTokens(&txParam, common.HexToAddress(b.CurUser.Account().Addr), big.NewInt(10000000)); err != nil {
        err = errors.Wrap(err, "Transfer token from deployer failed. ")
        return 
    }
    
    payload = true

    return
}

func (b *Business) Logout(_ *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return 
    }

    for i := range b.EventNames {
        if err = b.CurUser.UnSubscribeEvent(b.EventNames[i]); err != nil {
            err = errors.Wrap(err, "Unsubscribe failed, event:  "+b.EventNames[i]+" . ")
            return 
        }
    }
    
    payload = true

    return
}

func (b *Business) Publish(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var pd definition.PublishData
    if err = json.Unmarshal(mi.Payload, &pd); err != nil {
        return
    }

    if payload, err = b.Bin.ChainWrapper().Publish(
        b.makeTxParams(pd.Password),
        big.NewInt(int64(pd.Price)),
        []byte(pd.IDs.MetaDataID),
        pd.IDs.ProofDataIDs,
        int32(len(pd.IDs.ProofDataIDs)),
        pd.IDs.DetailsID,
        pd.SupportVerify,
    ); err != nil {
            return
    }

    return
}

func (b *Business) Buy(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var bd definition.BuyData
    if err = json.Unmarshal(mi.Payload, &bd); err != nil {
        return
    }

    fee := int64(bd.SelectedData.Price)
    if bd.StartVerify {
        fee += int64(verifierNum*verifierBonus) + int64(arbitratorNum*arbitratorBonus)
    }

    if err = b.Bin.ChainWrapper().ApproveTransfer(b.makeTxParams(bd.Password),
        common.HexToAddress(b.Bin.Config().ProtocolContractAddr),
        big.NewInt(fee),
    ); err != nil {
        err = errors.Wrap(err, "Contract transfer token from buyer failed. ")
        return
    }

    if err = b.Bin.ChainWrapper().PrepareToBuy(b.makeTxParams(bd.Password), bd.SelectedData.PublishID, bd.StartVerify); err != nil {
        err = errors.Wrap(err, "Transaction create failed. ")
        return
    }

    payload = true

    return
}

func (b *Business) Extensions(mi *ci.MessageIn) (payload interface{}, err error) {
    var p definition.Prepared
    if err = json.Unmarshal(mi.Payload, &p); err != nil {
        return
    }
    b.extChan <- p.Extensions
    payload = true

    return
}

func (b *Business) Purchase(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var pd definition.PurchaseData
    if err = json.Unmarshal(mi.Payload, &pd); err != nil {
        return
    }

    tID, ok := new(big.Int).SetString(pd.SelectedTx.TransactionID, 10)
    if !ok {
        err = errors.New("Set to *big.Int failed. ")
        return
    }

    if err = b.Bin.ChainWrapper().BuyData(b.makeTxParams(pd.Password), tID); err != nil {
        err = errors.Wrap(err, "Buy data failed. ")
        return
    }

    payload = true

    return
}

func (b *Business) ReEncrypt(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var re definition.ReEncryptData
    if err = json.Unmarshal(mi.Payload, &re); err != nil {
        return
    }

    txParam := b.makeTxParams(re.Password)
    tID, ok := new(big.Int).SetString(re.SelectedTx.TransactionID, 10)
    if !ok {
        err = errors.New("Set to *big.Int failed. ")
        return
    }

    if err = b.Bin.ChainWrapper().ReEncryptMetaDataId(txParam, tID, re.SelectedTx.MetaDataIDEncWithSeller); err != nil {
        err = errors.Wrap(err, "Submit encrypted ID with buyer failed. ")
        return
    }

    payload = true

    return
}

func (b *Business) Cancel(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var pd definition.PurchaseData
    if err = json.Unmarshal(mi.Payload, &pd); err != nil {
        return
    }

    tID, ok := new(big.Int).SetString(pd.SelectedTx.TransactionID, 10)
    if !ok {
        err = errors.New("Set to *big.Int failed. ")
        return
    }

    if err = b.Bin.ChainWrapper().CancelTransaction(b.makeTxParams(pd.Password), tID); err != nil {
        err = errors.Wrap(err, "Cancel transaction failed. ")
        return
    }

    payload = true

    return
}

func (b *Business) Decrypt(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var dd definition.DecryptData
    if err = json.Unmarshal(mi.Payload, &dd); err != nil {
        return
    }

    var oldFileName string
    {
        var metaDataIDByte []byte
        if metaDataIDByte, err = b.Bin.Account.Decrypt(dd.SelectedTx.MetaDataIDEncrypt, dd.SelectedTx.User, dd.Password); err != nil {
            return "", errors.Wrap(err, "Decrypt encrypted meta data ID failed. ")
        }
        outDir := b.config.IPFSOutDir
        if err = b.Bin.Storage.Get(string(metaDataIDByte), outDir); err != nil {
            return "", errors.Wrap(err, "Get meta data from IPFS failed. ")
        }
        oldFileName = outDir + "/" + string(metaDataIDByte)
    }

    newFileName := oldFileName + dd.SelectedTx.MetaDataExtension
    if err = os.Rename(oldFileName, newFileName); err != nil {
        return "", errors.Wrap(err, "Add extension to meta data failed. ")
    }

    payload = newFileName

    return
}

func (b *Business) Confirm(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var cd definition.ConfirmData
    if err = json.Unmarshal(mi.Payload, &cd); err != nil {
        return
    }

    tID, ok := new(big.Int).SetString(cd.SelectedTx.TransactionID, 10)
    if !ok {
        err = errors.New("Set to *big.Int failed. ")
        return
    }

    if err = b.Bin.ChainWrapper().ConfirmDataTruth(b.makeTxParams(cd.Password), tID, cd.Truth); err != nil {
        err = errors.Wrap(err, "Confirm data truth failed. ")
        return
    }

    payload = true

    return
}

func (b *Business) Register(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var rvd definition.RegisterVerifierData
    if err = json.Unmarshal(mi.Payload, &rvd); err != nil {
        return
    }

    if err = b.Bin.ChainWrapper().ApproveTransfer(b.makeTxParams(rvd.Password),
        common.HexToAddress(b.Bin.Config().ProtocolContractAddr),
        big.NewInt(registerAsVerifierCost),
    ); err != nil {
        err = errors.Wrap(err, "Contract transfer token from register failed. ")
        return
    }

    if err = b.Bin.ChainWrapper().RegisterAsVerifier(b.makeTxParams(rvd.Password)); err != nil {
        err = errors.Wrap(err, "Register as verifier failed. ")
        return
    }

    payload = true

    return
}

func (b *Business) Verify(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var vd definition.VerifyData
    if err = json.Unmarshal(mi.Payload, &vd); err != nil {
        return
    }

    tID, ok := new(big.Int).SetString(vd.TransactionID, 10)
    if !ok {
        err = errors.New("Set to *big.Int failed. ")
        return
    }

    if err = b.Bin.ChainWrapper().Vote(b.makeTxParams(vd.Password), tID, vd.Verify.Suggestion, vd.Verify.Comment); err != nil {
        err = errors.Wrap(err, "Vote failed. ")
        return
    }

    payload = true

    return
}

func (b *Business) Credit(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var cd definition.CreditData
    if err = json.Unmarshal(mi.Payload, &cd); err != nil {
        return
    }
    tID, ok := new(big.Int).SetString(cd.SelectedTx.TransactionID, 10)
    if !ok {
        err = errors.New("Set to *big.Int failed. ")
        return
    }

    txParam := b.makeTxParams(cd.Password)

    if cd.Credit.Verifier1Revert {
        credit := uint8(cd.Credit.Verifier1Credit)
        if err = b.Bin.ChainWrapper().CreditsToVerifier(txParam, tID, 0, credit); err != nil {
            err = errors.Wrap(err, "Credit verifier1 failed. ")
            return
        }
    }
    if cd.Credit.Verifier2Revert {
        credit := uint8(cd.Credit.Verifier2Credit)
        if err = b.Bin.ChainWrapper().CreditsToVerifier(txParam, tID, 1, credit); err != nil {
            err = errors.Wrap(err, "Credit verifier2 failed. ")
            return
        }
    }

    payload = true

    return
}

func (b *Business) Arbitrate(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var ad definition.ArbitrateData
    if err = json.Unmarshal(mi.Payload, &ad); err != nil {
        return
    }
    tID, ok := new(big.Int).SetString(ad.SelectedTx.TransactionId, 10)
    if !ok {
        err = errors.New("Set to *big.Int failed. ")
        return
    }

    if err = b.Bin.ChainWrapper().Arbitrate(b.makeTxParams(ad.Password), tID, ad.ArbitrateResult); err != nil {
        err = errors.Wrap(err, "Arbitrate failed. ")
        return
    }

    payload = true

    return
}

func (b *Business) GetEthBalance(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var pwd definition.AccInfo
    if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
        return
    }

    var balance *big.Int
    if balance, err = b.CurUser.GetEth(common.HexToAddress(b.CurUser.Account().Addr), b.Bin.ChainWrapper().Conn());err != nil {
        err = errors.Wrap(err, "Get eth balance failed. ")
        return
    }

    payload = balance.String() + sep + time.Now().String()

    return
}

func (b *Business) GetTokenBalance(mi *ci.MessageIn) (payload interface{}, err error) {
    if b.CurUser == nil {
        err = errors.New("Current user is nil. ")
        return
    }

    var pwd definition.AccInfo
    if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
        return
    }

    var balance *big.Int
    if balance, err = b.Bin.ChainWrapper().GetTokenBalance(b.makeTxParams(pwd.Password), common.HexToAddress(b.CurUser.Account().Addr));err != nil {
        err = errors.Wrap(err, "Get token balance failed. ")
        return
    }

    payload = balance.String() + sep + time.Now().String()

    return
}

func (b *Business) Backup(mi *ci.MessageIn) (interface{}, error) {
    return true, ioutil.WriteFile(b.config.AccsBackupFile, mi.Payload, 0777)
}

func (b *Business) Restore(_ *ci.MessageIn) (interface{}, error) {
    return ioutil.ReadFile(b.config.AccsBackupFile)
}

func (b *Business) makeTxParams(password string) *transaction.TxParams {
    return &transaction.TxParams{
        From:     common.HexToAddress(b.CurUser.Account().Addr),
        Password: password,
        Value:    big.NewInt(0),
        Pending:  false,
    }
}

func (b *Business) onPublish(event event.Event) bool {
   var op definition.OnPublish
   {
       var err error
       if op, err = b.getPubDataDetails(event.Data.Get("despDataId").(string)); err != nil {
           dot.Logger().Errorln("", zap.NamedError("onPublish: get publish data details failed. ", err))
       }
       op.Block = event.BlockNumber
       op.Price = event.Data.Get("price").(*big.Int).String()
       op.PublishID = event.Data.Get("publishId").(string)
       op.SupportVerify = event.Data.Get("supportVerify").(bool)
   }

   if err := b.WS.SendMessage("onPublish", op); err != nil {
       dot.Logger().Errorln("", zap.NamedError("onPublish"+ci.EventSendFailed, err))
   }

   return true
}

func (b *Business) getPubDataDetails(ipfsID string) (detailsData definition.OnPublish, err error) {
   defer func() {
       if er := recover(); er != nil {
           dot.Logger().Errorln("", zap.Any("onPublish.callback: get publish data details failed. ", er))
       }
   }()

   var fileName string
   {
       outDir := b.config.IPFSOutDir
       if err = b.Bin.Storage.Get(ipfsID, outDir); err != nil {
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

func (b *Business) onApprove(_ event.Event) bool {
   return true
}

func (b *Business) onVerifiersChosen(event event.Event) bool {
   var ovc definition.OnVerifiersChosen
   {
       ovc.PublishID = event.Data.Get("publishId").(string)
       if err := b.WS.SendMessage("onProofFilesExtensions", ovc.PublishID); err != nil {
           dot.Logger().Errorln("", zap.NamedError("onProofFilesExtensions"+ci.EventSendFailed, err))
       }

       ovc.Block = event.BlockNumber
       ovc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
       ovc.TxState = setTxState(event.Data.Get("state").(uint8))

       extensions := <- b.extChan
       var err error
       if ovc.ProofFileNames, err = b.getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
           dot.Logger().Errorln("", zap.NamedError("Node - onVC.callback: get and rename proof files failed. ", err))
       }
   }

   if err := b.WS.SendMessage("onVerifiersChosen", ovc); err != nil {
       dot.Logger().Errorln("", zap.NamedError("onVerifiersChosen"+ci.EventSendFailed, err))
   }

   return true
}

func (b *Business) onTransactionCreate(event event.Event) bool {
   var otc definition.OnTransactionCreate
   {
       otc.PublishID = event.Data.Get("publishId").(string)
       if err := b.WS.SendMessage("onProofFilesExtensions", otc.PublishID); err != nil {
           dot.Logger().Errorln("", zap.NamedError("onProofFilesExtensions"+ci.EventSendFailed, err))
       }

       otc.Block = event.BlockNumber
       otc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
       otc.Buyer = event.Data.Get("users").([]common.Address)[0].String()
       otc.StartVerify = event.Data.Get("needVerify").(bool)
       otc.TxState = setTxState(event.Data.Get("state").(uint8))

       extensions := <- b.extChan
       var err error
       if otc.ProofFileNames, err = b.getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
           dot.Logger().Errorln("", zap.NamedError("Node - onTC.callback: get and rename proof files failed. ", err))
       }
   }

   if err := b.WS.SendMessage("onTransactionCreate", otc); err != nil {
       dot.Logger().Errorln("", zap.NamedError("onTransactionCreate"+ci.EventSendFailed, err))
   }

   return true
}

func (b *Business) getAndRenameProofFiles(ipfsIDs [][32]byte, extensions []string) ([]string, error) {
   if len(ipfsIDs) != len(extensions) {
       return nil, errors.New("Quantity of IPFS IDs or extensions is wrong. ")
   }

   defer func() {
       if er := recover(); er != nil {
           dot.Logger().Errorln("", zap.Any("in callback: get and rename proof files failed. ", er))
       }
   }()

   var proofs = make([]string, len(ipfsIDs))

   outDir := b.config.IPFSOutDir
   for i := 0; i < len(ipfsIDs); i++ {
       ipfsID := ipfsBytes32ToHash(ipfsIDs[i])
       if err := b.Bin.Storage.Get(ipfsID, outDir); err != nil {
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

func (b *Business) onPurchase(event event.Event) bool {
   var op definition.OnPurchase
   {
       op.Block = event.BlockNumber
       op.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
       op.MetaDataIdEncWithSeller = event.Data.Get("metaDataIdEncSeller").([]byte)
       op.PublishID = event.Data.Get("publishId").(string)
       op.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
       op.TxState = setTxState(event.Data.Get("state").(uint8))
   }

   if err := b.WS.SendMessage("onPurchase", op); err != nil {
       dot.Logger().Errorln("", zap.NamedError("onPurchase"+ci.EventSendFailed, err))
   }

   return true
}

func (b *Business) onReadyForDownload(event event.Event) bool {
   var orfd definition.OnReadyForDownload
   {
       orfd.Block = event.BlockNumber
       orfd.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
       orfd.MetaDataIdEncWithBuyer = event.Data.Get("metaDataIdEncBuyer").([]byte)
       orfd.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
       orfd.TxState = setTxState(event.Data.Get("state").(uint8))
   }

   if err := b.WS.SendMessage("onReadyForDownload", orfd); err != nil {
       dot.Logger().Errorln("", zap.NamedError("onReadyForDownload"+ci.EventSendFailed, err))
   }

   return true
}

func (b *Business) onClose(event event.Event) bool {
   var oc definition.OnClose
   {
       oc.Block = event.BlockNumber
       oc.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
       oc.UserIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
       oc.TxState = setTxState(event.Data.Get("state").(uint8))
   }

   if err := b.WS.SendMessage("onClose", oc); err != nil {
       dot.Logger().Errorln("", zap.NamedError("onClose"+ci.EventSendFailed, err))
   }

   return true
}

func (b *Business) onRegisterAsVerifier(event event.Event) bool {
   var orav definition.OnRegisterAsVerifier
   {
       orav.Block = event.BlockNumber
   }

   if err := b.WS.SendMessage("onRegisterVerifier", orav); err != nil {
       dot.Logger().Errorln("", zap.NamedError("onRegisterVerifier"+ci.EventSendFailed, err))
   }

   return true
}

func (b *Business) onVote(event event.Event) bool {
   var ov definition.OnVote
   {
       ov.Block = event.BlockNumber
       ov.VerifierIndex = strconv.Itoa(int(event.Data.Get("index").(uint8)))
       ov.TransactionID = event.Data.Get("transactionId").(*big.Int).String()
       ov.TxState = setTxState(event.Data.Get("state").(uint8))

       judge := event.Data.Get("judge").(bool)
       comment := event.Data.Get("comments").(string)
       ov.VerifierResponse = setJudge(judge) + ", " + comment
   }

   if err := b.WS.SendMessage("onVote", ov); err != nil {
       dot.Logger().Errorln("", zap.NamedError("onVote"+ci.EventSendFailed, err))
   }

   return true
}

func (b *Business) onVerifierDisable(event event.Event) bool {
   var ovd definition.OnVerifierDisable
   {
       ovd.Block = event.BlockNumber
   }

   if err := b.WS.SendMessage("onVerifierDisable", ovd); err != nil {
       dot.Logger().Errorln("", zap.NamedError("onVerifierDisable"+ci.EventSendFailed, err))
   }

   return true
}

func (b *Business) onArbitrationBegin(event event.Event) bool {
   var oab definition.OnArbitrationBegin
   {
       oab.PublishId = event.Data.Get("publishId").(string)
       if err := b.WS.SendMessage("onProofFilesExtensions", oab.PublishId); err != nil {
           dot.Logger().Errorln("", zap.NamedError("onProofFilesExtensions"+ci.EventSendFailed, err))
       }

       oab.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
       oab.MetaDataIdEncWithArbitrator = event.Data.Get("metaDataIdEncArbitrator").([]byte)
       oab.Block = event.BlockNumber

       extensions := <-b.extChan
       var err error
       if oab.ProofFileNames, err = b.getAndRenameProofFiles(event.Data.Get("proofIds").([][32]uint8), extensions); err != nil {
           dot.Logger().Errorln("", zap.NamedError("Node - onVC.callback: get and rename proof files failed. ", err))
       }
   }

   if err := b.WS.SendMessage("onArbitrationBegin", oab); err != nil {
       dot.Logger().Errorln("", zap.NamedError("onArbitrationBegin"+ci.EventSendFailed, err))
   }

   return true
}

func (b *Business) onArbitrationResult(event event.Event) bool {
   var oar definition.OnArbitrationResult
   {
       oar.TransactionId = event.Data.Get("transactionId").(*big.Int).String()
       oar.ArbitrateResult = setArbitrateResult(event.Data.Get("judge").(bool))
       oar.User = strconv.Itoa(int(event.Data.Get("identify").(uint8)))
       oar.Block = event.BlockNumber
   }

   if err := b.WS.SendMessage("onArbitrationResult", oar); err != nil {
       dot.Logger().Errorln("", zap.NamedError("onArbitrationResult"+ci.EventSendFailed, err))
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
