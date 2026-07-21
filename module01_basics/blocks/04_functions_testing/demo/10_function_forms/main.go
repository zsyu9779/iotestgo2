package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Combiner func(int, int) int

var ErrInvalidNumber = errors.New("invalid number")

func add(a, b int) int {
	return a + b
}

// PublicFunctionDemo 因首字母大写而导出；下面的辅助函数首字母小写，
// 因此只在当前包内可见。
func PublicFunctionDemo() (int, error) {
	return addTextNamed("10", 20, 30)
}

func addText(text string, values ...int) (int, error) {
	base, err := strconv.Atoi(text)
	if err != nil {
		return 0, ErrInvalidNumber
	}
	return base + sum(values...), nil
}

func addTextNamed(text string, values ...int) (result int, err error) {
	base, parseErr := strconv.Atoi(text)
	if parseErr != nil {
		err = ErrInvalidNumber
		return
	}
	result = base + sum(values...)
	return
}

func sum(values ...int) (total int) {
	for _, value := range values {
		total += value
	}
	return
}

func apply(a, b int, combine Combiner) int {
	return combine(a, b)
}

func applyNamed(a, b int, combine Combiner) int {
	return combine(a, b)
}

// 指向函数值的指针是合法 Go 写法，但直接传递函数值更符合惯用写法。
// 保留这个指针示例是为了说明旧项目中的语法，不把它称为独立的“函数指针”特性。
func applyFunctionPointer(a, b int, combine *Combiner) int {
	return (*combine)(a, b)
}

func addWithOffset(offset int) func(int) int {
	return func(value int) int {
		return offset + value
	}
}

func mutateFunctionArguments(values []int, extra ...int) []int {
	values[0] = 99
	extra[0] = 88
	return append(values, extra...)
}

func parseError(value string) error {
	if _, err := strconv.Atoi(value); err != nil {
		return ErrInvalidNumber
	}
	return nil
}

func counter(start int) func() int {
	current := start
	return func() int {
		current++
		return current
	}
}

func main() {
	result, err := addText("1", 2, 3)
	fmt.Println("private multi-value return:", result, err)
	ignored, _ := addText("2", 3)
	fmt.Println("discarded error with blank identifier:", ignored)
	result, err = PublicFunctionDemo()
	fmt.Println("public function:", result, err)
	result, err = addTextNamed("4", 5, 6)
	fmt.Println("named return:", result, err)

	var combine Combiner = add
	fmt.Println("function type:", combine(2, 3))
	rawCombine := add
	converted := Combiner(rawCombine)
	fmt.Println("explicit function type conversion:", converted(3, 4))
	fmt.Println("higher-order function:", apply(4, 5, add))
	fmt.Println("function value pointer:", applyFunctionPointer(6, 7, &combine))
	allocated := new(Combiner)
	*allocated = add
	fmt.Println("new function type pointer:", applyFunctionPointer(8, 9, allocated))
	fmt.Println("variadic:", sum(1, 2, 3))

	values := []int{1, 2, 3}
	fmt.Println("variadic slice expansion:", sum(values...))
	args := []int{7, 8}
	extra := []int{9, 10}
	fmt.Println("slice and variadic arguments:", mutateFunctionArguments(args, extra...), "original:", args, "expanded:", extra)

	next := counter(10)
	fmt.Println("closure state:", next(), next())
	another := counter(10)
	fmt.Println("independent closure state:", another())
}
