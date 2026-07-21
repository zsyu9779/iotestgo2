package main

import "fmt"

// 链表节点定义
type Node struct {
	Value int
	Next  *Node
}

// 链表定义
type LinkedList struct {
	Head *Node
	Size int
}

// Add（追加）
func (list *LinkedList) Add(val int) {
	newNode := &Node{Value: val}
	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	list.Size++
}

// Remove（删除）
func (list *LinkedList) Remove(val int) bool {
	if list.Head == nil {
		return false
	}
	if list.Head.Value == val {
		list.Head = list.Head.Next
		list.Size--
		return true
	}
	current := list.Head
	for current.Next != nil {
		if current.Next.Value == val {
			current.Next = current.Next.Next
			list.Size--
			return true
		}
		current = current.Next
	}
	return false
}

// Print（打印）
func (list *LinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	ll := &LinkedList{}
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Print() // 10 -> 20 -> 30 -> nil

	ll.Remove(20)
	ll.Print() // 10 -> 30 -> nil

	// 使用 Slice 实现栈
	stack := []int{}
	stack = append(stack, 1) // 入栈
	stack = append(stack, 2)
	top := stack[len(stack)-1] // 查看栈顶
	fmt.Println("Top:", top)
	stack = stack[:len(stack)-1] // 出栈
	fmt.Println("Stack after pop:", stack)
}
