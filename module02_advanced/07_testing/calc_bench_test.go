package main

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}

func BenchmarkAddLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1000000, 2000000)
	}
}
