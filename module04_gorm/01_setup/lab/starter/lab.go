package starter

import "errors"

var ErrNotImplemented = errors.New("not implemented")

func ValidatePool(maxOpen, maxIdle int) error { return ErrNotImplemented }
