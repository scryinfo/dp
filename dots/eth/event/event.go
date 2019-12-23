// Scry Info.  All rights reserved.
// license that can be found in the license file.

package event

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// Event define event structure go scanned from tx.log
type Event struct {
	BlockNumber uint64
	TxHash      common.Hash
	Address     common.Address
	Name        string
	Data        JSONObj
}

// Progress progress
type Progress struct {
	From uint64
	To   uint64
}

func (evt Event) String() string {
	return fmt.Sprintf(
		`block: %v,tx: %s,address: %s,event: %s,data: %s`,
		evt.BlockNumber,
		evt.TxHash.Hex(),
		evt.Address.Hex(),
		evt.Name,
		evt.Data.String(),
	)
}
