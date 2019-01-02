package chainevents

import "../ethereum/events"

type EventCallback func(event events.Event) bool

type SubscribeInfo struct {
	clientAddr string
	eventCallback EventCallback
}

type SubscribeInfoList struct {
	subscribeInfos [](*SubscribeInfo)
}

type EventRepository struct {
	mapEventExecutor map[string]*SubscribeInfoList
}

func NewEventRepository() (*EventRepository) {
	return &EventRepository{
		mapEventExecutor: make(map[string]*SubscribeInfoList),
	}
}