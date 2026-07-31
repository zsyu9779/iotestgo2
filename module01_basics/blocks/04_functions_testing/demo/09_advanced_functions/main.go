package main

import "fmt"

func atLeast(min int) func(int) bool {
	return func(value int) bool {
		return value >= min
	}
}

func atMost(max int) func(int) bool {
	return func(value int) bool { return value < max }
}

//func atPrime(max int) func(int) bool {
//	return func(value int) bool {
//
//	}
//}

func filter(values []int, keep func(int) bool) []int {
	filtered := make([]int, 0, len(values))
	for _, value := range values {
		if keep(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func filter2(values []int, keeps... func(int) bool) []int {
	filtered := make([]int, 0, len(values))
	for _, value := range values {
		for _, keep := range keeps {
			//if keep(value) {
			//	filtered = append(filtered, value)
			//}
			if !keep(value) {
				continue
			}
		}
	}
	return filtered
}


func main() {
	passed := atLeast(60)
	passed2 := atMost(90)
	fmt.Println("75 passed:", passed(75))
	fmt.Println("filtered:", filter([]int{59, 60, 75}, passed))
	filter2([]int{59, 60, 75},passed,passed2)


	defer fmt.Println("end")
	fmt.Println("start")
}
