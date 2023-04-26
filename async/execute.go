package async

import (
	"context"
)

type (

	// FutureEmpty is an interface that represents a Future that returns an error
	FutureEmpty interface {
		Future

		// Get returns the error returned by the function that was executed
		Get() error
	}

	future struct {
		base
	}
)

// Get returns the error returned by the executed function
func (f *future) Get() error {
	f.Wait()
	return f.err
}

// Execute executes the function fn in a separate goroutine and returns a FutureEmpty
func Execute(ctx context.Context, fn func() error) FutureEmpty {
	f := &future{
		base: base{
			wait: make(chan bool),
			ctx:  ctx,
		},
	}
	go func(f *future) {
		defer close(f.wait)
		defer f.setReady()
		defer recoverFromPanic(&f.base)
		f.err = fn()
	}(f)

	return f
}
