package main

import "fmt"

func deferLIFO() {
	defer fmt.Println("defer first registered")
	defer fmt.Println("defer second registered")
	fmt.Println("operation")
}

func deferArgumentEvaluation() {
	value := 1
	defer fmt.Println("argument evaluated now:", value)
	value = 2
	defer func() {
		fmt.Println("closure evaluated later:", value)
	}()
}

func main() {
	fmt.Println("LIFO:")
	deferLIFO()
	fmt.Println("argument timing:")
	deferArgumentEvaluation()
}
