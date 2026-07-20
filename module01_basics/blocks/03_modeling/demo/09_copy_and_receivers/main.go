package main

import "fmt"

type Audit struct {
	CreatedBy string
}

type Account struct {
	Audit
	ID int
}

type Student struct {
	ID    int
	Name  string
	Score int
}

func (s Student) Label() string {
	return fmt.Sprintf("%d:%s=%d", s.ID, s.Name, s.Score)
}

func (s Student) RenameCopy(name string) {
	s.Name = name
}

func (s *Student) Rename(name string) {
	s.Name = name
}

func main() {
	student := Student{ID: 1, Name: "Alice", Score: 80}
	fmt.Println("value receiver:", student.Label())

	student.RenameCopy("Nobody")
	fmt.Println("value receiver mutation keeps:", student.Name)

	student.Rename("Bob")
	fmt.Println("pointer receiver:", student.Label())

	snapshot := student
	snapshot.Score = 100
	fmt.Printf("value snapshot: original=%d copy=%d\n", student.Score, snapshot.Score)

	account := Account{Audit: Audit{CreatedBy: "teacher"}, ID: 7}
	fmt.Printf("embedding promotes field access: id=%d created-by=%s\n", account.ID, account.CreatedBy)
	// Embedding is composition and promotion, not Java-style inheritance.
}
