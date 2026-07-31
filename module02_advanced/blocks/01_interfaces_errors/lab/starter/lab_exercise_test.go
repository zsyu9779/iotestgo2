//go:build exercise

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

func TestParsePortAcceptsValidPort(t *testing.T) {
	got, err := ParsePort("8080")
	if err != nil || got != 8080 {
		t.Fatalf("ParsePort() = %d, %v; want 8080, nil", got, err)
	}
}

func TestParsePortWrapsValidationError(t *testing.T) {
	for _, value := range []string{"invalid", "0", "65536"} {
		_, err := ParsePort(value)
		var validationErr *ValidationError
		if !errors.As(err, &validationErr) {
			t.Fatalf("ParsePort(%q) error = %v, want ValidationError in chain", value, err)
		}
	}
}
