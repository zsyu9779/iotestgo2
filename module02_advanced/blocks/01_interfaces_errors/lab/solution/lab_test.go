package lab

import (
	"errors"
	"testing"
)

func TestTotalAreaUsesImplicitInterface(t *testing.T) {
	got := TotalArea([]Shape{Rectangle{Width: 2, Height: 3}, Circle{Radius: 1}})
	if got < 9.14 || got > 9.15 {
		t.Fatalf("TotalArea() = %v, want approximately 9.1415", got)
	}
}

func TestParsePortRejectsInvalidInput(t *testing.T) {
	if _, err := ParsePort("not-a-port"); err == nil {
		t.Fatal("ParsePort() error = nil, want an error")
	}
}

func TestParsePortWrapsValidationError(t *testing.T) {
	_, err := ParsePort("70000")
	var validationErr *ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("ParsePort() error = %v, want ValidationError", err)
	}
}
