package chainevents

import (
	"fmt"
	"github.com/qjpcpu/ethereum/events"
)

func ExecuteEvents(dataChannel chan events.Event, eventRepo *EventRepository) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Failed to execute event, error: ", err)
		}
	}()

	for {
		select {
		//expected: will block until has data in dataChannel
		case event := <- dataChannel:
			executeEvent(event, eventRepo)
			break
		}
	}
}

func executeEvent(event events.Event, eventRepo *EventRepository) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Failed to execute event " + event.Name + " because of error: ", err)
		}
	}()

	eventCallback, exist := eventRepo.mapEventExecutor[event.Name]
	if !exist {
		fmt.Println("Event not subscribed: " + event.Name)
		return false
	}

	return eventCallback(event)
}