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
	// 1. Mutex
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

	// 2. Atomic
	var ops uint64
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&ops, 1)
		}()
	}
	wg.Wait()
	fmt.Println("Ops Value (Atomic):", ops)

	// Race condition example (if run with go run -race)
	// var x int
	// go func() { x++ }()
	// fmt.Println(x)
	time.Sleep(100 * time.Millisecond)
}
