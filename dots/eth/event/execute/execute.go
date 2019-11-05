// Scry Info.  All rights reserved.
// license that can be found in the license file.

package execute

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/eth/event"
	"go.uber.org/zap"
	"sync"
)

// const
const (
	ExecTypeId       = "4313210a-5824-4ff1-8dd8-c71ccad711db"
	BroadcastToAll   = "0x00"
	TargetUsers      = "users"
	TargetOwner      = "owner"
	AppSeqNo         = "seqNo"
	TokenEvtApproval = "Approval"
)

// Executor
type Executor struct {
	eventChan chan event.Event
	repo      *event.Repository
	appId     string
}

//construct dot
func newExecutorDot() (dot.Dot, error) {
	var err error
	d := &Executor{}

	return d, err
}

// ExecutorTypeLive Data structure needed when generating newer component
func ExecutorTypeLive() *dot.TypeLives {
	return &dot.TypeLives{
		Meta: dot.Metadata{TypeId: ExecTypeId,
			NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
				return newExecutorDot()
			}},
	}
}

// Create
func (c *Executor) Create(l dot.Line) error {
	return nil
}

// ExecuteEvents
func (c *Executor) ExecuteEvents(ce chan event.Event, r *event.Repository, appId string) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("Executor::ExecuteEvents", zap.Any("failed to execute event, error: ", er))
		}
	}()

	c.eventChan = ce
	c.repo = r
	c.appId = appId

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
			dot.Logger().Errorln("", zap.Any("failed to execute event "+e.Name+" because of error: ", er))
		}
	}()

	var subs sync.Map
	if rv, ok := c.repo.MapEventCallback.Load(e.Name); ok {
		subs = rv.(sync.Map)
	} else {
		dot.Logger().Warnln("no event was executed, event:" + e.Name)
		return false
	}

	seqNo := e.Data.Get(AppSeqNo)
	if seqNo != c.appId && e.Name != TokenEvtApproval {
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
			dot.Logger().Warnln("unknown e type, e:" + e.Name)
		}
	}

	return true
}

func (c *Executor) executeMatchedEvent(
	//sim map[common.Address]event.Callback,
	m sync.Map,
	users []common.Address, e event.Event,
) {
	m.Range(func(k, v interface{}) bool {
		cb, ok1 := v.(event.Callback)
		ca, ok2 := k.(common.Address)
		if ok1 && ok2 {
			if c.containUser(users, ca) && !cb(e) {
				dot.Logger().Warnln("event execute error, event:" + e.Name)
			}
		} else {
			dot.Logger().Warnln("parameters error, event:" + e.Name)
		}

		return true
	})
}

func (c *Executor) executeAllEvent(
	m sync.Map,
	e event.Event,
) {
	m.Range(func(k, v interface{}) bool {
		if cb, ok := v.(event.Callback); ok {
			if !cb(e) {
				dot.Logger().Warnln("event execute error, event:" + e.Name)
			}
		} else {
			dot.Logger().Warnln("parameters error, event:" + e.Name)
		}

		return true
	})
}

func (c *Executor) containUser(ul []common.Address, user common.Address) bool {
	for _, u := range ul {
		if u == user {
			return true
		}
	}

	return false
}
