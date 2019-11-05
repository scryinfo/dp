// Scry Info.  All rights reserved.
// license that can be found in the license file.

package redo

import (
	"os"
	"sync"
)

// StopType
type StopType string

// JobState
type JobState int

// const
const (
	STOP_SYS  StopType = "SYS"
	STOP_USER          = "USER"

	JobRunning JobState = iota
	JobStopping
)

// Receipt
type Receipt struct {
	done        chan struct{}
	plsExit     chan struct{}
	sigchan     chan os.Signal
	catchSignal bool
	state       JobState
	stopType    StopType // signal or user request stop
	*sync.Mutex
}

func newRecipet() *Receipt {
	return &Receipt{
		plsExit:     make(chan struct{}, 1),
		done:        make(chan struct{}, 1),
		state:       JobRunning,
		sigchan:     make(chan os.Signal, 1),
		catchSignal: false,
		Mutex:       new(sync.Mutex),
	}
}

// Stop
func (m *Receipt) Stop() bool {
	return m.stopWithRequest(STOP_USER)
}

func (m *Receipt) stopWithRequest(stopType StopType) (op bool) {
	if m.state != JobRunning {
		return op
	}
	m.Lock()
	if m.state == JobRunning {
		m.state = JobStopping
		if stopType == STOP_USER {
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

// WaitChan
func (m *Receipt) WaitChan() <-chan struct{} {
	return m.done
}

// Wait
func (m *Receipt) Wait() StopType {
	<-m.WaitChan()
	return m.stopType
}

// Concat
func (m *Receipt) Concat(others ...*Receipt) *CombiReceipt {
	rs := append([]*Receipt{m}, others...)
	return NewCombiRecipt(rs...)
}

// CombiReceipt
type CombiReceipt struct {
	receipts []*Receipt
	alldone  chan struct{}
	once     *sync.Once
}

// NewCombiRecipt
func NewCombiRecipt(list ...*Receipt) *CombiReceipt {
	var unsafeRecipets []*Receipt
	for i, r := range list {
		if !r.catchSignal {
			unsafeRecipets = append(unsafeRecipets, list[i])
		}
	}
	if len(unsafeRecipets) > 0 && len(unsafeRecipets) < len(list) {
		for i := range unsafeRecipets {
			unsafeRecipets[i].catchSignal = true
			batchCatchSignals(unsafeRecipets[i].sigchan)
		}
	}
	return &CombiReceipt{
		receipts: list,
		alldone:  make(chan struct{}, 1),
		once:     new(sync.Once),
	}
}

// Stop
func (cr *CombiReceipt) Stop() bool {
	var ok bool
	for _, r := range cr.receipts {
		ok = r.Stop()
	}
	return ok
}

// Wait
func (cr *CombiReceipt) Wait() StopType {
	<-cr.WaitChan()
	return cr.receipts[0].stopType
}

// WaitChan
func (cr *CombiReceipt) WaitChan() <-chan struct{} {
	cr.once.Do(func() {
		go func() {
			for _, r := range cr.receipts {
				r.Wait()
			}
			close(cr.alldone)
		}()
	})
	return cr.alldone
}
