package chainevents

import (
	"errors"
	"fmt"
	"github.com/qjpcpu/ethereum/events"
)

var (
	maxChannelEventNum     = 10000
	internalEventRepo              = NewEventRepository()
	externalEventRepo              = NewEventRepository()
	dataChannel            = make(chan events.Event, maxChannelEventNum)
	errorChannel           = make(chan error, 1)
)

func StartEventProcessing()  {
	fmt.Println("start event processing...")

	//configuration

	go ExecuteEvents(dataChannel, internalEventRepo, externalEventRepo)

	go ListenEvent("", "", "", "", 60, dataChannel, errorChannel)

	fmt.Println("finished event processing.")
}

func SubscribeInternal(eventName string, eventCallback EventCallback) error {
	return subscribe(eventName, eventCallback, internalEventRepo)
}

func SubscribeExternal(eventName string, eventCallback EventCallback) error {
	return subscribe(eventName, eventCallback, externalEventRepo)
}

func subscribe(eventName string, eventCallback EventCallback, eventRepo *EventRepository) error {
	if eventCallback == nil {
		return errors.New("couldn't subscribe event as eventCallback is null. ")
	}

	eventRepo.mapEventExecutor[eventName] = eventCallback

	return nil
}
