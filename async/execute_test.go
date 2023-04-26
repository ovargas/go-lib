package async

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {

	f := Execute(context.Background(), func() error {
		return nil
	})

	err := f.Get()

	assert.NoError(t, err)
}

func TestExecuteNoWait(t *testing.T) {

	f := Execute(context.Background(), func() error {
		return nil
	})

	// simulate other processing
	time.Sleep(10 * time.Millisecond)

	assert.True(t, f.IsReady())

	err := f.Get()

	assert.NoError(t, err)
}

func TestExecuteFail(t *testing.T) {

	f := Execute(context.Background(), func() error {
		return fmt.Errorf("error")
	})

	err := f.Get()

	assert.Error(t, err)
	assert.Equal(t, "error", err.Error())
}

func TestExecutePanic(t *testing.T) {

	f := Execute(context.Background(), func() error {
		panic("panic error")
	})

	err := f.Get()

	assert.Error(t, err)
	assert.Equal(t, "panic error", err.Error())
}

func TestExecuteNotReadyYet(t *testing.T) {

	f := Execute(context.Background(), func() error {
		time.Sleep(300 * time.Millisecond)
		return nil
	})

	assert.False(t, f.IsReady())

	err := f.Get()

	assert.NoError(t, err)
}

func TestExecuteCancel(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	f := Execute(ctx, func() error {
		time.Sleep(3 * time.Second)
		return nil
	})

	cancel()

	err := f.Get()

	assert.Error(t, err)
}

func TestExecuteTimeout(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	f := Execute(ctx, func() error {
		time.Sleep(3 * time.Second)
		return nil
	})

	err := f.Get()

	assert.Error(t, err)
}
