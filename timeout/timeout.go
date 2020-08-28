package timeout

import (
	"context"
	"time"
)

// SetTimeout sets a timer which executes a function once the timer expires.
// To abort timer, call returned function.
func SetTimeout(callback func(), delay time.Duration) (cancel func()) {
	return SetTimeoutWithContext(context.Background(), callback, delay)
}

// SetTimeoutWithContext sets a timer which executes a function once the timer expires.
// When ctx is cancelled, timer is aborted.
func SetTimeoutWithContext(ctx context.Context, callback func(), delay time.Duration) func() {
	childCtx, cancel := context.WithCancel(ctx)
	timer := time.NewTimer(delay)
	go func(f func()) {
		select {
		case <-childCtx.Done():
			return
		case <-timer.C:
			f()
		}
	}(callback)
	return cancel
}
