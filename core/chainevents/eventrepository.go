package chainevents

import "github.com/qjpcpu/ethereum/events"

type EventCallback func(event events.Event) bool

type EventRepository struct {
	mapEventExecutor map[string]EventCallback
}


func NewEventRepository() (*EventRepository) {
	return &EventRepository{
		mapEventExecutor: make(map[string]EventCallback),
	}
}