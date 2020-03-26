package scan

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"math/big"
	"time"
)

const ScanApiKeyTypeId = "fae9587e-21c6-4b92-bc74-837f1ce977c3"

type configScanApiKey struct {
	IntervalScanCall uint32 `json:"intervalScanCall"` //重试调用ScanCall出错的间隔
	IntervalBlock    uint32 `json:"intervalBlock"`    //重试取得区块出错的间隔
	IntervalTx       uint32 `json:"intervalTx"`       //重试取交易出错的间隔
	SafeBlockDiffer  uint32 `json:"safeBlockDiffer"`  //在遍历区块时，只遍历 最大区块- SafeHeightDiffer的区块， 这样是为了安全考虑，这个值如果是公链的话是 6
	ApiKey           string `json:"apiKey"`           //cn.etherscan.com 的key
	Url              string `json:"url"`              //
}
type ScanApiKey struct {
	ScanCall ScanCall `dot:""`

	ethConnect      EthClientI
	conf            configScanApiKey
	stopped         atomic.Bool
	stopChanel      chan bool
	initBlockNumber *big.Int
}

func (c *ScanApiKey) startScan() {
	logger := dot.Logger()
	c.initBlockNumber = big.NewInt(-1)

	for { //get the  init block number
		if c.stopped.Load() {
			return
		}
		t := c.ScanCall.StartBlockNumber()
		if t.Cmp(big.NewInt(0)) >= 0 {
			c.initBlockNumber.Set(t)
			break
		}
		select {
		case <-c.stopChanel:
			return
		case <-time.After(time.Second * time.Duration(c.conf.IntervalScanCall)):
		}
	}

	cur := big.NewInt(0)
	cur.Add(c.initBlockNumber, big.NewInt(1))
	eth := c.ethConnect
	max := big.NewInt(0)

	ctx, _ := context.WithCancel(context.Background())

ForMaxBlocks:
	for { //取最大区块号，遍历区块，取区块中的所有交易
		for { //get the max block number
			if c.stopped.Load() {
				break ForMaxBlocks
			}
			b, err := eth.BlockByNumber(ctx, nil)
			if err == nil {
				max.Sub(b.Header().Number, big.NewInt(int64(c.conf.SafeBlockDiffer)))
				break
			} else {
				logger.Debugln("ScanWeb3", zap.Error(err))
			}
			select { //等待后再重试
			case <-c.stopChanel:
				break ForMaxBlocks
			case <-time.After(time.Second * time.Duration(c.conf.IntervalBlock)):
			}
		}

		if cur.Cmp(max) >= 0 { //如果没有新的区块; 这个条件不要放到 ForBlocks， 因为那样会每一次完成都等 2s
			select { //等待后再取最大区块号
			case <-c.stopChanel:
				break ForMaxBlocks
			case <-time.After(time.Second * time.Duration(c.conf.IntervalBlock)):
			}
			continue ForMaxBlocks
		}

	ForBlocks:
		for { //从当前一直到max区块号。不能在这里++,及作判断，因为cur是多次使用的
			var b *types.Block
			var err error
			for {
				if c.stopped.Load() {
					break ForMaxBlocks
				}
				b, err = eth.BlockByNumber(ctx, cur)
				if err == nil && b != nil {
					break
				} else {
					logger.Debugln("ScanWeb3", zap.Error(err))
				}
				select { //等待后再重试
				case <-c.stopChanel:
					break ForMaxBlocks
				case <-time.After(time.Second * time.Duration(c.conf.IntervalBlock)):
				}
			}

			for !c.ScanCall.Block(b) { //一直到返回true, 可能没有准备好
				select { //等待后再重试
				case <-c.stopChanel:
					break ForMaxBlocks
				case <-time.After(time.Second * time.Duration(c.conf.IntervalScanCall)):
				}
			}

			for _, htx := range b.Transactions() { //处理区块中的交易
				if c.stopped.Load() {
					break ForMaxBlocks
				}
				for {
					tx, _, err := eth.TransactionByHash(ctx, htx.Hash())
					receipt, err2 := eth.TransactionReceipt(ctx, htx.Hash())
					if err == nil && err2 == nil && tx != nil && receipt != nil {
						for !c.ScanCall.Tx(b, tx, receipt) { //一直到返回true, 可能没有准备好
							select { //等待后再重试
							case <-c.stopChanel:
								break ForMaxBlocks
							case <-time.After(time.Second * time.Duration(c.conf.IntervalScanCall)):
							}
						}
						break
					} else {
						logger.Debugln("ScanWeb3", zap.Error(err))
						logger.Debugln("ScanWeb3", zap.Error(err2))
					}

					select { //等待后再重试
					case <-c.stopChanel:
						break ForMaxBlocks
					case <-time.After(time.Second * time.Duration(c.conf.IntervalTx)):
					}
				}
			}

			for !c.ScanCall.DoneBlock(b) { //一直到返回true, 可能没有准备好
				select { //等待后再重试
				case <-c.stopChanel:
					break ForMaxBlocks
				case <-time.After(time.Second * time.Duration(c.conf.IntervalScanCall)):
				}
			}

			if cur.Cmp(max) >= 0 { //检测是否已到最大区块
				break ForBlocks
			}
			cur.Add(cur, big.NewInt(1))
		}
	}
}

//func (c *ScanApiKey) Create(l dot.Line) error {
//	//todo add
//}
//func (c *ScanApiKey) Injected(l dot.Line) error {
//	//todo add
//}
//func (c *ScanApiKey) AfterAllInject(l dot.Line) {
//	//todo add
//}

func (c *ScanApiKey) Start(ignore bool) error {
	c.stopped.Store(false)
	if c.ScanCall == nil {
		return errors.New("the ScanCall is nil")
	}
	go c.startScan()
	return nil
}

func (c *ScanApiKey) Stop(ignore bool) error {
	return nil
}

//func (c *ScanApiKey) Destroy(ignore bool) error {
//	//todo add
//}

//construct dot
func newScanApiKey(conf []byte) (dot.Dot, error) {
	dconf := &configScanApiKey{}
	err := dot.UnMarshalConfig(conf, dconf)
	if err != nil {
		return nil, err
	}
	d := &ScanApiKey{conf: *dconf}
	d.ethConnect = NewEthClientApiKeyImp(d.conf.Url, d.conf.ApiKey, "")
	return d, nil
}

//ScanApiKeyTypeLives
func ScanApiKeyTypeLives() []*dot.TypeLives {
	tl := &dot.TypeLives{
		Meta: dot.Metadata{TypeId: ScanApiKeyTypeId, NewDoter: func(conf []byte) (dot.Dot, error) {
			return newScanApiKey(conf)
		}},
		//Lives: []dot.Live{
		//	{
		//		LiveId:    ScanApiKeyTypeId,
		//		RelyLives: map[string]dot.LiveId{"some field": "some id"},
		//	},
		//},
	}

	lives := []*dot.TypeLives{tl}

	return lives
}

//ScanApiKeyConfigTypeLive
func ScanApiKeyConfigTypeLive() *dot.ConfigTypeLives {
	paths := make([]string, 0)
	paths = append(paths, "")
	return &dot.ConfigTypeLives{
		TypeIdConfig: ScanApiKeyTypeId,
		ConfigInfo:   &configScanApiKey{
			//todo
		},
	}
}
