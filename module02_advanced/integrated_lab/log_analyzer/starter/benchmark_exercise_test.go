//go:build exercise

package main

import (
	"strconv"
	"testing"
)

func BenchmarkLogPipeline(b *testing.B) {
	for _, workers := range []int{1, 5, 10} {
		b.Run(strconv.Itoa(workers)+"-workers", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RunPipeline(workers, 1000)
			}
		})
	}
}
