package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Task struct
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// TaskManager handles tasks
type TaskManager struct {
	tasks  []*Task
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  make([]*Task, 0),
		nextID: 1,
	}
}

func (tm *TaskManager) Add(title string) {
	task := &Task{
		ID:        tm.nextID,
		Title:     title,
		Completed: false,
	}
	tm.tasks = append(tm.tasks, task)
	tm.nextID++
	fmt.Printf("Task added with ID: %d\n", task.ID)
}

func (tm *TaskManager) List() {
	if len(tm.tasks) == 0 {
		fmt.Println("No tasks.")
		return
	}
	fmt.Println("Tasks:")
	for _, t := range tm.tasks {
		status := "[ ]"
		if t.Completed {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", t.ID, status, t.Title)
	}
}

func (tm *TaskManager) Complete(id int) {
	for _, t := range tm.tasks {
		if t.ID == id {
			t.Completed = true
			fmt.Println("Task marked as completed.")
			return
		}
	}
	fmt.Println("Task not found.")
}

func (tm *TaskManager) Delete(id int) {
	index := -1
	for i, t := range tm.tasks {
		if t.ID == id {
			index = i
			break
		}
	}
	if index != -1 {
		// Efficient delete for slice if order matters: copy elements
		tm.tasks = append(tm.tasks[:index], tm.tasks[index+1:]...)
		fmt.Println("Task deleted.")
	} else {
		fmt.Println("Task not found.")
	}
}

func main() {
	tm := NewTaskManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- CLI Task Manager ---")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose option: ")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title := scanner.Text()
			tm.Add(title)
		case "2":
			tm.List()
		case "3":
			fmt.Print("Enter task ID to complete: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			tm.Complete(id)
		case "4":
			fmt.Print("Enter task ID to delete: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			tm.Delete(id)
		case "5":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
