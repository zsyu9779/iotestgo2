package main

import (
	"strconv"
	"strings"
	"testing"
)

// Benchmark: += vs strings.Builder vs bytes.Buffer

func BenchmarkStringPlusEqual(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += strconv.Itoa(i)
	}
	_ = s
}

func BenchmarkStringsBuilder(b *testing.B) {
	var sb strings.Builder
	sb.Grow(1024) // 预分配内存
	for i := 0; i < b.N; i++ {
		sb.WriteString(strconv.Itoa(i))
	}
	_ = sb.String()
}

func BenchmarkStringsBuilderNoGrow(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(strconv.Itoa(i))
	}
	_ = sb.String()
}

// 运行：go test -bench=. -benchmem
//
// 预期结果：
//   BenchmarkStringPlusEqual    极慢（大量分配）
//   BenchmarkStringsBuilder     最快（预分配内存）
//   BenchmarkStringsBuilderNoGrow 中间
//
// 教学要点：
//   1. 循环中不要用 += 拼接字符串
//   2. 优先用 strings.Builder（单线程）或 bytes.Buffer（多线程安全）
//   3. 能预估大小的用 sb.Grow() 预分配
