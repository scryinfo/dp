// Scry Info.  All rights reserved.
// license that can be found in the license file.

package settings

import (
	"github.com/scryinfo/dp/util"
)

var appId = "Scry-App-" + util.GenerateUUID()

func SetAppId(v string) {
	appId = v
}

func GetAppId() string {
	return appId
}
