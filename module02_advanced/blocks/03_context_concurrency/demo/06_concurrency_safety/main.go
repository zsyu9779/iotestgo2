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

	// 3. sync.Once 保证初始化函数至多执行一次。
	var once sync.Once
	initializations := 0
	var onceWG sync.WaitGroup
	for i := 0; i < 10; i++ {
		onceWG.Add(1)
		go func() {
			defer onceWG.Done()
			once.Do(func() { initializations++ })
		}()
	}
	onceWG.Wait()
	fmt.Println("Initializations (Once):", initializations)

	// 竞态案例放在 teaching_failures 中隔离运行。
	time.Sleep(10 * time.Millisecond)

	fmt.Println()
	RunRWMutexDemo()
	fmt.Println()
	RunSyncMapDemo()
	fmt.Println()
	RunMapConcurrentDemo()
}
