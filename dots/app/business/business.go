package business

import (
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/app/business/preset"
	"github.com/scryinfo/dp/dots/binary"
	"go.uber.org/zap"
)

// Business encapsulate system functions
type Business struct {
	Pre *preset.Preset `dot:""`
}

// const
const (
	BusTypeId = "64a3ff50-50de-447c-b0b9-401fff8c4fa4"
	BusLiveId = "64a3ff50-50de-447c-b0b9-401fff8c4fa4"
)

// Start load preset msg handlers and start web server
func (b *Business) Start(ignore bool) error {
	if err := b.Pre.CBs.WS.PresetMsgHandleFuncs(b.Pre.PresetMsgNames, b.Pre.PresetMsgHandlers); err != nil {
		return err
	}

	if err := b.Pre.CBs.WS.ListenAndServe(); err != nil {
		dot.Logger().Errorln("Start http web server failed. ", zap.NamedError("error", err))
		return errors.New("Start http web server failed. ")
	}

	return nil
}

//construct dot
func newBusDot(_ interface{}) (dot.Dot, error) {
	var err error
	d := &Business{}

	return d, err
}

// BusTypeLive add a dot component to dot.line with 'line.PreAdd()'
func BusTypeLive() []*dot.TypeLives {
	t := []*dot.TypeLives{
		{
			Meta: dot.Metadata{
				TypeId: BusTypeId,
				NewDoter: func(conf interface{}) (dot.Dot, error) {
					return newBusDot(conf)
				},
			},
			Lives: []dot.Live{{
				LiveId:    BusLiveId,
				RelyLives: map[string]dot.LiveId{"binary": binary.BinLiveId},
			}},
		},
	}

	t = append(t, preset.PreTypeLive()...)

	return t
}
