package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Go 语言高级函数特性完整演示 ===\n")

	// 1. 函数作为一等公民（函数变量）
	fmt.Println("1. 函数变量:")
	demoFunctionVariables()

	// 2. 匿名函数和闭包
	fmt.Println("\n2. 匿名函数和闭包:")
	demoAnonymousFunctions()

	// 3. 高阶函数和函数式编程
	fmt.Println("\n3. 高阶函数和函数式编程:")
	demoHigherOrderFunctions()

	// 4. 延迟执行和错误处理
	fmt.Println("\n4. 延迟执行和错误处理:")
	demoDeferAndErrorHandling()

	// 5. 函数组合和柯里化
	fmt.Println("\n5. 函数组合和柯里化:")
	demoFunctionComposition()

	// 6. 并发编程中的函数
	fmt.Println("\n6. 并发编程中的函数:")
	demoConcurrentFunctions()
}

// ================== 函数变量示例 ==================
func demoFunctionVariables() {
	// 函数类型声明
	type MathFunc func(int, int) int
	
	var operation MathFunc
	
	// 函数赋值
	operation = add
	fmt.Printf("加法: %d\n", operation(10, 5))
	
	operation = multiply
	fmt.Printf("乘法: %d\n", operation(10, 5))
	
	operation = func(a, b int) int {
		return a*a + b*b
	}
	fmt.Printf("自定义函数: %d\n", operation(3, 4))
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// ================== 匿名函数和闭包示例 ==================
func demoAnonymousFunctions() {
	// 立即执行匿名函数
	result := func(x, y int) int {
		return x*x + y*y
	}(3, 4)
	fmt.Printf("立即执行: %d\n", result)

	// 闭包：计数器
	counter := createCounter()
	fmt.Printf("计数器: %d, %d, %d\n", counter(), counter(), counter())

	// 闭包：配置生成器
	createLogger := func(prefix string) func(string) {
		return func(message string) {
			fmt.Printf("[%s] %s: %s\n", time.Now().Format("15:04:05"), prefix, message)
		}
	}
	
	infoLog := createLogger("INFO")
	errorLog := createLogger("ERROR")
	
	infoLog("应用程序启动")
	errorLog("发生了一个错误")
	
	// 闭包：状态保持
	bankAccount := createBankAccount(1000)
	fmt.Printf("余额: $%d\n", bankAccount())
	fmt.Printf("存款后: $%d\n", bankAccount(500))
	fmt.Printf("取款后: $%d\n", bankAccount(-200))
}

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func createBankAccount(initialBalance int) func(...int) int {
	balance := initialBalance
	return func(amounts ...int) int {
		if len(amounts) > 0 {
			balance += amounts[0]
		}
		return balance
	}
}

// ================== 高阶函数示例 ==================
func demoHigherOrderFunctions() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Map: 转换每个元素
	doubled := mapSlice(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("Map 加倍: %v\n", doubled)
	
	// Filter: 过滤元素
	evens := filterSlice(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Filter 偶数: %v\n", evens)
	
	// Reduce: 聚合计算
	sum := reduceSlice(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Printf("Reduce 求和: %d\n", sum)
	
	// 函数组合: Map + Filter
	result := mapSlice(
		filterSlice(numbers, func(n int) bool { return n > 5 }),
		func(n int) int { return n * 10 },
	)
	fmt.Printf("组合操作: %v\n", result)
}

func mapSlice[T any, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = fn(item)
	}
	return result
}

func filterSlice[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func reduceSlice[T any](slice []T, initial T, fn func(T, T) T) T {
	result := initial
	for _, item := range slice {
		result = fn(result, item)
	}
	return result
}

// ================== 延迟执行和错误处理 ==================
func demoDeferAndErrorHandling() {
	// 资源清理模式
	fmt.Println("打开资源...")
	defer fmt.Println("资源清理完成")
	
	// 多个 defer，LIFO 顺序执行
	defer fmt.Println("第三个 defer")
	defer fmt.Println("第二个 defer")
	defer fmt.Println("第一个 defer")
	
	// 错误处理函数
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %d\n", result)
	}
	
	// 带参数的 defer
	value := "初始值"
	defer func(val string) {
		fmt.Printf("Defer 捕获的值: %s\n", val)
	}(value)
	value = "修改后的值"
	fmt.Printf("最终值: %s\n", value)
}

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return a / b, nil
}

// ================== 函数组合和柯里化 ==================
func demoFunctionComposition() {
	// 函数柯里化
	add := curryAdd(5)
	fmt.Printf("柯里化加法: %d\n", add(10))
	
	multiply := curryMultiply(3)
	fmt.Printf("柯里化乘法: %d\n", multiply(4))
	
	// 函数组合
	processor := compose(
		func(s string) string { return strings.ToUpper(s) },
		func(s string) string { return "🚀 " + s + " 🚀" },
		func(s string) string { return s + "!" },
	)
	fmt.Printf("函数组合: %s\n", processor("hello world"))
	
	// 管道模式
	pipeline := createPipeline(
		func(n int) int { return n * 2 },
		func(n int) int { return n + 10 },
		func(n int) int { return n - 5 },
	)
	fmt.Printf("管道处理: %d\n", pipeline(8))
}

func curryAdd(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func curryMultiply(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}

func compose(functions ...func(string) string) func(string) string {
	return func(s string) string {
		result := s
		for i := len(functions) - 1; i >= 0; i-- {
			result = functions[i](result)
		}
		return result
	}
}

func createPipeline(functions ...func(int) int) func(int) int {
	return func(n int) int {
		result := n
		for _, fn := range functions {
			result = fn(result)
		}
		return result
	}
}

// ================== 并发编程中的函数 ==================
func demoConcurrentFunctions() {
	// Goroutine 中的匿名函数
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("来自 Goroutine 的消息")
	}()
	
	// 带参数的 Goroutine
	for i := 0; i < 3; i++ {
		go func(id int) {
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Goroutine %d 完成\n", id)
		}(i)
	}
	
	// 函数作为通信消息
	ch := make(chan func() string)
	
	go func() {
		ch <- func() string { return "消息1" }
		ch <- func() string { return "消息2" }
		close(ch)
	}()
	
	for fn := range ch {
		fmt.Printf("接收到: %s\n", fn())
	}
	
	time.Sleep(500 * time.Millisecond) // 等待 Goroutine 完成
}