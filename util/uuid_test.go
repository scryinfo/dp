package util

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
