package main

import "fmt"

func deferTimingValues() (arguments, closures []int) {
	value := 0
	for i := 0; i < 2; i++ {
		value = i
		defer func(captured int) {
			arguments = append(arguments, captured)
		}(value)
		defer func() {
			closures = append(closures, value)
		}()
	}
	return
}

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
	//fmt.Println("LIFO:")
	//deferLIFO()
	//fmt.Println("argument timing:")
	//deferArgumentEvaluation()
	arguments, closures := deferTimingValues()
	fmt.Println("deferred arguments:", arguments)
	fmt.Println("deferred closures:", closures)
}
