// Scry Info.  All rights reserved.
// license that can be found in the license file.

package chainevents

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/scryinfo/dp/demo/src/sdk/core/ethereum/events"
)

// EventCallback
type EventCallback func(event events.Event) bool

// EventRepository
type EventRepository struct {
	mapEventSubscribe map[string]map[common.Address]EventCallback
}

// NewEventRepository
func NewEventRepository() *EventRepository {
	return &EventRepository{
		mapEventSubscribe: make(map[string]map[common.Address]EventCallback),
	}
}
