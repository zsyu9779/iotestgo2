//go:build exercise
// +build exercise

package textstats

import "testing"

func TestAnalyze(t *testing.T) {
	got := Analyze("Go go 你好")
	if got.Bytes != 12 {
		t.Fatalf("Bytes = %d, want 12", got.Bytes)
	}
	if got.Runes != 8 {
		t.Fatalf("Runes = %d, want 8", got.Runes)
	}
	if got.Words != 3 {
		t.Fatalf("Words = %d, want 3", got.Words)
	}
	if got.Frequencies["go"] != 2 || got.Frequencies["你好"] != 1 {
		t.Fatalf("Frequencies = %#v", got.Frequencies)
	}
}

func TestAnalyzeEmptyText(t *testing.T) {
	got := Analyze("")
	if got.Bytes != 0 || got.Runes != 0 || got.Words != 0 || len(got.Frequencies) != 0 {
		t.Fatalf("Analyze empty = %#v", got)
	}
}
