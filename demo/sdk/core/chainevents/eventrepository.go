package chainevents

import "../ethereum/events"

type EventCallback func(event events.Event) bool

type EventRepository struct {
	//mapEventExecutor map[string][]SubscribeInfo
	mapEventSubscribe map[string]map[string]EventCallback
}

func NewEventRepository() (*EventRepository) {
	return &EventRepository{
		mapEventSubscribe: make(map[string]map[string]EventCallback),
	}
}