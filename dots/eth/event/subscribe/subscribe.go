package subscribe

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
)

const (
    SubsTypeId = "5535a065-0d90-46f4-9776-26630676c4c5"
)

type Subscribe struct {
    eventRepo *Repository
    config    subscribeConfig
}

type subscribeConfig struct {
}

func (c *Subscribe) Create(l dot.Line) error {
    c.eventRepo = NewRepository()
    return nil
}

//construct dot
func newSubsDot(conf interface{}) (dot.Dot, error) {
    var err error
    var bs []byte
    if bt, ok := conf.([]byte); ok {
        bs = bt
    } else {
        return nil, dot.SError.Parameter
    }

    dConf := &subscribeConfig{}
    err = dot.UnMarshalConfig(bs, dConf)
    if err != nil {
        return nil, err
    }

    d := &Subscribe{config: *dConf}

    return d, err
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
    eventCallback Callback,
) error {
    if eventCallback == nil || eventName == "" {
        return errors.New("couldn't subscribe event because of null eventCallback or empty event name")
    }

    subscribeInfoMap := c.eventRepo.mapEventCallback[eventName]
    if subscribeInfoMap == nil {
        subscribeInfoMap = make(map[common.Address]Callback)
        c.eventRepo.mapEventCallback[eventName] = subscribeInfoMap
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

    subscribeInfoMap := c.eventRepo.mapEventCallback[eventName]
    if subscribeInfoMap == nil || subscribeInfoMap[clientAddr] == nil {
        return errors.New("couldn't find corresponding event to unsubscribe:" + eventName)
    }

    delete(subscribeInfoMap, clientAddr)
    if len(subscribeInfoMap) == 0 {
        subscribeInfoMap = nil
        delete(c.eventRepo.mapEventCallback, eventName)
    }

    return nil
}
