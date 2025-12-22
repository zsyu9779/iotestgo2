package main

import (
	"testing"
)

func BenchmarkLogPipeline(b *testing.B) {
	// We run the pipeline b.N times? 
	// No, RunPipeline takes a count. We should probably set a fixed count and benchmark how long it takes,
	// OR we treat b.N as the number of logs to process.
	// Let's treat b.N as the number of logs.
	
	// Since RunPipeline creates channels and goroutines, we should be careful.
	// Let's just benchmark a fixed large number of logs to see throughput.
	
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
