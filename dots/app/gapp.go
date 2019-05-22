// Scry Info.  All rights reserved.
// license that can be found in the license file.

package app

import (
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
}

func GetGapp() *Gapp {
	return &app
}
