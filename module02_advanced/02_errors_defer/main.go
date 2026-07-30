package main

import (
	"errors"
	"fmt"
)

// MyError 是一个自定义错误类型。
type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

func doTask(fail bool) error {
	if fail {
		return &MyError{Code: 500, Msg: "Task failed"}
	}
	return nil
}

func main() {
	// 1. defer 按后进先出顺序执行。
	defer fmt.Println("Deferred 1: Cleanup resources")
	defer fmt.Println("Deferred 2: Closing file")
	fmt.Println("Main execution started")

	// 2. 显式处理错误返回值。
	err := doTask(true)
	if err != nil {
		fmt.Println("Handled error:", err)
		// 检查具体错误类型。
		var myErr *MyError
		if errors.As(err, &myErr) {
			fmt.Println("Custom error code:", myErr.Code)
		}
	}

	// 3. panic 和 recover 只作为边界兜底。
	safeFunction()
	fmt.Println("Main continues after recover")
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("About to panic...")
	panic("Something went terribly wrong!")
	// fmt.Println("这行不会执行")
}
