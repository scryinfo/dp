package preset

import (
    "encoding/json"
    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/app/business/definition"
    "github.com/scryinfo/dp/dots/app/business/preset/chain_event"
    "github.com/scryinfo/dp/dots/app/server"
    "github.com/scryinfo/dp/dots/binary"
    scry2 "github.com/scryinfo/dp/dots/binary/scry"
    "github.com/scryinfo/dp/dots/eth/transaction"
    "io/ioutil"
    "math/big"
    "os"
    "time"
)

type Preset struct {
    PresetMsgNames    []string
    PresetMsgHandlers []server.PresetFunc
    CurUser           scry2.Client
    Deployer          *definition.AccInfo
    config            presetConfig
    Bin               *binary.Binary `dot:""`
    CBs               *cec.Callbacks `dot:""`
}

type presetConfig struct {
    MetaDataOutDir string `json:"metaDataOutDir"`
    AccsBackupFile string `json:"accsInfoBackupFile"`
}

const (
    PreTypeId = "13d73d73-da19-4d39-9dca-3018fbf0ec30"

    verifierNum            = 2
    verifierBonus          = 300
    registerAsVerifierCost = 10000

    arbitratorNum   = 1
    arbitratorBonus = 500

    sep = "|"
)

func (p *Preset) Create(l dot.Line) error {
    p.Deployer = &definition.AccInfo{
        Account:  "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8",
        Password: "111111",
    }

    p.PresetMsgNames = []string{
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

    p.PresetMsgHandlers = []server.PresetFunc{
        p.LoginVerify,
        p.CreateNewAccount,
        p.BlockSet,
        p.Logout,
        p.Publish,
        p.Buy,
        p.Extensions,
        p.Purchase,
        p.ReEncrypt,
        p.Cancel,
        p.Decrypt,
        p.Confirm,
        p.Register,
        p.Verify,
        p.Credit,
        p.Arbitrate,
        p.GetEthBalance,
        p.GetTokenBalance,
        p.Backup,
        p.Restore,
    }
    return nil
}

func newPresetDot(conf interface{}) (dot.Dot, error) {
    var err error
    var bs []byte
    if bt, ok := conf.([]byte); ok {
        bs = bt
    } else {
        return nil, dot.SError.Parameter
    }

    dConf := &presetConfig{}
    err = dot.UnMarshalConfig(bs, dConf)
    if err != nil {
        return nil, err
    }

    d := &Preset{config: *dConf}

    return d, err
}

func PreTypeLive() []*dot.TypeLives {
    t := []*dot.TypeLives{
        {
            Meta: dot.Metadata{
                TypeId: PreTypeId,
                NewDoter: func(conf interface{}) (dot.Dot, error) {
                    return newPresetDot(conf)
                },
            },
        },
    }
    
    t = append(t, binary.BinTypeLive()...)
    t = append(t, cec.CBsTypeLive()...)
    
    return t
}

func (p *Preset) LoginVerify(mi *server.MessageIn) (payload interface{}, err error) {
   var ai definition.AccInfo
   if err = json.Unmarshal(mi.Payload, &ai); err != nil {
       return
   }

   var client scry2.Client
   if client = scry2.NewScryClient(ai.Account, p.Bin.ChainWrapper()); client == nil {
       err = errors.New("Call NewScryClient failed. ")
       return
   }

   var login bool
   if login, err = client.Authenticate(ai.Password); err != nil {
       err = errors.Wrap(err, "Authenticate user information failed. ")
       return
   }
   if login {
       p.CurUser = client
   } else {
       err = errors.New("Login verify failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) CreateNewAccount(mi *server.MessageIn) (payload interface{}, err error) {
   var pwd definition.AccInfo
   if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
       return
   }

   var client scry2.Client
   if client, err = scry2.CreateScryClient(pwd.Password, p.Bin.ChainWrapper()); err != nil {
       err = errors.Wrap(err, "Create new user failed. ")
       return
   }

   p.CurUser = client

   payload = client.Account().Addr

   return
}

func (p *Preset) BlockSet(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var sid definition.SDKInitData
   if err = json.Unmarshal(mi.Payload, &sid); err != nil {
       return
   }

   if len(p.CBs.EventNames) != len(p.CBs.EventHandler) {
       err = errors.New("Quantities of name and function are not matched. ")
       return
   }

   for i := range p.CBs.EventNames {
       if err = p.CurUser.SubscribeEvent(p.CBs.EventNames[i], p.CBs.EventHandler[i]); err != nil {
           err = errors.Wrap(err, "Subscribe event failed. ")
           return
       }
   }

   p.Bin.Listener.SetFromBlock(uint64(sid.FromBlock))

   // when an user login success, he will get 10,000,000 eth and tokens for test. in 'block.set' case.
   if err = p.CurUser.TransferEthFrom(common.HexToAddress(p.Deployer.Account),
       p.Deployer.Password,
       big.NewInt(10000000),
       p.Bin.ChainWrapper().Conn(),
   ); err != nil {
       err = errors.Wrap(err, "Transfer eth from Deployer failed. ")
       return
   }

   txParam := transaction.TxParams{
       From:     common.HexToAddress(p.Deployer.Account),
       Password: p.Deployer.Password,
       Value:    big.NewInt(0),
       Pending:  false,
   }
   if err = p.Bin.ChainWrapper().TransferTokens(&txParam, common.HexToAddress(p.CurUser.Account().Addr), big.NewInt(10000000)); err != nil {
       err = errors.Wrap(err, "Transfer token from Deployer failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) Logout(_ *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   for i := range p.CBs.EventNames {
       if err = p.CurUser.UnSubscribeEvent(p.CBs.EventNames[i]); err != nil {
           err = errors.Wrap(err, "Unsubscribe failed, event:  "+p.CBs.EventNames[i]+" . ")
           return
       }
   }

   p.CurUser = nil

   payload = true

   return
}

func (p *Preset) Publish(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var pd definition.PublishData
   if err = json.Unmarshal(mi.Payload, &pd); err != nil {
       return
   }

   if payload, err = p.Bin.ChainWrapper().Publish(
       p.makeTxParams(pd.Password),
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

func (p *Preset) Buy(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
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

   if err = p.Bin.ChainWrapper().ApproveTransfer(p.makeTxParams(bd.Password),
       common.HexToAddress(p.Bin.Config().ProtocolContractAddr),
       big.NewInt(fee),
   ); err != nil {
       err = errors.Wrap(err, "Contract transfer token from buyer failed. ")
       return
   }

   if err = p.Bin.ChainWrapper().PrepareToBuy(p.makeTxParams(bd.Password), bd.SelectedData.PublishID, bd.StartVerify); err != nil {
       err = errors.Wrap(err, "Transaction create failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) Extensions(mi *server.MessageIn) (payload interface{}, err error) {
   var ppd definition.Prepared
   if err = json.Unmarshal(mi.Payload, &p); err != nil {
       return
   }
   p.CBs.ExtChan <- ppd.Extensions
   payload = true

   return
}

func (p *Preset) Purchase(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
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

   if err = p.Bin.ChainWrapper().BuyData(p.makeTxParams(pd.Password), tID); err != nil {
       err = errors.Wrap(err, "Buy data failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) ReEncrypt(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var re definition.ReEncryptData
   if err = json.Unmarshal(mi.Payload, &re); err != nil {
       return
   }

   txParam := p.makeTxParams(re.Password)
   tID, ok := new(big.Int).SetString(re.SelectedTx.TransactionID, 10)
   if !ok {
       err = errors.New("Set to *big.Int failed. ")
       return
   }

   if err = p.Bin.ChainWrapper().ReEncryptMetaDataId(txParam, tID, re.SelectedTx.MetaDataIDEncWithSeller); err != nil {
       err = errors.Wrap(err, "Submit encrypted ID with buyer failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) Cancel(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
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

   if err = p.Bin.ChainWrapper().CancelTransaction(p.makeTxParams(pd.Password), tID); err != nil {
       err = errors.Wrap(err, "Cancel transaction failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) Decrypt(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
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
       if metaDataIDByte, err = p.Bin.Account.Decrypt(dd.SelectedTx.MetaDataIDEncrypt, dd.SelectedTx.User, dd.Password); err != nil {
           return "", errors.Wrap(err, "Decrypt encrypted meta data ID failed. ")
       }
       outDir := p.config.MetaDataOutDir
       if err = p.Bin.Storage.Get(string(metaDataIDByte), outDir); err != nil {
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

func (p *Preset) Confirm(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
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

   if err = p.Bin.ChainWrapper().ConfirmDataTruth(p.makeTxParams(cd.Password), tID, cd.Truth); err != nil {
       err = errors.Wrap(err, "Confirm data truth failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) Register(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var rvd definition.RegisterVerifierData
   if err = json.Unmarshal(mi.Payload, &rvd); err != nil {
       return
   }

   if err = p.Bin.ChainWrapper().ApproveTransfer(p.makeTxParams(rvd.Password),
       common.HexToAddress(p.Bin.Config().ProtocolContractAddr),
       big.NewInt(registerAsVerifierCost),
   ); err != nil {
       err = errors.Wrap(err, "Contract transfer token from register failed. ")
       return
   }

   if err = p.Bin.ChainWrapper().RegisterAsVerifier(p.makeTxParams(rvd.Password)); err != nil {
       err = errors.Wrap(err, "Register as verifier failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) Verify(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
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

   if err = p.Bin.ChainWrapper().Vote(p.makeTxParams(vd.Password), tID, vd.Verify.Suggestion, vd.Verify.Comment); err != nil {
       err = errors.Wrap(err, "Vote failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) Credit(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
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

   txParam := p.makeTxParams(cd.Password)

   if cd.Credit.Verifier1Revert {
       credit := uint8(cd.Credit.Verifier1Credit)
       if err = p.Bin.ChainWrapper().CreditsToVerifier(txParam, tID, 0, credit); err != nil {
           err = errors.Wrap(err, "Credit verifier1 failed. ")
           return
       }
   }
   if cd.Credit.Verifier2Revert {
       credit := uint8(cd.Credit.Verifier2Credit)
       if err = p.Bin.ChainWrapper().CreditsToVerifier(txParam, tID, 1, credit); err != nil {
           err = errors.Wrap(err, "Credit verifier2 failed. ")
           return
       }
   }

   payload = true

   return
}

func (p *Preset) Arbitrate(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
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

   if err = p.Bin.ChainWrapper().Arbitrate(p.makeTxParams(ad.Password), tID, ad.ArbitrateResult); err != nil {
       err = errors.Wrap(err, "Arbitrate failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) GetEthBalance(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var pwd definition.AccInfo
   if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
       return
   }

   var balance *big.Int
   if balance, err = p.CurUser.GetEth(common.HexToAddress(p.CurUser.Account().Addr), p.Bin.ChainWrapper().Conn());err != nil {
       err = errors.Wrap(err, "Get eth balance failed. ")
       return
   }

   payload = balance.String() + sep + time.Now().String()

   return
}

func (p *Preset) GetTokenBalance(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var pwd definition.AccInfo
   if err = json.Unmarshal(mi.Payload, &pwd); err != nil {
       return
   }

   var balance *big.Int
   if balance, err = p.Bin.ChainWrapper().GetTokenBalance(p.makeTxParams(pwd.Password), common.HexToAddress(p.CurUser.Account().Addr));err != nil {
       err = errors.Wrap(err, "Get token balance failed. ")
       return
   }

   payload = balance.String() + sep + time.Now().String()

   return
}

func (p *Preset) Backup(mi *server.MessageIn) (interface{}, error) {
   return true, ioutil.WriteFile(p.config.AccsBackupFile, mi.Payload, 0777)
}

func (p *Preset) Restore(_ *server.MessageIn) (interface{}, error) {
   return ioutil.ReadFile(p.config.AccsBackupFile)
}

func (p *Preset) makeTxParams(password string) *transaction.TxParams {
   return &transaction.TxParams{
       From:     common.HexToAddress(p.CurUser.Account().Addr),
       Password: password,
       Value:    big.NewInt(0),
       Pending:  false,
   }
}
