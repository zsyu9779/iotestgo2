package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 本示例介绍 Go runtime 控制，对应原项目中的 myparallel/myruntime 部分。

func main() {
	fmt.Println("=== 1. GOMAXPROCS Control ===")
	// runtime.NumCPU 返回当前进程可使用的逻辑 CPU 数量。
	numCPU := runtime.NumCPU()
	fmt.Printf("NumCPU: %d\n", numCPU)

	// runtime.GOMAXPROCS 设置同时执行 Go 代码的最大 CPU 数量，并返回旧值。
	prev := runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Previous GOMAXPROCS: %d, Set to: %d\n", prev, numCPU)

	fmt.Println("\n=== 2. Goroutine Scheduling (Gosched) ===")
	// runtime.Gosched 主动让出处理器，使其他 goroutine 有机会运行。

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 1 working...")
			// 让出 CPU，让另一个 goroutine 运行。
			runtime.Gosched()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 2 working...")
			// 让出 CPU。
			runtime.Gosched()
		}
	}()

	wg.Wait()

	fmt.Println("\n=== 3. Stack Trace / Caller Info ===")
	printCallerInfo()
}

func printCallerInfo() {
	// runtime.Caller 返回调用栈中的文件、行号和程序计数器。
	pc, file, line, ok := runtime.Caller(1) // 1 表示当前函数的调用者 main。
	if ok {
		fmt.Printf("Called from %s:%d (PC: %v)\n", file, line, pc)
	}
}
