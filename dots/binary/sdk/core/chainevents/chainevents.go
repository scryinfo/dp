// Scry Info.  All rights reserved.
// license that can be found in the license file.

package chainevents

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryinfo/dot/dot"
	events2 "github.com/scryinfo/dp/dots/binary/sdk/core/ethereum/events"
)

var (
	maxChannelEventNum = 10000
	externalEventRepo  = NewEventRepository()
	dataChannel        = make(chan events2.Event, maxChannelEventNum)
	errorChannel       = make(chan error, 1)
)

func StartEventProcessing(
	conn *ethclient.Client,
	contracts []ContractInfo,
) {
	dot.Logger().Infoln("start event processing...")

	go ExecuteEvents(dataChannel, externalEventRepo)
	go ListenEvent(conn, contracts, 0, 60, dataChannel, errorChannel)

	dot.Logger().Infoln("finished event processing.")
}

func SubscribeExternal(
	clientAddr common.Address,
	eventName string,
	eventCallback EventCallback,
) error {
	return subscribe(clientAddr, eventName, eventCallback, externalEventRepo)
}

func subscribe(
	clientAddr common.Address,
	eventName string,
	eventCallback EventCallback,
	eventRepo *EventRepository,
) error {
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

func UnSubscribeExternal(
	clientAddr common.Address,
	eventName string,
) error {
	return unsubscribe(clientAddr, eventName, externalEventRepo)
}

func unsubscribe(
	clientAddr common.Address,
	eventName string,
	eventRepo *EventRepository,
) error {
	if eventName == "" {
		return errors.New("couldn't unsubscribe event because of empty event name")
	}

	subscribeInfoMap := eventRepo.mapEventSubscribe[eventName]
	if subscribeInfoMap == nil || subscribeInfoMap[clientAddr] == nil {
		return errors.New("couldn't find corresponding event to unsubscribe:" + eventName)
	}

	delete(subscribeInfoMap, clientAddr)
	if len(subscribeInfoMap) == 0 {
		subscribeInfoMap = nil
		delete(eventRepo.mapEventSubscribe, eventName)
	}

	return nil
}
