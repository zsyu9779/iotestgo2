package starter

import "errors"

var ErrNotImplemented = errors.New("not implemented")

func Transfer(from, to, amount int) (int, int, error) { return from, to, ErrNotImplemented }
