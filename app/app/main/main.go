package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/scryInfo/dot/dot"
	"github.com/scryInfo/dot/dots/line"
	"github.com/scryInfo/dp/app/app"
	"github.com/scryInfo/dp/app/app/sdkinterface"
	settings2 "github.com/scryInfo/dp/app/app/settings"
	WSConnect2 "github.com/scryInfo/dp/app/app/websocket"
	sdk2 "github.com/scryInfo/dp/dots/binary/sdk"
	"github.com/scryInfo/scryg/sutils/ssignal"
	rlog "github.com/sirupsen/logrus"
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

	ssignal.WatiCtrlC(func(s os.Signal) bool {
		return false //退出
	})

	line.StopAndDestroy(l, true)
}

func Init(l dot.Line) error {
	logger := dot.Logger()
	var err error = nil
	conf := & settings2.ScryInfo{}
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
		logger.Errorln("", zap.Any("", err))
	}
	l.ToInjecter().ReplaceOrAddByType(app.GetGapp().ChainWrapper)

	logger.Infoln("inited ChainWrapper")

	WSConnect2.SetCurUser(sdkinterface.NewSDKWrapperImp())
	WSConnect2.MessageHandlerInit()

	if err = WSConnect2.ConnectWithProtocolWebsocket(conf.Config.WSPort); err != nil { //todo do not block
		rlog.Error(errors.Wrap(err, "WebSocket Connect failed. "))
	}

	return err
}
