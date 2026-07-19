package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. Variable Declaration
	var age int = 30
	fmt.Println("Age:", age)

	// Type Inference
	name := "Gopher" // Short declaration (only inside functions)
	fmt.Printf("Name: %s, Type: %T\n", name, name)

	// 2. Constants & Iota
	const pi = 3.14159
	// pi = 3.14 // Error: cannot assign to pi

	const (
		StatusPending  = iota // 0
		StatusActive          // 1
		StatusInactive        // 2
	)
	fmt.Println("Status:", StatusPending, StatusActive, StatusInactive)

	// 3. Basic Types
	// Go has specific sized integers: int8, int16, int32, int64, uint...
	var maxInt32 int32 = math.MaxInt32
	var overflow int32 = maxInt32 + 1 // This wraps around in runtime if not constant, but let's just show types
	fmt.Println("Max Int32:", maxInt32)
	fmt.Println("Overflow example (be careful):", overflow)

	// 4. Type Conversion (Explicit only!)
	var i int = 42
	var f float64 = float64(i)
	fmt.Println("Float:", f)
}
