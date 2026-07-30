package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i // 发送数据。
		fmt.Println("Sent:", i)
		time.Sleep(500 * time.Millisecond)
	}
	close(ch) // 发送方关闭 channel。
}

func main() {
	// 1. 无缓冲 channel 是同步握手。
	// ch := make(chan int)

	// 2. 有缓冲 channel 可以暂存数据。
	ch := make(chan int, 2) // 缓冲容量为 2。

	go producer(ch)

	// 接收方读取到 channel 关闭。
	for val := range ch {
		fmt.Println("Received:", val)
	}

	// 3. 使用 select 多路复用。
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout")
		}
	}
}
