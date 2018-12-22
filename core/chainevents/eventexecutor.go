package chainevents

import (
	"fmt"
	"github.com/qjpcpu/ethereum/events"
)

func ExecuteEvents(dataChannel chan events.Event,
	internalEventRepo *EventRepository, externalEventRepo *EventRepository) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("failed to execute event, error: ", err)
		}
	}()

	for {
		select {
		// block until data exist in dataChannel
		case event := <- dataChannel:
			executeEvent(event, internalEventRepo)
			executeEvent(event, externalEventRepo)
			break
		}
	}
}

func executeEvent(event events.Event, eventRepo *EventRepository) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("failed to execute event " + event.Name + " because of error: ", err)
		}
	}()

	eventCallback, exist := eventRepo.mapEventExecutor[event.Name]
	if !exist {
		fmt.Println("event not subscribed: " + event.Name)
		return false
	}

	return eventCallback(event)
}