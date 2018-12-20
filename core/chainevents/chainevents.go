package chainevents

import (
	"errors"
	"fmt"
	"github.com/qjpcpu/ethereum/events"
)

var (
	maxChannelEventNum     = 10000
	eventRepo              = NewEventRepository()
	dataChannel            = make(chan events.Event, maxChannelEventNum)
	errorChannel           = make(chan error, 1)
)

func StartEventProcessing()  {
	fmt.Println("Start event processing...")

	//configuration

	//event execution routine
	go ExecuteEvents(dataChannel, eventRepo)

	//event listener routine
	go ListenEvent("", "", "", "", 60, dataChannel, errorChannel)

	fmt.Println("Finished event processing.")
}

func Subscribe(eventName string, eventCallback EventCallback) error {
	if eventCallback == nil {
		return errors.New("Couldn't subscribe event as eventCallback is null. ")
	}

	eventRepo.mapEventExecutor[eventName] = eventCallback

	return nil
}
