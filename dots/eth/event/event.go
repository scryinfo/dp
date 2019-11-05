// Scry Info.  All rights reserved.
// license that can be found in the license file.

package event

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// ModifyContractParam
type Event struct {
	BlockNumber uint64
	TxHash      common.Hash
	Address     common.Address
	Name        string
	Data        JSONObj
}

// Progress
type Progress struct {
	From uint64
	To   uint64
}

// String
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
