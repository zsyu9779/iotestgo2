package main

import "testing"

func TestTaskManager(t *testing.T) {
	tm := NewTaskManager()

	// Test Add
	tm.Add("Task 1")
	if len(tm.tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(tm.tasks))
	}
	if tm.tasks[0].Title != "Task 1" {
		t.Errorf("expected title 'Task 1', got '%s'", tm.tasks[0].Title)
	}

	// Test Complete
	tm.Complete(1)
	if !tm.tasks[0].Completed {
		t.Error("expected task to be completed")
	}

	// Test Delete
	tm.Delete(1)
	if len(tm.tasks) != 0 {
		t.Errorf("expected 0 tasks, got %d", len(tm.tasks))
	}
}
