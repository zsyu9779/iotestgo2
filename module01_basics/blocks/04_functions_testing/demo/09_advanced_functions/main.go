package main

import "fmt"

func atLeast(min int) func(int) bool {
	return func(value int) bool {
		return value >= min
	}
}

func main() {
	passed := atLeast(60)
	fmt.Println("75 passed:", passed(75))

	defer fmt.Println("end")
	fmt.Println("start")
}
