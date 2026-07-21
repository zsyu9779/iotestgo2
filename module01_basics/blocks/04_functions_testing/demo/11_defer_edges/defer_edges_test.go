package main

import (
	"reflect"
	"testing"
)

func TestDeferArgumentAndClosureTiming(t *testing.T) {
	arguments, closures := deferTimingValues()
	if !reflect.DeepEqual(arguments, []int{1, 0}) {
		t.Fatalf("deferred arguments = %v, want [1 0]", arguments)
	}
	if !reflect.DeepEqual(closures, []int{1, 1}) {
		t.Fatalf("deferred closures = %v, want [1 1]", closures)
	}
}
