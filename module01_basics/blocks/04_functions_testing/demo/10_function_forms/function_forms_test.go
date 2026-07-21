package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestFunctionDeclarationForms(t *testing.T) {
	if got, err := addText("1", 2, 3); err != nil || got != 6 {
		t.Fatalf("addText valid input = (%d, %v), want (6, nil)", got, err)
	}
	if _, err := addText("bad", 2, 3); err == nil {
		t.Fatal("addText invalid input returned nil error")
	}

	got, err := addTextNamed("4", 5, 6)
	if err != nil || got != 15 {
		t.Fatalf("addTextNamed valid input = (%d, %v), want (15, nil)", got, err)
	}
	if got := sum(1, 2, 3); got != 6 {
		t.Fatalf("sum = %d, want 6", got)
	}
	values := []int{4, 5, 6}
	if got := sum(values...); got != 15 {
		t.Fatalf("sum expanded slice = %d, want 15", got)
	}

	if got := apply(2, 3, add); got != 5 {
		t.Fatalf("apply = %d, want 5", got)
	}
	if got := addWithOffset(10)(5); got != 15 {
		t.Fatalf("returned closure = %d, want 15", got)
	}

	var typed Combiner = add
	if reflect.TypeOf(typed).Name() != "Combiner" {
		t.Fatalf("typed function has type %v, want named Combiner", reflect.TypeOf(typed))
	}
	if got := applyNamed(2, 3, typed); got != 5 {
		t.Fatalf("applyNamed = %d, want 5", got)
	}
	if got := applyFunctionPointer(2, 3, &typed); got != 5 {
		t.Fatalf("applyFunctionPointer = %d, want 5", got)
	}

	valuesForExpansion := []int{1, 2}
	extraValues := []int{3, 4}
	if got := mutateFunctionArguments(valuesForExpansion, extraValues...); !reflect.DeepEqual(got, []int{99, 2, 88, 4}) {
		t.Fatalf("mutateFunctionArguments = %v, want [99 2 88 4]", got)
	}
	if extraValues[0] != 88 {
		t.Fatalf("expanded variadic slice = %v, want first element 88", extraValues)
	}
	if !errors.Is(parseError("bad"), ErrInvalidNumber) {
		t.Fatalf("parseError did not return ErrInvalidNumber")
	}
}
