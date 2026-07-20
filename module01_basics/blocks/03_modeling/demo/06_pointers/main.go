package main

import "fmt"

type Counter struct {
	Value int
}

func main() {
	// 1. Basic Pointers
	x := 10
	ptr := &x // Address of x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", ptr)
	fmt.Printf("Value via pointer: %d\n", *ptr)

	// Modify via pointer
	*ptr = 20
	fmt.Println("New value of x:", x)

	// 2. Nil Pointer
	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("Pointer is nil")
	}
	// *nilPtr = 1 // Panic: invalid memory address or nil pointer dereference

	// 3. Struct pointer field access automatically dereferences the pointer.
	counter := &Counter{Value: 1}
	counter.Value = 2
	(*counter).Value = 3
	fmt.Println("pointer field sugar:", counter.Value)

	// 4. Pointer parameter vs value parameter. Receiver semantics are shown in
	// the Structs and Methods demo; both cases still pass a value to a function.
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
