package student

import (
	"errors"
	"strings"
)

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
	if id <= 0 {
		return nil, ErrInvalidID
	}

	trimmedName, err := validateName(name)
	if err != nil {
		return nil, err
	}
	if err := validateScore(score); err != nil {
		return nil, err
	}

	return &Student{ID: id, Name: trimmedName, Score: score}, nil
}

func (s *Student) Rename(name string) error {
	trimmedName, err := validateName(name)
	if err != nil {
		return err
	}

	s.Name = trimmedName
	return nil
}

func (s *Student) UpdateScore(score int) error {
	if err := validateScore(score); err != nil {
		return err
	}

	s.Score = score
	return nil
}

func (s Student) Snapshot() Student {
	return s
}

func validateName(name string) (string, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return "", ErrInvalidName
	}

	return trimmedName, nil
}

func validateScore(score int) error {
	if score < 0 || score > 100 {
		return ErrInvalidScore
	}

	return nil
}
