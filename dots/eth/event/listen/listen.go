// Scry Info.  All rights reserved.
// license that can be found in the license file.

package listen

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/eth/event"
    "go.uber.org/zap"
    "time"
)

const (
	ListenerTypeId = "9ff2cb44-e73a-4a53-add4-3166954983d7"
)

type Listener struct {
	builder *Builder
	config  listenerConfig
}

type listenerConfig struct {
}

//construct dot
func newListenerDot(conf interface{}) (dot.Dot, error) {
	var err error
	var bs []byte
	if bt, ok := conf.([]byte); ok {
		bs = bt
	} else {
		return nil, dot.SError.Parameter
	}

	dConf := &listenerConfig{}
	err = dot.UnMarshalConfig(bs, dConf)
	if err != nil {
		return nil, err
	}

	d := &Listener{config: *dConf}

	return d, err
}

//Data structure needed when generating newer component
func ListenerTypeLive() *dot.TypeLives {
	return &dot.TypeLives{
		Meta: dot.Metadata{TypeId: ListenerTypeId,
			NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
				return newListenerDot(conf)
			}},
	}
}

func (c *Listener) Create(l dot.Line) error {
	c.builder = NewScanBuilder()
	return nil
}



func (c *Listener) ListenEvent(
	conn *ethclient.Client,
	contracts []event.ContractInfo,
	fromBlock uint64,
	interval time.Duration,
	dataChannel chan event.Event,
	errorChannel chan error,
) bool {
	logger := dot.Logger()
	logger.Infoln("start listening events...")

	defer func() {
		if er := recover(); er != nil {
			logger.Errorln("", zap.Any("Failed to listen event. error:", er))
		}
	}()

	if len(contracts) == 0 {
		logger.Errorln("invalid contracts parameter")
		return false
	}

	for _, v := range contracts {
		c.builder.SetContract(common.HexToAddress(v.Address), v.Abi, v.Events...)
	}

	r, err := c.builder.SetClient(conn).
		SetFrom(fromBlock).
		SetTo(0).
		SetGracefulExit(true).
		SetDataChan(dataChannel, errorChannel).
		SetInterval(interval).
		BuildAndRun()
	if err != nil {
		logger.Errorln("", zap.NamedError("failed to listen to events.", err))
		return false
	}

	r.WaitChan()

	return true
}

func (c *Listener) SetFromBlock(from uint64) {
	if c.builder != nil {
		c.builder.SetFrom(from)
	} else {
		dot.Logger().Warnln("Failed to set from block because of nil builder.")
	}
}
