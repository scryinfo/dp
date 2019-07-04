// Scry Info.  All rights reserved.
// license that can be found in the license file.

package subscribe

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/eth/event"
)

const (
	SubsTypeId = "5535a065-0d90-46f4-9776-26630676c4c5"
)

type Subscribe struct {
	eventRepo *event.Repository
}

func (c *Subscribe) Create(l dot.Line) error {
	return nil
}

func (c *Subscribe) SetRepo(r *event.Repository) {
	c.eventRepo = r
}

//construct dot
func newSubsDot(conf interface{}) (dot.Dot, error) {
	d := &Subscribe{}

	return d, nil
}

//Data structure needed when generating newer component
func SubsTypeLive() *dot.TypeLives {
	return &dot.TypeLives{
		Meta: dot.Metadata{TypeId: SubsTypeId,
			NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
				return newSubsDot(conf)
			}},
	}
}

func (c *Subscribe) Subscribe(
	clientAddr common.Address,
	eventName string,
	eventCallback event.Callback,
) error {
	if eventCallback == nil || eventName == "" {
		return errors.New("couldn't subscribe event because of null eventCallback or empty event name")
	}

	subscribeInfoMap := c.eventRepo.MapEventCallback[eventName]
	if subscribeInfoMap == nil {
		subscribeInfoMap = make(map[common.Address]event.Callback)
		c.eventRepo.MapEventCallback[eventName] = subscribeInfoMap
	}

	subscribeInfoMap[clientAddr] = eventCallback

	return nil
}

func (c *Subscribe) UnSubscribe(
	clientAddr common.Address,
	eventName string,
) error {
	if eventName == "" {
		return errors.New("couldn't unsubscribe event because of empty event name")
	}

	subscribeInfoMap := c.eventRepo.MapEventCallback[eventName]
	if subscribeInfoMap == nil || subscribeInfoMap[clientAddr] == nil {
		return errors.New("couldn't find corresponding event to unsubscribe:" + eventName)
	}

	delete(subscribeInfoMap, clientAddr)
	if len(subscribeInfoMap) == 0 {
		subscribeInfoMap = nil
		delete(c.eventRepo.MapEventCallback, eventName)
	}

	return nil
}
