//go:build exercise
// +build exercise

package scorebook

import (
	"errors"
	"fmt"
	"testing"
)

func TestScorebookWorkflow(t *testing.T) {
	book := New()
	alice, err := book.Add(" Alice ", 90)
	if err != nil {
		t.Fatal(err)
	}
	bob, err := book.Add("Bob", 70)
	if err != nil {
		t.Fatal(err)
	}

	if alice.ID != 1 || bob.ID != 2 {
		t.Fatalf("IDs = %d, %d; want 1, 2", alice.ID, bob.ID)
	}
	if alice.Name != "Alice" {
		t.Fatalf("Alice name = %q, want %q", alice.Name, "Alice")
	}

	alice.Name = "changed returned value"
	alice.Score = 0
	storedAlice, err := book.Find(alice.ID)
	if err != nil {
		t.Fatal(err)
	}
	if storedAlice.Name != "Alice" || storedAlice.Score != 90 {
		t.Fatalf("stored Alice = %#v after changing Add result; want unchanged", storedAlice)
	}

	storedAlice.Name = "changed found value"
	again, err := book.Find(alice.ID)
	if err != nil {
		t.Fatal(err)
	}
	if again.Name != "Alice" {
		t.Fatalf("stored Alice name = %q after changing Find result; want Alice", again.Name)
	}

	if err := book.UpdateScore(bob.ID, 80); err != nil {
		t.Fatal(err)
	}
	average, err := book.Average()
	if err != nil || average != 85 {
		t.Fatalf("Average = %v, %v; want 85, nil", average, err)
	}

	counts := book.CountByGrade()
	if counts["A"] != 1 || counts["B"] != 1 || len(counts) != 2 {
		t.Fatalf("CountByGrade = %#v; want map[A:1 B:1]", counts)
	}
}

func TestAddRejectsInvalidFieldsWithoutConsumingID(t *testing.T) {
	book := New()
	tests := []struct {
		name      string
		student   string
		score     int
		wantError error
	}{
		{name: "blank name", student: " \t\n ", score: 80, wantError: ErrInvalidName},
		{name: "score below range", student: "Alice", score: -1, wantError: ErrInvalidScore},
		{name: "score above range", student: "Alice", score: 101, wantError: ErrInvalidScore},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := book.Add(tt.student, tt.score); !errors.Is(err, tt.wantError) {
				t.Fatalf("Add(%q, %d) error = %v, want %v", tt.student, tt.score, err, tt.wantError)
			}
		})
	}

	student, err := book.Add("Alice", 80)
	if err != nil {
		t.Fatal(err)
	}
	if student.ID != 1 {
		t.Fatalf("first valid student ID = %d, want 1", student.ID)
	}
}

func TestFindAndUpdateReturnStudentNotFound(t *testing.T) {
	book := New()
	for _, id := range []int{0, 42} {
		t.Run(fmt.Sprintf("id_%d", id), func(t *testing.T) {
			if _, err := book.Find(id); !errors.Is(err, ErrStudentNotFound) {
				t.Fatalf("Find(%d) error = %v, want ErrStudentNotFound", id, err)
			}
			if err := book.UpdateScore(id, 80); !errors.Is(err, ErrStudentNotFound) {
				t.Fatalf("UpdateScore(%d, 80) error = %v, want ErrStudentNotFound", id, err)
			}
		})
	}
}

func TestUpdateScoreRejectsInvalidScoresWithoutMutation(t *testing.T) {
	for _, score := range []int{-1, 101} {
		t.Run(fmt.Sprintf("score_%d", score), func(t *testing.T) {
			book := New()
			student, err := book.Add("Alice", 80)
			if err != nil {
				t.Fatal(err)
			}

			if err := book.UpdateScore(student.ID, score); !errors.Is(err, ErrInvalidScore) {
				t.Fatalf("UpdateScore(%d, %d) error = %v, want ErrInvalidScore", student.ID, score, err)
			}
			stored, err := book.Find(student.ID)
			if err != nil {
				t.Fatal(err)
			}
			if stored.Score != 80 {
				t.Fatalf("score after rejected update = %d, want 80", stored.Score)
			}
		})
	}
}

func TestAverageReturnsErrorForEmptyScorebook(t *testing.T) {
	average, err := New().Average()
	if !errors.Is(err, ErrEmptyScorebook) {
		t.Fatalf("Average error = %v, want ErrEmptyScorebook", err)
	}
	if average != 0 {
		t.Fatalf("Average = %v, want 0 when scorebook is empty", average)
	}
}

func TestAverageUsesFloatingPointDivision(t *testing.T) {
	book := New()
	for _, entry := range []struct {
		name  string
		score int
	}{
		{name: "Alice", score: 80},
		{name: "Bob", score: 81},
	} {
		if _, err := book.Add(entry.name, entry.score); err != nil {
			t.Fatal(err)
		}
	}

	average, err := book.Average()
	if err != nil || average != 80.5 {
		t.Fatalf("Average = %v, %v; want 80.5, nil", average, err)
	}
}

func TestCountByGradeCoversEveryBoundary(t *testing.T) {
	book := New()
	scores := []int{100, 90, 89, 80, 79, 70, 69, 60, 59, 0}
	for i, score := range scores {
		if _, err := book.Add(fmt.Sprintf("Student %d", i+1), score); err != nil {
			t.Fatal(err)
		}
	}

	counts := book.CountByGrade()
	for _, grade := range []string{"A", "B", "C", "D", "F"} {
		if counts[grade] != 2 {
			t.Errorf("CountByGrade()[%q] = %d, want 2", grade, counts[grade])
		}
	}
	if len(counts) != 5 {
		t.Errorf("CountByGrade = %#v; want exactly five grade keys", counts)
	}
}
