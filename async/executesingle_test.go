package async

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExecuteSingle(t *testing.T) {

	f := ExecuteSingle(context.Background(), func() (int, error) {
		return 1, nil
	})

	v, err := f.Get()

	assert.NoError(t, err)
	assert.Equal(t, 1, v)
}

func TestExecuteSingleFail(t *testing.T) {

	f := ExecuteSingle(context.Background(), func() (int, error) {
		return 0, fmt.Errorf("error")
	})

	v, err := f.Get()

	assert.Error(t, err)
	assert.Equal(t, "error", err.Error())
	assert.Equal(t, 0, v)
}

func TestExecuteSinglePanic(t *testing.T) {

	f := ExecuteSingle(context.Background(), func() (int, error) {
		panic("panic error")
	})

	v, err := f.Get()

	assert.Error(t, err)
	assert.Equal(t, "panic error", err.Error())
	assert.Equal(t, 0, v)
}

func TestExecuteSingleNotReadyYet(t *testing.T) {

	f := ExecuteSingle(context.Background(), func() (int, error) {
		time.Sleep(300 * time.Millisecond)
		return 1, nil
	})

	assert.False(t, f.IsReady())

	v, err := f.Get()

	assert.NoError(t, err)
	assert.Equal(t, 1, v)
}

func TestExecuteSingleCancel(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	f := ExecuteSingle(ctx, func() (int, error) {
		time.Sleep(3 * time.Second)
		return 1, nil
	})

	cancel()

	v, err := f.Get()

	assert.Error(t, err)
	assert.Equal(t, 0, v)
}

func TestExecuteSingleTimeout(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	f := ExecuteSingle(ctx, func() (int, error) {
		time.Sleep(3 * time.Second)
		return 1, nil
	})

	v, err := f.Get()

	assert.Error(t, err)
	assert.Equal(t, 0, v)
}
