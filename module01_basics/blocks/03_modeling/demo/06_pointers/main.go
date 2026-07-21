package main

import "fmt"

type Counter struct {
	Value int
}

func main() {
	// 1. 指针基础
	x := 10
	ptr := &x // x 的地址

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", ptr)
	fmt.Printf("Value via pointer: %d\n", *ptr)

	// 通过指针修改
	*ptr = 20
	fmt.Println("New value of x:", x)

	// 2. nil 指针
	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("Pointer is nil")
	}
	// *nilPtr = 1 // panic：无效内存地址或 nil 指针解引用

	// 3. 访问 Struct 指针字段时会自动解引用。
	counter := &Counter{Value: 1}
	counter.Value = 2
	(*counter).Value = 3
	fmt.Println("pointer field sugar:", counter.Value)

	// 4. 指针参数与值参数。接收者语义见 Struct 与方法 Demo；两种情况本质上都向函数传递一个值。
	val := 5
	modifyValue(val)
	fmt.Println("After value pass:", val) // 5

	modifyPointer(&val)
	fmt.Println("After pointer pass:", val) // 100
}

func modifyValue(n int) {
	n = 100
}

func modifyPointer(n *int) {
	*n = 100
}
