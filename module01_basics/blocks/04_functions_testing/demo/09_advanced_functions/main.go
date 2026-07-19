package main

import "fmt"

func atLeast(min int) func(int) bool {
	return func(value int) bool {
		return value >= min
	}
}

func filter(values []int, keep func(int) bool) []int {
	filtered := make([]int, 0, len(values))
	for _, value := range values {
		if keep(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func main() {
	passed := atLeast(60)
	fmt.Println("75 passed:", passed(75))
	fmt.Println("filtered:", filter([]int{59, 60, 75}, passed))

	defer fmt.Println("end")
	fmt.Println("start")
}
