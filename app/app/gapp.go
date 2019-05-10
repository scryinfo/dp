package app

import (
	"github.com/scryInfo/dp/app/app/settings"
	"github.com/scryInfo/dp/dots/binary/sdk/scry"
)

var app Gapp

//todo 减少全局变量的个数
type Gapp struct {
	ChainWrapper scry.ChainWrapper
	Deployer     scry.Client
	ScryInfo     *settings.ScryInfo
}

func GetGapp() *Gapp {
	return &app
}
