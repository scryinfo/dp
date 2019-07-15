package grpc

import (
    "context"
    "errors"
    "github.com/ethereum/go-ethereum/common"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dot/dots/grpc/gserver"
    "github.com/scryinfo/dp/api/go"
    "github.com/scryinfo/dp/dots/binary/scry"
    "github.com/scryinfo/dp/dots/eth/event"
    "github.com/scryinfo/dp/dots/eth/event/subscribe"
    "github.com/scryinfo/dp/dots/eth/transaction"
    "go.uber.org/zap"
    "math/big"
    "time"
)

const BinaryGrpcServerTypeId = "96a6e2b5-f0b6-48dc-b0ff-2d9f2c5c9f1d"

type ChanEvent chan event.Event

type BinaryGrpcServer struct {
    config       binaryGrpcServerConfig
    eventChan    map[string]ChanEvent
    chainWrapper scry.ChainWrapper
    Subscriber   *subscribe.Subscribe `dot:""`
    ServerNobl   gserver.ServerNobl   `dot:""`
}

type binaryGrpcServerConfig struct {
    EventChanCap uint64
}

func newBinaryGrpcServerDot(conf interface{}) (dot.Dot, error) {
    var err error
    var bs []byte
    if bt, ok := conf.([]byte); ok {
        bs = bt
    } else {
        return nil, dot.SError.Parameter
    }

    dConf := &binaryGrpcServerConfig{}
    err = dot.UnMarshalConfig(bs, dConf)
    if err != nil {
        return nil, err
    }

    d := &BinaryGrpcServer{config: *dConf}

    return d, err
}

//Data structure needed when generating newer component
func BinaryGrpcServerTypeLive() []*dot.TypeLives {
    t := []*dot.TypeLives{
        &dot.TypeLives{
            Meta: dot.Metadata{TypeId: BinaryGrpcServerTypeId,
                NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
                    return newBinaryGrpcServerDot(conf)
                },
            },
            Lives: []dot.Live{
                dot.Live{
                    LiveId:    BinaryGrpcServerTypeId,
                    RelyLives: map[string]dot.LiveId{"ServerNobl": gserver.ServerNoblTypeId},
                },
            },
        },
        subscribe.SubsTypeLive(),
    }

    t = append(t, gserver.HttpNoblTypeLives()...)

    return t
}

func (c *BinaryGrpcServer) Create(l dot.Line) error {
    c.eventChan = make(map[string]ChanEvent)

    return nil
}

func (c *BinaryGrpcServer) Start(ignore bool) error {
    api.RegisterBinaryServiceServer(c.ServerNobl.Server(), c)
    return nil
}

func (c *BinaryGrpcServer) Stop(ignore bool) error {
    return nil
}

func (c *BinaryGrpcServer) Destroy(ignore bool) error {
    return nil
}

func (c *BinaryGrpcServer) SetChainWrapper(w scry.ChainWrapper) {
    c.chainWrapper = w
}

func makeResult(s bool, e string) *api.Result {
    return &api.Result{Success: s, ErrMsg: e}
}

func (c *BinaryGrpcServer) SubscribeEvent(ctx context.Context, info *api.SubscribeInfo) (*api.Result, error) {
    rs := makeResult(true,"")

    hexAddr := info.GetAddress()
    if hexAddr == "" {
        errMsg := "client address can not be empty"
        dot.Logger().Errorln("BinaryGrpcServer::SubscribeEvent", zap.String("error:", errMsg))
        rs.ErrMsg = errMsg
        return rs, errors.New(errMsg)
    }

    //data channel for server streaming
    ce := c.eventChan[hexAddr]
    if ce == nil {
        errMsg := "failed to subscribe event as no server streaming channel could be found"
        dot.Logger().Errorln("BinaryGrpcServer::SubscribeEvent", zap.String("error:", errMsg))
        rs.ErrMsg = errMsg
        return rs, errors.New(errMsg)
    }

    addr := common.HexToAddress(info.GetAddress())
    err := c.Subscriber.Subscribe(addr, info.GetEvent(), func(event event.Event) bool {
        ce <- event
        return true
    })

    if err != nil {
        dot.Logger().Errorln("BinaryGrpcServer::SubscribeEvent", zap.Error(err))
        rs.ErrMsg = err.Error()
        return rs, err
    }

    return rs, nil
}

func (c *BinaryGrpcServer) UnSubscribeEvent(context.Context, *api.SubscribeInfo) (*api.Result, error) {
    return nil, nil
}

