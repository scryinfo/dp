// Scry Info.  All rights reserved.
// license that can be found in the license file.

package subscribe

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/eth/event"
	"sync"
)

const (
	// SubsTypeId
	SubsTypeId = "5535a065-0d90-46f4-9776-26630676c4c5"
)

// Subscribe contains a event repository
type Subscribe struct {
	eventRepo *event.Repository
}

// SetRepo set repo
func (c *Subscribe) SetRepo(r *event.Repository) {
	c.eventRepo = r
}

//construct dot
func newSubsDot(conf interface{}) (dot.Dot, error) {
	d := &Subscribe{}

	return d, nil
}

// SubsTypeLive Data structure needed when generating newer component
func SubsTypeLive() *dot.TypeLives {
	return &dot.TypeLives{
		Meta: dot.Metadata{TypeId: SubsTypeId,
			NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
				return newSubsDot(conf)
			}},
	}
}

// Subscribe subscribe
func (c *Subscribe) Subscribe(
	clientAddr common.Address,
	eventName string,
	eventCallback event.Callback,
) error {
	if eventCallback == nil || eventName == "" {
		return errors.New("couldn't subscribe event because of null eventCallback or empty event name")
	}

	var subscribeInfoMap sync.Map
	if rv, ok := c.eventRepo.MapEventCallback.Load(eventName); !ok {
		subscribeInfoMap.Store(clientAddr, eventCallback)
		c.eventRepo.MapEventCallback.Store(eventName, subscribeInfoMap)
	} else {
		subscribeInfoMap = rv.(sync.Map)
		subscribeInfoMap.Store(clientAddr, eventCallback)
	}

	return nil
}

// UnSubscribe unsubscribe
func (c *Subscribe) UnSubscribe(
	clientAddr common.Address,
	eventName string,
) error {
	if eventName == "" {
		return errors.New("couldn't unsubscribe event because of empty event name")
	}

	var subscribeInfoMap sync.Map
	if rv, ok := c.eventRepo.MapEventCallback.Load(eventName); ok {
		subscribeInfoMap = rv.(sync.Map)
		if _, ok = subscribeInfoMap.Load(clientAddr); !ok {
			return errors.New("couldn't find corresponding event to unsubscribe:" + clientAddr.String())
		}
	} else {
		return errors.New("couldn't find corresponding client to unsubscribe:" + eventName)
	}

	subscribeInfoMap.Delete(clientAddr)
	if getMapLen(subscribeInfoMap) == 0 {
		c.eventRepo.MapEventCallback.Delete(eventName)
	}

	return nil
}

func getMapLen(m sync.Map) int {
	l := 0
	m.Range(func(key, value interface{}) bool {
		l++
		return true
	})

	return l
}
