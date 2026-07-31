package scanner

import (
	"context"
	"errors"
)

var (
	ErrNotImplemented     = errors.New("not implemented")
	ErrInvalidRoot        = errors.New("root must be an existing directory")
	ErrEmptyNeedle        = errors.New("search text must not be empty")
	ErrInvalidWorkerCount = errors.New("worker count must be greater than zero")
)

type Result struct {
	Files         int
	MatchingLines int
}

func Scan(ctx context.Context, root, needle string, workers int) (Result, error) {
	return Result{}, ErrNotImplemented
}
