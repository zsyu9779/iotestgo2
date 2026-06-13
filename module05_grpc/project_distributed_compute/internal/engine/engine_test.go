package engine

import "testing"

func TestCompute(t *testing.T) {
	tests := []struct {
		name      string
		numbers   []int64
		operation Operation
		want      float64
		wantError bool
	}{
		{name: "sum", numbers: []int64{1, 2, 3}, operation: OperationSum, want: 6},
		{name: "avg", numbers: []int64{2, 4, 6}, operation: OperationAvg, want: 4},
		{name: "max", numbers: []int64{2, 9, 1}, operation: OperationMax, want: 9},
		{name: "min", numbers: []int64{2, 9, 1}, operation: OperationMin, want: 1},
		{name: "median even", numbers: []int64{3, 1, 4, 2}, operation: OperationMedian, want: 2.5},
		{name: "stddev", numbers: []int64{2, 4, 4, 4, 5, 5, 7, 9}, operation: OperationStddev, want: 2},
		{name: "empty", numbers: nil, operation: OperationSum, wantError: true},
		{name: "unknown", numbers: []int64{1}, operation: Operation("p99"), wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Compute(tt.operation, tt.numbers)
			if tt.wantError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %.4f, got %.4f", tt.want, got)
			}
		})
	}
}
