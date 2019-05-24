// Scry Info.  All rights reserved.
// license that can be found in the license file.

package main

import (
	"fmt"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/line"
	app2 "github.com/scryinfo/dp/dots/app"
	"github.com/scryinfo/dp/dots/app/connection"
	"github.com/scryinfo/dp/dots/app/connection/msg_handler"
	sdkinterface2 "github.com/scryinfo/dp/dots/app/sdkinterface"
	"github.com/scryinfo/dp/dots/app/settings"
	sdk2 "github.com/scryinfo/dp/dots/binary/sdk"
	"github.com/scryinfo/scryg/sutils/ssignal"
	"go.uber.org/zap"
	"os"
)

func main() {
	l, err := line.BuildAndStart(func(l dot.Line) error {
		//todo
		return Init(l)
	})

	if err != nil {
		fmt.Println(err)
		return
	}

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

	//todo
	app2.GetGapp().ChainWrapper, err = sdk2.Init(
		conf.Chain.Ethereum.EthNode,
		conf.Chain.Contracts.ProtocolAddr,
		conf.Chain.Contracts.TokenAddr,
		conf.Services.Keystore,
		conf.Services.Ipfs,
		conf.Config.AppId,
	)
	if err != nil {
		logger.Errorln("", zap.NamedError("", err))
	}
	l.ToInjecter().ReplaceOrAddByType(app2.GetGapp().ChainWrapper)

	logger.Infoln("ChainWrapper init finished. ")

	app2.GetGapp().CurUser = sdkinterface2.CreateSDKWrapperImp(app2.GetGapp().ChainWrapper, app2.GetGapp().ScryInfo)

	msg_handler.MessageHandlerInit()

	if err = app2.GetGapp().Connection.Connect(); err != nil {
		logger.Errorln("WebSocket Connect failed. ", zap.NamedError("", err))
	}

	logger.Infoln("Connect finished. ")

	return err
}
