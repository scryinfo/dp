package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/scryInfo/dot/dot"
	_ "github.com/scryInfo/dot/dot"
	"github.com/scryInfo/dot/dots/line"
	"github.com/scryInfo/dp/app/app"
	sdkinterface2 "github.com/scryInfo/dp/app/app/sdkinterface"
	settings2 "github.com/scryInfo/dp/app/app/settings"
	WSConnect2 "github.com/scryInfo/dp/app/app/wsconnect"
	sdk2 "github.com/scryInfo/dp/dots/binary/sdk"
	"github.com/scryInfo/scryg/sutils/ssignal"
	rlog "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)



func main()  {
	l, err := line.BuildAndStart(func(l dot.Line) error {
		//todo
		return Init()
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

func Init() error {
	setFile, logFile := "","" //todo config
	if ex, err := os.Executable(); err == nil {
		exPath := filepath.Dir(ex)
		setFile = filepath.Join(exPath, "definition.yaml")
		logFile = filepath.Join(exPath,"log.log")
	}

	scryInfo, err := settings2.LoadSettings(setFile)
	if err != nil {
		return err
	}
	app.GetGapp().ScryInfo = scryInfo
	//todo
	app.GetGapp().ChainWrapper, err = sdk2.Init(scryInfo.Chain.Ethereum.EthNode,
		scryInfo.Services.Keystore,
		scryInfo.Chain.Contracts.ProtocolAddr,
		scryInfo.Chain.Contracts.TokenAddr,
		scryInfo.Services.Ipfs,
		logFile,"App demo",
	)
	sdkinterface2.SetScryInfo(scryInfo)
	if err = WSConnect2.WebsocketConnect(scryInfo.Services.Wsport); err != nil {
		rlog.Error(errors.Wrap(err, "WebSocket Connect failed. "))
	}
	return err
}