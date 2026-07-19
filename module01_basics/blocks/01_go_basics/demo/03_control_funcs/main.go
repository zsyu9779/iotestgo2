package main

import "fmt"

func main() {
	// 1. If with initialization
	if score := 85; score >= 60 {
		fmt.Println("Passed with score:", score)
	}

	// 2. Switch (No break needed)
	role := "admin"
	switch role {
	case "admin":
		fmt.Println("Access granted")
		// fallthrough // Optional: continue to next case
	case "user":
		fmt.Println("Read only")
	default:
		fmt.Println("Access denied")
	}

	// 3. For loop (The only loop in Go)
	// Java: for (int i = 0; i < 5; i++)
	for i := 0; i < 3; i++ {
		fmt.Println("Loop:", i)
	}

	// While-like loop
	count := 0
	for count < 2 {
		fmt.Println("Count:", count)
		count++
	}

	// 4. Functions
	sum, diff := calculate(10, 5)
	fmt.Printf("Sum: %d, Diff: %d\n", sum, diff)

	// Anonymous function / Closure
	greet := func(n string) {
		fmt.Println("Hello,", n)
	}
	greet("Closure")
}

// Function with multiple return values
func calculate(a, b int) (int, int) {
	return a + b, a - b
}
