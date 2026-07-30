package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 通知任务已完成。
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // 模拟工作。
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// 查看可用 CPU 数量。
	fmt.Println("CPUs:", runtime.NumCPU())

	// 1. 启动 Goroutine。
	go func() {
		fmt.Println("Hello from detached goroutine")
	}()

	// 2. 使用 WaitGroup 等待任务完成。
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	fmt.Println("Main waiting for workers...")
	wg.Wait()
	fmt.Println("All workers done")
}
