package lab

import "errors"

var ErrNotImplemented = errors.New("not implemented")

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 { return 0 }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 { return 0 }

func TotalArea(shapes []Shape) float64 { return 0 }

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string { return e.Field + ": " + e.Msg }

func ParsePort(value string) (int, error) { return 0, ErrNotImplemented }
