package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. 使用 WithCancel 传播取消信号。
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine canceled")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	fmt.Println("Main: canceling context")
	cancel()
	time.Sleep(1 * time.Second)

	// 2. 使用 WithTimeout 控制最长等待时间。
	fmt.Println("\n--- Timeout Example ---")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel2()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Operation finished")
	case <-ctx2.Done():
		fmt.Println("Operation timed out:", ctx2.Err())
	}
}
