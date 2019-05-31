// Scry Info.  All rights reserved.
// license that can be found in the license file.

package subscribe

import (
	"github.com/ethereum/go-ethereum/common"
	events2 "github.com/scryinfo/dp/dots/binary/sdk/core/ethereum/events"
)

type Callback func(event events2.Event) bool

type Repository struct {
	mapEventCallback map[string]map[common.Address]Callback
}

func NewRepository() *Repository {
	return &Repository{
        mapEventCallback: make(map[string]map[common.Address]Callback),
	}
}
