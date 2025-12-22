# Module 02: Go 语言进阶

本模块包含 Go 语言的高级特性和并发编程，适合有一定基础的开发者深入学习。

## 目录结构

### 01_interfaces/
- **main.go**: 接口和抽象
- 学习内容：接口定义、接口实现、空接口、类型断言

### 02_errors_defer/
- **main.go**: 错误处理和延迟执行
- 学习内容：error 接口、errors.New、fmt.Errorf、defer、panic、recover

### 03_goroutines/
- **main.go**: 协程
- 学习内容：go 关键字、goroutine 创建、并发执行

### 04_channels/
- **main.go**: 通道
- 学习内容：channel 创建、发送和接收、缓冲通道、select 语句

### 05_context/
- **main.go**: 上下文
- 学习内容：context.Context、超时控制、取消传播、值传递

### 06_concurrency_safety/
- **main.go**: 并发安全
- 学习内容：竞态条件、互斥锁(sync.Mutex)、原子操作、sync.WaitGroup

### 07_testing/
- **calc.go**: 计算器实现
- **calc_test.go**: 单元测试
- 学习内容：testing 包、表格驱动测试、基准测试、代码覆盖率

### project_log_analyzer/
- **main.go**: 日志分析器项目
- **benchmark_test.go**: 性能基准测试
- 学习内容：综合应用并发编程、文件处理、性能优化

## 学习目标

1. 掌握接口的设计和使用
2. 熟练处理错误和使用 defer
3. 理解 goroutine 和并发模型
4. 掌握 channel 的使用和模式
5. 熟练使用 context 进行超时和取消控制
6. 确保并发程序的安全性
7. 编写高质量的单元测试和基准测试
8. 完成一个并发处理的综合项目

## 运行方式

每个目录下的程序都可以通过以下命令运行：
```bash
cd 目录名
go run main.go
```

对于测试相关的目录：
```bash
cd 07_testing/
go test -v

cd project_log_analyzer/
go test -bench=. -benchmem
```