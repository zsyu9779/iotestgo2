package main

import "fmt"

type Config struct {
	Port    int
	Enabled bool
	Name    string
}

func main() {
	var count int
	var ratio float64
	var enabled bool
	var name string

	fmt.Printf("scalar zero values: int=%d float=%.1f bool=%t string=%q\n", count, ratio, enabled, name)

	var cfg Config
	fmt.Printf("struct zero value: %#v\n", cfg)

	var numbers []int
	var scores map[string]int
	fmt.Printf("nil slice: nil=%t len=%d\n", numbers == nil, len(numbers))
	fmt.Printf("nil map: nil=%t read-missing=%d\n", scores == nil, scores["missing"])

	// A nil Slice can be appended to; a nil Map must be initialized before writing.
	numbers = append(numbers, 1)
	scores = make(map[string]int)
	scores["Alice"] = 95
	fmt.Printf("after initialization: numbers=%v scores=%v\n", numbers, scores)
}
