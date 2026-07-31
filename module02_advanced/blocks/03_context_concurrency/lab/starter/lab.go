package lab

import (
	"context"
	"errors"
)

var ErrNotImplemented = errors.New("not implemented")

func RunCancellableTask(ctx context.Context, started chan<- struct{}) error {
	close(started)
	return ErrNotImplemented
}

type SafeCounter struct {
	value int
}

func (c *SafeCounter) Inc() {}

func (c *SafeCounter) Value() int { return c.value }

func RunInParallel(counter *SafeCounter, goroutines, increments int) {}
