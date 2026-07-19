// 黑暗角落：循环变量是按迭代创建，还是被显式复用
package main

import (
	"fmt"
	"sync"
)

func valuesAt(pointers []*int) []int {
	values := make([]int, len(pointers))
	for index, pointer := range pointers {
		values[index] = *pointer
	}
	return values
}

// explicitReusePointerValues 预先声明 i，再用赋值形式的 for 显式复用它。
func explicitReusePointerValues() []int {
	var pointers []*int
	var i int
	for i = 0; i < 3; i++ {
		pointers = append(pointers, &i)
	}
	return valuesAt(pointers)
}

// loopDeclaredPointerValues 使用 := 让 Go 1.22+ 为每次迭代创建新的 i。
func loopDeclaredPointerValues() []int {
	var pointers []*int
	for i := 0; i < 3; i++ {
		pointers = append(pointers, &i)
	}
	return valuesAt(pointers)
}

// explicitReuseClosureValues 让 goroutine 在循环结束后读取同一个 i。
// start 保证循环已经结束，WaitGroup 保证所有结果写入完成，因此没有数据竞态。
func explicitReuseClosureValues() []int {
	values := make([]int, 3)
	start := make(chan struct{})
	var workers sync.WaitGroup

	var i int
	for i = 0; i < len(values); i++ {
		workers.Add(1)
		go func(slot int) {
			defer workers.Done()
			<-start
			values[slot] = i
		}(i)
	}

	close(start)
	workers.Wait()
	return values
}

// loopDeclaredClosureValues 捕获由循环声明的每轮 i。
func loopDeclaredClosureValues() []int {
	values := make([]int, 3)
	start := make(chan struct{})
	var workers sync.WaitGroup

	for i := 0; i < len(values); i++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			<-start
			values[i] = i
		}()
	}

	close(start)
	workers.Wait()
	return values
}

// 保留原有知识点：range 只接收一个变量时，该变量是索引。
func showRangeSingleVar() {
	words := []string{"one", "two", "three"}
	for index := range words {
		fmt.Printf("%d ", index)
	}
	fmt.Println("(单变量 range 返回索引，而不是值)")
}

func main() {
	fmt.Println("=== Go 1.22+ 循环变量语义 ===")
	fmt.Println("循环用 := 声明变量时，每次迭代都有新变量；预先声明后用 = 才会显式复用。")
	fmt.Println()
	fmt.Println("显式复用变量的地址值:", explicitReusePointerValues())
	fmt.Println("循环声明变量的地址值:", loopDeclaredPointerValues())
	fmt.Println("显式复用变量的闭包值:", explicitReuseClosureValues())
	fmt.Println("循环声明变量的闭包值:", loopDeclaredClosureValues())
	fmt.Print("单变量 range: ")
	showRangeSingleVar()
}
