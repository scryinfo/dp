// Scry Info.  All rights reserved.
// license that can be found in the license file.

package uuid

import (
	"github.com/chilts/sid"
)

// GenerateUUID
func GenerateUUID() string {
	return sid.Id()
}
