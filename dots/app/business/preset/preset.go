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
    Deployer          *definition.Preset
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
    p.Deployer = &definition.Preset{
        Address:  "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8",
        Password: "111111",
    }

    p.PresetMsgNames = []string{
        "loginVerify",
        "createNewAccount",
        "blockSet",
        "logout",
        "publish",
        "advancePurchase",
        "extensions",
        "confirmPurchase",
        "reEncrypt",
        "cancelPurchase",
        "decrypt",
        "confirmData",
        "register",
        "vote",
        "gradeToVerifier",
        "arbitrate",
        "getEthBalance",
        "getTokenBalance",
        "accountsBackup",
        "accountsRestore",
    }

    p.PresetMsgHandlers = []server.PresetFunc{
        p.LoginVerify,
        p.CreateNewAccount,
        p.BlockSet,
        p.Logout,
        p.Publish,
        p.AdvancePurchase,
        p.Extensions,
        p.ConfirmPurchase,
        p.ReEncrypt,
        p.CancelPurchase,
        p.Decrypt,
        p.ConfirmData,
        p.Register,
        p.Vote,
        p.GradeToVerifier,
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
    
    t = append(t, binary.BinTypeLiveWithoutGrpc()...)
    t = append(t, cec.CBsTypeLive()...)
    
    return t
}

