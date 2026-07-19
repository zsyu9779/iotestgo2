// Generics 泛型入门（Go 1.18+）
//
// Go 泛型特点：
// - 编译时类型安全（不像 Java 的类型擦除）
// - 通过类型约束（constraint）限制类型参数
// - 可运用于函数和类型
package main

import "fmt"

// ===== 1. 泛型函数 =====

// 类型约束：[T any] 接受任意类型
func First[T any](slice []T) T {
	return slice[0]
}

// 类型约束：[T comparable] 只能用于可比较类型（可用于 map key）
func Contains[T comparable](slice []T, val T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// 自定义约束接口：Number（int/float64）
type Number interface {
	int | int64 | float64
}

func Sum[T Number](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// 多类型参数
func Map[T any, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// ===== 2. 泛型类型 =====

// 泛型 Stack
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// 泛型 Set（用 map 实现）
type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func main() {
	fmt.Println("=== 泛型入门（Go 1.18+）===")
	fmt.Println()

	// 泛型函数
	fmt.Println("--- 泛型函数 ---")
	ints := []int{10, 20, 30}
	strings := []string{"a", "b", "c"}

	fmt.Printf("  First(%v) = %v\n", ints, First(ints))
	fmt.Printf("  First(%v) = %v\n", strings, First(strings))
	fmt.Printf("  Contains(%v, 20) = %v\n", ints, Contains(ints, 20))
	fmt.Printf("  Contains(%v, \"d\") = %v\n", strings, Contains(strings, "d"))
	fmt.Printf("  Sum(%v) = %v\n", ints, Sum(ints))

	// Map 函数
	squares := Map(ints, func(n int) int { return n * n })
	fmt.Printf("  Map(square, %v) = %v\n", ints, squares)

	// 泛型 Stack
	fmt.Println()
	fmt.Println("--- 泛型 Stack ---")
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	for val, ok := intStack.Pop(); ok; val, ok = intStack.Pop() {
		fmt.Printf("  Pop: %v\n", val)
	}

	// 泛型 Set
	fmt.Println()
	fmt.Println("--- 泛型 Set ---")
	set := NewSet[string]()
	set.Add("go")
	set.Add("rust")
	set.Add("go") // 重复不会报错
	fmt.Printf("  Set size: %d\n", set.Size())
	fmt.Printf("  Contains 'go': %v\n", set.Contains("go"))
	fmt.Printf("  Contains 'java': %v\n", set.Contains("java"))

	fmt.Println()
	fmt.Println("=== Go 泛型 vs Java 泛型 ===")
	fmt.Println()
	fmt.Println("  Go:")
	fmt.Println("    - 编译时具象化（monomorphization），不擦除类型")
	fmt.Println("    - 通过 interface 约束类型参数")
	fmt.Println("    - 无泛型继承（extends/super）")
	fmt.Println("    - 无通配符（? extends T, ? super T）")
	fmt.Println()
	fmt.Println("  Java:")
	fmt.Println("    - 类型擦除（编译后泛型信息丢失）")
	fmt.Println("    - 用 extends/super 限制类型参数")
	fmt.Println("    - 支持泛型继承和通配符")
	fmt.Println("    - 无法创建泛型数组（new T[]）→ Go 可以")
	fmt.Println()
	fmt.Println("  补充：Go 泛型不能做方法上单独泛型（方法接收者类型已确定）")
	fmt.Println("       Java 可以：<T> T getSomething()")
}
