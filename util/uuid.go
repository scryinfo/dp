package util

import (
	"github.com/chilts/sid"
)

func GenerateUUID() string {
	return sid.Id()
}
