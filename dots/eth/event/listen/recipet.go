// Scry Info.  All rights reserved.
// license that can be found in the license file.

package listen

import (
	"os"
	"sync"
)

const (
	StopSys  = "SYS"
	StopUser = "USER"
)

type JobState int

const (
	JobRunning JobState = iota
	JobStopping
)

type Receipt struct {
	done        chan struct{}
	plsExit     chan struct{}
	sigChan     chan os.Signal
	catchSignal bool
	state       JobState
	stopType    string // signal or user request stop
	*sync.Mutex
}

func newRecipet() *Receipt {
	return &Receipt{
		plsExit:     make(chan struct{}, 1),
		done:        make(chan struct{}, 1),
		state:       JobRunning,
		sigChan:     make(chan os.Signal, 1),
		catchSignal: false,
		Mutex:       new(sync.Mutex),
	}
}

func (m *Receipt) Stop() bool {
	return m.stopWithRequest(StopUser)
}

func (m *Receipt) stopWithRequest(stopType string) bool {
	var op bool = false
	if m.state != JobRunning {
		return op
	}
	m.Lock()
	if m.state == JobRunning {
		m.state = JobStopping
		if stopType == StopUser {
			m.plsExit <- struct{}{}
		}
		m.stopType = stopType
		op = true
	}
	m.Unlock()
	return op
}

func (m *Receipt) closeChannels() {
	close(m.plsExit)
	close(m.done)
}

func (m *Receipt) requestStopChan() <-chan struct{} {
	return m.plsExit
}

func (m *Receipt) WaitChan() <-chan struct{} {
	return m.done
}

func (m *Receipt) Wait() string {
	<-m.WaitChan()
	return m.stopType
}

func (m *Receipt) Concat(others ...*Receipt) *CombiRecipt {
	rs := append([]*Receipt{m}, others...)
	return NewCombiRecipt(rs...)
}

type CombiRecipt struct {
	recipets []*Receipt
	alldone  chan struct{}
	once     *sync.Once
}

func NewCombiRecipt(list ...*Receipt) *CombiRecipt {
	var unsafeRecipets []*Receipt
	for i, r := range list {
		if !r.catchSignal {
			unsafeRecipets = append(unsafeRecipets, list[i])
		}
	}
	if len(unsafeRecipets) > 0 && len(unsafeRecipets) < len(list) {
		for i := range unsafeRecipets {
			unsafeRecipets[i].catchSignal = true
			batchCatchSignals(unsafeRecipets[i].sigChan)
		}
	}
	return &CombiRecipt{
		recipets: list,
		alldone:  make(chan struct{}, 1),
		once:     new(sync.Once),
	}
}

func (cr *CombiRecipt) Stop() bool {
	var ok bool
	for _, r := range cr.recipets {
		ok = r.Stop()
	}
	return ok
}

func (cr *CombiRecipt) Wait() string {
	<-cr.WaitChan()
	return cr.recipets[0].stopType
}

func (cr *CombiRecipt) WaitChan() <-chan struct{} {
	cr.once.Do(func() {
		go func() {
			for _, r := range cr.recipets {
				r.Wait()
			}
			close(cr.alldone)
		}()
	})
	return cr.alldone
}
