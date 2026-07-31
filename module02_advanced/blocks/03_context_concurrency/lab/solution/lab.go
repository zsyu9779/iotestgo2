package lab

import (
	"context"
	"sync"
)

// RunCancellableTask 模拟一个可被 Context 打断的任务。
func RunCancellableTask(ctx context.Context, started chan<- struct{}) error {
	close(started)
	<-ctx.Done()
	return ctx.Err()
}

// SafeCounter 是一个使用互斥锁保护的计数器。
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Inc 将计数器加一。
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value 返回当前计数。
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// RunInParallel 启动指定数量的 goroutine，每个 goroutine 增加指定次数。
func RunInParallel(counter *SafeCounter, goroutines, increments int) {
	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				counter.Inc()
			}
		}()
	}
	wg.Wait()
}
