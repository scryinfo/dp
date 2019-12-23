// Scry Info.  All rights reserved.
// license that can be found in the license file.

package listen

import (
	"os"
	"sync"
)

// JobState job state
type JobState int

// const
const (
	StopSys  = "SYS"
	StopUser = "USER"

	JobRunning JobState = iota
	JobStopping
)

// Receipt receipt
type Receipt struct {
	done        chan struct{}
	plsExit     chan struct{}
	sigChan     chan os.Signal
	catchSignal bool
	state       JobState
	stopType    string // signal or user request stop
	*sync.Mutex
}

func newReceipt() *Receipt {
	return &Receipt{
		plsExit:     make(chan struct{}, 1),
		done:        make(chan struct{}, 1),
		state:       JobRunning,
		sigChan:     make(chan os.Signal, 1),
		catchSignal: false,
		Mutex:       new(sync.Mutex),
	}
}

// Stop stop
func (m *Receipt) Stop() bool {
	return m.stopWithRequest(StopUser)
}

func (m *Receipt) stopWithRequest(stopType string) (op bool) {
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

// WaitChan wait chan
func (m *Receipt) WaitChan() <-chan struct{} {
	return m.done
}

// Wait wait
func (m *Receipt) Wait() string {
	<-m.WaitChan()
	return m.stopType
}

// Concat concat
func (m *Receipt) Concat(others ...*Receipt) *CombiReceipt {
	rs := append([]*Receipt{m}, others...)
	return NewCombiReceipt(rs...)
}

// CombiReceipt combi receipt
type CombiReceipt struct {
	receipts []*Receipt
	allDone  chan struct{}
	once     *sync.Once
}

// NewCombiReceipt new combi receipt
func NewCombiReceipt(list ...*Receipt) *CombiReceipt {
	var unsafeReceipts []*Receipt
	for i, r := range list {
		if !r.catchSignal {
			unsafeReceipts = append(unsafeReceipts, list[i])
		}
	}
	if len(unsafeReceipts) > 0 && len(unsafeReceipts) < len(list) {
		for i := range unsafeReceipts {
			unsafeReceipts[i].catchSignal = true
			batchCatchSignals(unsafeReceipts[i].sigChan)
		}
	}
	return &CombiReceipt{
		receipts: list,
		allDone:  make(chan struct{}, 1),
		once:     new(sync.Once),
	}
}

// Stop stop
func (cr *CombiReceipt) Stop() bool {
	var ok bool
	for _, r := range cr.receipts {
		ok = r.Stop()
	}
	return ok
}

// Wait wait
func (cr *CombiReceipt) Wait() string {
	<-cr.WaitChan()
	return cr.receipts[0].stopType
}

// WaitChan wait chan
func (cr *CombiReceipt) WaitChan() <-chan struct{} {
	cr.once.Do(func() {
		go func() {
			for _, r := range cr.receipts {
				r.Wait()
			}
			close(cr.allDone)
		}()
	})
	return cr.allDone
}
