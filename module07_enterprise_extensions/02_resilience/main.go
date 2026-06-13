package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func callDownstream(ctx context.Context, latency time.Duration) error {
	select {
	case <-time.After(latency):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func retry(ctx context.Context, maxAttempts int, fn func(context.Context) error) error {
	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if err := fn(ctx); err != nil {
			lastErr = err
			if ctx.Err() != nil {
				return ctx.Err()
			}
			continue
		}
		return nil
	}
	return lastErr
}

type rateLimiter struct {
	limit int
	used  int
}

func newRateLimiter(limit int) *rateLimiter {
	return &rateLimiter{limit: limit}
}

func (l *rateLimiter) allow() bool {
	if l.used >= l.limit {
		return false
	}
	l.used++
	return true
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	err := callDownstream(ctx, 500*time.Millisecond)
	fmt.Printf("timeout result: %v\n", err)

	attempts := 0
	err = retry(context.Background(), 3, func(context.Context) error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary unavailable")
		}
		return nil
	})
	fmt.Printf("retry result: %v after %d attempts\n", err, attempts)

	limiter := newRateLimiter(2)
	fmt.Printf("rate limit results: %v %v %v\n", limiter.allow(), limiter.allow(), limiter.allow())
}
