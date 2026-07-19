package student

import "errors"

var (
	ErrInvalidID    = errors.New("invalid student ID")
	ErrInvalidName  = errors.New("invalid student name")
	ErrInvalidScore = errors.New("invalid student score")
)

type Student struct {
	ID    int
	Name  string
	Score int
}

func New(id int, name string, score int) (*Student, error) {
	return &Student{}, nil
}

func (s *Student) Rename(name string) error {
	return nil
}

func (s *Student) UpdateScore(score int) error {
	return nil
}

func (s Student) Snapshot() Student {
	return s
}
