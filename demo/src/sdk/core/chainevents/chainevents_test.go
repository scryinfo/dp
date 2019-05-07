package chainevents

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/scryInfo/dp/demo/src/sdk/core/ethereum/events"
    "testing"
)

func TestUnSubscribeExternal(t *testing.T) {
    addr1 := common.HexToAddress("0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8")
    addr2 := common.HexToAddress("0xd280b60c38bc8db9d309fa5a540ffec499f0a3e9")
    evtName1 := "testEvent1"
    evtName2 := "testEvent2"

    repo := NewEventRepository()

    subMap1 := make(map[common.Address]EventCallback)
    subMap1[addr1] = callback
    subMap1[addr2] = callback

    subMap2 := make(map[common.Address]EventCallback)
    subMap2[addr1] = callback
    subMap2[addr2] = callback

    repo.mapEventSubscribe[evtName1] = subMap1
    repo.mapEventSubscribe[evtName2] = subMap2

    unsubscribe(addr1, evtName1, repo)

    if len(repo.mapEventSubscribe[evtName1]) != 1 {
        t.Fail()
    }

    if len(repo.mapEventSubscribe[evtName2]) != 2 {
        t.Fail()
    }

    unsubscribe(addr2, evtName1, repo)

    if len(repo.mapEventSubscribe[evtName1]) != 0 {
        t.Fail()
    }
}

func callback(event events.Event) bool {
    return true
}