package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// pprof 性能分析初探：
// 生成 CPU profile、Memory profile、Goroutine profile
// 使用 go tool pprof 查看分析结果

func main() {
	fmt.Println("=== pprof 性能分析初探 ===")
	fmt.Println()

	// 1. CPU Profile
	fmt.Println("--- 1. CPU Profile ---")
	showCPUProfile()

	// 2. Memory Profile
	fmt.Println()
	fmt.Println("--- 2. Memory Profile ---")
	showMemoryProfile()

	// 3. Goroutine Profile
	fmt.Println()
	fmt.Println("--- 3. Goroutine Profile ---")
	showGoroutineProfile()

	fmt.Println()
	fmt.Println("=== 分析命令 ===")
	fmt.Println("  go tool pprof cpu.prof")
	fmt.Println("  go tool pprof mem.prof")
	fmt.Println("  go tool pprof goroutine.prof")
	fmt.Println()
	fmt.Println("  pprof 常用命令（进入交互后）：")
	fmt.Println("    top        - 按消耗排序前 N 项")
	fmt.Println("    list func  - 显示函数源码和耗时")
	fmt.Println("    web        - 生成调用图（需安装 graphviz）")
	fmt.Println("    png        - 导出为 PNG 图片")
	fmt.Println()
	fmt.Println("  启动 Web UI：")
	fmt.Println("    go tool pprof -http=:8080 cpu.prof")
}

func showCPUProfile() {
	// 创建 CPU profile 文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("  创建 cpu.prof 失败:", err)
		return
	}
	defer f.Close()

	// 开始 CPU profiling
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// 执行一些 CPU 密集操作
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += i * i
	}
	fmt.Printf("  计算结果: %d\n", sum)
	fmt.Println("  CPU profile 已写入 cpu.prof")
}

func showMemoryProfile() {
	// 分配一些内存
	var slices [][]int
	for i := 0; i < 100; i++ {
		s := make([]int, 10000)
		for j := range s {
			s[j] = j
		}
		slices = append(slices, s)
	}

	// 写入 memory profile
	f, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("  创建 mem.prof 失败:", err)
		return
	}
	defer f.Close()

	pprof.WriteHeapProfile(f)
	fmt.Printf("  Memory profile 已写入 mem.prof (%d slices)\n", len(slices))
	_ = slices
}

func showGoroutineProfile() {
	// 启动多个 goroutine
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(id int) {
			time.Sleep(500 * time.Millisecond)
			done <- true
		}(i)
	}

	time.Sleep(50 * time.Millisecond) // 确保 goroutine 已启动

	// 写入 goroutine profile
	f, err := os.Create("goroutine.prof")
	if err != nil {
		fmt.Println("  创建 goroutine.prof 失败:", err)
		return
	}
	defer f.Close()

	pprof.Lookup("goroutine").WriteTo(f, 0)
	fmt.Printf("  Goroutine profile 已写入 goroutine.prof (当前 %d goroutines)\n", runtime.NumGoroutine())

	// 等待 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}
}
