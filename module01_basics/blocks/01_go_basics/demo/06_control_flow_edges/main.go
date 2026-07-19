package main

import "fmt"

func main() {
	if score := 85; score >= 60 {
		fmt.Println("if initializer: passed", score)
	}

	for i := 0; i < 4; i++ {
		if i == 1 {
			continue
		}
		if i == 3 {
			break
		}
		fmt.Println("for break/continue:", i)
	}

	count := 0
	for count < 2 {
		fmt.Println("condition-only for:", count)
		count++
	}

	role := "owner"
	switch role {
	case "admin", "owner":
		fmt.Println("multi-value case: full access")
	case "user", "viewer":
		fmt.Println("multi-value case: read only")
	default:
		fmt.Println("multi-value case: denied")
	}

	switch {
	case count == 0:
		fmt.Println("expressionless switch: empty")
	case count < 3:
		fmt.Println("expressionless switch: small")
	default:
		fmt.Println("expressionless switch: large")
	}

	switch role {
	case "owner":
		fmt.Println("fallthrough: owner")
		fallthrough
	case "admin":
		// fallthrough does not re-check the next case condition.
		fmt.Println("fallthrough: admin branch")
	}

	// Go does not fall through automatically. Explicit fallthrough is unusual
	// in business code because it couples neighboring cases unconditionally.
}
