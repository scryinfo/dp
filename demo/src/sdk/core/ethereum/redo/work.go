// Scry Info.  All rights reserved.
// license that can be found in the license file.

package redo

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// Job
type Job func(*RedoCtx)

// RedoCtx
type RedoCtx struct {
	delayBeforeNextLoop time.Duration
	stopRedo            bool
}

func newCtx(duration time.Duration) *RedoCtx {
	return &RedoCtx{
		delayBeforeNextLoop: duration,
		stopRedo:            false,
	}
}

// SetDelayBeforeNext
func (ctx *RedoCtx) SetDelayBeforeNext(newDuration time.Duration) {
	ctx.delayBeforeNextLoop = newDuration
}

// StartNextRightNow
func (ctx *RedoCtx) StartNextRightNow() {
	ctx.SetDelayBeforeNext(time.Duration(0))
}

// StopRedo
func (ctx *RedoCtx) StopRedo() {
	ctx.stopRedo = true
}

// WrapFunc
func WrapFunc(work func()) Job {
	return func(ctx *RedoCtx) {
		work()
	}
}

// Perform perform job without gracefully exit
func Perform(once Job, duration time.Duration) *Receipt {
	return performWork(once, duration, false)
}

// PerformSafe perform job with gracefully exit
func PerformSafe(once Job, duration time.Duration) *Receipt {
	return performWork(once, duration, true)
}

func performWork(once Job, duration time.Duration, catchSignal bool) *Receipt {
	onceFunc := func(ctx *RedoCtx) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 1<<16)
				runtime.Stack(buf, false)
				log.Printf("panic occur:%+v\nstacktrace:%s\n", r, string(buf))
			}
		}()
		once(ctx)
	}
	recipet := newRecipet()
	recipet.catchSignal = catchSignal
	go func(m *Receipt) {
		if catchSignal {
			batchCatchSignals(recipet.sigchan)
		}
		plsExit := m.requestStopChan()
		for {
			ctx := newCtx(duration)
			onceFunc(ctx)
			if ctx.stopRedo {
				m.Stop()
			}

			select {
			case <-plsExit:
				m.closeChannels()
				return
			case <-recipet.sigchan:
				signal.Stop(recipet.sigchan)
				m.stopWithRequest(STOP_SYS)
				m.closeChannels()
				return
			case <-time.After(ctx.delayBeforeNextLoop):
			}
		}
	}(recipet)
	return recipet
}

func batchCatchSignals(sigchan chan os.Signal) {
	signal.Notify(sigchan, syscall.SIGABRT, syscall.SIGALRM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
}