//the function should be called firstly to create server stream channel
func (c *BinaryGrpcServer) RecvEvents(client *api.ClientInfo,srv api.BinaryService_RecvEventsServer) error {
    //create channel for server streaming
    if client.Address == "" {
        errMsg := "client address can not be empty"
        dot.Logger().Errorln("BinaryGrpcServer::RecvEvents", zap.String("error:", errMsg))
        return errors.New(errMsg)
    }

    ce := c.eventChan[client.Address]
    if ce == nil {
        ce = make(chan event.Event, c.config.EventChanCap)
        c.eventChan[client.Address] = ce
    }

    //push stream
    for {
        select {
        case e := <- ce:
            dot.Logger().Debugln("BinaryGrpcServer::RecvEvents", zap.String("event:", e.Name))
            err := srv.Send(makeProtoEvent(&e))
            if err != nil {
                dot.Logger().Errorln("BinaryGrpcServer::RecvEvents", zap.String("error:", err.Error()))
            }
        case <-time.After(time.Microsecond * 50):
        }
    }
}

func makeProtoEvent(e *event.Event) *api.Event {
    pe := &api.Event{
        BlockNumber: e.BlockNumber,
        Address: e.Address.String(),
        Name: e.Name,
        TxHash: e.TxHash.String(),
        JsonData: e.Data.String(),
    }

    return pe
}

func (c *BinaryGrpcServer) Publish(ctx context.Context, params *api.PublishParams) (*api.PublishResult, error) {
    var pr *api.PublishResult
    makePublishResult(&pr, "", "", true)

    if c.chainWrapper == nil {
        errMsg := "invalid scry chain interface"
        makePublishResult(&pr, "", errMsg, false)
        return pr, errors.New(errMsg)
    }

    if params == nil || params.TxParam == nil {
        errMsg := "null publish parameters"
        makePublishResult(&pr, "", errMsg, false)
        return pr, errors.New(errMsg)
    }

    pid, err := c.chainWrapper.Publish(
        makeTxParams(params.TxParam),
        big.NewInt(params.Price),
        params.MetaDataID,
        params.ProofDataIDs,
        params.ProofNum,
        params.DetailsID,
        params.SupportVerify,
    )
    if err != nil {
        e := err.Error()
        makePublishResult(&pr, "", e, false)
        return pr, err
    }

    makePublishResult(&pr, pid, "", true)
    return pr, nil
}

func makeTxParams(p *api.TxParams) *transaction.TxParams {
    t := &transaction.TxParams{
        From: common.HexToAddress(p.From),
        Password: p.Password,
        Value: big.NewInt(p.Value),
        Pending: p.Pending,
        GasLimit: p.GasLimit,
        GasPrice: big.NewInt(p.GasPrice),
    }

    return t
}

func makePublishResult(r **api.PublishResult, pid, e string, s bool)  {
    if *r == nil {
        *r = &api.PublishResult{
            PublishId: pid,
            Result: makeResult(s, e),
        }
    } else {
        (*r).PublishId = pid
        (*r).Result = makeResult(s, e)
    }
}

func (c *BinaryGrpcServer) CreateAccount(
    ctx context.Context,
    in *api.CreateAccountParams,
) (*api.AccountResult, error) {
    var ar *api.AccountResult
    makeAccountResult(&ar, "", "", true)

    if c.chainWrapper == nil {
        e := "invalid scry chain interface"
        makeAccountResult(&ar, "", e, false)
        return ar, errors.New(e)
    }

    client, err:= scry.CreateScryClient(in.Password, c.chainWrapper)
    if err != nil {
        makeAccountResult(&ar, "", err.Error(), false)
        return ar, err
    }

    makeAccountResult(&ar, client.Account().Addr, "", true)
    return ar, nil
}

func makeAccountResult(r **api.AccountResult, aid, e string, s bool)  {
    if *r == nil {
        *r = &api.AccountResult{
            AccountId: aid,
            Result: makeResult(s, e),
        }
    } else {
        (*r).AccountId = aid
        (*r).Result = makeResult(s, e)
    }
}

func (c *BinaryGrpcServer) Authenticate(
    ctx context.Context,
    in *api.ClientInfo,
) (*api.Result, error) {
    return nil, nil
}

func (c *BinaryGrpcServer) TransferEth(
    ctx context.Context,
    in *api.TransferEthParams,
) (*api.Result, error) {
    if c.chainWrapper == nil {
        e := "invalid scry chain interface"
        return makeResult(false, e), errors.New(e)
    }

    client := scry.NewScryClient(in.To, c.chainWrapper)
    if client == nil {
        e := "failed to create scry client"
        return makeResult(false, e), errors.New(e)
    }

    err := client.TransferEthFrom(
        common.HexToAddress(in.From),
        in.Password,
        big.NewInt(in.Value),
        c.chainWrapper.Conn(),
    )
    if err != nil {
        return makeResult(false, err.Error()), err
    }

    return makeResult(true, ""), nil
}

