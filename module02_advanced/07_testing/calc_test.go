package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}

// Table-Driven Test
func TestAddTable(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 1, 2},
		{10, 20, 30},
		{-1, 1, 0},
	}

	for _, tt := range tests {
		got := Add(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

// Subtests：使用 t.Run() 组织子测试
func TestAddSubtests(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		if got := Add(2, 3); got != 5 {
			t.Errorf("2 + 3 = %d; want 5", got)
		}
	})
	t.Run("negative", func(t *testing.T) {
		if got := Add(-1, -1); got != -2 {
			t.Errorf("-1 + -1 = %d; want -2", got)
		}
	})
	t.Run("zero", func(t *testing.T) {
		if got := Add(0, 5); got != 5 {
			t.Errorf("0 + 5 = %d; want 5", got)
		}
	})
}

// Parallel subtests
func TestAddParallel(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"negative", -1, -1, -2},
		{"zero", 0, 5, 5},
		{"mixed", -5, 3, -2},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := Add(tc.a, tc.b); got != tc.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
