package lab

import "errors"

var ErrNotImplemented = errors.New("not implemented")

type User struct {
	ID   int
	Name string
}

func ReadFieldName(value User, fieldName string) (string, error) {
	return "", ErrNotImplemented
}
