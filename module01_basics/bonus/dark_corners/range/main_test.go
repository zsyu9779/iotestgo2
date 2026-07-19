package main

import (
	"reflect"
	"testing"
)

func TestExplicitReuseSharesOneAddress(t *testing.T) {
	want := []int{3, 3, 3}
	if got := explicitReusePointerValues(); !reflect.DeepEqual(got, want) {
		t.Fatalf("explicitReusePointerValues() = %v, want %v", got, want)
	}
}

func TestLoopDeclarationCreatesPerIterationAddresses(t *testing.T) {
	want := []int{0, 1, 2}
	if got := loopDeclaredPointerValues(); !reflect.DeepEqual(got, want) {
		t.Fatalf("loopDeclaredPointerValues() = %v, want %v", got, want)
	}
}

func TestExplicitReuseClosuresObserveFinalValue(t *testing.T) {
	want := []int{3, 3, 3}
	if got := explicitReuseClosureValues(); !reflect.DeepEqual(got, want) {
		t.Fatalf("explicitReuseClosureValues() = %v, want %v", got, want)
	}
}

func TestLoopDeclaredClosuresKeepPerIterationValues(t *testing.T) {
	want := []int{0, 1, 2}
	if got := loopDeclaredClosureValues(); !reflect.DeepEqual(got, want) {
		t.Fatalf("loopDeclaredClosureValues() = %v, want %v", got, want)
	}
}
