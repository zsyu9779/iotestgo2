// 黑暗角落：defer 参数求值、recover 局限性、log.Fatal 副作用、Goroutine Panic
package main

import (
	"fmt"
	"log"
	"time"
)

// 演示 defer 参数立即求值
func showDeferEval() {
	s := "defer"
	defer fmt.Println("defer 参数:", s) // 此时 s="defer" 已被求值
	s = "original"
	fmt.Println("函数内值:", s)
	// 输出：
	//   函数内值: original
	//   defer 参数: defer   ← 注意！不是 original
}

// 演示 defer 只在函数返回时触发（不是离开代码块！）
func showDeferScope() {
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			defer func(i int) {
				fmt.Printf("%d ", i)
			}(i) // defer 一直等到 main 返回才执行
		}
	}
	// 输出: 6 3 0 (LIFO)
	fmt.Print("exiting → ")
}

// 演示 recover 只在 defer 内生效
func showRecoverOnlyInDefer() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recovered:", r)
		}
	}()
	panic("boom") // 后面的代码不会执行
	// fmt.Println("unreachable")
}

// 演示 goroutine panic 崩全局
func showGoroutinePanic() {
	go func() {
		panic("goroutine panicked!") // 会崩掉整个进程（不建议在教学环境运行）
	}()
}

// 修复模板：GoSafe 保护 goroutine
func GoSafe(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("GoSafe: recovered panic: %v\n", r)
			}
		}()
		fn()
	}()
}

// 演示 log.Fatal 不跑 defer
func showLogFatal() {
	fmt.Print("log.Fatal 演示: ")
	// 生产代码中：log.Fatal 调用 os.Exit(1)，不执行 defer
	// defer fmt.Println("never executed")
	// log.Fatal("fatal error")
	fmt.Println("(见注释：log.Fatal 会 os.Exit，不跑 defer！)")
	fmt.Println()
	fmt.Println("常见问题：")
	fmt.Println("  - defer 中的资源释放不会执行（文件关闭、连接归还等）")
	fmt.Println("  - defer 中的 Unlock 不会执行（导致死锁）")
	fmt.Println("  - 替代方案：用 panic + recover，或先 log 后 return")
}

func main() {
	fmt.Println("=== Defer/Panic 黑暗角落 ===")
	fmt.Println()

	fmt.Print("1. defer 参数立即求值: ")
	showDeferEval()
	fmt.Println()

	fmt.Print("2. defer 作用域: ")
	showDeferScope()
	fmt.Println()
	fmt.Println()

	fmt.Print("3. recover 只在 defer 内: ")
	showRecoverOnlyInDefer()
	fmt.Println()

	fmt.Println("4. goroutine panic 防护:")
	GoSafe(func() {
		panic("bad thing")
	})
	time.Sleep(50 * time.Millisecond)
	fmt.Println("   （进程继续运行，因为 GoSafe 保护了 goroutine）")
	fmt.Println()

	fmt.Println("5. log.Fatal 陷阱:")
	showLogFatal()
}
