package definition

import (
	"github.com/scryinfo/dot/dot"
)

// Account accounts
type Account struct {
	Address    string `gorm:"primary_key"`
	Nickname   string
	FromBlock  int64
	IsVerifier bool
	Verify     []byte // []string, json serialize
	Arbitrate  []byte // []string, json serialize
}

// DataList data_lists
type DataList struct {
	PublishId           string `gorm:"primary_key"`
	Title               string
	Price               string
	Keys                string
	Description         string
	Seller              string
	SupportVerify       bool
	MetaDataExtension   string `json:"-"`
	ProofDataExtensions []byte `json:"-"` // []string, json serialize

	CreatedTime int64 `json:"-"`
}

// Transaction transactions
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
	ProofDataExtensions []byte `json:"-"` // []string, json serialize

	Identify int

	CreatedTime int64 `json:"-"`
}

// Event events
type Event struct {
	Id          int    `gorm:"primary_key"`
	NotifyTo    []byte // []string, json serialize
	EventName   string
	EventBody   []byte // json string, json serialize
	EventStatus int // 0: unread, 1: already read

	CreatedTime int64 `json:"-"`
}

// AfterUpdate show use for hooks
func (dl *DataList) AfterUpdate() error {
	dot.Logger().Infoln("-------Show Use For Hooks (After Update).-------")
	return nil
}
