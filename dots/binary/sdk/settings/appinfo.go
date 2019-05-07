package settings

import (
	"github.com/scryInfo/dp/util"
	"os"
)

var (
    appId = "Scry-App-" + util.GenerateUUID()
    logDir, _ = os.Getwd()
    logPath = logDir + "scry_sdk.log"
)

func SetAppId(v string)  {
    appId = v
}

func GetAppId() string {
    return appId
}

func SetLogPath(d string)  {
    logPath = d
}

func GetLogPath() string {
    return logPath
}