func (c *BinaryGrpcServer) GetEthBalance(
    ctx context.Context,
    in *api.EthBalanceParams,
) (*api.EthBalanceResult, error) {
    var r *api.EthBalanceResult
    makeEthBalanceResult(&r, "", true, 0)

    if c.chainWrapper == nil {
        e := "invalid scry chain interface"
        makeEthBalanceResult(&r, e, false, 0)
        return r, errors.New(e)
    }

    client := scry.NewScryClient(in.Owner, c.chainWrapper)
    if client == nil {
        e := "failed to create scry client"
        makeEthBalanceResult(&r, e, false, 0)
        return r, errors.New(e)
    }

    b, err := client.GetEth(
        common.HexToAddress(in.Owner),
        c.chainWrapper.Conn(),
    )
    if err != nil {
        makeEthBalanceResult(&r, err.Error(), false, 0)
        return r, err
    }

    makeEthBalanceResult(&r, "", true, b.Int64())
    return r, err
}

func makeEthBalanceResult(r **api.EthBalanceResult, e string, s bool, b int64)  {
    if *r == nil {
        *r = &api.EthBalanceResult{
            Balance: b,
            Result: makeResult(s, e),
        }
    } else {
        (*r).Balance = b
        (*r).Result = makeResult(s, e)
    }
}

func (c *BinaryGrpcServer) TransferTokens(
    ctx context.Context,
    params *api.TransferTokenParams,
) (*api.Result, error) {
    if c.chainWrapper == nil {
        e := "invalid scry chain interface"
        return makeResult(false, e), errors.New(e)
    }

    if params == nil || params.TxParam == nil {
        e := "null publish parameters"
        return makeResult(false, e), errors.New(e)
    }

    err := c.chainWrapper.TransferTokens(
        makeTxParams(params.TxParam),
        common.HexToAddress(params.To),
        big.NewInt(params.Value),
    )
    if err != nil {
        e := err.Error()
        return makeResult(false, e), err
    }

    return makeResult(true, ""), nil
}

func (c *BinaryGrpcServer) GetTokenBalance(
    ctx context.Context,
    params *api.TokenBalanceParams,
) (*api.TokenBalanceResult, error) {
    var r *api.TokenBalanceResult
    makeTokenBalanceResult(&r, "", true, 0)

    if c.chainWrapper == nil {
        e := "invalid scry chain interface"
        makeTokenBalanceResult(&r, e, false, 0)
        return r, errors.New(e)
    }

    b, err := c.chainWrapper.GetTokenBalance(
        makeTxParams(params.TxParam),
        common.HexToAddress(params.Owner),
    )
    if err != nil {
        makeTokenBalanceResult(&r, err.Error(), false, 0)
        return r, err
    }

    makeTokenBalanceResult(&r, "", true, b.Int64())
    return r, err
}

func makeTokenBalanceResult(r **api.TokenBalanceResult, e string, s bool, b int64)  {
    if *r == nil {
        *r = &api.TokenBalanceResult{
            Balance: b,
            Result: makeResult(s, e),
        }
    } else {
        (*r).Balance = b
        (*r).Result = makeResult(s, e)
    }
}

func (c *BinaryGrpcServer) PrepareToBuy(context.Context, *api.PrepareParams) (*api.Result, error) {
    return nil, nil
}

func (c *BinaryGrpcServer) BuyData(context.Context, *api.BuyParams) (*api.Result, error) {
    return nil, nil
}

func (c *BinaryGrpcServer) CancelTransaction(context.Context, *api.CancelTxParams) (*api.Result, error) {
    return nil, nil
}

func (c *BinaryGrpcServer) SubmitMetaDataIdEncWithBuyer(context.Context, *api.SubmitMetaDataIdParams) (*api.Result, error) {
    return nil, nil
}

func (c *BinaryGrpcServer) ConfirmDataTruth(context.Context, *api.DataConfirmParams) (*api.Result, error) {
    return nil, nil
}

func (c *BinaryGrpcServer) ApproveTransfer(context.Context, *api.ApproveTransferParams) (*api.Result, error) {
    return nil, nil
}

func (c *BinaryGrpcServer) Vote(context.Context, *api.VoteParams) (*api.Result, error) {
    return nil, nil
}

func (c *BinaryGrpcServer) RegisterAsVerifier(context.Context, *api.RegisterVerifierParams) (*api.Result, error) {
    return nil, nil
}

func (c *BinaryGrpcServer) CreditsToVerifier(context.Context, *api.CreditVerifierParams) (*api.Result, error) {
    return nil, nil
}




