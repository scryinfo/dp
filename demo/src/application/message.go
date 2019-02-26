package main

import (
	"encoding/json"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"math/big"
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
		err     error = nil
	)

	switch m.Name {
	case "get.accounts":
		payload = []string{"0x123456789012345678901234567890", "0x152478321257865520236547951123"}
		return payload, err
	case "create.new.account":
		payload = "0x123123123123123123123123123123"
		return payload, err
	case "login.verify":
		payload = true
		return payload, err
	case "save.keystore":
		payload = true
		return payload, err
	case "get.datalist":
		payload = []Datalist{
			{"Qm461", "title1", 1, "test tags461", "test description461", "0x123456789012345678901234567890"},
			{"Qm462", "title2", 2, "test tags462", "test description462", "0x123456789012345678901234567890"},
			{"Qm463", "title3", 3, "test tags463", "test description463", "0x123456789012345678901234567890"},
			{"Qm464", "title4", 4, "test tags464", "test description464", "0x123456789012345678901234567890"},
			{"Qm465", "title5", 5, "test tags465", "test description465", "0x123456789012345678901234567890"},
			{"Qm466", "title6", 6, "test tags466", "test description466", "0x123456789012345678901234567890"},
			{"Qm467", "title7", 7, "test tags467", "test description467", "0x123456789012345678901234567890"},
			{"Qm468", "title8", 8, "test tags468", "test description468", "0x123456789012345678901234567890"},
			{"Qm469", "title9", 9, "test tags469", "test description469", "0x123456789012345678901234567890"},
			{"Qm4610", "title10", 10, "test tags4610", "test description4610", "0x123456789012345678901234567890"},
			{"Qm4611", "title11", 11, "test tags4611", "test description4611", "0x123456789012345678901234567890"},
			{"Qm4612", "title12", 12, "test tags4612", "test description4612", "0x123456789012345678901234567890"},
			{"Qm4613", "title13", 13, "test tags4613", "test description4613", "0x123456789012345678901234567890"},
		}
		return payload, err
	case "get.transaction":
		payload = []Transaction{
			{"title1", 1, "0x123456789012345678901234567890", "0x152478321257865520236547951123", Created},
		}
		return payload, err
	case "buy":
		payload = true
		return payload, err
	case "publish":
		var dl PubData = PubData{}
		err = json.Unmarshal(m.Payload, &dl)
		if err != nil {
			break
		}
		payload, iserr := SellerPublishData(dl)
		if !iserr {
			break
		}
		return payload, err
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

type PubData struct {
	MetaData  string   `json:"Data"`
	ProofData []string `json:"Proofs"`
	DespData  string   `json:"Description"`
	Price     *big.Int `json:"Price"`
	Seller    string   `json:"Owner"`
}

type Transaction struct {
	Title         string
	TransactionID int
	Seller        string
	Buyer         string
	State         byte
}
