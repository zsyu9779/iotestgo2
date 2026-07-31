package main

import (
	"fmt"
)

func main() {
	// 1. Map
	// Java 中相当于：Map<String, Integer> map = new HashMap<>();
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 88

	// 读取
	// val, exists = map[key]
	if val, ok := scores["Alice"]; ok {
		fmt.Println("Alice's score:", val)
	} else {
		fmt.Println("Alice not found")
	}

	delete(scores, "Bob")
	fmt.Println("Map:", scores)

	var nilScores map[string]int
	fmt.Println("Nil map read:", nilScores["Alice"])
	// nilScores["Alice"] = 1 // panic：向 nil map 写入元素

	for name, score := range scores {
		fmt.Println("Map iteration (order is not guaranteed):", name, score)
	}

	nested := make(map[int]map[int]string)
	nested[1] = make(map[int]string)
	nested[1][2] = "ready"
	fmt.Println("nested map value:", nested[1][2])

}
