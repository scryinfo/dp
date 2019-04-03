package chainevents

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/scryinfo/iscap/demo/src/sdk/core/ethereum/events"
    "github.com/scryinfo/iscap/demo/src/sdk/settings"
    rlog "github.com/sirupsen/logrus"
)

const (
    BROADCAST_TO_USERS = "0x00"
    TARGET_USERS       = "users"
    TARGET_OWNER       = "_owner"
    APP_SEQ_NO             = "seqNo"
    TOKEN_EVT_APPROVAL = "Approval"
)

func ExecuteEvents(dataChannel chan events.Event, externalEventRepo *EventRepository) {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("Error: failed to execute event, error: ", err)
		}
	}()

	for {
		select {
		case event := <-dataChannel:
			rlog.Debug("event coming:", event.String())
			executeEvent(event, externalEventRepo)
		}
	}
}

func executeEvent(event events.Event, eventRepo *EventRepository) bool {
    defer func() {
        if err := recover(); err != nil {
            rlog.Error("error: failed to execute event "+event.Name+" because of error: ", err)
        }
    }()

    subscribeInfoMap := eventRepo.mapEventSubscribe[event.Name]
    if subscribeInfoMap == nil {
        rlog.Error("warning: no event was executed, event:", event.Name)
        return false
    }

    seqNo := event.Data.Get(APP_SEQ_NO)
    if seqNo != settings.GetAppId() && event.Name != TOKEN_EVT_APPROVAL {
        return true
    }

    objUsers := event.Data.Get(TARGET_USERS)
    if objUsers != nil {
        users := objUsers.([]common.Address)
        if len(users) == 1 && users[0] == common.HexToAddress(BROADCAST_TO_USERS) {
            executeAllEvent(subscribeInfoMap, event)
        } else {
            executeMatchedEvent(subscribeInfoMap, users, event)
        }

    } else {
        obj, ok := event.Data.Get(TARGET_OWNER).(string)
        if ok {
            owner := common.HexToAddress(obj)
            executeMatchedEvent(subscribeInfoMap, []common.Address{owner}, event)
        } else {
            rlog.Error("Warning: unknown event type, event:", event.Name)
        }
    }

    return true
}

func executeMatchedEvent(subscribeInfoMap map[common.Address]EventCallback,
    users []common.Address, event events.Event) {
    for k, v := range subscribeInfoMap {
        if containUser(users, k) {
            if v != nil {
                EventCallback(v)(event)
            }
        }
    }
}

func executeAllEvent(subscribeInfoMap map[common.Address]EventCallback, event events.Event) {
    for _, v := range subscribeInfoMap {
        EventCallback(v)(event)
    }
}

func containUser(userList []common.Address, user common.Address) bool {
    for _, u := range userList {
        if u == user {
            return true
        }
    }

    return false
}
