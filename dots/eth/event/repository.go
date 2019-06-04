// Scry Info.  All rights reserved.
// license that can be found in the license file.

package event

import (
	"github.com/ethereum/go-ethereum/common"
)

type Callback func(event Event) bool

type Repository struct {
	MapEventCallback map[string]map[common.Address]Callback
}

func NewRepository() *Repository {
	return &Repository{
        MapEventCallback: make(map[string]map[common.Address]Callback),
	}
}
