package main

import "fmt"

const primeCount = 100

// GenerateNatural 返回一条从 2 开始、按需生成自然数的只读数据流。
func GenerateNatural() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// PrimeFilter 返回一条新数据流，过滤掉 prime 的所有倍数。
func PrimeFilter(in <-chan int, prime int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-in
			if i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural()
	for i := 0; i < primeCount; i++ {
		prime := <-ch
		fmt.Printf("%d: %d\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
}
