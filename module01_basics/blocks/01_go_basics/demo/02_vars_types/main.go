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
	// 1. Variable Declaration
	var age int = 30
	fmt.Println("Age:", age)

	// Type Inference
	name := "Gopher" // Short declaration (only inside functions)
	fmt.Printf("Name: %s, Type: %T\n", name, name)

	left, _, right := 1, 2, 3
	fmt.Println("blank identifier:", left, right)

	// 2. Constants & Iota
	const pi = 3.14159
	// pi = 3.14 // Error: cannot assign to pi
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

	// 3. Basic Types
	// Go has specific sized integers: int8, int16, int32, int64, uint...
	var maxInt32 int32 = math.MaxInt32
	var overflow int32 = maxInt32 + 1 // This wraps around in runtime if not constant, but let's just show types
	fmt.Println("Max Int32:", maxInt32)
	fmt.Println("Overflow example (be careful):", overflow)

	// 4. Type Conversion (Explicit only!)
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

	// 5. Zero values: every declared variable starts with a useful default.
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
