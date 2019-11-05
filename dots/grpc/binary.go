package grpc

import (
	"context"
	"encoding/json"
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
	"sync"
	"time"
)

// const
const (
	BinaryGrpcServerTypeId = "96a6e2b5-f0b6-48dc-b0ff-2d9f2c5c9f1d"
	ScanEventInterval      = 200 //milli seconds
)

// ChanEvent
type ChanEvent chan event.Event

// BinaryGrpcServerImp
type BinaryGrpcServerImp struct {
	config       binaryGrpcServerConfig
	eventChanMap sync.Map
	chainWrapper scry.ChainWrapper
	Subscriber   *subscribe.Subscribe `dot:""`
	ServerNobl   gserver.ServerNobl   `dot:""`
}

type binaryGrpcServerConfig struct {
	EventChanCapacity uint64
}

var _ BinaryGrpcServer = (*BinaryGrpcServerImp)(nil)

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

	d := &BinaryGrpcServerImp{config: *dConf}

	return d, err
}

// BinaryGrpcServerImpTypeLive Data structure needed when generating newer component
func BinaryGrpcServerImpTypeLive() []*dot.TypeLives {
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

// Create
func (c *BinaryGrpcServerImp) Create(l dot.Line) error {
	//c.eventChan = make(map[string]ChanEvent)

	return nil
}

// Start
func (c *BinaryGrpcServerImp) Start(ignore bool) error {
	api.RegisterBinaryServiceServer(c.ServerNobl.Server(), c)
	return nil
}

// Stop
func (c *BinaryGrpcServerImp) Stop(ignore bool) error {
	return nil
}

// Destroy
func (c *BinaryGrpcServerImp) Destroy(ignore bool) error {
	return nil
}

// SetChainWrapper
func (c *BinaryGrpcServerImp) SetChainWrapper(w scry.ChainWrapper) {
	c.chainWrapper = w
}

func makeResult(s bool, e string) *api.Result {
	return &api.Result{Success: s, ErrMsg: e}
}

// SubscribeEvent
func (c *BinaryGrpcServerImp) SubscribeEvent(ctx context.Context, info *api.SubscribeInfo) (*api.Result, error) {
	rs := makeResult(true, "")

	hexAddr := info.GetAddress()
	if hexAddr == "" {
		errMsg := "client address can not be empty"
		dot.Logger().Errorln("BinaryGrpcServerImp::SubscribeEvent", zap.String("error:", errMsg))
		rs.ErrMsg = errMsg
		return rs, errors.New(errMsg)
	}

	//data channel for server streaming

	//event channel stored event data between grpc client and grpc server
	var ce ChanEvent
	if rv, ok := c.eventChanMap.Load(hexAddr); ok {
		ce = rv.(ChanEvent)
	} else {
		errMsg := "failed to subscribe event since no server streaming channel found"
		dot.Logger().Errorln("BinaryGrpcServerImp::SubscribeEvent", zap.String("error:", errMsg))
		rs.ErrMsg = errMsg
		return rs, errors.New(errMsg)
	}

	//grpc client addr
	addr := common.HexToAddress(info.GetAddress())
	for _, ev := range info.GetEvent() {
		err := c.Subscriber.Subscribe(addr, ev, func(event event.Event) bool {
			ce <- event
			return true
		})

		if err != nil {
			dot.Logger().Errorln("BinaryGrpcServerImp::SubscribeEvent", zap.Error(err))
			rs.ErrMsg = err.Error()
			return rs, err
		}
	}

	return rs, nil
}

// UnSubscribeEvent
func (c *BinaryGrpcServerImp) UnSubscribeEvent(
	ctx context.Context,
	info *api.SubscribeInfo,
) (*api.Result, error) {
	rs := makeResult(true, "")

	hexAddr := info.GetAddress()
	if hexAddr == "" {
		errMsg := "client address can not be empty"
		dot.Logger().Errorln("BinaryGrpcServerImp::UnSubscribeEvent", zap.String("error:", errMsg))
		rs.ErrMsg = errMsg
		return rs, errors.New(errMsg)
	}

	addr := common.HexToAddress(info.GetAddress())
	for _, ev := range info.GetEvent() {
		err := c.Subscriber.UnSubscribe(addr, ev)
		if err != nil {
			dot.Logger().Errorln("BinaryGrpcServerImp::UnSubscribeEvent", zap.Error(err))
		}
	}

	return rs, nil
}

// RecvEvents RecvEvents the function should be called firstly to create server stream channel
func (c *BinaryGrpcServerImp) RecvEvents(client *api.ClientInfo, srv api.BinaryService_RecvEventsServer) error {
	defer func() {
		if err := recover(); err != nil {
			dot.Logger().Errorln("BinaryGrpcServerImp::RecvEvents", zap.Any("error:", err))
		}
	}()

	//create channel for server streaming
	if client == nil || client.Address == "" {
		errMsg := "client address can not be empty"
		dot.Logger().Errorln("BinaryGrpcServerImp::RecvEvents", zap.String("error:", errMsg))
		return errors.New(errMsg)
	}

	//event channel stored event data between grpc client and grpc server
	var ce ChanEvent
	if rv, ok := c.eventChanMap.Load(client.Address); !ok {
		ce = make(ChanEvent, c.config.EventChanCapacity)
		c.eventChanMap.Store(client.Address, ce)
	} else {
		ce = rv.(ChanEvent)
	}

	//channel created event
	ce <- *makeChannelCreatedEvent()

	//push stream
	for {
		select {
		case e := <-ce:
			dot.Logger().Debugln("BinaryGrpcServerImp::RecvEvents", zap.String("event:", e.Name))

			ev, err := makeProtoEvent(&e)
			if err != nil {
				dot.Logger().Errorln("BinaryGrpcServerImp::RecvEvents", zap.String("error:", err.Error()))
				break
			}

			err = srv.Send(ev)
			if err != nil {
				dot.Logger().Errorln("BinaryGrpcServerImp::RecvEvents", zap.String("error:", err.Error()))
				ce <- e
				return err
			}

		case <-time.After(time.Microsecond * ScanEventInterval):
		}
	}
}

func makeChannelCreatedEvent() *event.Event {
	return &event.Event{
		Name: "ChannelCreated",
	}
}

func makeProtoEvent(e *event.Event) (*api.Event, error) {
	obj := map[string]interface{}{
		"BlockNumber":     e.BlockNumber,
		"ContractAddress": e.Address.String(),
		"EventName":       e.Name,
		"TxHash":          e.TxHash.String(),
		"EventData":       e.Data.String(),
	}

	jsonEvent, err := json.Marshal(obj)
	if err != nil {
		dot.Logger().Errorln("BinaryGrpcServerImp::makeProtoEvent", zap.String("error:", err.Error()))
		return nil, err
	}

	pe := &api.Event{
		Time:     time.Now().Unix(),
		JsonData: string(jsonEvent),
	}

	return pe, nil
}

// Publish
func (c *BinaryGrpcServerImp) Publish(ctx context.Context, params *api.PublishParams) (*api.PublishResult, error) {
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
		From:     common.HexToAddress(p.From),
		Password: p.Password,
		Value:    big.NewInt(p.Value),
		Pending:  p.Pending,
		GasLimit: p.GasLimit,
		GasPrice: big.NewInt(p.GasPrice),
	}

	return t
}

