// 黑暗角落：并发 map 读写会 fatal、sync.Map 适用场景
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ⚠️ 演示并发写 map 会 fatal（仅演示，课堂谨慎运行）
// fatal error: concurrent map writes
func showConcurrentMapFatal() {
	fmt.Println("--- 并发写 map 会 fatal ---")
	fmt.Println("  运行: go run -race map_concurrent_demo.go")
	fmt.Println()

	// 这个代码如果取消注释运行，会触发 fatal：
	// m := make(map[int]int)
	// var wg sync.WaitGroup
	// for i := 0; i < 8; i++ {
	//     wg.Add(1)
	//     go func() {
	//         defer wg.Done()
	//         for j := 0; j < 1000; j++ {
	//             k := rand.Int()
	//             m[k] = m[k] + 1
	//         }
	//     }()
	// }
	// wg.Wait()
	fmt.Println("  并发读写 map 会触发 fatal error: concurrent map writes")
	fmt.Println("  Go runtime 检测到并发 map 访问就直接崩溃（不像 Java 的 ConcurrentModificationException）")
}

// 修复方案 1：map + Mutex（常规场景）
type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

func (s *SafeMap) Inc(key int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key]++
}

func (s *SafeMap) Get(key int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.m[key]
}

// 修复方案 2：map + RWMutex（读多写少场景）
type SafeMapRW struct {
	mu sync.RWMutex
	m  map[int]int
}

func (s *SafeMapRW) Get(key int) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m[key]
}

func (s *SafeMapRW) Inc(key int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key]++
}

// 修复方案 3：sync.Map（读多写少、键不相交场景）
func showSyncMapUsage() {
	var sm sync.Map

	sm.Store("key1", 100)
	sm.Store("key2", 200)

	if val, ok := sm.Load("key1"); ok {
		fmt.Printf("  key1 = %v\n", val)
	}

	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v = %v\n", key, value)
		return true
	})
}

func RunMapConcurrentDemo() {
	fmt.Println("=== 并发 Map 黑暗角落 ===")
	fmt.Println()

	showConcurrentMapFatal()

	fmt.Println("--- 修复方案对比 ---")
	fmt.Println()
	fmt.Println("  方案 1: map + sync.Mutex")
	fmt.Println("    适用：读写均衡的常规场景")
	fmt.Println()
	fmt.Println("  方案 2: map + sync.RWMutex")
	fmt.Println("    适用：读多写少场景（读不互斥，写互斥）")
	fmt.Println()
	fmt.Println("  方案 3: sync.Map")
	fmt.Println("    适用：读多写少 + 不同 goroutine 读写不同键")
	fmt.Println("    不适用：需要全局锁、需要批量操作")

	fmt.Println()
	showSyncMapUsage()

	fmt.Println()
	fmt.Println("结论：")
	fmt.Println("  - 千万不能裸用 map 做并发读写")
	fmt.Println("  - 首选 map + Mutex/RWMutex")
	fmt.Println("  - sync.Map 只在特定场景更优，不要盲目使用")
	fmt.Println()
	fmt.Println("Java 对比：")
	fmt.Println("  Java: ConcurrentHashMap 分片锁")
	fmt.Println("  Go:   sync.Map 内部也是分片（read map + dirty map）")

	_ = rand.Int
	_ = time.Now
}
