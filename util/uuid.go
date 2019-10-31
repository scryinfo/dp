// Scry Info.  All rights reserved.
// license that can be found in the license file.

package util

import (
	"github.com/chilts/sid"
)

func GenerateUUID() string {
	return sid.Id()
}
