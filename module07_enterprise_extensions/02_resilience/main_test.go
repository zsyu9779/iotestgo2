package main

import (
	"context"
	"errors"
	"testing"
)

func TestRetrySucceedsBeforeLimit(t *testing.T) {
	attempts := 0

	err := retry(context.Background(), 3, func(context.Context) error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary unavailable")
		}
		return nil
	})

	if err != nil {
		t.Fatalf("expected retry to succeed, got %v", err)
	}
	if attempts != 3 {
		t.Fatalf("expected 3 attempts, got %d", attempts)
	}
}

func TestRateLimiterRejectsBeyondLimit(t *testing.T) {
	limiter := newRateLimiter(2)

	if !limiter.allow() {
		t.Fatal("expected first request to pass")
	}
	if !limiter.allow() {
		t.Fatal("expected second request to pass")
	}
	if limiter.allow() {
		t.Fatal("expected third request to be rate limited")
	}
}
