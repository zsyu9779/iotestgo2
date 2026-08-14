package solution

import "fmt"

func ValidatePool(maxOpen, maxIdle int) error {
	if maxOpen <= 0 || maxIdle < 0 || maxIdle > maxOpen {
		return fmt.Errorf("invalid pool: max_open=%d max_idle=%d", maxOpen, maxIdle)
	}
	return nil
}
