package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "Account":
		if string(m.Payload) == "administrator" {
			payload = "0x123456789012345678901234567890"
			return
		}
	}
	payload = err.Error()
	return
}
