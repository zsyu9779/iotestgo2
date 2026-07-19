package main

import "fmt"

func main() {
	// 1. If with initialization
	if score := 85; score >= 60 {
		fmt.Println("Passed with score:", score)
	}

	// 2. Switch (No break needed; a case may contain multiple values)
	role := "owner"
	switch role {
	case "admin", "owner":
		fmt.Println("Full access")
	case "user", "viewer":
		fmt.Println("Read only")
	default:
		fmt.Println("Access denied")
	}

	// fallthrough 不会自动发生，只会把执行继续交给紧邻的下一个 case。
	switch role {
	case "owner":
		fmt.Println("Owner branch")
		fallthrough
	case "admin":
		fmt.Println("Admin branch reached by fallthrough")
	}

	// 3. For loop (The only loop in Go)
	// Java: for (int i = 0; i < 5; i++)
	for i := 0; i < 3; i++ {
		if i == 1 {
			continue
		}
		fmt.Println("Loop:", i)
	}

	// While-like loop
	count := 0
	for count < 2 {
		fmt.Println("Count:", count)
		count++
	}

	switch {
	case count == 0:
		fmt.Println("Expressionless switch: empty")
	case count < 3:
		fmt.Println("Expressionless switch: small")
	default:
		fmt.Println("Expressionless switch: large")
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
