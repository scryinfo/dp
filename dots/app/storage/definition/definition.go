package definition

import (
	"fmt"
	"time"
)

type Account struct {
	Address    string `gorm:"primary_key"`
	NickName   string
	FromBlock  int64
	IsVerifier bool
	Verify     []byte // []string, json serialize
}

type DataList struct {
	PublishId           string `gorm:"primary_key"`
	Title               string
	Price               string
	Keys                string
	Description         string
	Seller              string
	SupportVerify       bool
	MetaDataExtension   string
	ProofDataExtensions []byte // []string, json serialize

	CreatedAt time.Time // from gorm
}

type Transaction struct {
	TransactionId               string `gorm:"primary_key"`
	Buyer                       string
	State                       int
	StartVerify                 bool
	MetaDataIdEncWithSeller     string
	MetaDataIdEncWithBuyer      string
	MetaDataIdEncWithArbitrator string
	Verifier1Response           string
	Verifier2Response           string
	ArbitrateResult             bool

	PublishId           string
	Title               string
	Price               string
	Keys                string
	Description         string
	Seller              string
	SupportVerify       bool
	MetaDataExtension   string
	ProofDataExtensions []byte // []string, json serialize

	CreatedAt time.Time // from gorm
}

type Event struct {
	Id           int    `gorm:"primary_key"`
	NotifyTo     []byte // []string, json serialize
	EventName    int    // enum
	EventKeyword string
	EventPayload string

	CreatedAt time.Time // from gorm
}

// show use for hooks
func (dl *DataList) AfterCreate() error {
	fmt.Println("-------Show Use For Hooks (After Create).-------")
	return nil
}

func (dl *DataList) AfterUpdate() error {
	fmt.Println("-------Show Use For Hooks (After Update).-------")
	return nil
}
