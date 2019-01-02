package chainevents

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"../ethereum/events"
)

var (
	maxChannelEventNum     = 10000
	internalEventRepo              = NewEventRepository()
	externalEventRepo              = NewEventRepository()
	dataChannel            = make(chan events.Event, maxChannelEventNum)
	errorChannel           = make(chan error, 1)
	settingPath            = "../../settings/setting.yaml"
)

func StartEventProcessing(conn *ethclient.Client,
	                      protocolContractAddr string,
	                      protocolContractABI string)  {
	fmt.Println("start event processing...")

	go ExecuteEvents(dataChannel, internalEventRepo, externalEventRepo)

	go ListenEvent(conn, protocolContractAddr, protocolContractABI, "Published", 60, dataChannel, errorChannel)

	fmt.Println("finished event processing.")
}

func SubscribeInternal(eventName string, eventCallback EventCallback) error {
	return subscribe("", eventName, eventCallback, internalEventRepo)
}

func SubscribeExternal(clientAddr string, eventName string, eventCallback EventCallback) error {
	return subscribe(clientAddr, eventName, eventCallback, externalEventRepo)
}

func subscribe(clientAddr string, eventName string,
					eventCallback EventCallback, eventRepo *EventRepository) error {
	if eventCallback == nil || clientAddr == "" || eventName == "" {
		return errors.New("couldn't subscribe event because of null eventCallback or empty client address or empty event name")
	}

	subscribeInfo := SubscribeInfo{
		clientAddr: clientAddr,
		eventCallback: eventCallback,
	}

	subscribeInfoList := eventRepo.mapEventExecutor[eventName]
	if subscribeInfoList == nil {
		subscribeInfoList = &SubscribeInfoList{
			subscribeInfos: make(*SubscribeInfo, 5)
		}
	}


	return nil
}
