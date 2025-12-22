package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i // Send
		fmt.Println("Sent:", i)
		time.Sleep(500 * time.Millisecond)
	}
	close(ch) // Close channel
}

func main() {
	// 1. Unbuffered Channel (Synchronous)
	// ch := make(chan int)

	// 2. Buffered Channel
	// Java: BlockingQueue
	ch := make(chan int, 2) // Can hold 2 items without blocking sender

	go producer(ch)

	// Receiver
	for val := range ch {
		fmt.Println("Received:", val)
	}

	// 3. Select (Multiplexing)
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
