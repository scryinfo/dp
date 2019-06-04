// Scry Info.  All rights reserved.
// license that can be found in the license file.

package execute

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/eth/event"
    "go.uber.org/zap"
)

const (
    ExecTypeId       = "4313210a-5824-4ff1-8dd8-c71ccad711db"
    BroadcastToAll   = "0x00"
    TargetUsers      = "users"
    TargetOwner      = "owner"
    AppSeqNo         = "seqNo"
    TokenEvtApproval = "Approval"
)

type Executor struct {
    eventChan chan event.Event
    repo      *event.Repository
    config    executorConfig
}

type executorConfig struct {
    appId string
}

//construct dot
func newExecutorDot(conf interface{}) (dot.Dot, error) {
    var err error
    var bs []byte

    if bt, ok := conf.([]byte); ok {
        bs = bt
    } else {
        return nil, dot.SError.Parameter
    }

    dConf := &executorConfig{}
    err = dot.UnMarshalConfig(bs, dConf)
    if err != nil {
        return nil, err
    }

    d := &Executor{config: *dConf}

    return d, err
}

//Data structure needed when generating newer component
func ExecutorTypeLive() *dot.TypeLives {
    return &dot.TypeLives{
        Meta: dot.Metadata{TypeId: ExecTypeId,
            NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
                return newExecutorDot(conf)
            }},
    }
}


func (c *Executor) Create(l dot.Line) error {
    return nil
}

func (c *Executor) SetEventChan(ce chan event.Event)  {
    c.eventChan = ce
}

func (c *Executor) SetSubsRepo(r *event.Repository)  {
    c.repo = r
}


func (c *Executor) ExecuteEvents() {
    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("Error: failed to execute event, error: ", er))
        }
    }()

    for {
        select {
        case e := <-c.eventChan:
            dot.Logger().Debugln("event coming:" + e.String())
            c.executeEvent(e)
        }
    }
}

func (c *Executor) executeEvent(e event.Event) bool {
    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("error: failed to execute e "+e.Name+" because of error: ", er))
        }
    }()

    subs := c.repo.MapEventCallback[e.Name]
    if subs == nil {
        dot.Logger().Warnln("warning: no e was executed, e:" + e.Name)
        return false
    }

    seqNo := e.Data.Get(AppSeqNo)
    if seqNo != c.config.appId && e.Name != TokenEvtApproval {
        return true
    }

    objUsers := e.Data.Get(TargetUsers)
    if objUsers != nil {
        users := objUsers.([]common.Address)
        if len(users) == 1 && users[0] == common.HexToAddress(BroadcastToAll) {
            c.executeAllEvent(subs, e)
        } else {
            c.executeMatchedEvent(subs, users, e)
        }

    } else {
        obj, ok := e.Data.Get(TargetOwner).(string)
        if ok {
            owner := common.HexToAddress(obj)
            c.executeMatchedEvent(subs, []common.Address{owner}, e)
        } else {
            dot.Logger().Warnln("Warning: unknown e type, e:" + e.Name)
        }
    }

    return true
}

func (c *Executor) executeMatchedEvent(
        sim map[common.Address]event.Callback,
        users []common.Address, e event.Event,
    ) {
    for k, v := range sim {
        if c.containUser(users, k) {
            if v != nil {
                event.Callback(v)(e)
            }
        }
    }
}

func (c *Executor) executeAllEvent(
        sim map[common.Address]event.Callback,
        e event.Event,
    ) {
    for _, v := range sim {
        event.Callback(v)(e)
    }
}

func (c *Executor) containUser(ul []common.Address, user common.Address) bool {
    for _, u := range ul {
        if u == user {
            return true
        }
    }

    return false
}