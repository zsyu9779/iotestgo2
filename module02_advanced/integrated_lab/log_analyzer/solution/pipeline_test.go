package main

import "testing"

func TestCountErrors(t *testing.T) {
	entries := []LogEntry{
		{ID: 1, Level: "INFO"},
		{ID: 2, Level: "ERROR"},
		{ID: 3, Level: "WARN"},
		{ID: 4, Level: "ERROR"},
	}

	got := CountErrors(entries, 2)
	if got != 2 {
		t.Fatalf("expected 2 errors, got %d", got)
	}
}
