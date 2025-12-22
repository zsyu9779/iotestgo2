package main

import "fmt"

func main() {
	// 1. Arrays (Fixed size, Value type)
	var arr [3]int = [3]int{1, 2, 3}
	// Copy happens here!
	arrCopy := arr
	arrCopy[0] = 100
	fmt.Println("Original Array:", arr) // [1 2 3]
	fmt.Println("Copy Array:", arrCopy) // [100 2 3]

	// 2. Slices (Dynamic size, Reference-like)
	// Java: ArrayList<Integer> list = new ArrayList<>();
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

	// 3. Making Slices
	// make([]type, len, cap)
	dynamicSlice := make([]int, 0, 5)
	dynamicSlice = append(dynamicSlice, 1)
	fmt.Println("Dynamic Slice:", dynamicSlice)
}
