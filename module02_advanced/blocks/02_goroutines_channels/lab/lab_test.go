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

func TestCollectSquaresWithNoWorkersUsesOneWorker(t *testing.T) {
	got := CollectSquares([]int{2}, 0)
	if !reflect.DeepEqual(got, []int{4}) {
		t.Fatalf("CollectSquares() = %v, want [4]", got)
	}
}
