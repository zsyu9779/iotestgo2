package taskmanager

import (
	"errors"
	"reflect"
	"testing"
)

func TestAddTrimsTitleAndReturnsCopy(t *testing.T) {
	m := NewManager()
	added, err := m.Add(" write tests ")
	if err != nil {
		t.Fatal(err)
	}
	if added.ID != 1 || added.Title != "write tests" || added.Completed {
		t.Fatalf("added task = %#v", added)
	}

	added.Title = "changed returned value"
	stored := m.List()
	if len(stored) != 1 || stored[0].Title != "write tests" {
		t.Fatalf("stored task changed through Add result: %#v", stored)
	}
}

func TestAddRejectsBlankTitleWithoutConsumingID(t *testing.T) {
	m := NewManager()
	for _, title := range []string{"", " \t\n"} {
		if _, err := m.Add(title); !errors.Is(err, ErrEmptyTitle) {
			t.Fatalf("Add(%q) error = %v; want ErrEmptyTitle", title, err)
		}
	}

	added, err := m.Add("first valid task")
	if err != nil {
		t.Fatal(err)
	}
	if added.ID != 1 {
		t.Fatalf("first valid task ID = %d; want 1", added.ID)
	}
}

func TestDeletePreservesOrderAndIDsRemainMonotonic(t *testing.T) {
	m := NewManager()
	first, err := m.Add("first")
	if err != nil {
		t.Fatal(err)
	}
	second, err := m.Add("second")
	if err != nil {
		t.Fatal(err)
	}
	third, err := m.Add("third")
	if err != nil {
		t.Fatal(err)
	}

	if err := m.Delete(second.ID); err != nil {
		t.Fatal(err)
	}
	remaining := m.List()
	if len(remaining) != 2 {
		t.Fatalf("tasks after deletion = %#v; want two tasks", remaining)
	}
	if got, want := []int{remaining[0].ID, remaining[1].ID}, []int{first.ID, third.ID}; !reflect.DeepEqual(got, want) {
		t.Fatalf("IDs after deletion = %v; want %v", got, want)
	}

	fourth, err := m.Add("fourth")
	if err != nil {
		t.Fatal(err)
	}
	if fourth.ID != 4 || fourth.ID == second.ID {
		t.Fatalf("ID after deletion = %d; want 4 and no reuse of %d", fourth.ID, second.ID)
	}
	if got, want := m.List(), []Task{first, third, fourth}; !reflect.DeepEqual(got, want) {
		t.Fatalf("tasks after adding again = %#v; want %#v", got, want)
	}
}

func TestCompleteReturnsCopy(t *testing.T) {
	m := NewManager()
	added, err := m.Add("keep this title")
	if err != nil {
		t.Fatal(err)
	}

	completed, err := m.Complete(added.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !completed.Completed {
		t.Fatal("completed task is not marked completed")
	}
	completed.Title = "changed returned value"
	completed.Completed = false

	stored := m.List()
	if stored[0].Title != "keep this title" || !stored[0].Completed {
		t.Fatalf("stored task changed through Complete result: %#v", stored[0])
	}
}

func TestListReturnsIndependentSliceAndValues(t *testing.T) {
	m := NewManager()
	first, err := m.Add("first")
	if err != nil {
		t.Fatal(err)
	}
	second, err := m.Add("second")
	if err != nil {
		t.Fatal(err)
	}

	listed := m.List()
	if len(listed) != 2 {
		t.Fatalf("List() = %#v; want two tasks", listed)
	}
	listed[0].Title = "changed outside manager"
	listed[1] = Task{ID: 99, Title: "replacement"}
	listed = append(listed, Task{ID: 100, Title: "extra"})

	if got, want := m.List(), []Task{first, second}; !reflect.DeepEqual(got, want) {
		t.Fatalf("stored tasks changed through List result: %#v; want %#v", got, want)
	}
}

func TestMissingIDsReturnErrTaskNotFound(t *testing.T) {
	m := NewManager()
	if _, err := m.Complete(99); !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("Complete(99) error = %v; want ErrTaskNotFound", err)
	}
	if err := m.Delete(99); !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("Delete(99) error = %v; want ErrTaskNotFound", err)
	}
}
