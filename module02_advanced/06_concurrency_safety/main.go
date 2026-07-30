package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// 1. 使用 Mutex 保护共享计数器。
	counter := SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println("Counter Value (Mutex):", counter.Value())

	// 2. 使用 Atomic 完成简单计数。
	var ops uint64
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			atomic.AddUint64(&ops, 1)
		}()
	}
	wg2.Wait()
	fmt.Println("Ops Value (Atomic):", ops)

	// 竞态示例：可使用 go run -race 检测。
	// var x int
	// go func() { x++ }()
	// fmt.Println(x)
	time.Sleep(100 * time.Millisecond)

	fmt.Println()
	RunRWMutexDemo()
	fmt.Println()
	RunSyncMapDemo()
	fmt.Println()
	RunMapConcurrentDemo()
}
