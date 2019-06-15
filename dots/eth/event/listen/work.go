// Scry Info.  All rights reserved.
// license that can be found in the license file.

package listen

import (
    "github.com/scryinfo/dot/dot"
    "go.uber.org/zap"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type Job func(*RedoCtx)

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

func (ctx *RedoCtx) SetDelayBeforeNext(new_duration time.Duration) {
	ctx.delayBeforeNextLoop = new_duration
}

func (ctx *RedoCtx) StartNextRightNow() {
	ctx.SetDelayBeforeNext(time.Duration(0))
}

func (ctx *RedoCtx) StopRedo() {
	ctx.stopRedo = true
}

func WrapFunc(work func()) Job {
	return func(ctx *RedoCtx) {
		work()
	}
}

// perform job without graceful exit
func Perform(once Job, duration time.Duration) *Receipt {
	return performWork(once, duration, false)
}

// perform job with graceful exit
func PerformSafe(once Job, duration time.Duration) *Receipt {
	return performWork(once, duration, true)
}

func performWork(once Job, duration time.Duration, catchSignal bool) *Receipt {
	onceFunc := func(ctx *RedoCtx) {
		defer func() {
			if r := recover(); r != nil {
				dot.Logger().Errorln("failed to perform work, panic error", zap.Stack("onceFunc"))
			}
		}()
		once(ctx)
	}

	r := newRecipet()
	r.catchSignal = catchSignal

	go func(m *Receipt) {
		if catchSignal {
			batchCatchSignals(r.sigChan)
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
			case <-r.sigChan:
				signal.Stop(r.sigChan)
				m.stopWithRequest(StopSys)
				m.closeChannels()
				return
			case <-time.After(ctx.delayBeforeNextLoop):
			}
		}
	}(r)
	return r
}

func batchCatchSignals(sigchan chan os.Signal) {
	signal.Notify(sigchan, syscall.SIGABRT, syscall.SIGALRM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
}
