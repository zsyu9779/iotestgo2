package main

import "fmt"

func main() {
	// 1. If with initialization
	outerScore := 50
	if score := 85; score >= 60 {
		fmt.Println("Passed with score:", score)
	}
	fmt.Println("outer score:", outerScore)

	// 2. Switch initialization (No break needed; a case may contain multiple values)
	role := "owner"
	switch currentRole := role; currentRole {
	case "admin", "owner":
		fmt.Println("switch init:", currentRole, "has full access")
	case "user", "viewer":
		fmt.Println("switch init:", currentRole, "has read only access")
	default:
		fmt.Println("switch init:", currentRole, "is denied")
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
	for i := 0; i < 5; i++ {
		if i == 1 {
			continue
		}
		if i == 3 {
			break
		}
		fmt.Println("loop body:", i)
	}

	attempts := 0
	for {
		attempts++
		if attempts == 3 {
			break
		}
	}
	fmt.Println("infinite for stopped at:", attempts)

	count := 2
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
