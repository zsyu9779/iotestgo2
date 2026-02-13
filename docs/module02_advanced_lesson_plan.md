# Module 02: Go 语言进阶 - 教师备课教案

**适用对象**: 已掌握 Go 基础语法，希望进阶并发编程与工程化实践的学员  
**总课时**: 预计 8-10 小时  
**教学目标**: 深入理解 Go 的并发模型 (GMP)，掌握接口抽象与错误处理哲学，能够编写高并发、可测试的生产级代码。

---

## 目录
1. [第 1 课: 接口 (Interfaces) 与多态 (01_interfaces)](#第-1-课-接口-interfaces-与多态-01_interfaces)
2. [第 2 课: 错误处理与 Defer 机制 (02_errors_defer)](#第-2-课-错误处理与-defer-机制-02_errors_defer)
3. [第 3 课: Goroutines 协程基础 (03_goroutines)](#第-3-课-goroutines-协程基础-03_goroutines)
4. [第 4 课: Channels 通道深度解析 (04_channels)](#第-4-课-channels-通道深度解析-04_channels)
5. [第 5 课: Context 上下文管理 (05_context)](#第-5-课-context-上下文管理-05_context)
6. [第 6 课: 并发安全与锁 (06_concurrency_safety)](#第-6-课-并发安全与锁-06_concurrency_safety)
7. [第 7 课: 单元测试与基准测试 (07_testing)](#第-7-课-单元测试与基准测试-07_testing)
8. [第 8 课: 系统交互与信号处理 (08_os_interaction)](#第-8-课-系统交互与信号处理-08_os_interaction)
9. [第 9 课: 文件 I/O 操作 (09_file_io)](#第-9-课-文件-io-操作-09_file_io)
10. [第 10 课: 反射机制 (10_reflection)](#第-10-课-反射机制-10_reflection)
11. [第 11 课: 运行时控制 (11_runtime_control)](#第-11-课-运行时控制-11_runtime_control)
12. [项目实战: 并发日志分析器 (project_log_analyzer)](#项目实战-并发日志分析器-project_log_analyzer)

---

## 第 1 课: 接口 (Interfaces) 与多态 (01_interfaces)

**源码路径**: `module02_advanced/01_interfaces/main.go`

### 1. 教学目标
- 理解 **隐式接口实现** (Duck Typing).
- 掌握空接口 `interface{}` (Any Type).
- 理解类型断言 (Type Assertion) 和 Type Switch.

### 2. 理论讲解
- **Go vs Java**: Java 需要 `implements` 关键字, Go 不需要. 只要实现了方法集, 就实现了接口.
- **接口值**: 底层包含 `(type, value)`. 即使 value 是 nil, 接口本身可能不为 nil.
- **组合**: 接口可以嵌入其他接口 (如 `io.ReadWriter` 嵌入了 `Reader` 和 `Writer`).

### 3. 代码演示脚本
- **多态**: 
  ```go
  type Animal interface { Speak() string }
  // Dog 和 Cat 都实现了 Speak(), 所以都是 Animal
  ```
- **空接口**: 类似 Java `Object`. `func Print(v interface{})`.
- **类型断言**:
  ```go
  s, ok := val.(string) // 安全断言
  if !ok { ... }
  ```

### 4. 实操指南
- **练习**: 定义 `Geometry` 接口 (Area, Perimeter), 让 `Rect` 和 `Circle` 实现它.
- **陷阱**: 将一个 nil 指针赋值给接口变量, 该接口变量 **不等于** nil. (这是一个经典坑).

---

## 第 2 课: 错误处理与 Defer 机制 (02_errors_defer)

**源码路径**: `module02_advanced/02_errors_defer/main.go`

### 1. 教学目标
- 理解 Go 的错误处理哲学: Errors are values.
- 掌握 `panic` 和 `recover`.
- 深入理解 `defer` 栈执行顺序.

### 2. 理论讲解
- **Error**: 只是一个接口 `type error interface { Error() string }`.
- **Panic**: 类似异常抛出, 但仅用于不可恢复的错误 (如数组越界, 空指针).
- **Recover**: 类似 `catch`, 只能在 `defer` 中使用.

### 3. 代码演示脚本
- **自定义错误**:
  ```go
  type MyError struct { Msg string; Code int }
  func (e *MyError) Error() string { return ... }
  ```
- **Defer 陷阱**: 在循环中使用 defer (可能导致资源耗尽).
- **Wrap Error**: 简要介绍 Go 1.13 `fmt.Errorf("%w", err)` 和 `errors.Is/As`.

### 4. 实操指南
- **任务**: 编写一个除法函数, 分母为 0 时 panic. 在 main 中捕获该 panic 并打印日志, 确保程序不崩溃.

---

## 第 3 课: Goroutines 协程基础 (03_goroutines)

**源码路径**: `module02_advanced/03_goroutines/main.go`

### 1. 教学目标
- **核心**: 理解 Goroutine vs OS Thread. (M:N 模型).
- 掌握 `go` 关键字启动协程.
- 使用 `sync.WaitGroup` 等待协程结束.

### 2. 理论讲解
- **轻量级**: 初始栈仅 2KB, 可动态伸缩. 单机可启动百万个.
- **并发 vs 并行**: 并发是结构, 并行是执行.
- **WaitGroup**: 计数器模式. `Add`, `Done`, `Wait`.

### 3. 代码演示脚本
- **主线程退出**: 如果 main 结束, 所有协程会被强制终止. 演示不加 WaitGroup 时没有任何输出的情况.
- **闭包陷阱**:
  ```go
  for i := 0; i < 3; i++ {
      go func() { fmt.Println(i) }() // 可能都打印 3
  }
  // 修正: go func(val int) { ... }(i)
  ```

### 4. 实操指南
- **练习**: 启动 10 个协程, 每个睡眠 1 秒后打印自己的 ID. 确保主程序等待所有协程完成.

---

## 第 4 课: Channels 通道深度解析 (04_channels)

**源码路径**: `module02_advanced/04_channels/main.go`

### 1. 教学目标
- 理解 "Don't communicate by sharing memory; share memory by communicating."
- 掌握无缓冲 (Unbuffered) vs 有缓冲 (Buffered) Channel.
- 掌握 `select` 多路复用.

### 2. 理论讲解
- **阻塞特性**: 无缓冲通道的读写是同步阻塞的 (握手).
- **Channel 状态**: nil (阻塞), open (正常), closed (读返回零值, 写 panic).
- **Select**: 类似 IO 多路复用, 随机选择一个可用的 case 执行.

### 3. 代码演示脚本
- **死锁 (Deadlock)**: 
  ```go
  ch := make(chan int)
  ch <- 1 // main goroutine 阻塞在这里, 没有接收者 -> 死锁
  ```
- **Range 遍历**: 必须先 `close(ch)` 才能结束 range 循环, 否则死锁.

### 4. 实操指南
- **任务**: 实现一个生产者-消费者模型. 1 个生产者发 10 个数, 3 个消费者抢着处理.

---

## 第 5 课: Context 上下文管理 (05_context)

**源码路径**: `module02_advanced/05_context/main.go`

### 1. 教学目标
- 理解 Context 树状结构.
- 掌握 `WithCancel`, `WithTimeout`, `WithValue`.
- 场景: HTTP 请求全链路超时控制.

### 2. 理论讲解
- **生命周期**: Context 应该是函数的一参数.
- **传递性**: 父 Context 取消, 所有子 Context 也会收到 Done 信号.

### 3. 代码演示脚本
- **超时控制**:
  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
  defer cancel()
  select {
  case <-time.After(3*time.Second): fmt.Println("Task done")
  case <-ctx.Done(): fmt.Println("Timeout!")
  }
  ```
- **Value**: 仅用于传递请求范围的数据 (Request ID, Token), 不要滥用.

### 4. 实操指南
- **任务**: 模拟一个长耗时任务 (5s). 使用 Context 设置 2s 超时, 确保任务能响应取消信号并退出.

---

## 第 6 课: 并发安全与锁 (06_concurrency_safety)

**源码路径**: `module02_advanced/06_concurrency_safety/main.go`

### 1. 教学目标
- 识别 **Race Condition** (竞态条件).
- 使用 `sync.Mutex` 和 `sync.RWMutex`.
- 了解 `sync/atomic` 原子操作.

### 2. 理论讲解
- **Data Race**: 多个协程同时读写同一内存且无同步.
- **Mutex**: 互斥锁. `Lock()` / `Unlock()`. 建议用 `defer Unlock()`.
- **RWMutex**: 读写锁. 适合读多写少的场景.

### 3. 代码演示脚本
- **并发计数器**:
  ```go
  count := 0
  // 启动 1000 个协程 count++
  // 结果通常 < 1000
  ```
- **检测**: 使用 `go run -race main.go` 检测竞态.

### 4. 实操指南
- **练习**: 修复上述并发计数器的 Bug (加锁或使用 atomic).

---

## 第 7 课: 单元测试与基准测试 (07_testing)

**源码路径**: `module02_advanced/07_testing`

### 1. 教学目标
- 掌握 `testing` 包. 文件名 `_test.go`.
- 表格驱动测试 (Table Driven Tests).
- 基准测试 (Benchmarks).

### 2. 理论讲解
- **Test**: `func TestXxx(t *testing.T)`.
- **Benchmark**: `func BenchmarkXxx(b *testing.B)`. `b.N` 自动调整.

### 3. 代码演示脚本
- **表格驱动**:
  ```go
  tests := []struct{in, want int}{ {1, 2}, {2, 4} }
  for _, tt := range tests { ... }
  ```
- **Subtests**: `t.Run("case1", func(t *testing.T){...})`.

### 4. 实操指南
- **任务**: 为之前的 "斐波那契数列" 函数编写测试和基准测试, 比较递归版和迭代版的性能差异.

---

## 第 8 课: 系统交互与信号处理 (08_os_interaction)

**源码路径**: `module02_advanced/08_os_interaction/main.go`

### 1. 教学目标
- 读取命令行参数 `os.Args`.
- 监听系统信号 (Ctrl+C, SIGTERM) 实现优雅退出 (Graceful Shutdown).

### 2. 代码演示脚本
- **优雅退出**:
  ```go
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)
  <-c // 阻塞直到收到信号
  // 执行清理逻辑
  ```

---

## 第 9 课: 文件 I/O 操作 (09_file_io)

**源码路径**: `module02_advanced/09_file_io/main.go`

### 1. 教学目标
- `os.Open` vs `os.ReadFile`.
- `bufio` 缓冲读写 (处理大文件).
- `io.Reader` / `io.Writer` 接口的通用性.

### 2. 实操指南
- **任务**: 编写一个程序, 逐行读取一个文本文件, 给每一行加上行号后写入新文件.

---

## 第 10 课: 反射机制 (10_reflection)

**源码路径**: `module02_advanced/10_reflection/main.go`

### 1. 教学目标
- 理解反射的三大定律 (Interface <-> Reflection Object).
- `reflect.TypeOf` 和 `reflect.ValueOf`.

### 2. 警告
- **Rob Pike**: "Reflection is never clear." 除非必要, 勿用反射 (性能差, 代码可读性低).
- **用途**: JSON 序列化库, ORM 框架.

---

## 第 11 课: 运行时控制 (11_runtime_control)

**源码路径**: `module02_advanced/11_runtime_control/main.go`

### 1. 教学目标
- `GOMAXPROCS`: 控制使用的 CPU 核心数.
- `Gosched`: 让出 CPU 时间片.

---

## 项目实战: 并发日志分析器 (project_log_analyzer)

**源码路径**: `module02_advanced/project_log_analyzer`

### 1. 项目概述
- 模拟处理海量日志文件.
- 统计日志中的 Error 数量, 或最频繁的 IP. 

### 2. 架构设计 (Pipeline Pattern)
- **Generator**: 读取文件行 -> 发送给 Channel A.
- **Worker Pool**: 多个 Worker 从 Channel A 取日志 -> 解析 -> 发送结果给 Channel B.
- **Aggregator**: 从 Channel B 汇总结果.

### 3. 教学重点
- **Fan-out / Fan-in**: 扇出(分发)与扇入(汇总)模式.
- **优雅停止**: 如何通知所有 Worker 停止工作 (关闭 Channel 或 Context).

### 4. 实操挑战
- **性能优化**: 使用 Benchmark 测试 Worker 数量对处理速度的影响 (1 vs 10 vs 100).
- **作业**: 修改程序, 支持从标准输入 (Stdin) 读取日志流 (类似 `grep`).
