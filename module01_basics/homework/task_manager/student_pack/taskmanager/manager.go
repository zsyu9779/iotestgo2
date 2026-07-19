package taskmanager

import "errors"

var (
	ErrEmptyTitle     = errors.New("task title cannot be empty")
	ErrTaskNotFound   = errors.New("task not found")
	ErrNotImplemented = errors.New("not implemented")
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
	return Task{}, ErrNotImplemented
}

func (m *Manager) List() []Task {
	return make([]Task, 0)
}

func (m *Manager) Complete(id int) (Task, error) {
	return Task{}, ErrNotImplemented
}

func (m *Manager) Delete(id int) error {
	return ErrNotImplemented
}
