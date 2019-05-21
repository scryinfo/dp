package main

import (
	"fmt"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/line"
	"github.com/scryinfo/dp/app/app"
	"github.com/scryinfo/dp/app/app/sdkinterface"
	settings2 "github.com/scryinfo/dp/app/app/settings"
	WSConnect2 "github.com/scryinfo/dp/app/app/websocket"
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
	conf := &settings2.ScryInfo{}
	l.SConfig().UnmarshalKey("app", conf)
	app.GetGapp().ScryInfo = conf
	//todo
	app.GetGapp().ChainWrapper, err = sdk2.Init(
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
	l.ToInjecter().ReplaceOrAddByType(app.GetGapp().ChainWrapper)

	logger.Infoln("ChainWrapper init finished. ")

	app.GetGapp().CurUser = sdkinterface.CreateSDKWrapperImp(app.GetGapp().ChainWrapper, app.GetGapp().ScryInfo)

	WSConnect2.MessageHandlerInit()

	if err = WSConnect2.ConnectWithProtocolWebsocket(conf.Config.WSPort); err != nil { //todo do not block
		logger.Errorln("", zap.NamedError("WebSocket Connect failed. ", err))
	}

	return err
}
