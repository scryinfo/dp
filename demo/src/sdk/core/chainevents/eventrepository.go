package chainevents

import (
	"sdk/core/ethereum/events"
	"github.com/ethereum/go-ethereum/common"
)

type EventCallback func(event events.Event) bool

type EventRepository struct {
	mapEventSubscribe map[string]map[common.Address]EventCallback
}

func NewEventRepository() (*EventRepository) {
	return &EventRepository{
		mapEventSubscribe: make(map[string]map[common.Address]EventCallback),
	}
}