// Scry Info.  All rights reserved.
// license that can be found in the license file.

package main

import (
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/line"
	app2 "github.com/scryinfo/dp/dots/app"
	"github.com/scryinfo/dp/dots/app/connection"
	"github.com/scryinfo/dp/dots/app/connection/msg_handler"
	sdkinterface2 "github.com/scryinfo/dp/dots/app/sdkinterface"
	"github.com/scryinfo/dp/dots/app/settings"
	"github.com/scryinfo/dp/dots/binary"
	"github.com/scryinfo/scryg/sutils/ssignal"
	"go.uber.org/zap"
	"os"
)

func main() {
	l, err := line.BuildAndStart(func(l dot.Line) error {
		//todo
		l.PreAdd(binary.BinTypeLive()...)
		l.PreAdd(settings.ConfTypeLive())
		return nil
	})
	if err != nil {
		dot.Logger().Debugln("Line init failed. ", zap.NamedError("", err))
		return
	}

	if err := Init(l); err != nil {
		dot.Logger().Errorln("App init failed. ", zap.NamedError("err", err))
		return
	}

	dot.SetDefaultLine(l)

	defer line.StopAndDestroy(l, true)

	ssignal.WaitCtrlC(func(s os.Signal) bool {
		return false //quit
	})
}

func Init(l dot.Line) (err error) {
	logger := dot.Logger()

	{
		d, err := l.ToInjecter().GetByLiveId(dot.LiveId(binary.BinLiveId))
		if err != nil {
			logger.Errorln("load Binary component failed.")
			return nil
		}

		if g, ok := d.(*binary.Binary); ok {
			app2.GetGapp().ChainWrapper = g.ChainWrapper()
		} else {
			logger.Errorln("load Binary component failed.", zap.Any("d", d))
			return nil
		}
	}

	l.ToInjecter().ReplaceOrAddByType(app2.GetGapp().ChainWrapper)

	logger.Infoln("ChainWrapper init finished. ")

	{
		d, err := l.ToInjecter().GetByLiveId(dot.LiveId(settings.ConfLiveId))
		if err != nil {
			logger.Errorln("load Config component failed.")
			return nil
		}

		if g, ok := d.(*settings.Config); ok {
			app2.GetGapp().ScryInfo = g
		} else {
			logger.Errorln("load Config component failed.", zap.Any("d", d))
			return nil
		}
	}

	app2.GetGapp().Connection = connection.CreateConnetion(app2.GetGapp().ScryInfo.WSPort, app2.GetGapp().ScryInfo.UIResourcesDir)

	app2.GetGapp().CurUser = sdkinterface2.CreateSDKWrapperImp(app2.GetGapp().ChainWrapper)

	msg_handler.MessageHandlerInit()

	if err = app2.GetGapp().Connection.Connect(); err != nil {
		logger.Errorln("WebSocket Connect failed. ", zap.NamedError("", err))
	}

	return
}
