package main

import "fmt"

func main() {
	// 1. 数组（固定长度，值类型）
	var arr [3]int = [3]int{1, 2, 3}
	var arr2 = [...]int{1, 2, 3}
	var arr3 = [...]int{1: 2, 2: 3} // [0 2 3]
	fmt.Printf("arr2: %v\narr3: %v\n", arr2, arr3)
	// 这里发生数组元素复制！
	arrCopy := arr
	arrCopy[0] = 100
	fmt.Println("Original Array:", arr) // [1 2 3]
	fmt.Println("Copy Array:", arrCopy) // [100 2 3]

	// 2. Slice（描述底层数组视图的一个值）
	slice := []int{1, 2, 3}
	fmt.Printf("Slice: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	// append（可能触发重新分配）
	slice = append(slice, 4, 5)
	fmt.Printf("After Append: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	// 切片
	subSlice := slice[1:3] // 下标 1 包含、3 不包含，结果为 [2, 3]
	fmt.Println("SubSlice:", subSlice)

	// 修改会影响底层数组
	subSlice[0] = 999
	fmt.Println("Original Slice after sub-slice mod:", slice) // [1 999 3 4 5]

	// append 可能复用底层数组，也可能分配新数组；讲解时不要依赖某个固定的扩容倍数。

	// 3. 创建 Slice
	// make([]type, len, cap)
	dynamicSlice := make([]int, 0, 5)
	dynamicSlice = append(dynamicSlice, 1)
	fmt.Println("Dynamic Slice:", dynamicSlice)
	for i := 0; i < 6; i++ {
		dynamicSlice = append(dynamicSlice, i)
	}
	fmt.Println("Dynamic Slice:", dynamicSlice)

	source := []int{10, 20, 30}
	destination := make([]int, 5)
	copied := copy(destination, source)
	fmt.Printf("copy count=%d dst=%v src=%v\n", copied, destination, source)
}
