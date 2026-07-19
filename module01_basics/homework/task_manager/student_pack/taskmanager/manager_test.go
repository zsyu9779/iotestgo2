package taskmanager

import "testing"

func TestManagerLifecycle(t *testing.T) {
	m := NewManager()
	first, err := m.Add(" write tests ")
	if err != nil {
		t.Fatal(err)
	}
	if first.ID != 1 || first.Title != "write tests" || first.Completed {
		t.Fatalf("first task = %#v", first)
	}
	completed, err := m.Complete(first.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !completed.Completed {
		t.Fatal("task was not completed")
	}
	if err := m.Delete(first.ID); err != nil {
		t.Fatal(err)
	}
	if len(m.List()) != 0 {
		t.Fatalf("tasks after delete = %#v", m.List())
	}
}

func TestManagerAssignsSequentialIDs(t *testing.T) {
	m := NewManager()
	first, err := m.Add("first")
	if err != nil {
		t.Fatal(err)
	}
	second, err := m.Add("second")
	if err != nil {
		t.Fatal(err)
	}
	if first.ID != 1 || second.ID != 2 {
		t.Fatalf("task IDs = %d, %d; want 1, 2", first.ID, second.ID)
	}
}

func TestListReturnsCopies(t *testing.T) {
	m := NewManager()
	added, err := m.Add("keep this title")
	if err != nil {
		t.Fatal(err)
	}

	tasks := m.List()
	if len(tasks) != 1 {
		t.Fatalf("List() = %#v; want one task", tasks)
	}
	tasks[0].Title = "changed outside manager"

	completed, err := m.Complete(added.ID)
	if err != nil {
		t.Fatal(err)
	}
	completed.Title = "changed returned copy"

	stored := m.List()
	if stored[0].Title != "keep this title" || !stored[0].Completed {
		t.Fatalf("stored task changed through returned value: %#v", stored[0])
	}
}

func TestUnknownIDsReturnErrors(t *testing.T) {
	m := NewManager()
	if _, err := m.Complete(99); err == nil {
		t.Fatal("Complete(99) error = nil; want an error")
	}
	if err := m.Delete(99); err == nil {
		t.Fatal("Delete(99) error = nil; want an error")
	}
}
