// Scry Info.  All rights reserved.
// license that can be found in the license file.

package settings

import (
	"github.com/scryinfo/dp/demo/src/sdk/util/uuid"
	"os"
)

var (
	appId     = "Scry-App-" + uuid.GenerateUUID()
	logDir, _ = os.Getwd()
	logPath   = logDir + "scry_sdk.log"
)

func SetAppId(v string) {
	appId = v
}

func GetAppId() string {
	return appId
}

func SetLogPath(d string) {
	logPath = d
}

func GetLogPath() string {
	return logPath
}