func (p *Preset) LoginVerify(mi *server.MessageIn) (payload interface{}, err error) {
   var lv definition.Preset
   if err = json.Unmarshal(mi.Payload, &lv); err != nil {
       return
   }

   var client scry2.Client
   if client = scry2.NewScryClient(lv.Address, p.Bin.ChainWrapper()); client == nil {
       err = errors.New("Call NewScryClient failed. ")
       return
   }

   var login bool
   if login, err = client.Authenticate(lv.Password); err != nil {
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
   var cna definition.Preset
   if err = json.Unmarshal(mi.Payload, &cna); err != nil {
       return
   }

   var client scry2.Client
   if client, err = scry2.CreateScryClient(cna.Password, p.Bin.ChainWrapper()); err != nil {
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

   var bs definition.Preset
   if err = json.Unmarshal(mi.Payload, &bs); err != nil {
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

   p.Bin.Listener.SetFromBlock(uint64(bs.FromBlock))

   // when an user login success, he will get 10,000,000 eth and tokens for test. in 'block.set' case.
   if err = p.CurUser.TransferEthFrom(common.HexToAddress(p.Deployer.Address),
       p.Deployer.Password,
       big.NewInt(10000000),
       p.Bin.ChainWrapper().Conn(),
   ); err != nil {
       err = errors.Wrap(err, "Transfer eth from Deployer failed. ")
       return
   }

   txParam := transaction.TxParams{
       From:     common.HexToAddress(p.Deployer.Address),
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

   var publish definition.Preset
   if err = json.Unmarshal(mi.Payload, &publish); err != nil {
       return
   }

   if payload, err = p.Bin.ChainWrapper().Publish(
       p.makeTxParams(publish.Password),
       big.NewInt(int64(publish.Price)),
       []byte(publish.Ids.MetaDataId),
       publish.Ids.ProofDataIds,
       int32(len(publish.Ids.ProofDataIds)),
       publish.Ids.DetailsId,
       publish.SupportVerify,
   ); err != nil {
       return
   }

   return
}

func (p *Preset) AdvancePurchase(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var ap definition.Preset
   if err = json.Unmarshal(mi.Payload, &ap); err != nil {
       return
   }

   fee := int64(ap.Price)
   if ap.StartVerify {
       fee += int64(verifierNum*verifierBonus) + int64(arbitratorNum*arbitratorBonus)
   }

   if err = p.Bin.ChainWrapper().ApproveTransfer(p.makeTxParams(ap.Password),
       common.HexToAddress(p.Bin.Config().ProtocolContractAddr),
       big.NewInt(fee),
   ); err != nil {
       err = errors.Wrap(err, "Contract transfer token from buyer failed. ")
       return
   }

   time.Sleep(5 * time.Second)

   if err = p.Bin.ChainWrapper().AdvancePurchase(p.makeTxParams(ap.Password), ap.PublishId, ap.StartVerify); err != nil {
       err = errors.Wrap(err, "Advance purchase failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) Extensions(mi *server.MessageIn) (payload interface{}, err error) {
   var ext definition.Preset
   if err = json.Unmarshal(mi.Payload, &ext); err != nil {
       return
   }

   p.CBs.ExtChan <- ext.Extensions.ProofDataExtensions
   payload = true

   return
}

func (p *Preset) ConfirmPurchase(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var cp definition.Preset
   if err = json.Unmarshal(mi.Payload, &cp); err != nil {
       return
   }

   tId, ok := new(big.Int).SetString(cp.TransactionId, 10)
   if !ok {
       err = errors.New("Set to *big.Int failed. ")
       return
   }

   if err = p.Bin.ChainWrapper().ConfirmPurchase(p.makeTxParams(cp.Password), tId); err != nil {
       err = errors.Wrap(err, "Confirm purchase failed. ")
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

   var re definition.Preset
   if err = json.Unmarshal(mi.Payload, &re); err != nil {
       return
   }

   txParam := p.makeTxParams(re.Password)
   tId, ok := new(big.Int).SetString(re.TransactionId, 10)
   if !ok {
       err = errors.New("Set to *big.Int failed. ")
       return
   }

   if err = p.Bin.ChainWrapper().ReEncrypt(txParam, tId, re.EncryptedId.EncryptedId); err != nil {
       err = errors.Wrap(err, "Re-encrypt failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) CancelPurchase(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var cp definition.Preset
   if err = json.Unmarshal(mi.Payload, &cp); err != nil {
       return
   }

   tId, ok := new(big.Int).SetString(cp.TransactionId, 10)
   if !ok {
       err = errors.New("Set to *big.Int failed. ")
       return
   }

   if err = p.Bin.ChainWrapper().CancelPurchase(p.makeTxParams(cp.Password), tId); err != nil {
       err = errors.Wrap(err, "Cancel purchase failed. ")
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

   var decrypt definition.Preset
   if err = json.Unmarshal(mi.Payload, &decrypt); err != nil {
       return
   }

   var oldFileName string
   {
       var metaDataIdByte []byte
       if metaDataIdByte, err = p.Bin.Account.Decrypt(decrypt.EncryptedId.EncryptedId, decrypt.Address, decrypt.Password); err != nil {
           return "", errors.Wrap(err, "Decrypt encrypted meta data Id failed. ")
       }
       outDir := p.config.MetaDataOutDir
       if err = p.Bin.Storage.Get(string(metaDataIdByte), outDir); err != nil {
           return "", errors.Wrap(err, "Get meta data from IPFS failed. ")
       }
       oldFileName = outDir + "/" + string(metaDataIdByte)
   }

   newFileName := oldFileName + decrypt.Extensions.MetaDataExtension
   if err = os.Rename(oldFileName, newFileName); err != nil {
       return "", errors.Wrap(err, "Add extension to meta data failed. ")
   }

   payload = newFileName

   return
}

func (p *Preset) ConfirmData(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var cd definition.Preset
   if err = json.Unmarshal(mi.Payload, &cd); err != nil {
       return
   }

   tId, ok := new(big.Int).SetString(cd.TransactionId, 10)
   if !ok {
       err = errors.New("Set to *big.Int failed. ")
       return
   }

   if err = p.Bin.ChainWrapper().ConfirmData(p.makeTxParams(cd.Password), tId, cd.Confirm.Truth); err != nil {
       err = errors.Wrap(err, "Confirm data failed. ")
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

   var register definition.Preset
   if err = json.Unmarshal(mi.Payload, &register); err != nil {
       return
   }

   if err = p.Bin.ChainWrapper().ApproveTransfer(p.makeTxParams(register.Password),
       common.HexToAddress(p.Bin.Config().ProtocolContractAddr),
       big.NewInt(registerAsVerifierCost),
   ); err != nil {
       err = errors.Wrap(err, "Contract transfer token from register failed. ")
       return
   }

   time.Sleep(5 * time.Second)

   if err = p.Bin.ChainWrapper().RegisterAsVerifier(p.makeTxParams(register.Password)); err != nil {
       err = errors.Wrap(err, "Register as verifier failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) Vote(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var vote definition.Preset
   if err = json.Unmarshal(mi.Payload, &vote); err != nil {
       return
   }

   tId, ok := new(big.Int).SetString(vote.TransactionId, 10)
   if !ok {
       err = errors.New("Set to *big.Int failed. ")
       return
   }

   if err = p.Bin.ChainWrapper().Vote(p.makeTxParams(vote.Password), tId, vote.Verify.Suggestion, vote.Verify.Comment); err != nil {
       err = errors.Wrap(err, "Vote failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) GradeToVerifier(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var gtv definition.Preset
   if err = json.Unmarshal(mi.Payload, &gtv); err != nil {
       return
   }
   tId, ok := new(big.Int).SetString(gtv.TransactionId, 10)
   if !ok {
       err = errors.New("Set to *big.Int failed. ")
       return
   }

   txParam := p.makeTxParams(gtv.Password)

   if gtv.Grade.Verifier1Revert {
       credit := uint8(gtv.Grade.Verifier1Grade)
       if err = p.Bin.ChainWrapper().GradeToVerifier(txParam, tId, 0, credit); err != nil {
           err = errors.Wrap(err, "Grade verifier1 failed. ")
           return
       }
   }
   if gtv.Grade.Verifier2Revert {
       credit := uint8(gtv.Grade.Verifier2Grade)
       if err = p.Bin.ChainWrapper().GradeToVerifier(txParam, tId, 1, credit); err != nil {
           err = errors.Wrap(err, "Grade verifier2 failed. ")
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

   var ad definition.Preset
   if err = json.Unmarshal(mi.Payload, &ad); err != nil {
       return
   }
   tId, ok := new(big.Int).SetString(ad.TransactionId, 10)
   if !ok {
       err = errors.New("Set to *big.Int failed. ")
       return
   }

   if err = p.Bin.ChainWrapper().Arbitrate(p.makeTxParams(ad.Password), tId, ad.Arbitrate.ArbitrateResult); err != nil {
       err = errors.Wrap(err, "Arbitrate failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) GetEthBalance(_ *server.MessageIn) (payload interface{}, err error) {
   if p.CurUser == nil {
       err = errors.New("Current user is nil. ")
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

   var gtb definition.Preset
   if err = json.Unmarshal(mi.Payload, &gtb); err != nil {
       return
   }

   var balance *big.Int
   if balance, err = p.Bin.ChainWrapper().GetTokenBalance(p.makeTxParams(gtb.Password), common.HexToAddress(p.CurUser.Account().Addr));err != nil {
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
