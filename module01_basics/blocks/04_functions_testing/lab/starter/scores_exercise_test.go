//go:build exercise
// +build exercise

package scores

import (
	"reflect"
	"testing"
)

func TestFilterWithClosure(t *testing.T) {
	got := Filter([]int{59, 60, 75, 90}, AtLeast(60))
	want := []int{60, 75, 90}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Filter = %v, want %v", got, want)
	}
}

func TestWithAuditRecordsEndAfterOperation(t *testing.T) {
	var events []string
	WithAudit("average", func(event string) {
		events = append(events, event)
	}, func() {
		events = append(events, "operation")
	})
	want := []string{"start:average", "operation", "end:average"}
	if !reflect.DeepEqual(events, want) {
		t.Fatalf("events = %v, want %v", events, want)
	}
}
