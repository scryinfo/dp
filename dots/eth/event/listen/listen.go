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
	// ListenerTypeId
	ListenerTypeId = "9ff2cb44-e73a-4a53-add4-3166954983d7"
)

// Listener listen event
type Listener struct {
	builder *Builder
}

//construct dot
func newListenerDot(conf interface{}) (dot.Dot, error) {
	var err error
	d := &Listener{}

	return d, err
}

// ListenerTypeLive add a dot component to dot.line with 'line.PreAdd()'
func ListenerTypeLive() *dot.TypeLives {
	return &dot.TypeLives{
		Meta: dot.Metadata{TypeId: ListenerTypeId,
			NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
				return newListenerDot(conf)
			}},
	}
}

// Create
func (c *Listener) Create(l dot.Line) error {
	c.builder = NewScanBuilder()
	return nil
}

// ListenEvent listen event
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
			logger.Errorln("failed to listen event, panic error", zap.Any("", er))
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
		logger.Errorln("failed to listen to events", zap.Error(err))
		return false
	}

	r.WaitChan()

	return true
}

// SetFromBlock
func (c *Listener) SetFromBlock(from uint64) {
	if c.builder != nil {
		c.builder.SetFrom(from)
	} else {
		dot.Logger().Warnln("failed to set from block because of null builder")
	}
}
