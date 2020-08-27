package timeout_test

import (
	"context"
	"testing"
	"time"

	"github.com/progfay/go-timeout/timeout"
)

type Callback struct {
	isCalled bool
}

func (c *Callback) Call() {
	c.isCalled = true
}

func (c *Callback) IsCalled() bool {
	return c.isCalled
}

func Test_SetTimeout(t *testing.T) {
	t.Run("Callback function called after the delay", func(t *testing.T) {
		c := Callback{isCalled: false}
		cancel := timeout.SetTimeout(c.Call, 100*time.Microsecond)
		defer cancel()
		time.Sleep(300 * time.Microsecond)
		if !c.IsCalled() {
			t.Error("First argument of SetTimeout must be called after the delay.")
		}
	})

	t.Run("Canceling SetTimeout", func(t *testing.T) {
		c := Callback{isCalled: false}
		cancel := timeout.SetTimeout(c.Call, 100*time.Microsecond)
		cancel()
		time.Sleep(300 * time.Microsecond)
		if c.IsCalled() {
			t.Error("First argument of SetTimeout must be not called when cancelled.")
		}
	})
}

func Test_SetTimeoutWithContext(t *testing.T) {
	t.Run("Canceling SetTimeout when parent context is canceled", func(t *testing.T) {
		c := Callback{isCalled: false}
		ctx, cancel := context.WithCancel(context.Background())
		_ = timeout.SetTimeoutWithContext(ctx, c.Call, 100*time.Microsecond)
		cancel()
		time.Sleep(300 * time.Microsecond)
		if c.IsCalled() {
			t.Error("First argument of SetTimeout must be not called when parent context is cancelled.")
		}
	})
}
