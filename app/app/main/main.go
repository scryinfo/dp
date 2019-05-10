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
	dir, _ := os.Getwd() // todo: is it necessary to handle error here? read origin code choose a time.
	scryInfo, err := settings2.LoadSettings(dir + "/config.yaml")
	if err != nil {
		return err
	}

	app.GetGapp().ScryInfo = scryInfo
	//todo
	app.GetGapp().ChainWrapper, err = sdk2.Init(
		scryInfo.Chain.Ethereum.EthNode,
		scryInfo.Chain.Contracts.ProtocolAddr,
		scryInfo.Chain.Contracts.TokenAddr,
		scryInfo.Services.Keystore,
		scryInfo.Services.Ipfs,
		scryInfo.Config.LogPath,
		scryInfo.Config.AppId,
	)
	l.ToInjecter().ReplaceOrAddByType(app.GetGapp().ChainWrapper)

	WSConnect2.SetCurUser(sdkinterface.NewSDKWrapperImp())
	WSConnect2.MessageHandlerInit()

	if err = WSConnect2.ConnectWithProtocolWebsocket(scryInfo.Config.WSPort); err != nil { //todo do not block
		rlog.Error(errors.Wrap(err, "WebSocket Connect failed. "))
	}

	return err
}
