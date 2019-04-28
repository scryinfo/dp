package main

import (
	"flag"
	"github.com/pkg/errors"
	"github.com/scryinfo/iscap/demo/src/application/WSConnect"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface"
	"github.com/scryinfo/iscap/demo/src/application/sdkinterface/settings"
	"github.com/scryinfo/iscap/demo/src/sdk"
	rlog "github.com/sirupsen/logrus"
)

const logPath = "D:/EnglishRoad/workspace/Go/src/github.com/scryinfo/iscap/demo/src/log/scry_sdk.log"

var (
	scryInfo *settings.ScryInfo
	err      error = nil
	port           = flag.String("port", "9822", "")
)

func init() {
	scryInfo, err = settings.LoadSettings()
	err = sdk.Init(scryInfo.Chain.Ethereum.EthNode,
		scryInfo.Services.Keystore,
		scryInfo.Chain.Contracts.ProtocolAddr,
		scryInfo.Chain.Contracts.TokenAddr,
		scryInfo.Services.Ipfs,
		logPath,
		"App demo",
	)
	sdkinterface.SetScryInfo(scryInfo)
}

func main() {
	flag.Parse()
	if err = WSConnect.WebsocketConnect(*port); err != nil {
		rlog.Error(errors.Wrap(err, "WebSocket Connect failed. "))
	}
}
