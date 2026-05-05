package main

import (
	"fmt"
	"sync"
)

// 演示 sync.Map：适合读多写少、不同 goroutine 读写不同键的场景
func RunSyncMapDemo() {
	fmt.Println("=== sync.Map 演示 ===")
	fmt.Println()
	fmt.Println("sync.Map vs map+Mutex：")
	fmt.Println("  sync.Map 适合：")
	fmt.Println("    1. 读多写少（缓存场景）")
	fmt.Println("    2. 多个 goroutine 读写不相交的键")
	fmt.Println("  map+Mutex 适合：")
	fmt.Println("    1. 读写均衡")
	fmt.Println("    2. 需要批量操作或范围锁")
	fmt.Println()

	var sm sync.Map

	// Store：写入
	sm.Store("app_name", "iotestgo2")
	sm.Store("version", "1.0")
	sm.Store("author", "Gopher")

	// Load：读取
	if val, ok := sm.Load("app_name"); ok {
		fmt.Printf("  app_name = %s\n", val)
	}

	// LoadOrStore：有则返回，无则存储
	actual, loaded := sm.LoadOrStore("version", "2.0")
	fmt.Printf("  version: actual=%s, already_loaded=%v\n", actual, loaded)

	actual, loaded = sm.LoadOrStore("new_key", "new_value")
	fmt.Printf("  new_key: actual=%s, already_loaded=%v\n", actual, loaded)

	// Delete
	sm.Delete("author")

	// Range：遍历（不保证顺序）
	fmt.Println()
	fmt.Println("  遍历所有键值：")
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("    %s = %s\n", key, value)
		return true // 返回 true 继续遍历
	})

	// Load（已删除的 key 不存在了）
	if _, ok := sm.Load("author"); !ok {
		fmt.Println("  author 已被删除")
	}

	fmt.Println()
	fmt.Println("--- 并发场景演示 ---")
	fmt.Println()

	// 并发写入不同 key
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("goroutine_%d", id)
			sm.Store(key, id)
			fmt.Printf("  goroutine %d stored\n", id)
		}(i)
	}
	wg.Wait()

	count := 0
	sm.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	fmt.Printf("  总共 %d 个键（含初始 3 个 + 并发写入 5 个）\n", count)
}
