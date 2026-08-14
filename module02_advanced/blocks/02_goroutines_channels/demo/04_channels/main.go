package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	defer close(ch) // 唯一发送方负责关闭 channel。
	for i := 0; i < 5; i++ {
		ch <- i // 发送数据。
	}
}

func showUnbufferedHandoff() {
	ch := make(chan int)
	sent := make(chan struct{})
	go func() {
		ch <- 42 // 直到接收方取值后，发送才完成。
		close(sent)
	}()
	value := <-ch
	<-sent
	fmt.Println("Unbuffered handoff:", value)
}

func showBufferedCapacity() {
	ch := make(chan string, 2)
	ch <- "first"
	ch <- "second"
	//ch <- "third"
	// 第三次发送会阻塞，直到有接收方；因此这里只发送两个值。
	fmt.Printf("Buffered before receive: len=%d cap=%d\n", len(ch), cap(ch))
	fmt.Println("Buffered values:", <-ch, <-ch)
}

func showClosedChannel() {
	ch := make(chan int, 1)
	ch <- 7
	close(ch)
	first, firstOK := <-ch
	second, secondOK := <-ch
	fmt.Printf("Closed channel: first=%d/%v second=%d/%v\n", first, firstOK, second, secondOK)
}

func showNilChannelDisablesCase() {
	ready := make(chan string, 1)
	ready <- "ready"
	var disabled <-chan string
	select {
	case <-disabled:
		fmt.Println("nil channel should not be selected")
	case message := <-ready:
		fmt.Println("Nil channel disabled case:", message)
	}
}

func showMultipleReadyCases() {
	left := make(chan string, 1)
	right := make(chan string, 1)
	left <- "left"
	right <- "right"
	select {
	case value := <-left:
		fmt.Println("One ready case selected:", value)
	case value := <-right:
		fmt.Println("One ready case selected:", value)
	}
}

func showNonBlockingSelect() {
	ch := make(chan int)
	select {
	case value := <-ch:
		fmt.Println(value)
	default:
		fmt.Println("default makes select non-blocking; loops must avoid busy spinning")
	}
}

func main() {
	//showUnbufferedHandoff()
	//showBufferedCapacity()
	//
	//ch := make(chan int)
	//go producer(ch)
	//
	////接收方读取到 channel 关闭；关闭责任仍由 producer 承担。
	//for val := range ch {
	//	fmt.Println("Received from producer:", val)
	//}
	//showClosedChannel()
	showNilChannelDisablesCase()
	showMultipleReadyCases()
	showNonBlockingSelect()
	//
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

	//for i := 0; i < 2; i++ {
	//select {
	//case msg1 := <-ch1:
	//	fmt.Println("Received from ch1:", msg1)
	//case msg2 := <-ch2:
	//	fmt.Println("Received from ch2:", msg2)
	//case <-time.After(3 * time.Second):
	//	fmt.Println("Timeout")
	//}
	//}

	//for {
	//	select {
	//	case msg1 := <-ch1:
	//		fmt.Println("Received from ch1:", msg1)
	//	case msg2 := <-ch2:
	//		fmt.Println("Received from ch2:", msg2)
	//	}
	//}
}
