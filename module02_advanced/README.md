# Module 02: Go 语言进阶

本模块深入探讨 Go 语言的高级特性，超越基本语法，重点关注并发编程、系统交互和运行时控制。

## 目录结构

### 01_interfaces/
- **main.go**: 接口定义与实现
- 学习内容：接口的多态性、空接口、类型断言、组合

### 02_errors_defer/
- **main.go**: 错误处理与 Defer
- 学习内容：自定义错误、panic 与 recover、defer 执行顺序

### 03_goroutines/
- **main.go**: 协程基础
- 学习内容：go 关键字、协程生命周期、并发执行

### 04_channels/
- **main.go**: 通道操作
- 学习内容：无缓冲与缓冲通道、select 多路复用、关闭通道

### 05_context/
- **main.go**: 上下文管理
- 学习内容：Context 超时控制、取消信号、值传递

### 06_concurrency_safety/
- **main.go**: 并发安全
- 学习内容：sync.Mutex、sync.RWMutex、sync.WaitGroup、原子操作

### 07_testing/
- **calc.go**, **calc_test.go**: 单元测试
- 学习内容：Go 测试框架、表格驱动测试、基准测试 (Benchmark)

### 08_os_interaction/
- **main.go**: 系统交互
- 学习内容：信号处理 (Signal)、执行外部命令、环境变量

### 09_file_io/
- **main.go**: 文件 I/O 操作 (对应原 myio)
- 学习内容：os.Open/Create, bufio 读写, io.Copy, Seek 与 ReadAt

### 10_reflection/
- **main.go**: 反射机制 (对应原 myref)
- 学习内容：reflect.TypeOf/ValueOf, 动态修改字段值, 动态调用方法

### 11_runtime_control/
- **main.go**: 运行时控制 (对应原 myruntime)
- 学习内容：runtime.GOMAXPROCS, runtime.Gosched, CPU 核心数获取

### 12_stdlib_utils/
- **main.go**: 标准库工具
- 学习内容：常用标准库 (sort, time, json) 的使用技巧

### project_log_analyzer/
- **main.go**, **benchmark_test.go**: 日志分析器项目
- 学习内容：综合运用 Goroutines 和 Channels 实现并发日志处理流水线

## 学习目标

1. 深入理解 Go 接口与多态
2. 掌握健壮的错误处理机制
3. 熟练运用 Goroutine 和 Channel 进行并发编程
4. 学会使用 Context 管理并发任务
5. 理解并发安全与锁机制
6. 掌握 Go 测试与性能基准测试
7. 熟悉系统级交互与文件 I/O 操作
8. 理解反射与运行时控制的高级特性
