package scorebook

import (
	"errors"
	"strings"
)

var (
	ErrInvalidName     = errors.New("invalid student name")
	ErrInvalidScore    = errors.New("invalid student score")
	ErrStudentNotFound = errors.New("student not found")
	ErrEmptyScorebook  = errors.New("scorebook is empty")
)

type Student struct {
	ID    int
	Name  string
	Score int
}

type Scorebook struct {
	students map[int]*Student
	nextID   int
}

func New() *Scorebook {
	return &Scorebook{
		students: make(map[int]*Student),
		nextID:   1,
	}
}

func (s *Scorebook) Add(name string, score int) (Student, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return Student{}, ErrInvalidName
	}
	if !validScore(score) {
		return Student{}, ErrInvalidScore
	}

	student := &Student{ID: s.nextID, Name: trimmedName, Score: score}
	s.students[student.ID] = student
	s.nextID++
	return *student, nil
}

func (s *Scorebook) Find(id int) (Student, error) {
	student, ok := s.students[id]
	if !ok {
		return Student{}, ErrStudentNotFound
	}

	return *student, nil
}

func (s *Scorebook) UpdateScore(id int, score int) error {
	student, ok := s.students[id]
	if !ok {
		return ErrStudentNotFound
	}
	if !validScore(score) {
		return ErrInvalidScore
	}

	student.Score = score
	return nil
}

func (s *Scorebook) Average() (float64, error) {
	if len(s.students) == 0 {
		return 0, ErrEmptyScorebook
	}

	total := 0
	for _, student := range s.students {
		total += student.Score
	}

	return float64(total) / float64(len(s.students)), nil
}

func (s *Scorebook) CountByGrade() map[string]int {
	counts := make(map[string]int)
	for _, student := range s.students {
		counts[grade(student.Score)]++
	}

	return counts
}

func validScore(score int) bool {
	return score >= 0 && score <= 100
}

func grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}