func makePublishResult(r **api.PublishResult, pid, e string, s bool) {
	if *r == nil {
		*r = &api.PublishResult{
			PublishId: pid,
			Result:    makeResult(s, e),
		}
	} else {
		(*r).PublishId = pid
		(*r).Result = makeResult(s, e)
	}
}

// CreateAccount
func (c *BinaryGrpcServerImp) CreateAccount(
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

	client, err := scry.CreateScryClient(in.Password, c.chainWrapper)
	if err != nil {
		makeAccountResult(&ar, "", err.Error(), false)
		return ar, err
	}

	makeAccountResult(&ar, client.Account().Addr, "", true)
	return ar, nil
}

func makeAccountResult(r **api.AccountResult, aid, e string, s bool) {
	if *r == nil {
		*r = &api.AccountResult{
			AccountId: aid,
			Result:    makeResult(s, e),
		}
	} else {
		(*r).AccountId = aid
		(*r).Result = makeResult(s, e)
	}
}

// Authenticate
func (c *BinaryGrpcServerImp) Authenticate(
	ctx context.Context,
	in *api.ClientInfo,
) (*api.Result, error) {
	if c.chainWrapper == nil {
		e := "invalid scry chain interface"
		return makeResult(false, e), errors.New(e)
	}

	client := scry.NewScryClient(in.Address, c.chainWrapper)
	if client == nil {
		e := "failed to new scry client"
		return makeResult(false, e), errors.New(e)
	}

	_, err := client.Authenticate(in.Password)
	if err != nil {
		return makeResult(false, err.Error()), err
	}

	return makeResult(true, ""), nil
}

