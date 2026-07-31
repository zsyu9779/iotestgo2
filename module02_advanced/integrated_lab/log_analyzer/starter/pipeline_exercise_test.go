//go:build exercise

package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestCountErrors(t *testing.T) {
	entries := []LogEntry{{Level: "INFO"}, {Level: "ERROR"}, {Level: "ERROR"}}
	if got := CountErrors(entries, 2); got != 2 {
		t.Fatalf("CountErrors() = %d, want 2", got)
	}
}

func TestAnalyzeScenarios(t *testing.T) {
	t.Run("invalid workers", func(t *testing.T) {
		source := make(chan LogEntry)
		close(source)
		if _, err := Analyze(context.Background(), source, 0); err == nil {
			t.Fatal("Analyze() error = nil, want invalid worker error")
		}
	})

	t.Run("empty input", func(t *testing.T) {
		source := make(chan LogEntry)
		close(source)
		got, err := Analyze(context.Background(), source, 2)
		if err != nil || got != 0 {
			t.Fatalf("Analyze() = %d, %v; want 0, nil", got, err)
		}
	})

	t.Run("pre-canceled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		source := make(chan LogEntry)
		if _, err := Analyze(ctx, source, 2); !errors.Is(err, context.Canceled) {
			t.Fatalf("Analyze() error = %v, want context.Canceled", err)
		}
	})
}

func TestAnalyzeStopsDuringCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	source := make(chan LogEntry, 1)
	done := make(chan error, 1)
	go func() {
		_, err := Analyze(ctx, source, 4)
		done <- err
	}()
	source <- LogEntry{Level: "ERROR"}
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
