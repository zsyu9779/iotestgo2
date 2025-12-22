package main

import (
	"errors"
	"fmt"
)

// Custom Error
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
	// 1. Defer (LIFO)
	// Java: finally { cleanup(); }
	defer fmt.Println("Deferred 1: Cleanup resources")
	defer fmt.Println("Deferred 2: Closing file")
	fmt.Println("Main execution started")

	// 2. Error Handling
	err := doTask(true)
	if err != nil {
		fmt.Println("Handled error:", err)
		// Check specific error type
		var myErr *MyError
		if errors.As(err, &myErr) {
			fmt.Println("Custom error code:", myErr.Code)
		}
	}

	// 3. Panic & Recover
	// Java: throw new RuntimeException() ... catch
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
	// fmt.Println("This won't be executed")
}
