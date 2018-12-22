package transactions

import "github.com/ethereum/go-ethereum/common"

const (
	Start int = iota
	Created
	Voted
	Purchasing
	ReadyForDownload
	Arbitrating
	Payed
	Closed
)


type Transaction struct {
	id common.Hash
	state int
	publishId string
}

func NewTransaction(txId common.Hash) (*Transaction)  {
	return &Transaction{
		id: txId,
		state: Start,
	}
}
