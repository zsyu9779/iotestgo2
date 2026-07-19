package scorebook

import "errors"

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

type Scorebook struct{}

func New() *Scorebook {
	return &Scorebook{}
}

func (s *Scorebook) Add(name string, score int) (Student, error) {
	return Student{}, nil
}

func (s *Scorebook) Find(id int) (Student, error) {
	return Student{}, ErrStudentNotFound
}

func (s *Scorebook) UpdateScore(id int, score int) error {
	return ErrStudentNotFound
}

func (s *Scorebook) Average() (float64, error) {
	return 0, ErrEmptyScorebook
}

func (s *Scorebook) CountByGrade() map[string]int {
	return nil
}
