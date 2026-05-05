// 黑暗角落：相等性 — slice/map 不能 ==、DeepEqual 的坑与性能
package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

// 演示：slice 不能 ==
func showSliceNotComparable() {
	a := []int{1}
	b := []int{1}
	_ = a
	_ = b
	// _ = (a == b) // 编译错误：slice can only be compared to nil
	fmt.Println("slice 不能 == 比较（只能和 nil 比）")
}

// 演示：bytes.Equal vs DeepEqual 的语义差异（nil vs empty）
func showNilVsEmptyByteSlice() {
	var a []byte   // nil
	b := []byte{}  // empty
	fmt.Printf("bytes.Equal(nil, empty) = %v\n", bytes.Equal(a, b))     // true
	fmt.Printf("DeepEqual(nil, empty) = %v\n", reflect.DeepEqual(a, b)) // false!
	fmt.Println("→ bytes.Equal 认为 nil 和 empty 相等，DeepEqual 认为不等")
}

// ===== Benchmark：手写 Equals vs DeepEqual =====

type S struct {
	A int
	B string
	C []int
}

func (s S) Equals(o S) bool {
	if s.A != o.A || s.B != o.B || len(s.C) != len(o.C) {
		return false
	}
	for i := 0; i < len(s.C); i++ {
		if s.C[i] != o.C[i] {
			return false
		}
	}
	return true
}

var s1 = S{A: 1, B: "x", C: []int{1, 2, 3}}
var s2 = S{A: 1, B: "x", C: []int{1, 2, 3}}

func BenchmarkManualEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = s1.Equals(s2)
	}
}

func BenchmarkDeepEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reflect.DeepEqual(s1, s2)
	}
}

// 运行：go test -bench=. -benchmem
//
// 预期结果：
//   ManualEquals  > 10x faster than DeepEqual
//   DeepEqual 使用反射，有大量分配
//
// 教学要点：
//   1. slice/map 不能用 == 比较
//   2. reflect.DeepEqual 通用但慢，且语义有坑
//   3. []byte 用 bytes.Equal
//   4. 关键路径手写 Equals
//   5. NaN != NaN（DeepEqual 视为相等，手写 Equals 视为不等—看你需求）

func TestEqualityDemo(t *testing.T) {
	showSliceNotComparable()
	showNilVsEmptyByteSlice()
}
