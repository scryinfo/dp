// Scry Info.  All rights reserved.
// license that can be found in the license file.

package app

import (
	"github.com/scryinfo/dp/dots/app/connection"
	sdkinterface2 "github.com/scryinfo/dp/dots/app/sdkinterface"
	settings2 "github.com/scryinfo/dp/dots/app/settings"
	"github.com/scryinfo/dp/dots/binary/sdk/scry"
)

var app Gapp

//todo decrease the number of global variables
type Gapp struct {
	ChainWrapper scry.ChainWrapper
	Deployer     scry.Client
	CurUser      sdkinterface2.SDKWrapper
	ScryInfo     *settings2.ScryInfo
	Connection   *connection.WSServer
}

func GetGapp() *Gapp {
	return &app
}
