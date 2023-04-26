package async

import (
	"context"
	"fmt"
	"sync/atomic"
)

type (
	// Future is an interface that represents the result of an asynchronous operation
	Future interface {
		// IsReady returns true if the result is ready to be read without blocking
		IsReady() bool

		// Wait blocks until execution is complete or the context is cancelled
		Wait()
	}

	base struct {
		err      error
		complete atomic.Bool
		wait     chan bool
		ctx      context.Context
	}
)

// IsReady returns true if the result is ready to be read without blocking
func (f *base) IsReady() bool {
	return f.complete.Load()
}

func (f *base) setReady() {
	f.complete.Swap(true)
	f.wait <- true
}

// Wait blocks until execution is complete or the context is cancelled
func (f *base) Wait() {
	if f.IsReady() {
		return
	}
	select {
	case <-f.wait:
		return
	case <-f.ctx.Done():
		f.err = f.ctx.Err()
	}
}

func recoverFromPanic(f *base) {
	if r := recover(); r != nil {
		f.err = fmt.Errorf("%v", r)
	}
}
