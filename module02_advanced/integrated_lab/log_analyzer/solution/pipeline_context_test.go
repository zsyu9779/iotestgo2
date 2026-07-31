package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestAnalyzeReturnsErrorForInvalidWorkerCount(t *testing.T) {
	source := make(chan LogEntry)
	close(source)
	_, err := Analyze(context.Background(), source, 0)
	if err == nil {
		t.Fatal("Analyze() error = nil, want invalid worker error")
	}
}

func TestAnalyzeHonorsCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	source := make(chan LogEntry)
	defer close(source)

	_, err := Analyze(ctx, source, 2)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Analyze() error = %v, want context.Canceled", err)
	}
}

func TestAnalyzeCountsErrorsFromDeterministicSource(t *testing.T) {
	source := make(chan LogEntry, 3)
	source <- LogEntry{ID: 1, Level: "ERROR"}
	source <- LogEntry{ID: 2, Level: "INFO"}
	source <- LogEntry{ID: 3, Level: "ERROR"}
	close(source)

	got, err := Analyze(context.Background(), source, 2)
	if err != nil {
		t.Fatalf("Analyze() error = %v, want nil", err)
	}
	if got != 2 {
		t.Fatalf("Analyze() = %d, want 2", got)
	}
}

func TestAnalyzeHandlesEmptyInputWithMultipleWorkers(t *testing.T) {
	source := make(chan LogEntry)
	close(source)
	got, err := Analyze(context.Background(), source, 8)
	if err != nil || got != 0 {
		t.Fatalf("Analyze() = %d, %v; want 0, nil", got, err)
	}
}

func TestAnalyzeStopsDuringCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	source := make(chan LogEntry)
	done := make(chan error, 1)
	go func() {
		_, err := Analyze(ctx, source, 4)
		done <- err
	}()

	source <- LogEntry{ID: 1, Level: "ERROR"}
	cancel()

	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("Analyze() error = %v, want context.Canceled", err)
		}
	case <-time.After(time.Second):
		t.Fatal("Analyze() did not return after cancellation")
	}
}
