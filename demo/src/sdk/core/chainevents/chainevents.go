package chainevents

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
	rlog "github.com/sirupsen/logrus"
)

var (
	maxChannelEventNum = 10000
	externalEventRepo  = NewEventRepository()
	dataChannel        = make(chan events.Event, maxChannelEventNum)
	errorChannel       = make(chan error, 1)
	settingPath        = "../../settings/definition.yaml"
)

func StartEventProcessing(conn *ethclient.Client,
                            contracts []ContractInfo,
                            fromBlock uint64) {
	rlog.Info("start event processing...")

	go ExecuteEvents(dataChannel, externalEventRepo)
	go ListenEvent(conn, contracts, fromBlock, 60, dataChannel, errorChannel)

	rlog.Info("finished event processing.")
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
