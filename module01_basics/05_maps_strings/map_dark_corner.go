// 黑暗角落：Map 迭代顺序、nil map、值不可寻址
package main

import "fmt"

// 演示 map 迭代顺序未定义
func showMapIterationOrder() {
	m := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	fmt.Println("迭代顺序（5 次，注意每次可能不同）：")
	for t := 0; t < 5; t++ {
		for k := range m {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	// 添加新元素后，迭代顺序可能改变
	m[6] = 6
	m[7] = 7
	m[8] = 8
	fmt.Println("添加元素后：")
	for t := 0; t < 3; t++ {
		for k := range m {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}
}

// 演示 nil map 行为
func showNilMap() {
	var m map[int]int
	fmt.Println("nil map 长度:", len(m))     // OK: 0
	fmt.Println("nil map 读取:", m[10])      // OK: 0（不 panic）
	// m[10] = 1 // PANIC: assignment to entry in nil map

	// 正确初始化
	m = make(map[int]int)
	m[10] = 11
	fmt.Println("初始化后写入:", m[10])

	// comma-ok 检查存在性
	if val, ok := m[10]; ok {
		fmt.Printf("存在: %d\n", val)
	}
	if _, ok := m[99]; !ok {
		fmt.Println("不存在: 99")
	}
}

// 演示 map 值不可寻址
func showMapValueNotAddressable() {
	type Item struct {
		Value string
	}
	m := map[int]Item{1: {Value: "one"}}

	// 编译错误：cannot assign to struct field m[1].Value in map
	// m[1].Value = "two"

	// 正确方式：取出 → 修改 → 写回
	tmp := m[1]
	tmp.Value = "two"
	m[1] = tmp

	fmt.Println("修改后:", m[1].Value) // "two"

	// 或者用指针 map
	m2 := map[int]*Item{1: {Value: "one"}}
	m2[1].Value = "two" // OK：指针的值可寻址
	fmt.Println("指针 map 修改:", m2[1].Value)
}

// 演示 map 传参的"替换"误解
func showMapPassByValue() {
	m := map[int]int{1: 1}

	func(m map[int]int) {
		m[2] = 2 // 修改元素影响外部
	}(m)
	fmt.Println("添加后:", m) // {1:1, 2:2}

	// 但外部变量不会被替换
	func(m map[int]int) {
		m = make(map[int]int)
		m[3] = 3
	}(m)
	fmt.Println("替换后:", m) // 还是 {1:1, 2:2}，不是 {3:3}
}

func main() {
	fmt.Println("=== Map 黑暗角落 ===")
	fmt.Println()
	showMapIterationOrder()
	fmt.Println()
	showNilMap()
	fmt.Println()
	showMapValueNotAddressable()
	fmt.Println()
	showMapPassByValue()
}
