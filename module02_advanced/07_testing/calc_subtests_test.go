package main

import "testing"

func TestAddSubtests(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 1, 2, 3},
		{"negative numbers", -1, -2, -3},
		{"zero", 0, 0, 0},
		{"mixed", -5, 10, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestAddEdgeCases(t *testing.T) {
	t.Run("large numbers", func(t *testing.T) {
		if Add(1000000, 2000000) != 3000000 {
			t.Error("large number addition failed")
		}
	})

	t.Run("max int", func(t *testing.T) {
		// Verify Add works with max int values
		result := Add(int(^uint(0)>>1)-1, 1)
		if result != int(^uint(0)>>1) {
			t.Error("expected correct addition near max int")
		}
	})
}
