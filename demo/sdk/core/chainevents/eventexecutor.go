package chainevents

import (
	"../ethereum/events"
	"fmt"
)

func ExecuteEvents(dataChannel chan events.Event,
	internalEventRepo *EventRepository, externalEventRepo *EventRepository) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: failed to execute event, error: ", err)
		}
	}()

	for {
		select {
		case event := <- dataChannel:
			eventUsers := event.Data.Get("users").(string)
			if eventUsers == "" {
				fmt.Println("Error: no users in event data")
				break
			}

			executeEvent(event, internalEventRepo)
			executeEvent(event, externalEventRepo)
		}
	}
}

func executeEvent(event events.Event, eventRepo *EventRepository) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error: failed to execute event " + event.Name + " because of error: ", err)
		}
	}()

	callbackInfo, exist := eventRepo.mapEventExecutor[event.Name]
	if !exist {
		fmt.Println("error: event not subscribed: " + event.Name)
		return false
	}



	return eventCallback(event)
}