package main

import (
	"testing"
)

func BenchmarkLogPipeline(b *testing.B) {
	// 每个子测试固定处理一批日志，比较不同 worker 数量的吞吐。

	b.Run("1000 logs, 1 processor", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RunPipeline(1, 1000)
		}
	})

	b.Run("1000 logs, 5 processors", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RunPipeline(5, 1000)
		}
	})

	b.Run("1000 logs, 10 processors", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RunPipeline(10, 1000)
		}
	})
}
