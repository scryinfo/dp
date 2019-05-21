// Scry Info.  All rights reserved.
// license that can be found in the license file.

package main

import (
	"fmt"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/line"
	app2 "github.com/scryinfo/dp/dots/app"
	sdkinterface2 "github.com/scryinfo/dp/dots/app/sdkinterface"
	"github.com/scryinfo/dp/dots/app/settings"
	"github.com/scryinfo/dp/dots/app/websocket"
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
		return false //退出
	})

}

func Init(l dot.Line) (err error) {
	logger := dot.Logger()
	conf := &settings.ScryInfo{}
	l.SConfig().UnmarshalKey("app", conf)
	app2.GetGapp().ScryInfo = conf
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

	websocket.MessageHandlerInit()

	if err = websocket.ConnectWithProtocolWebsocket(conf.Config.WSPort); err != nil { //todo do not block
		logger.Errorln("", zap.NamedError("WebSocket Connect failed. ", err))
	}

	return err
}