// TransferEth
func (c *BinaryGrpcServerImp) TransferEth(
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

// GetEthBalance
func (c *BinaryGrpcServerImp) GetEthBalance(
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

func makeEthBalanceResult(r **api.EthBalanceResult, e string, s bool, b int64) {
	if *r == nil {
		*r = &api.EthBalanceResult{
			Balance: b,
			Result:  makeResult(s, e),
		}
	} else {
		(*r).Balance = b
		(*r).Result = makeResult(s, e)
	}
}

// TransferTokens
func (c *BinaryGrpcServerImp) TransferTokens(
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

// GetTokenBalance
func (c *BinaryGrpcServerImp) GetTokenBalance(
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

func makeTokenBalanceResult(r **api.TokenBalanceResult, e string, s bool, b int64) {
	if *r == nil {
		*r = &api.TokenBalanceResult{
			Balance: b,
			Result:  makeResult(s, e),
		}
	} else {
		(*r).Balance = b
		(*r).Result = makeResult(s, e)
	}
}

// PrepareToBuy
func (c *BinaryGrpcServerImp) PrepareToBuy(
	ctx context.Context,
	params *api.PrepareParams,
) (*api.Result, error) {
	if c.chainWrapper == nil {
		e := "invalid scry chain interface"
		return makeResult(false, e), errors.New(e)
	}

	if params == nil || params.TxParam == nil {
		e := "null publish parameters"
		return makeResult(false, e), errors.New(e)
	}

	err := c.chainWrapper.AdvancePurchase(
		makeTxParams(params.TxParam),
		params.PublishId,
		params.StartVerify,
	)
	if err != nil {
		e := err.Error()
		return makeResult(false, e), err
	}

	return makeResult(true, ""), nil
}

// BuyData
func (c *BinaryGrpcServerImp) BuyData(
	ctx context.Context,
	params *api.BuyParams,
) (*api.Result, error) {
	if c.chainWrapper == nil {
		e := "invalid scry chain interface"
		return makeResult(false, e), errors.New(e)
	}

	if params == nil || params.TxParam == nil {
		e := "null publish parameters"
		return makeResult(false, e), errors.New(e)
	}

	err := c.chainWrapper.ConfirmPurchase(
		makeTxParams(params.TxParam),
		big.NewInt(params.TxId),
	)
	if err != nil {
		e := err.Error()
		return makeResult(false, e), err
	}

	return makeResult(true, ""), nil
}

// CancelTransaction
func (c *BinaryGrpcServerImp) CancelTransaction(
	ctx context.Context,
	params *api.CancelTxParams,
) (*api.Result, error) {
	if c.chainWrapper == nil {
		e := "invalid scry chain interface"
		return makeResult(false, e), errors.New(e)
	}

	if params == nil || params.TxParam == nil {
		e := "null publish parameters"
		return makeResult(false, e), errors.New(e)
	}

	err := c.chainWrapper.CancelPurchase(
		makeTxParams(params.TxParam),
		big.NewInt(params.TxId),
	)
	if err != nil {
		e := err.Error()
		return makeResult(false, e), err
	}

	return makeResult(true, ""), nil
}

// ReEncryptMetaDataId re-encrypt meta data id
func (c *BinaryGrpcServerImp) ReEncryptMetaDataId(
	ctx context.Context,
	params *api.ReEncryptDataParams,
) (*api.Result, error) {
	if c.chainWrapper == nil {
		e := "invalid scry chain interface"
		return makeResult(false, e), errors.New(e)
	}

	if params == nil || params.TxParam == nil {
		e := "null publish parameters"
		return makeResult(false, e), errors.New(e)
	}

	//get buyer address and arbitrators address
	err := c.chainWrapper.ReEncrypt(
		makeTxParams(params.TxParam),
		big.NewInt(params.TxId),
		params.EncodedDataWithSeller,
	)
	if err != nil {
		e := err.Error()
		return makeResult(false, e), err
	}

	return makeResult(true, ""), nil
}

// ConfirmDataTruth
func (c *BinaryGrpcServerImp) ConfirmDataTruth(
	ctx context.Context,
	params *api.DataConfirmParams,
) (*api.Result, error) {
	if c.chainWrapper == nil {
		e := "invalid scry chain interface"
		return makeResult(false, e), errors.New(e)
	}

	if params == nil || params.TxParam == nil {
		e := "null publish parameters"
		return makeResult(false, e), errors.New(e)
	}

	err := c.chainWrapper.ConfirmData(
		makeTxParams(params.TxParam),
		big.NewInt(params.TxId),
		params.Truth,
	)
	if err != nil {
		e := err.Error()
		return makeResult(false, e), err
	}

	return makeResult(true, ""), nil
}

// ApproveTransfer
func (c *BinaryGrpcServerImp) ApproveTransfer(
	ctx context.Context,
	params *api.ApproveTransferParams,
) (*api.Result, error) {
	if c.chainWrapper == nil {
		e := "invalid scry chain interface"
		return makeResult(false, e), errors.New(e)
	}

	if params == nil || params.TxParam == nil {
		e := "null publish parameters"
		return makeResult(false, e), errors.New(e)
	}

	err := c.chainWrapper.ApproveTransfer(
		makeTxParams(params.TxParam),
		common.HexToAddress(params.SpenderAddr),
		big.NewInt(params.Value),
	)
	if err != nil {
		e := err.Error()
		return makeResult(false, e), err
	}

	return makeResult(true, ""), nil
}

// Vote
func (c *BinaryGrpcServerImp) Vote(
	ctx context.Context,
	params *api.VoteParams,
) (*api.Result, error) {
	if c.chainWrapper == nil {
		e := "invalid scry chain interface"
		return makeResult(false, e), errors.New(e)
	}

	if params == nil || params.TxParam == nil {
		e := "null publish parameters"
		return makeResult(false, e), errors.New(e)
	}

	err := c.chainWrapper.Vote(
		makeTxParams(params.TxParam),
		big.NewInt(params.TxId),
		params.Judge,
		params.Comments,
	)
	if err != nil {
		e := err.Error()
		return makeResult(false, e), err
	}

	return makeResult(true, ""), nil
}

// RegisterAsVerifier
func (c *BinaryGrpcServerImp) RegisterAsVerifier(
	ctx context.Context,
	params *api.RegisterVerifierParams,
) (*api.Result, error) {

	if c.chainWrapper == nil {
		e := "invalid scry chain interface"
		return makeResult(false, e), errors.New(e)
	}

	if params == nil || params.TxParam == nil {
		e := "null publish parameters"
		return makeResult(false, e), errors.New(e)
	}

	err := c.chainWrapper.RegisterAsVerifier(
		makeTxParams(params.TxParam),
	)
	if err != nil {
		e := err.Error()
		return makeResult(false, e), err
	}

	return makeResult(true, ""), nil
}

// CreditsToVerifier
func (c *BinaryGrpcServerImp) CreditsToVerifier(
	ctx context.Context,
	params *api.CreditVerifierParams,
) (*api.Result, error) {
	if c.chainWrapper == nil {
		e := "invalid scry chain interface"
		return makeResult(false, e), errors.New(e)
	}

	if params == nil || params.TxParam == nil {
		e := "null publish parameters"
		return makeResult(false, e), errors.New(e)
	}

	err := c.chainWrapper.GradeToVerifier(
		makeTxParams(params.TxParam),
		big.NewInt(params.TxId),
		uint8(params.Index),
		uint8(params.Credit),
	)
	if err != nil {
		e := err.Error()
		return makeResult(false, e), err
	}

	return makeResult(true, ""), nil
}
