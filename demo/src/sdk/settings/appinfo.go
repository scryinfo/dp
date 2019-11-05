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

// SetAppId
func SetAppId(v string) {
	appId = v
}

// GetAppId
func GetAppId() string {
	return appId
}

// SetLogPath
func SetLogPath(d string) {
	logPath = d
}

// GetLogPath
func GetLogPath() string {
	return logPath
}
