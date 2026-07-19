// 黑暗角落：range 变量复用 — 取地址、闭包捕获、goroutine 陷阱
package main

import (
	"fmt"
	"time"
)

// 陷阱 1：取地址——全指向同一个变量
func showRangeAddressTrap() {
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	// 全打印同一个值（最后一次迭代的 i 值）
	for _, p := range out {
		fmt.Printf("%v ", *p)
	}
	fmt.Printf("  (期望 0 1 2，实际全部相同！因为 &i 始终指向同一个变量)\n")
}

// 修复 1：循环体内复制一份
func fixRangeAddressTrap() {
	var out []*int
	for i := 0; i < 3; i++ {
		i := i // 关键：复制一份，新变量有自己的地址
		out = append(out, &i)
	}
	for _, p := range out {
		fmt.Printf("%v ", *p)
	}
	fmt.Printf("  (修复后正确)\n")
}

// 陷阱 2：goroutine 中捕获循环变量
func showGoroutineTrap() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("%d ", i) // 全部打印最后一个值
		}()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("  (goroutine 闭包捕获，全为同一个值)\n")
}

// 修复 2：通过参数传入
func fixGoroutineTrap() {
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Printf("%d ", n)
		}(i) // 值复制传入
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("  (通过参数传入后正确)\n")
}

// 陷阱 3：range 单变量是索引
func showRangeSingleVar() {
	s := []string{"one", "two", "three"}
	for v := range s {
		fmt.Printf("%d ", v) // 0 1 2（索引，而非值！）
	}
	fmt.Printf("  (单变量 range 返回的是索引，不是值)\n")
}

func main() {
	fmt.Println("=== Range 黑暗角落 ===")
	fmt.Println()
	fmt.Print("陷阱 1 (取地址): ")
	showRangeAddressTrap()
	fmt.Print("修复 1: ")
	fixRangeAddressTrap()
	fmt.Println()
	fmt.Print("陷阱 2 (goroutine): ")
	showGoroutineTrap()
	fmt.Print("修复 2: ")
	fixGoroutineTrap()
	fmt.Println()
	fmt.Print("陷阱 3 (单变量 range): ")
	showRangeSingleVar()
}
