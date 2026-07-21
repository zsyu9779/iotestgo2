package main

import (
	"fmt"
	"math"
)

type Config struct {
	Port    int
	Enabled bool
	Name    string
}

type UserID int
type UserIDAlias = int

func main() {
	// 1. 变量声明
	var age int = 30
	fmt.Println("Age:", age)

	// 类型推断
	name := "Gopher" // 短变量声明只能出现在函数体内
	fmt.Printf("Name: %s, Type: %T\n", name, name)

	left, _, right := 1, 2, 3
	fmt.Println("blank identifier:", left, right)

	// 2. 常量与 iota
	const pi = 3.14159
	// pi = 3.14 // 编译错误：不能给常量 pi 重新赋值
	const message = "abc"
	const messageLength = len(message)
	fmt.Println("const expression length:", messageLength)

	const (
		StatusPending  = iota // 0
		StatusActive          // 1
		StatusInactive        // 2
	)
	fmt.Println("Status:", StatusPending, StatusActive, StatusInactive)

	// 仅供参考，不想挨打请勿模仿：省略声明会复用上一行完整表达式。
	const (
		consA = iota // 0
		consB        // 1，复用 = iota
		consC        // 2，复用 = iota
		consD = 250  // iota 仍递增到 3，但当前表达式为 250
		consE        // 250，复用上一行完整表达式
		consF = iota // 5，显式恢复使用当前行的 iota
		consG        // 6，复用 = iota
	)
	fmt.Println("iota edge cases:", consA, consB, consC, consD, consE, consF, consG)

	const (
		resetA = iota
		resetB
	)
	fmt.Println("iota reset:", resetA, resetB)

	// 3. 基本类型
	// Go 提供明确位数的整数类型：int8、int16、int32、int64、uint 等。
	var maxInt32 int32 = math.MaxInt32
	var overflow int32 = maxInt32 + 1 // 非常量运行时计算会发生环绕，这里只观察类型
	fmt.Println("Max Int32:", maxInt32)
	fmt.Println("Overflow example (be careful):", overflow)

	// 4. 类型转换（必须显式转换）
	var i int = 42
	var f float64 = float64(i)
	fmt.Println("Float:", f)

	var rawID int = 12
	var userID UserID = UserID(rawID)
	var aliasID UserIDAlias = rawID
	fmt.Println("defined type conversion:", userID, "alias assignment:", aliasID)

	floatValue := 1.9
	truncated := int(floatValue)
	fmt.Println("truncated float:", truncated)

	// 5. 零值：每个声明的变量都会从一个有意义的默认值开始。
	var count int
	var ratio float64
	var enabled bool
	var label string
	fmt.Printf("Zero values: int=%d float=%.1f bool=%t string=%q\n", count, ratio, enabled, label)

	var cfg Config
	fmt.Printf("Struct zero value: %#v\n", cfg)

	var numbers []int
	var scores map[string]int
	fmt.Printf("Nil collections: slice nil=%t len=%d map nil=%t missing=%d\n",
		numbers == nil, len(numbers), scores == nil, scores["missing"])
	numbers = append(numbers, 1)
	scores = make(map[string]int)
	scores["Alice"] = 95
	fmt.Printf("After initialization: numbers=%v scores=%v\n", numbers, scores)
}
