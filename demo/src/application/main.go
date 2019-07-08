// Scry Info.  All rights reserved.
// license that can be found in the license file.

package main

import (
    "flag"
    "github.com/pkg/errors"
    "github.com/scryinfo/dp/demo/src/application/WSConnect"
    "github.com/scryinfo/dp/demo/src/application/sdkinterface"
    "github.com/scryinfo/dp/demo/src/application/sdkinterface/settings"
    "github.com/scryinfo/dp/demo/src/sdk"
    rlog "github.com/sirupsen/logrus"
)

const logPath = "D:/EnglishRoad/workspace/Go/src/github.com/scryinfo/dp/demo/src/log/scry_sdk.log"

var (
    scryinfo *settings.scryinfo
    err      error = nil
    port           = flag.String("port", "9822", "")
)

func init() {
    scryinfo, err = settings.LoadSettings()
    err = sdk.Init(scryinfo.Chain.Ethereum.EthNode,
        scryinfo.Services.Keystore,
        scryinfo.Chain.Contracts.ProtocolAddr,
        scryinfo.Chain.Contracts.TokenAddr,
        scryinfo.Services.Ipfs,
        logPath,
        "App demo",
    )
    sdkinterface.SetScryInfo(scryinfo)
}

func main() {
    flag.Parse()
    if err = WSConnect.WebsocketConnect(*port); err != nil {
        rlog.Error(errors.Wrap(err, "WebSocket Connect failed. "))
    }
}
