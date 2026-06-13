package main

import (
	"context"
	"fmt"
	"time"
)

func callDownstream(ctx context.Context, latency time.Duration) error {
	select {
	case <-time.After(latency):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	err := callDownstream(ctx, 500*time.Millisecond)
	fmt.Printf("downstream result: %v\n", err)
}
