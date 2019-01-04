package chainevents

import (
	"../ethereum/events"
	"fmt"
	"strings"
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
			fmt.Println("error: failed to execute event "+event.Name+" because of error: ", err)
		}
	}()

	subscribeInfoMap := eventRepo.mapEventSubscribe[event.Name]
	if subscribeInfoMap == nil {
		fmt.Println("Warning: no event was executed")
		return false
	}

	users := event.Data.Get("users").(string)
	if users == "" {
		fmt.Println("Warning: users domain is empty in event data")
		return false
	} else if users == "*" {
		for _, v := range subscribeInfoMap {
			if v != nil {
				EventCallback(v)(event)
			}
		}
	} else {
		for k, v := range subscribeInfoMap {
			if strings.Contains(users, k) {
				if v != nil {
					EventCallback(v)(event)
				}
			}
		}
	}

	return true
}