package definition

import (
	"fmt"
	"github.com/scryinfo/dot/dot"
	"time"
)

type Account struct {
	Address    string `gorm:"primary_key"`
	Nickname   string
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
	MetaDataExtension   string `json:"-"`
	ProofDataExtensions []byte  `json:"-"`// []string, json serialize

	CreatedAt time.Time `json:"-"` // from gorm
}

type Transaction struct {
	TransactionId               string `gorm:"primary_key"`
	Buyer                       string
	State                       int
	StartVerify                 bool
	MetaDataIdEncWithSeller     string `json:"-"`
	MetaDataIdEncWithBuyer      string `json:"-"`
	MetaDataIdEncWithArbitrator string `json:"-"`
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
	MetaDataExtension   string `json:"-"`
	ProofDataExtensions []byte   `json:"-"`// []string, json serialize

	Identify int

	CreatedAt time.Time `json:"-"` // from gorm
}

type Event struct {
	Id           int    `gorm:"primary_key"`
	NotifyTo     []byte // []string, json serialize
	EventName    int    // enum(iota)
	EventKeyword string
	EventPayload string

	CreatedAt time.Time `json:"-"`// from gorm
}

// show use for hooks
func (dl *DataList) AfterUpdate() error {
	fmt.Println("-------Show Use For Hooks (After Update).-------")
	dot.Logger().Infoln("-------Show Use For Hooks (After Update).-------")
	return nil
}
