package lab

import (
	"context"
	"testing"
	"time"
)

func TestRunCancellableTaskStopsAfterCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	started := make(chan struct{})
	done := make(chan error, 1)
	go func() {
		done <- RunCancellableTask(ctx, started)
	}()
	<-started
	cancel()

	select {
	case err := <-done:
		if err != context.Canceled {
			t.Fatalf("RunCancellableTask() error = %v, want context.Canceled", err)
		}
	case <-time.After(time.Second):
		t.Fatal("RunCancellableTask() did not stop after cancellation")
	}
}

func TestSafeCounterIsCorrectUnderConcurrency(t *testing.T) {
	var counter SafeCounter
	RunInParallel(&counter, 100, 10)
	if got := counter.Value(); got != 1000 {
		t.Fatalf("SafeCounter.Value() = %d, want 1000", got)
	}
}
