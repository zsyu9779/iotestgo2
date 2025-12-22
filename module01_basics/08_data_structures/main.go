package main

import "fmt"

// Node definition for LinkedList
type Node struct {
	Value int
	Next  *Node
}

// LinkedList definition
type LinkedList struct {
	Head *Node
	Size int
}

// Add (Append)
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

// Remove
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

// Print
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

	// Stack using Slice
	stack := []int{}
	stack = append(stack, 1) // Push
	stack = append(stack, 2)
	top := stack[len(stack)-1] // Peek
	fmt.Println("Top:", top)
	stack = stack[:len(stack)-1] // Pop
	fmt.Println("Stack after pop:", stack)
}
