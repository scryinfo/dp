package preset

import (
    "encoding/json"
    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    def "github.com/scryinfo/dp/dots/app/business/definition"
    "github.com/scryinfo/dp/dots/app/business/preset/chain_event"
    "github.com/scryinfo/dp/dots/app/server"
    "github.com/scryinfo/dp/dots/app/storage/definition"
    "github.com/scryinfo/dp/dots/binary"
    scry2 "github.com/scryinfo/dp/dots/binary/scry"
    "github.com/scryinfo/dp/dots/eth/transaction"
    "go.uber.org/zap"
    "math/big"
    "os"
    "strconv"
    "time"
)

type Preset struct {
    PresetMsgNames    []string
    PresetMsgHandlers []server.PresetFunc
    Deployer          *def.Preset
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

    // todo: params get from contract.
    verifierNum            = 2
    verifierBonus          = 300
    registerAsVerifierCost = 10000

    arbitratorNum   = 1
    arbitratorBonus = 500

    sep = "|"
)

func (p *Preset) Create(l dot.Line) error {
    p.Deployer = &def.Preset{
        Address:  "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8",
        Password: "111111",
    }

    p.PresetMsgNames = []string{
        "loginVerify",
        "createNewAccount",
        "currentUserDataUpdate",
        "logout",

        "publish",
        "advancePurchase",
        "confirmPurchase",
        "reEncrypt",
        "finishPurchase",
        "decrypt",
        "confirmData",
        "register",
        "vote",
        "gradeToVerifier",
        "arbitrate",

        "getEthBalance",
        "getTokenBalance",
        "isVerifier",
        "modifyNickname",
        "modifyContractParam",
    }

    p.PresetMsgHandlers = []server.PresetFunc{
        p.LoginVerify,
        p.CreateNewAccount,
        p.CurrentUserDataUpdate,
        p.Logout,

        p.Publish,
        p.AdvancePurchase,
        p.ConfirmPurchase,
        p.ReEncrypt,
        p.FinishPurchase,
        p.Decrypt,
        p.ConfirmData,
        p.Register,
        p.Vote,
        p.GradeToVerifier,
        p.Arbitrate,

        p.GetEthBalance,
        p.GetTokenBalance,
        p.IsVerifier,
        p.ModifyNickname,
        p.ModifyContractParam,
    }

    if len(p.PresetMsgNames) != len(p.PresetMsgHandlers) {
        return errors.New("Quantities of name and function are not matched. (preset) ")
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
   var lv def.Preset
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
       p.CBs.CurUser = client
   } else {
       err = errors.New("Login verify failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) CreateNewAccount(mi *server.MessageIn) (payload interface{}, err error) {
   var cna def.Preset
   if err = json.Unmarshal(mi.Payload, &cna); err != nil {
       return
   }

   var client scry2.Client
   if client, err = scry2.CreateScryClient(cna.Password, p.Bin.ChainWrapper()); err != nil {
       err = errors.Wrap(err, "Create new user failed. ")
       return
   }

   p.CBs.CurUser = client

   payload = client.Account().Addr

   var num int64
   if num, err = p.CBs.DB.Insert(&definition.Account{
       Address: client.Account().Addr,
       Nickname: client.Account().Addr,
       FromBlock: 1,
       IsVerifier: false,
       Verify: nil,
   }); num != 1 || err != nil {
       err = errors.Wrap(err, "in create new account")
       return
   }

   return
}

func (p *Preset) CurrentUserDataUpdate(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var bs def.Preset
   if err = json.Unmarshal(mi.Payload, &bs); err != nil {
       return
   }

   for i := range p.CBs.EventNames {
       if err = p.CBs.CurUser.SubscribeEvent(p.CBs.EventNames[i], p.CBs.EventHandler[i]); err != nil {
           err = errors.Wrap(err, "Subscribe event failed. ")
           return
       }
   }

   var acc definition.Account
   // set from block
   {
       var num int64
       if num, err = p.CBs.DB.Read(&acc, "", "address = ?", bs.Address); num != 1 || err != nil {
           err = errors.Wrap(err, "db read failed")
           return
       }
       p.Bin.Listener.SetFromBlock(uint64(acc.FromBlock))
   }

   // when an user login success, he will get 10,000,000 eth and tokens for test.
   {
       if err = p.CBs.CurUser.TransferEthFrom(common.HexToAddress(p.Deployer.Address),
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
       if err = p.Bin.ChainWrapper().TransferTokens(&txParam, common.HexToAddress(p.CBs.CurUser.Account().Addr), big.NewInt(10000000)); err != nil {
           err = errors.Wrap(err, "Transfer token from Deployer failed. ")
           return
       }
   }

   payload = acc.Nickname

   return
}

func (p *Preset) Logout(_ *server.MessageIn) (payload interface{}, err error) {
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   for i := range p.CBs.EventNames {
       if err = p.CBs.CurUser.UnSubscribeEvent(p.CBs.EventNames[i]); err != nil {
           err = errors.Wrap(err, "Unsubscribe failed, event:  "+p.CBs.EventNames[i]+" . ")
           return
       }
   }

   p.CBs.CurUser = nil

   p.CBs.FlagChan = nil

   payload = true

   return
}

func (p *Preset) Publish(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var publish def.Preset
   if err = json.Unmarshal(mi.Payload, &publish); err != nil {
       return
   }

   price, err := strconv.Atoi(publish.Price)
   if err != nil {
       return
   }

   if payload, err = p.Bin.ChainWrapper().Publish(
       p.makeTxParams(publish.Password),
       big.NewInt(int64(price)),
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
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var ap def.Preset
   if err = json.Unmarshal(mi.Payload, &ap); err != nil {
       return
   }

   price, err := strconv.Atoi(ap.Price)
   if err != nil {
       return
   }
   fee := int64(price)
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

    <- p.CBs.FlagChan

   if err = p.Bin.ChainWrapper().AdvancePurchase(p.makeTxParams(ap.Password), ap.PublishId, ap.StartVerify); err != nil {
       err = errors.Wrap(err, "Advance purchase failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) ConfirmPurchase(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var cp def.Preset
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
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var re def.Preset
   if err = json.Unmarshal(mi.Payload, &re); err != nil {
       return
   }

   txParam := p.makeTxParams(re.Password)
   tId, ok := new(big.Int).SetString(re.TransactionId, 10)
   if !ok {
       err = errors.New("Set to *big.Int failed. ")
       return
   }

   var tx definition.Transaction
   // get meta data id enc with seller
   {
       var num int64
       if num, err = p.CBs.DB.Read(&tx, "", "transaction_id = ?", re.TransactionId); num != 1 || err != nil {
           dot.Logger().Errorln("db read failed", zap.Int64("affect rows number", num) ,zap.NamedError("error", err))
           return
       }
   }

   if err = p.Bin.ChainWrapper().ReEncrypt(txParam, tId, []byte(tx.MetaDataIdEncWithSeller)); err != nil {
       err = errors.Wrap(err, "Re-encrypt failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) FinishPurchase(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var cp def.Preset
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
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var decrypt def.Preset
   if err = json.Unmarshal(mi.Payload, &decrypt); err != nil {
       return
   }

    var tx definition.Transaction
    // get meta data id enc with seller
    {
        var num int64
        if num, err = p.CBs.DB.Read(&tx, "", "transaction_id = ?", decrypt.TransactionId); num != 1 || err != nil {
            dot.Logger().Errorln("db read failed", zap.Int64("affect rows number", num) ,zap.NamedError("error", err))
            return
        }
    }

   var oldFileName string
   {
       var metaDataIdByte []byte
       {
           if p.CBs.CurUser.Account().Addr == tx.Buyer {
               metaDataIdByte = []byte(tx.MetaDataIdEncWithBuyer)
           } else {
               metaDataIdByte = []byte(tx.MetaDataIdEncWithArbitrator)
           }
       }
       if metaDataIdByte, err = p.Bin.Account.Decrypt(metaDataIdByte, p.CBs.CurUser.Account().Addr, decrypt.Password); err != nil {
           return "", errors.Wrap(err, "Decrypt encrypted meta data Id failed. ")
       }
       outDir := p.config.MetaDataOutDir
       if err = p.Bin.Storage.Get(string(metaDataIdByte), outDir); err != nil {
           return "", errors.Wrap(err, "Get meta data from IPFS failed. ")
       }
       oldFileName = outDir + "/" + string(metaDataIdByte)
   }

   newFileName := oldFileName + tx.MetaDataExtension
   if err = os.Rename(oldFileName, newFileName); err != nil {
       return "", errors.Wrap(err, "Add extension to meta data failed. ")
   }

   payload = newFileName

   return
}

func (p *Preset) ConfirmData(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var cd def.Preset
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
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var register def.Preset
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

    <- p.CBs.FlagChan

   if err = p.Bin.ChainWrapper().RegisterAsVerifier(p.makeTxParams(register.Password)); err != nil {
       err = errors.Wrap(err, "Register as verifier failed. ")
       return
   }

   payload = true

   return
}

func (p *Preset) Vote(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var vote def.Preset
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
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var gtv def.Preset
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
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var ad def.Preset
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
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var balance *big.Int
   if balance, err = p.CBs.CurUser.GetEth(common.HexToAddress(p.CBs.CurUser.Account().Addr), p.Bin.ChainWrapper().Conn());err != nil {
       err = errors.Wrap(err, "Get eth balance failed. ")
       return
   }

   payload = balance.String() + sep + time.Now().String()

   return
}

func (p *Preset) GetTokenBalance(mi *server.MessageIn) (payload interface{}, err error) {
   if p.CBs.CurUser == nil {
       err = errors.New("Current user is nil. ")
       return
   }

   var gtb def.Preset
   if err = json.Unmarshal(mi.Payload, &gtb); err != nil {
       return
   }

   var balance *big.Int
   if balance, err = p.Bin.ChainWrapper().GetTokenBalance(p.makeTxParams(gtb.Password), common.HexToAddress(p.CBs.CurUser.Account().Addr));err != nil {
       err = errors.Wrap(err, "Get token balance failed. ")
       return
   }

   payload = balance.String() + sep + time.Now().String()

   return
}

func (p *Preset) IsVerifier(_ *server.MessageIn) (payload interface{}, err error) {
    var (
        acc definition.Account
        num int64
    )
    if num, err = p.CBs.DB.Read(&acc, "", "address = ?", p.CBs.CurUser.Account().Addr); num != 1 || err != nil {
        dot.Logger().Errorln("db read failed", zap.Int64("affect rows number", num) ,zap.NamedError("error", err))
        return
    }

    payload = acc.IsVerifier

    return
}

func (p *Preset) ModifyNickname(mi *server.MessageIn) (payload interface{}, err error) {
    var mn def.User
    if err = json.Unmarshal(mi.Payload, &mn); err != nil {
        return
    }

    var acc definition.Account
    num, err := p.CBs.DB.Update(&acc, map[string]interface{}{"nickname": mn.Nickname}, "address = ?", mn.Address)
    if num != 1 || err != nil {
        err = errors.Wrap(err, "Modify nickname failed. ")
        return
    }

    payload = true

    return
}

func (p *Preset) ModifyContractParam(mi *server.MessageIn) (payload interface{}, err error) {
    var mcp def.Preset
    if err = json.Unmarshal(mi.Payload, &mcp); err != nil {
        return
    }

    if err = p.Bin.ChainWrapper().ModifyContractParam(&transaction.TxParams{
        From: common.HexToAddress(p.Deployer.Address),
        Password: p.Deployer.Password,
        Value: big.NewInt(0),
        Pending: false,
    }, mcp.Contract.ParamName, mcp.Contract.ParamValue); err != nil {
        err = errors.New("Unknown contract param or param is not allowed to modify. ")
        return
    }

    return
}

func (p *Preset) makeTxParams(password string) *transaction.TxParams {
   return &transaction.TxParams{
       From:     common.HexToAddress(p.CBs.CurUser.Account().Addr),
       Password: password,
       Value:    big.NewInt(0),
       Pending:  false,
   }
}
