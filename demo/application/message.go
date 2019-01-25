package main

import (
	"encoding/json"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "hello":
		var s string
		err = json.Unmarshal(m.Payload,&s)
		if err != nil {
			payload = err.Error()
			return
		}

		if s == "message from js" {
			payload = "message from go"
			return
		}
	}
	payload = err.Error()
	return
}
