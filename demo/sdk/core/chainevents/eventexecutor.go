package chainevents

import (
	"fmt"
	"../ethereum/events"
	"strings"
	"../../util/usermanager"
)

func ExecuteEvents(dataChannel chan events.Event,
	internalEventRepo *EventRepository, externalEventRepo *EventRepository) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: failed to execute event, error: ", err)
		}
	}()

	fmt.Println("start execute events")
	for {
		select {
		// block until data exist in dataChannel
		case event := <- dataChannel:
			fmt.Println("channel len:", len(dataChannel))
			eventUsers := event.Data.Get("users").(string)
			fmt.Println(event.Data)
			if eventUsers == "" {
				fmt.Println("Error: no users in event data")
				break
			}

			if eventUsers != "*" {
				curUser, err := usermanager.GetCurrentUser()
				if err != nil {
					fmt.Println("Error: failed to get current user.")
					break
				}

				if !strings.Contains(eventUsers, curUser.GetPublicKey()) {
					break
				}
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

	fmt.Println("start execute event")
	eventCallback, exist := eventRepo.mapEventExecutor[event.Name]
	if !exist {
		fmt.Println("error: event not subscribed: " + event.Name)
		return false
	}

	return eventCallback(event)
}