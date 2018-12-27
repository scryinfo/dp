package chainevents

import (
	"fmt"
	"github.com/qjpcpu/ethereum/events"
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

	for {
		select {
		// block until data exist in dataChannel
		case event := <- dataChannel:
			eventUsers := event.Data.Get("users").(string)
			if eventUsers == "" {
				fmt.Println("Error: no users in event data")
				break
			}

			curUser, err := usermanager.GetCurrentUser()
			if err != nil {
				fmt.Println("Error: failed to get current user.")
				break
			}

			if strings.Contains(eventUsers, curUser.GetPublicKey()) {
				executeEvent(event, internalEventRepo)
				executeEvent(event, externalEventRepo)
			}

			break
		}
	}
}

func executeEvent(event events.Event, eventRepo *EventRepository) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error: failed to execute event " + event.Name + " because of error: ", err)
		}
	}()

	eventCallback, exist := eventRepo.mapEventExecutor[event.Name]
	if !exist {
		fmt.Println("error: event not subscribed: " + event.Name)
		return false
	}

	return eventCallback(event)
}