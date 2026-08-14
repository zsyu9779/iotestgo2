package solution

import "fmt"

func Transfer(from, to, amount int) (int, int, error) {
	if amount <= 0 || from < amount {
		return from, to, fmt.Errorf("insufficient balance")
	}
	return from - amount, to + amount, nil
}
