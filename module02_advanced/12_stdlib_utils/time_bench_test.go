package main

import (
	"testing"
	"time"
)

// Benchmark: time.LoadLocation 缓存 vs 每次调用
//
// 教学要点：time.LoadLocation 每次调用可能读文件，要缓存 *time.Location

func BenchmarkLoadLocationEveryTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		_ = time.Now().In(loc)
	}
}

func BenchmarkLoadLocationCached(b *testing.B) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	for i := 0; i < b.N; i++ {
		_ = time.Now().In(loc)
	}
}

// 运行：go test -bench=. -benchmem
//
// 预期结果：
//   Cached 版本快 10-100 倍（无文件 IO）
//
// 修复模板：
//   var shanghaiLoc *time.Location
//   func init() {
//       shanghaiLoc, _ = time.LoadLocation("Asia/Shanghai")
//   }
//   // 后续直接用 shanghaiLoc，不必重复 LoadLocation
