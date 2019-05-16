package app

import (
	"github.com/scryinfo/dp/app/app/sdkinterface"
	"github.com/scryinfo/dp/app/app/settings"
	"github.com/scryinfo/dp/dots/binary/sdk/scry"
)

var app Gapp

//todo 减少全局变量的个数
type Gapp struct {
	ChainWrapper scry.ChainWrapper
	Deployer     scry.Client
	CurUser      sdkinterface.SDKWrapper
	ScryInfo     *settings.ScryInfo
}

func GetGapp() *Gapp {
	return &app
}
