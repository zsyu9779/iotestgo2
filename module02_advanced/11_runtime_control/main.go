package main

import (
	"fmt"
	"runtime"
	"sync"
)

// This module covers Runtime control in Go.
// It corresponds to the 'myparallel/myruntime' section from the original iotestgo.

func main() {
	fmt.Println("=== 1. GOMAXPROCS Control ===")
	// runtime.NumCPU() returns the number of logical CPUs usable by the current process.
	numCPU := runtime.NumCPU()
	fmt.Printf("NumCPU: %d\n", numCPU)

	// runtime.GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously.
	// Returning the previous setting.
	prev := runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Previous GOMAXPROCS: %d, Set to: %d\n", prev, numCPU)

	fmt.Println("\n=== 2. Goroutine Scheduling (Gosched) ===")
	// runtime.Gosched() yields the processor, allowing other goroutines to run.

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 1 working...")
			// Yield CPU to let Goroutine 2 run
			runtime.Gosched()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 2 working...")
			// Yield CPU
			runtime.Gosched()
		}
	}()

	wg.Wait()

	fmt.Println("\n=== 3. Stack Trace / Caller Info ===")
	printCallerInfo()
}

func printCallerInfo() {
	// runtime.Caller reports file and line number information about function invocations on the calling goroutine's stack.
	pc, file, line, ok := runtime.Caller(1) // 1 means the caller of this function (main)
	if ok {
		fmt.Printf("Called from %s:%d (PC: %v)\n", file, line, pc)
	}
}
