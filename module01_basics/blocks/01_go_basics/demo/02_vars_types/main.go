package main

import (
	"fmt"
	"math"
)

type Config struct {
	Port    int
	Enabled bool
	Name    string
}

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

	// 5. Zero values: every declared variable starts with a useful default.
	var count int
	var ratio float64
	var enabled bool
	var label string
	fmt.Printf("Zero values: int=%d float=%.1f bool=%t string=%q\n", count, ratio, enabled, label)

	var cfg Config
	fmt.Printf("Struct zero value: %#v\n", cfg)

	var numbers []int
	var scores map[string]int
	fmt.Printf("Nil collections: slice nil=%t len=%d map nil=%t missing=%d\n",
		numbers == nil, len(numbers), scores == nil, scores["missing"])
	numbers = append(numbers, 1)
	scores = make(map[string]int)
	scores["Alice"] = 95
	fmt.Printf("After initialization: numbers=%v scores=%v\n", numbers, scores)
}
