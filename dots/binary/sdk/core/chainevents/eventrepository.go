// Scry Info.  All rights reserved.
// license that can be found in the license file.

package chainevents

import (
	"github.com/ethereum/go-ethereum/common"
	events2 "github.com/scryinfo/dp/dots/binary/sdk/core/ethereum/events"
)

type EventCallback func(event events2.Event) bool

type EventRepository struct {
	mapEventSubscribe map[string]map[common.Address]EventCallback
}

func NewEventRepository() *EventRepository {
	return &EventRepository{
		mapEventSubscribe: make(map[string]map[common.Address]EventCallback),
	}
}
