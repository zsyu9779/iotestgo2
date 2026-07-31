package scanner

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestScanCountsFilesAndMatchingLines(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "one.log", "INFO start\nERROR failed\nERROR retry\n")
	writeFile(t, root, "nested/two.log", "WARN wait\nERROR stopped\n")
	writeFile(t, root, "nested/empty.log", "")

	got, err := Scan(context.Background(), root, "ERROR", 3)
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	want := Result{Files: 3, MatchingLines: 3}
	if got != want {
		t.Fatalf("Scan() = %+v, want %+v", got, want)
	}
}

func TestScanRejectsInvalidArguments(t *testing.T) {
	root := t.TempDir()
	tests := []struct {
		name    string
		root    string
		needle  string
		workers int
		want    error
	}{
		{name: "missing root", root: filepath.Join(root, "missing"), needle: "x", workers: 1, want: ErrInvalidRoot},
		{name: "empty needle", root: root, needle: "", workers: 1, want: ErrEmptyNeedle},
		{name: "zero workers", root: root, needle: "x", workers: 0, want: ErrInvalidWorkerCount},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Scan(context.Background(), tt.root, tt.needle, tt.workers)
			if !errors.Is(err, tt.want) {
				t.Fatalf("Scan() error = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestScanHonorsCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := Scan(ctx, t.TempDir(), "ERROR", 2)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Scan() error = %v, want context.Canceled", err)
	}
}

func writeFile(t *testing.T, root, name, content string) {
	t.Helper()
	filename := filepath.Join(root, name)
	if err := os.MkdirAll(filepath.Dir(filename), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filename, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
