package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify when done
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// GOMAXPROCS
	fmt.Println("CPUs:", runtime.NumCPU())
	
	// 1. Goroutine
	// Java: new Thread(() -> { ... }).start();
	go func() {
		fmt.Println("Hello from detached goroutine")
	}()

	// 2. WaitGroup
	// Java: CountDownLatch
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	fmt.Println("Main waiting for workers...")
	wg.Wait()
	fmt.Println("All workers done")
}
