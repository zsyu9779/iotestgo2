package grade

import (
	"errors"
	"testing"
)

func TestGrade(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{name: "excellent lower bound", score: 90, want: "A"},
		{name: "good", score: 82, want: "B"},
		{name: "average", score: 75, want: "C"},
		{name: "pass", score: 60, want: "D"},
		{name: "fail", score: 59, want: "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Grade(tt.score)
			if err != nil {
				t.Fatalf("Grade(%d) returned error: %v", tt.score, err)
			}
			if got != tt.want {
				t.Fatalf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestGradeRejectsOutOfRange(t *testing.T) {
	for _, score := range []int{-1, 101} {
		got, err := Grade(score)
		if !errors.Is(err, ErrScoreOutOfRange) {
			t.Fatalf("Grade(%d) error = %v, want ErrScoreOutOfRange", score, err)
		}
		if got != "" {
			t.Fatalf("Grade(%d) grade = %q, want empty string", score, got)
		}
	}
}
