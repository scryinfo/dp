// Scry Info.  All rights reserved.
// license that can be found in the license file.

package uuid

import (
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	id := GenerateUUID()
	if id == "" {
		t.Fail()
	} else {
		t.Log("GenerateUUID Successed! id:", id)
	}
}
