package main

import "fmt"

func main() {
	// 1. Arrays (Fixed size, Value type)
	var arr [3]int = [3]int{1, 2, 3}
	var arr2 = [...]int{1, 2, 3}
	var arr3 = [...]int{1: 2, 2: 3} // [0 2 3]
	fmt.Printf("arr2: %v\narr3: %v\n", arr2, arr3)
	// Copy happens here!
	arrCopy := arr
	arrCopy[0] = 100
	fmt.Println("Original Array:", arr) // [1 2 3]
	fmt.Println("Copy Array:", arrCopy) // [100 2 3]

	// 2. Slices (A value describing a view over an underlying array)
	slice := []int{1, 2, 3}
	fmt.Printf("Slice: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	// Append (might trigger reallocation)
	slice = append(slice, 4, 5)
	fmt.Printf("After Append: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	// Slicing
	subSlice := slice[1:3] // Index 1 inclusive, 3 exclusive -> [2, 3]
	fmt.Println("SubSlice:", subSlice)

	// Modification affects underlying array
	subSlice[0] = 999
	fmt.Println("Original Slice after sub-slice mod:", slice) // [1 999 3 4 5]

	// append may reuse the underlying array or allocate a new one; never rely on
	// a particular growth factor when explaining this behavior.

	// 3. Making Slices
	// make([]type, len, cap)
	dynamicSlice := make([]int, 0, 5)
	dynamicSlice = append(dynamicSlice, 1)
	fmt.Println("Dynamic Slice:", dynamicSlice)

	source := []int{10, 20, 30}
	destination := make([]int, 5)
	copied := copy(destination, source)
	fmt.Printf("copy count=%d dst=%v src=%v\n", copied, destination, source)
}
