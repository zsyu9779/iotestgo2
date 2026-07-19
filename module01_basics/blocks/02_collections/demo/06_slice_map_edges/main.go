package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Value string
}

func main() {
	array := [3]int{1, 2, 3}
	for _, value := range array {
		value = 99
		fmt.Print(value, " ")
	}
	fmt.Println()
	fmt.Println("array after changing range value:", array)
	for index := range array {
		array[index] = 99
	}
	fmt.Println("array after changing by index:", array)

	base := make([]int, 2, 4)
	base[0], base[1] = 1, 2
	view := base[:1]
	view[0] = 9
	view = append(view, 2, 3)
	fmt.Printf("shared slice: base=%v view=%v len=%d cap=%d\n", base, view, len(view), cap(view))

	full := make([]int, 2, 2)
	full[0], full[1] = 1, 2
	grown := append(full, 3)
	grown[0] = 9
	fmt.Printf("append after capacity: original=%v grown=%v\n", full, grown)

	var nilSlice []int
	nilSlice = append(nilSlice, 7)
	fmt.Printf("nil slice append: %v\n", nilSlice)

	var nilMap map[string]int
	fmt.Printf("nil map read: value=%d exists=%t\n", nilMap["missing"], false)
	nilMap = make(map[string]int)
	nilMap["done"] = 0
	value, exists := nilMap["done"]
	missing, missingExists := nilMap["missing"]
	fmt.Printf("comma-ok: done=%d/%t missing=%d/%t\n", value, exists, missing, missingExists)

	items := map[int]Item{1: {Value: "one"}}
	item := items[1]
	item.Value = "updated"
	items[1] = item
	itemPointers := map[int]*Item{1: {Value: "one"}}
	itemPointers[1].Value = "updated"
	fmt.Printf("map struct write-back=%q pointer-value=%q\n", items[1].Value, itemPointers[1].Value)

	unordered := map[int]string{2: "two", 1: "one", 3: "three"}
	keys := make([]int, 0, len(unordered))
	for key := range unordered {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	fmt.Printf("map keys shown stably=%v; raw iteration order is not a contract\n", keys)
}
