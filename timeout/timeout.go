package timeout

import (
	"context"
	"time"
)

func SetTimeout(callback func(), delay time.Duration) (cancel func()) {
	return SetTimeoutWithContext(context.Background(), callback, delay)
}

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
