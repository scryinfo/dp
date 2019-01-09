package chainevents

import (
	"../ethereum/events"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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
			broadcast := event.Data.Get("boardcast").(bool)
			eventUsers := event.Data.Get("users").([]common.Address)
			if !broadcast && len(eventUsers) == 0 {
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


	broadcast := event.Data.Get("boardcast").(bool)
	if broadcast {
		for _, v := range subscribeInfoMap {
			if v != nil {
				EventCallback(v)(event)
			}
		}
	} else {
		eventUsers := event.Data.Get("users").([]common.Address)
		for k, v := range subscribeInfoMap {
			if containUser(eventUsers, k) {
				if v != nil {
					EventCallback(v)(event)
				}
			}
		}
	}

	return true
}

func containUser(userList []common.Address, user common.Address) bool {
	for _, u := range userList {
		if u == user {
			return true
		}
	}

	return false
}