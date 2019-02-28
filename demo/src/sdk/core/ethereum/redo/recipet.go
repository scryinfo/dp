package redo

import (
	"os"
	"sync"
)

type StopType string

const (
	STOP_SYS  StopType = "SYS"
	STOP_USER          = "USER"
)

type JobState int

const (
	JobRunning JobState = iota
	JobStopping
)

type Recipet struct {
	done        chan struct{}
	pls_exit    chan struct{}
	sigchan     chan os.Signal
	catchSignal bool
	state       JobState
	stop_type   StopType // signal or user request stop
	*sync.Mutex
}

func newRecipet() *Recipet {
	return &Recipet{
		pls_exit:    make(chan struct{}, 1),
		done:        make(chan struct{}, 1),
		state:       JobRunning,
		sigchan:     make(chan os.Signal, 1),
		catchSignal: false,
		Mutex:       new(sync.Mutex),
	}
}

func (m *Recipet) Stop() bool {
	return m.stopWithRequest(STOP_USER)
}

func (m *Recipet) stopWithRequest(stop_type StopType) bool {
	var op bool = false
	if m.state != JobRunning {
		return op
	}
	m.Lock()
	if m.state == JobRunning {
		m.state = JobStopping
		if stop_type == STOP_USER {
			m.pls_exit <- struct{}{}
		}
		m.stop_type = stop_type
		op = true
	}
	m.Unlock()
	return op
}

func (m *Recipet) closeChannels() {
	close(m.pls_exit)
	close(m.done)
}

func (m *Recipet) requestStopChan() <-chan struct{} {
	return m.pls_exit
}

func (m *Recipet) WaitChan() <-chan struct{} {
	return m.done
}

func (m *Recipet) Wait() StopType {
	<-m.WaitChan()
	return m.stop_type
}

func (m *Recipet) Concat(others ...*Recipet) *CombiRecipt {
	rs := append([]*Recipet{m}, others...)
	return NewCombiRecipt(rs...)
}

type CombiRecipt struct {
	recipets []*Recipet
	alldone  chan struct{}
	once     *sync.Once
}

func NewCombiRecipt(list ...*Recipet) *CombiRecipt {
	var unsafeRecipets []*Recipet
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

func (cr *CombiRecipt) Wait() StopType {
	<-cr.WaitChan()
	return cr.recipets[0].stop_type
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
