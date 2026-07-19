package student

import (
	"fmt"
	"testing"
)

func TestStudentLifecycle(t *testing.T) {
	s, err := New(1, " Alice ", 80)
	if err != nil {
		t.Fatal(err)
	}
	if s.Name != "Alice" {
		t.Fatalf("Name = %q, want Alice", s.Name)
	}
	if s.ID != 1 {
		t.Fatalf("ID = %d, want 1", s.ID)
	}
	if err := s.UpdateScore(95); err != nil {
		t.Fatal(err)
	}
	if s.Score != 95 {
		t.Fatalf("Score = %d, want 95", s.Score)
	}
	copy := s.Snapshot()
	copy.Name = "changed copy"
	if s.Name != "Alice" {
		t.Fatalf("snapshot mutation leaked to original: %#v", s)
	}
}

func TestNewRejectsInvalidFields(t *testing.T) {
	tests := []struct {
		name      string
		id        int
		student   string
		score     int
		wantError error
	}{
		{name: "zero ID", id: 0, student: "Alice", score: 80, wantError: ErrInvalidID},
		{name: "negative ID", id: -1, student: "Alice", score: 80, wantError: ErrInvalidID},
		{name: "blank name", id: 1, student: " \t\n ", score: 80, wantError: ErrInvalidName},
		{name: "score below range", id: 1, student: "Alice", score: -1, wantError: ErrInvalidScore},
		{name: "score above range", id: 1, student: "Alice", score: 101, wantError: ErrInvalidScore},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.id, tt.student, tt.score)
			if err != tt.wantError {
				t.Fatalf("New(%d, %q, %d) error = %v, want %v", tt.id, tt.student, tt.score, err, tt.wantError)
			}
		})
	}
}

func TestNewAcceptsScoreBounds(t *testing.T) {
	for _, score := range []int{0, 100} {
		t.Run(fmt.Sprintf("score_%d", score), func(t *testing.T) {
			s, err := New(1, "Alice", score)
			if err != nil {
				t.Fatalf("New score %d: %v", score, err)
			}
			if s.Score != score {
				t.Fatalf("Score = %d, want %d", s.Score, score)
			}
		})
	}
}

func TestRenameTrimsAndMutatesStudent(t *testing.T) {
	s, err := New(1, "Alice", 80)
	if err != nil {
		t.Fatal(err)
	}

	if err := s.Rename(" Bob "); err != nil {
		t.Fatal(err)
	}
	if s.Name != "Bob" {
		t.Fatalf("Name = %q, want Bob", s.Name)
	}
}

func TestRenameRejectsBlankNameWithoutMutation(t *testing.T) {
	s, err := New(1, "Alice", 80)
	if err != nil {
		t.Fatal(err)
	}

	if err := s.Rename(" \t "); err != ErrInvalidName {
		t.Fatalf("Rename error = %v, want %v", err, ErrInvalidName)
	}
	if s.Name != "Alice" {
		t.Fatalf("Name = %q after invalid rename, want Alice", s.Name)
	}
}

func TestUpdateScoreAcceptsBoundsAndMutatesStudent(t *testing.T) {
	s, err := New(1, "Alice", 80)
	if err != nil {
		t.Fatal(err)
	}

	for _, score := range []int{0, 100} {
		if err := s.UpdateScore(score); err != nil {
			t.Fatalf("UpdateScore(%d): %v", score, err)
		}
		if s.Score != score {
			t.Fatalf("Score = %d, want %d", s.Score, score)
		}
	}
}

func TestUpdateScoreRejectsOutOfRangeWithoutMutation(t *testing.T) {
	for _, score := range []int{-1, 101} {
		t.Run(fmt.Sprintf("score_%d", score), func(t *testing.T) {
			s, err := New(1, "Alice", 80)
			if err != nil {
				t.Fatal(err)
			}

			if err := s.UpdateScore(score); err != ErrInvalidScore {
				t.Fatalf("UpdateScore(%d) error = %v, want %v", score, err, ErrInvalidScore)
			}
			if s.Score != 80 {
				t.Fatalf("Score = %d after invalid update, want 80", s.Score)
			}
		})
	}
}
