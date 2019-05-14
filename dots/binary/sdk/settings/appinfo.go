package settings

import (
	"github.com/scryInfo/dp/util"
)

var (
	appId     = "Scry-App-" + util.GenerateUUID()
)

func SetAppId(v string) {
	appId = v
}

func GetAppId() string {
	return appId
}
