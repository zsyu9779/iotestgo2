package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 通知任务已完成。
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(100 * time.Millisecond) // 模拟工作。
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// 使用 WaitGroup 等待每个 goroutine 完成；main 不会自动等待它们。
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	fmt.Println("Main waiting for workers...")
	wg.Wait()
	fmt.Println("All workers done")
}
