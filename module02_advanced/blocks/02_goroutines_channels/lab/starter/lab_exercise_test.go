//go:build exercise

package lab

import (
	"reflect"
	"testing"
)

func TestCollectSquaresWaitsForAllWorkers(t *testing.T) {
	got := CollectSquares([]int{1, 2, 3, 4}, 2)
	want := []int{1, 4, 9, 16}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CollectSquares() = %v, want %v", got, want)
	}
}

func TestCollectSquaresHandlesWorkerBounds(t *testing.T) {
	for _, workers := range []int{0, 1, 8} {
		got := CollectSquares([]int{2, 3}, workers)
		if !reflect.DeepEqual(got, []int{4, 9}) {
			t.Fatalf("CollectSquares(workers=%d) = %v, want [4 9]", workers, got)
		}
	}
}

func TestCollectSquaresHandlesEmptyInput(t *testing.T) {
	if got := CollectSquares(nil, 2); len(got) != 0 {
		t.Fatalf("CollectSquares(nil, 2) = %v, want empty result", got)
	}
}
