package chainevents

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/iscap/demo/src/sdk/core/ethereum/events"
)

type EventCallback func(event events.Event) bool

type EventRepository struct {
	mapEventSubscribe map[string]map[common.Address]EventCallback
}

func NewEventRepository() *EventRepository {
	return &EventRepository{
		mapEventSubscribe: make(map[string]map[common.Address]EventCallback),
	}
}
