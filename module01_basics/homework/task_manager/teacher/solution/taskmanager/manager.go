package taskmanager

import (
	"errors"
	"strings"
)

var (
	ErrEmptyTitle   = errors.New("task title cannot be empty")
	ErrTaskNotFound = errors.New("task not found")
)

type Manager struct {
	tasks  []*Task
	nextID int
}

func NewManager() *Manager {
	return &Manager{
		tasks:  make([]*Task, 0),
		nextID: 1,
	}
}

func (m *Manager) Add(title string) (Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, ErrEmptyTitle
	}

	task := &Task{
		ID:    m.nextID,
		Title: title,
	}
	m.tasks = append(m.tasks, task)
	m.nextID++
	return *task, nil
}

func (m *Manager) List() []Task {
	tasks := make([]Task, len(m.tasks))
	for i, task := range m.tasks {
		tasks[i] = *task
	}
	return tasks
}

func (m *Manager) Complete(id int) (Task, error) {
	for _, task := range m.tasks {
		if task.ID == id {
			task.Completed = true
			return *task, nil
		}
	}
	return Task{}, ErrTaskNotFound
}

func (m *Manager) Delete(id int) error {
	for i, task := range m.tasks {
		if task.ID == id {
			copy(m.tasks[i:], m.tasks[i+1:])
			m.tasks[len(m.tasks)-1] = nil
			m.tasks = m.tasks[:len(m.tasks)-1]
			return nil
		}
	}
	return ErrTaskNotFound
}
