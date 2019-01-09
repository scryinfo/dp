package chainevents

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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

	go ListenEvent(conn, protocolContractAddr, protocolContractABI, 60, dataChannel, errorChannel, "Publish", "TransactionCreate", "Purchase", "ReadyForDownload", "Close")

	fmt.Println("finished event processing.")
}

func SubscribeInternal(eventName string, eventCallback EventCallback) error {
	return subscribe(common.HexToAddress("0x00"), eventName, eventCallback, internalEventRepo)
}

func SubscribeExternal(clientAddr common.Address, eventName string, eventCallback EventCallback) error {
	return subscribe(clientAddr, eventName, eventCallback, externalEventRepo)
}

func subscribe(clientAddr common.Address, eventName string,
					eventCallback EventCallback, eventRepo *EventRepository) error {
	if eventCallback == nil || eventName == "" {
		return errors.New("couldn't subscribe event because of null eventCallback or empty event name")
	}

	subscribeInfoMap := eventRepo.mapEventSubscribe[eventName]
	if subscribeInfoMap == nil {
		subscribeInfoMap = make(map[common.Address]EventCallback)
		eventRepo.mapEventSubscribe[eventName] = subscribeInfoMap
	}

	subscribeInfoMap[clientAddr] = eventCallback

	return nil
}
