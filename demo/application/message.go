package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

const (
	Created = byte(iota)
	Voted
	Payed
	ReadyForDownload
	Closed
)

func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (interface{}, error) {
	var (
		payload interface{}
		err     error
	)

	switch m.Name {
	case "get.accounts":
		payload = []string{"0x123456789012345678901234567890", "0x152478321257865520236547951123"}
		return payload, nil
	case "login.verify":
		payload = true
		return payload, nil
	case "save.keystore":
		payload = true
		return payload, nil
	case "get.datalist":
		payload = []Datalist{
			{"Qm461", "title1", 100, "tag1,tag2,tag3", "test description", "0x123456789012345678901234567890"},
		}
		return payload, nil
	case "get.transaction":
		payload = []Transaction{
			{"title1", 1, "0x123456789012345678901234567890", "0x152478321257865520236547951123", Created},
		}
		return payload, nil
	case "buy":
		payload = true
		return payload, nil
	case "publish":
		payload = true
		return payload, nil
	}

	payload = err.Error()
	return payload, err
}

type Datalist struct {
	ID          string
	Title       string
	Price       int
	Keys        string
	Description string
	Owner       string
}

type Transaction struct {
	Title         string
	TransactionID int
	Seller        string
	Buyer         string
	State         byte
}
