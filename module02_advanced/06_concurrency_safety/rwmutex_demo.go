package main

import (
	"fmt"
	"sync"
	"time"
)

// 演示 sync.RWMutex 读写锁
// 读读不互斥、读写互斥、写写互斥

type SafeCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *SafeCache) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	time.Sleep(5 * time.Millisecond) // 模拟读耗时
	return c.data[key]
}

func (c *SafeCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	time.Sleep(20 * time.Millisecond) // 模拟写耗时
	c.data[key] = value
}

func main() {
	fmt.Println("=== RWMutex 读写锁演示 ===")
	fmt.Println()
	fmt.Println("特性：")
	fmt.Println("  - 多个读协程可以同时持有读锁（读读不互斥）")
	fmt.Println("  - 读锁和写锁互斥（读写互斥）")
	fmt.Println("  - 写锁和写锁互斥（写写互斥）")
	fmt.Println()

	c := &SafeCache{data: map[string]string{"hello": "world"}}

	// 场景 1：并发读（不互斥，速度快）
	fmt.Println("--- 场景 1：10 个并发读 ---")
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			val := c.Get("hello")
			fmt.Printf("  reader %d: %s\n", id, val)
		}(i)
	}
	wg.Wait()
	fmt.Printf("  10 个并发读耗时: %v (远小于 10×5ms，证明并发读)\n\n", time.Since(start))

	// 场景 2：写操作互斥
	fmt.Println("--- 场景 2：3 个写操作（写写互斥） ---")
	start = time.Now()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			c.Set(fmt.Sprintf("key%d", id), fmt.Sprintf("val%d", id))
			fmt.Printf("  writer %d done\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Printf("  3 个写操作耗时: %v (约等于 3×20ms，证明写写互斥)\n\n", time.Since(start))

	// 场景 3：读写混合（读被写阻塞）
	fmt.Println("--- 场景 3：读写混合 ---")
	start = time.Now()
	// 先启动一个写
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Set("blocking", "write")
		fmt.Println("  write done")
	}()

	time.Sleep(3 * time.Millisecond) // 确保写先获得锁

	// 再启动读
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Get("hello")
		fmt.Println("  read done")
	}()

	wg.Wait()
	fmt.Printf("  读写混合耗时: %v (读必须等待写完成)\n", time.Since(start))
	fmt.Println()
	fmt.Println("结论：读多写少的场景用 RWMutex 代替 Mutex，读操作性能大幅提升")
	fmt.Println("Java 对比：ReadWriteLock / ReentrantReadWriteLock")
}
