//go:build exercise
// +build exercise

package scores

import "testing"

func TestFilterPreservesInputOrderForAdditionalCase(t *testing.T) {
	got := Filter([]int{7, 2, 9, 4}, func(value int) bool {
		return value%2 == 1
	})
	want := []int{7, 9}
	if len(got) != len(want) {
		t.Fatalf("Filter length = %d, want %d", len(got), len(want))
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("Filter[%d] = %d, want %d", index, got[index], want[index])
		}
	}
}
