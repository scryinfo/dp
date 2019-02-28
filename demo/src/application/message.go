package main

import (
	"encoding/json"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/scryinfo/iscap/demo/src/sdk/scryclient"
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
	case "create.new.account":
		var pwd = Password{}
		err = json.Unmarshal(m.Payload, &pwd)
		if err != nil {
			break
		}
		payload, err = scryclient.CreateScryClient(pwd.Password)
		return payload, nil
	case "login.verify":
		payload = true
		return payload, nil
	case "save.keystore":
		payload = true
		return payload, nil
	case "get.datalist":
		payload = []Datalist{
			{"Qm461", "title1", 1, "test tags461", "test description461", "0x1234567890123456789012345678901234567890"},
			{"Qm462", "title2", 2, "test tags462", "test description462", "0x1234567890123456789012345678901234567890"},
			{"Qm463", "title3", 3, "test tags463", "test description463", "0x1234567890123456789012345678901234567890"},
			{"Qm464", "title4", 4, "test tags464", "test description464", "0x1234567890123456789012345678901234567890"},
			{"Qm465", "title5", 5, "test tags465", "test description465", "0x1234567890123456789012345678901234567890"},
			{"Qm466", "title6", 6, "test tags466", "test description466", "0x1234567890123456789012345678901234567890"},
			{"Qm467", "title7", 7, "test tags467", "test description467", "0x1234567890123456789012345678901234567890"},
			{"Qm468", "title8", 8, "test tags468", "test description468", "0x1234567890123456789012345678901234567890"},
			{"Qm469", "title9", 9, "test tags469", "test description469", "0x1234567890123456789012345678901234567890"},
			{"Qm4610", "title10", 10, "test tags4610", "test description4610", "0x1234567890123456789012345678901234567890"},
			{"Qm4611", "title11", 11, "test tags4611", "test description4611", "0x1234567890123456789012345678901234567890"},
			{"Qm4612", "title12", 12, "test tags4612", "test description4612", "0x1234567890123456789012345678901234567890"},
			{"Qm4613", "title13", 13, "test tags4613", "test description4613", "0x1234567890123456789012345678901234567890"},
		}
		return payload, nil
	case "get.transaction":
		payload = []Transaction{
			{"title1", 1, "0x1234567890123456789012345678901234567890", "0x1524783212578655202365479511235413256752", Created},
		}
		return payload, nil
	case "buy":
		payload = true
		return payload, nil
	case "publish":
		var dl PubData = PubData{}
		if err = json.Unmarshal(m.Payload, &dl);err != nil {
			break
		}
		if payload, err = SellerPublishData(dl);err != nil {
			break
		}
		return payload, nil
	}

	payload = err.Error()
	return payload, err
}

type Password struct { Password string `json:"password"` }

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
	Price     string   `json:"Price"`
	Seller    string   `json:"Owner"`
}

type Transaction struct {
	Title         string
	TransactionID int
	Seller        string
	Buyer         string
	State         byte
}
