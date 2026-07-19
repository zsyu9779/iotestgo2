package main

import "fmt"

type Combiner func(int, int) int

func add(a, b int) int {
	return a + b
}

func sum(values ...int) (total int) {
	for _, value := range values {
		total += value
	}
	return
}

func apply(a, b int, combine Combiner) int {
	return combine(a, b)
}

func counter(start int) func() int {
	current := start
	return func() int {
		current++
		return current
	}
}

func main() {
	var combine Combiner = add
	fmt.Println("function type:", combine(2, 3))
	fmt.Println("higher-order function:", apply(4, 5, add))
	fmt.Println("variadic:", sum(1, 2, 3))

	values := []int{1, 2, 3}
	fmt.Println("variadic slice expansion:", sum(values...))

	next := counter(10)
	fmt.Println("closure state:", next(), next())
	another := counter(10)
	fmt.Println("independent closure state:", another())
}
