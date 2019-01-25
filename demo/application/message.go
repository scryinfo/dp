package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (interface{}, error) {
	var (
		payload interface{}
		err     error
	)

	switch m.Name {
	case "get.accounts":
		payload = []string{"0x123456789012345678901234567890","0x152478321257865520236547951123"}
		return payload,nil
	case "login.verify":
		payload = true
		return payload,nil
	case "save.keystroe":
		payload = true
		return payload,nil
	}

	payload = err.Error()
	return payload, err
}
