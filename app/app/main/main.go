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
	"github.com/scryinfo/dp/dots/storage"
	"github.com/scryinfo/scryg/sutils/ssignal"
	"go.uber.org/zap"
	"os"
)

func main() {
	l, err := line.BuildAndStart(func(l dot.Line) error {
		//todo
        l.PreAdd(storage.IpfsTypeLive())
        l.PreAdd(binary.BinTypeLive())
        return Init(l)
	})

	if err != nil {
		dot.Logger().Debugln("Line init failed. ", zap.NamedError("", err))
		return
	}

	dot.SetDefaultLine(l)

	defer line.StopAndDestroy(l, true)

	ssignal.WatiCtrlC(func(s os.Signal) bool {
		return false //quit
	})
}

func Init(l dot.Line) (err error) {
	logger := dot.Logger()
	conf := &settings.ScryInfo{}
	l.SConfig().UnmarshalKey("app", conf)
	app2.GetGapp().ScryInfo = conf
	app2.GetGapp().Connection = connection.CreateConnetion(conf.Config.WSPort, conf.Config.UIResourcesDir)

	d, err := l.ToInjecter().GetByLiveId(dot.LiveId(binary.BinLiveId))
    if err != nil {
        logger.Errorln("load Binary component failed.")
        return nil
    }

    if g, ok := d.(*binary.Binary); ok {
        app2.GetGapp().ChainWrapper = g.ChainWrapper()
    } else {
        logger.Errorln("load Binary component failed.")
        return nil
    }

	l.ToInjecter().ReplaceOrAddByType(app2.GetGapp().ChainWrapper)

	logger.Infoln("ChainWrapper init finished. ")

	app2.GetGapp().CurUser = sdkinterface2.CreateSDKWrapperImp(app2.GetGapp().ChainWrapper, conf)

	msg_handler.MessageHandlerInit()

	if err = app2.GetGapp().Connection.Connect(); err != nil {
		logger.Errorln("WebSocket Connect failed. ", zap.NamedError("", err))
	}

	return
}
