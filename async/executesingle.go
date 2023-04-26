package async

import "context"

type (

	// FutureSingle is a FutureEmpty that returns a single value of type T
	FutureSingle[T any] interface {
		Future

		// Get returns the result of the function that was executed
		Get() (T, error)
	}

	futureOne[T any] struct {
		base
		result T
	}
)

// Get returns the result of the function that was executed
func (f *futureOne[T]) Get() (T, error) {
	f.Wait()
	return f.result, f.err
}

// ExecuteSingle executes the function fn in a separate goroutine and returns a Future tha returns a single result
func ExecuteSingle[T any](ctx context.Context, fn func() (T, error)) FutureSingle[T] {
	f := &futureOne[T]{
		base: base{
			wait: make(chan bool),
			ctx:  ctx,
		},
	}
	go func(f *futureOne[T]) {
		defer close(f.wait)
		defer f.setReady()
		defer recoverFromPanic(&f.base)
		f.result, f.err = fn()
	}(f)

	return f
}
