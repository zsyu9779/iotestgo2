# D1/D2 可控失败式教学 Case 设计

## 1. 定位

这份文档用于设计穿插在 Day 1 和 Day 2 的“可控失败”教学 case。

所谓可控失败，不是为了让学员出错，而是刻意让学员用刚学过的知识解决一个真实但略超出当前能力边界的问题。学员先遇到现象，再产生问题，最后由讲师引入新概念或新工具。

推荐课堂节奏：

```text
已有知识尝试 -> 出现失败或不适 -> 描述现象 -> 抽象问题 -> 引入新概念 -> 修复方案 -> 形成原则
```

这类 case 最适合放在 D1/D2：

- D1/D2 的知识颗粒度小，失败现象直接，容易复现。
- 学员还在建立 Go 语义模型，适合通过反例建立边界感。
- 到 D3 之后，Gin、GORM、gRPC、go-zero 会引入框架、网络、数据库和工程结构。此时故意挖坑容易让问题来源变得混杂，应该更多使用事故复盘、代码评审和工程案例，而不是基础语法式撞墙。

## 2. 设计原则

### 2.1 坑要短

单个 case 控制在 5-20 分钟。

如果一个失败需要半小时才能定位，它就不适合作为 D1/D2 的课堂坑。它可以保留为课后挑战或 D7 复盘案例。

### 2.2 坑要可复现

优先选择稳定复现的现象：

- 编译错误
- panic
- deadlock
- `go run -race` 报告 data race
- 输出明显不符合预期
- benchmark 差异明显

对“偶现型”问题要谨慎，比如普通竞态导致的错误结果。可以先用 race detector 定性，再讨论为什么结果不稳定。

### 2.3 坑要有明确出口

每个 case 必须有对应的修复路径：

- 语言规则
- 标准库工具
- 推荐写法
- 工程约束

不要只让学生记住“这里会炸”。必须让学生带走一句可复用原则。

### 2.4 坑不能破坏课堂信任

讲师可以提前给一点提示：

> 这个练习故意留了一个不舒服的地方。先用你们当前掌握的方式做，等会儿我们一起复盘为什么 Go 不鼓励这样写。

这样学员会把失败理解为教学设计，而不是被捉弄。

### 2.5 D1/D2 与后续课程分工

| 阶段 | 推荐形式 | 典型问题 |
|---|---|---|
| D1 | 语义边界小坑 | 零值、slice、map、string、指针、struct |
| D2 | 并发与工程化小坑 | interface typed nil、defer、panic/recover、goroutine、channel、race、context |
| D3-D6 | 事故复盘与代码评审 | middleware 顺序、事务边界、N+1、超时传播、状态码语义、缓存一致性 |
| D7 | 综合系统复盘 | 把前面的小坑映射到真实后端系统问题 |

## 3. Case 模板

每个可控失败 case 建议按这个模板准备：

| 项目 | 内容 |
|---|---|
| 插入位置 | 放在哪一天、哪一节、哪个知识点之后 |
| 前置知识 | 学员此刻已经学过什么 |
| 任务描述 | 让学员尝试完成什么 |
| 预期失败 | 会出现什么错误、panic、死锁、异常输出或不适感 |
| 讲师追问 | 用哪些问题引导学生说出现象 |
| 引入概念 | 这个失败自然引出什么知识点 |
| 修复方案 | 用什么方式改正 |
| 带走原则 | 学员最后应该记住的一句话 |
| 时间控制 | 推荐耗时 |
| 风险控制 | 课堂上要避免什么 |

## 4. Day 1 Case 设计

### Case D1-01：未使用变量与导入

**插入位置**：Day 1 环境、`go run`、包和导入之后。

**前置知识**：`package main`、`import`、`func main`、`go run`。

**任务描述**：让学员在 Hello World 里多声明一个变量，或者多导入一个包但不使用。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	name := "go"
	fmt.Println("hello")
}
```

**预期失败**：编译报错：变量或包未使用。

**讲师追问**：

- 这在 Python/JavaScript 里通常只是 warning，为什么 Go 直接不让过？
- Go 这样设计是在限制你，还是在帮团队保持代码干净？

**引入概念**：Go 的编译器即工程约束，未使用代码不是风格问题，而是构建失败。

**修复方案**：删除未使用导入和变量；临时调试可以用 `_ = name`，但要说明这不是常规业务写法。

**带走原则**：Go 把一部分工程规范前置到了编译期。

**时间控制**：5 分钟。

**风险控制**：不要在这里展开太多语言哲学，点到为止。

### Case D1-02：短变量声明不能在函数外使用

**插入位置**：变量声明三种方式之后。

**前置知识**：`var`、`:=`、包级变量。

**任务描述**：让学员把函数内变量移动到文件顶部。

```go
package main

name := "task-manager"

func main() {}
```

**预期失败**：编译报错：函数体外不能使用非声明语句。

**讲师追问**：

- `:=` 到底是“声明语法”，还是“函数内的快捷语法”？
- 包级变量为什么要更显式？

**引入概念**：包级声明、函数作用域、短变量声明的使用边界。

**修复方案**：

```go
var name = "task-manager"
```

**带走原则**：`:=` 是函数内部的局部变量快捷声明，不是万能声明语法。

**时间控制**：5 分钟。

### Case D1-03：`if err := ...` 的作用域误判

**插入位置**：`if` 初始化语句、多返回值、错误处理入门之后。

**前置知识**：`if`、多返回值、`error` 初步概念。

**任务描述**：让学员在 `if err := do(); err != nil` 外部继续打印 `err`。

```go
if err := doSomething(); err != nil {
	fmt.Println(err)
}
fmt.Println(err)
```

**预期失败**：`err` 超出作用域。

**讲师追问**：

- 这个写法为什么适合处理“只在这里关心”的错误？
- 如果后面还要用 `err`，变量应该声明在哪里？

**引入概念**：局部作用域、错误就近处理。

**修复方案**：如果后续需要使用，提前声明；如果不需要，保持 `if` 内部短作用域。

**带走原则**：Go 倾向于让错误变量的生命周期尽量短。

**时间控制**：5-8 分钟。

### Case D1-04：slice append 后是否还共享底层数组

**插入位置**：数组、slice、`len/cap`、`append` 之后。

**前置知识**：slice 是数组视图，`append` 会改变长度和容量。

**任务描述**：让学员预测下面代码输出。

```go
arr := [4]int{1, 2, 3, 4}
s1 := arr[:2]
s2 := append(s1, 99)
s2[0] = 100
fmt.Println(arr, s1, s2)
```

随后改成：

```go
arr := [2]int{1, 2}
s1 := arr[:2]
s2 := append(s1, 99)
s2[0] = 100
fmt.Println(arr, s1, s2)
```

**预期失败**：学员对是否影响原数组判断不稳定。

**讲师追问**：

- `append` 一定会创建新数组吗？
- 什么时候还是原底层数组，什么时候换新底层数组？
- 为什么只看代码表面很难判断共享关系？

**引入概念**：slice header、`len`、`cap`、扩容。

**修复方案**：需要隔离时显式 `copy`；函数传参时不要偷偷保留并修改共享 slice。

**带走原则**：slice 是视图，`append` 是否扩容决定了后续修改会不会影响原数据。

**时间控制**：10-15 分钟。

**风险控制**：不要陷入 runtime 扩容细节，只讲“容量不足会换底层数组”。

### Case D1-05：nil map 可读不可写

**插入位置**：map 声明、`make`、读写之后。

**前置知识**：map 基本读写。

**任务描述**：让学员声明一个 map 后直接写入。

```go
var scores map[string]int
fmt.Println(scores["alice"])
scores["alice"] = 100
```

**预期失败**：读取返回零值，写入 panic。

**讲师追问**：

- 为什么读 nil map 不 panic，写 nil map 会 panic？
- `var m map[...]...` 和 `make(map[...]...)` 差在哪里？

**引入概念**：nil 引用类型、map 初始化。

**修复方案**：

```go
scores := make(map[string]int)
scores["alice"] = 100
```

**带走原则**：map 零值可读不可写；要写入必须先 `make` 或字面量初始化。

**时间控制**：8 分钟。

### Case D1-06：用零值判断 map key 是否存在

**插入位置**：map 统计词频之前。

**前置知识**：map 读取、零值。

**任务描述**：让学员判断某个用户是否存在。

```go
ages := map[string]int{
	"tom": 0,
}

if ages["tom"] == 0 {
	fmt.Println("not found")
}
```

**预期失败**：把真实存在的零值误判为不存在。

**讲师追问**：

- 读不存在的 key 返回什么？
- 如果业务值本身可能是零值，该怎么区分？

**引入概念**：comma-ok 模式。

**修复方案**：

```go
age, ok := ages["tom"]
if !ok {
	fmt.Println("not found")
}
fmt.Println(age)
```

**带走原则**：判断 key 是否存在，用 `value, ok := m[k]`，不要用零值猜。

**时间控制**：8 分钟。

### Case D1-07：string 的 `len` 不是字符数

**插入位置**：字符串、`range`、`rune` 之后。

**前置知识**：string、`len`、for range。

**任务描述**：让学员写一个函数统计 `"你好Go"` 的字符数。

```go
s := "你好Go"
fmt.Println(len(s))
```

**预期失败**：输出 8，而不是 4。

**讲师追问**：

- `len` 数的是字符，还是字节？
- `s[i]` 得到的是什么？
- `range` 字符串时为什么能正确遍历中文？

**引入概念**：UTF-8、byte、rune。

**修复方案**：

```go
fmt.Println(len([]rune(s)))
```

或使用 `for range` 计数。

**带走原则**：Go 字符串是只读字节序列；中文字符数要按 rune 处理。

**时间控制**：10 分钟。

### Case D1-08：map 中 struct 值不可直接改字段

**插入位置**：struct 与 map 结合之后，或者任务管理器项更新前。

**前置知识**：map、struct、字段赋值。

**任务描述**：让学员维护任务状态。

```go
type Task struct {
	Title string
	Done  bool
}

tasks := map[int]Task{
	1: {Title: "learn go"},
}

tasks[1].Done = true
```

**预期失败**：编译错误，不能给 map 索引表达式的字段赋值。

**讲师追问**：

- `tasks[1]` 返回的是一个可长期持有的地址，还是一个临时值？
- 如果 map 扩容或搬迁，直接拿元素地址会有什么问题？

**引入概念**：map 元素不可寻址、值语义。

**修复方案 1**：取出、修改、写回。

```go
task := tasks[1]
task.Done = true
tasks[1] = task
```

**修复方案 2**：map 存指针。

```go
tasks := map[int]*Task{
	1: {Title: "learn go"},
}
tasks[1].Done = true
```

**带走原则**：map 里放 struct 值时，修改字段要取出改完再写回；需要频繁原地修改可考虑存指针。

**时间控制**：10-15 分钟。

### Case D1-09：指针接收器和值接收器混用导致状态没变

**插入位置**：struct method、值接收器 vs 指针接收器之后。

**前置知识**：struct、method、指针。

**任务描述**：让学员给 Task 写一个完成任务的方法。

```go
type Task struct {
	Done bool
}

func (t Task) Finish() {
	t.Done = true
}

func main() {
	t := Task{}
	t.Finish()
	fmt.Println(t.Done)
}
```

**预期失败**：输出 `false`。

**讲师追问**：

- 方法接收器也是参数吗？
- 值接收器拿到的是原对象，还是副本？

**引入概念**：Go 只有值传递，接收器也是参数。

**修复方案**：

```go
func (t *Task) Finish() {
	t.Done = true
}
```

**带走原则**：方法需要修改对象状态时，用指针接收器。

**时间控制**：8-10 分钟。

## 5. Day 2 Case 设计

### Case D2-01：typed nil interface

**插入位置**：interface 底层 `(type, value)` 之后。

**前置知识**：interface、指针、nil。

**任务描述**：让学员判断下面的 `err` 是否为 nil。

```go
type MyError struct{}

func (e *MyError) Error() string {
	return "boom"
}

func do() error {
	var e *MyError = nil
	return e
}

func main() {
	err := do()
	fmt.Println(err == nil)
}
```

**预期失败**：输出 `false`，与直觉冲突。

**讲师追问**：

- `error` 接口内部此时有没有动态类型？
- 一个接口为 nil，需要 type 和 value 分别是什么状态？

**引入概念**：interface 内部结构、typed nil。

**修复方案**：无错误时直接 `return nil`，不要返回带具体类型的 nil 指针。

**带走原则**：接口 nil 不是看里面的指针值，而是看动态类型和值是否都为空。

**时间控制**：12-15 分钟。

### Case D2-02：defer 参数立即求值

**插入位置**：defer 基本语义之后。

**前置知识**：函数调用、defer。

**任务描述**：让学员预测输出。

```go
x := 1
defer fmt.Println(x)
x = 2
fmt.Println("end")
```

**预期失败**：学员以为 defer 打印 2，实际打印 1。

**讲师追问**：

- defer 延迟的是函数执行，还是参数求值？
- 如果想打印最新值，该怎么写？

**引入概念**：defer 参数求值时机。

**修复方案**：使用闭包。

```go
defer func() {
	fmt.Println(x)
}()
```

**带走原则**：defer 语句出现时，参数已经求值；延迟的是调用执行。

**时间控制**：8 分钟。

### Case D2-03：循环中 defer 导致资源释放太晚

**插入位置**：defer 和资源释放之后，文件 I/O 前也可复用。

**前置知识**：defer、循环、资源释放。

**任务描述**：让学员在循环里打开多个文件并 `defer Close()`。

```go
for _, name := range names {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	// process file
}
```

**预期失败**：短 demo 不一定炸，但语义上所有文件都等函数结束才关闭。

**讲师追问**：

- defer 是在循环本轮结束执行，还是函数返回执行？
- 如果有 10000 个文件，会发生什么？

**引入概念**：defer 作用域、资源生命周期。

**修复方案**：提取函数，让每次循环的 defer 绑定到小函数。

```go
func process(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
```

**带走原则**：defer 跟函数作用域绑定，不跟代码块或循环轮次绑定。

**时间控制**：10 分钟。

### Case D2-04：goroutine 还没执行 main 就退出

**插入位置**：`go` 关键字之后，WaitGroup 之前。

**前置知识**：函数调用、`go` 启动 goroutine。

**任务描述**：让学员启动 goroutine 打印 10 行。

```go
for i := 0; i < 10; i++ {
	go fmt.Println(i)
}
```

**预期失败**：可能没有输出或输出不完整。

**讲师追问**：

- main goroutine 退出时，其他 goroutine 会怎样？
- 我们需要一种什么机制等待它们？

**引入概念**：goroutine 生命周期、`sync.WaitGroup`。

**修复方案**：使用 WaitGroup。

**带走原则**：启动 goroutine 不等于等待 goroutine；main 退出会结束整个进程。

**时间控制**：8-10 分钟。

### Case D2-05：WaitGroup Add 放在 goroutine 内

**插入位置**：WaitGroup 基本用法之后。

**前置知识**：WaitGroup `Add/Done/Wait`。

**任务描述**：展示一个看起来合理但有竞态的写法。

```go
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
	go func() {
		wg.Add(1)
		defer wg.Done()
		fmt.Println("work")
	}()
}
wg.Wait()
```

**预期失败**：`Wait` 可能先执行，主流程提前结束。

**讲师追问**：

- `Add(1)` 是在任务开始前登记，还是任务内部才登记？
- 如果主 goroutine 先跑到 `Wait`，它看到的计数是多少？

**引入概念**：并发代码里的 happens-before 直觉。

**修复方案**：在启动 goroutine 之前 `Add(1)`。

```go
wg.Add(1)
go func() {
	defer wg.Done()
}()
```

**带走原则**：WaitGroup 要先登记任务，再启动任务。

**时间控制**：8 分钟。

### Case D2-06：goroutine 捕获外部复用变量

**插入位置**：goroutine 与 for 循环结合之后。

**前置知识**：for、闭包、goroutine。

**任务描述**：让学员并发打印编号。注意：当前项目使用 Go 1.25，Go 1.22 之后 `for` 循环变量语义已经改进，经典的 `for i := 0; i < n; i++` 捕获问题不再适合作为稳定课堂坑。这里改用外部复用变量来演示闭包捕获“变量本身”。

```go
i := 0
for ; i < 5; i++ {
	go func() {
		fmt.Println(i)
	}()
}
```

**预期失败**：多个 goroutine 可能打印相同或不符合预期的值，因为它们捕获的是同一个外部变量 `i`。

**讲师追问**：

- 闭包捕获的是变量，还是当时的值？
- 怎么显式把当时的值传进去？

**引入概念**：闭包捕获、goroutine 调度。

**修复方案**：

```go
for i := 0; i < 5; i++ {
	go func() {
		fmt.Println(i)
	}()
}
```

在 Go 1.22+，上面的写法已经能为每轮循环创建独立变量。为了兼容旧版本或表达得更显式，也可以写成：

```go
go func(v int) {
	fmt.Println(v)
}(i)
```

**带走原则**：并发闭包里要明确捕获值，不要依赖外部循环变量的时序。

**时间控制**：10 分钟。

### Case D2-07：用锁和 slice 手写生产-消费者

**插入位置**：goroutine、WaitGroup、Mutex 之后，channel 之前。

**前置知识**：goroutine、WaitGroup、Mutex、slice。

**任务描述**：实现 1 个生产者、3 个消费者，生产者生成 10 个任务，消费者处理任务。

第一轮要求暂时不能用 channel，只能用 `Mutex + []int + WaitGroup`。

**预期失败或不适**：

- 队列为空时消费者怎么办？
- 消费者什么时候退出？
- 是否需要轮询？
- 是否需要额外的 done 标记？
- 锁的范围如何控制？

**讲师追问**：

- 我们现在真正想共享的是内存，还是想传递任务？
- 消费者等待任务这件事，用锁表达自然吗？
- 有没有一种结构天然表达“发送、接收、关闭”？

**引入概念**：channel。

**修复方案**：改成 channel。

```go
jobs := make(chan int)

go func() {
	defer close(jobs)
	for i := 0; i < 10; i++ {
		jobs <- i
	}
}()

for w := 0; w < 3; w++ {
	wg.Add(1)
	go func(id int) {
		defer wg.Done()
		for job := range jobs {
			fmt.Println(id, job)
		}
	}(w)
}
wg.Wait()
```

**带走原则**：需要传递任务和表达等待时，channel 往往比共享队列加锁更贴近问题本身。

**时间控制**：20 分钟。

**风险控制**：第一轮不要要求学生写出完美版本，重点是让他们感受到“锁能做，但表达很别扭”。

### Case D2-08：channel 无接收者导致 deadlock

**插入位置**：channel 基本发送接收之后。

**前置知识**：channel、阻塞。

**任务描述**：让学员运行。

```go
ch := make(chan int)
ch <- 1
fmt.Println(<-ch)
```

**预期失败**：fatal error: all goroutines are asleep - deadlock。

**讲师追问**：

- 无缓冲 channel 的发送什么时候能完成？
- 谁在同时接收？

**引入概念**：无缓冲 channel 是同步交接，不是普通队列。

**修复方案**：

```go
go func() {
	ch <- 1
}()
fmt.Println(<-ch)
```

或使用缓冲 channel：

```go
ch := make(chan int, 1)
ch <- 1
fmt.Println(<-ch)
```

**带走原则**：无缓冲 channel 的发送和接收必须同时配对。

**时间控制**：8-10 分钟。

### Case D2-09：range channel 但没人 close

**插入位置**：channel close、range 之后。

**前置知识**：channel、`range`。

**任务描述**：让消费者 `range jobs`，但生产者发送完不关闭。

```go
jobs := make(chan int, 3)
jobs <- 1
jobs <- 2
jobs <- 3

for job := range jobs {
	fmt.Println(job)
}
```

**预期失败**：打印完已有值后 deadlock。

**讲师追问**：

- `range channel` 怎么知道没有下一个值了？
- close 应该由发送方还是接收方负责？

**引入概念**：关闭 channel 的语义。

**修复方案**：

```go
close(jobs)
```

**带走原则**：发送方负责关闭 channel；`range channel` 依赖 close 结束循环。

**时间控制**：8 分钟。

### Case D2-10：从多个 goroutine close 同一个 channel

**插入位置**：channel close 之后，worker pool 之前。

**前置知识**：channel close。

**任务描述**：让多个 worker 在处理完后尝试关闭结果 channel。

```go
results := make(chan int)
for i := 0; i < 3; i++ {
	go func() {
		results <- 1
		close(results)
	}()
}
```

**预期失败**：send on closed channel 或 close of closed channel。

**讲师追问**：

- 谁拥有 channel 的关闭权？
- 多个发送者时，怎么知道所有发送都结束了？

**引入概念**：channel 所有权、WaitGroup + 单点 close。

**修复方案**：worker 只发送结果，另起一个 goroutine 等待所有 worker 完成后关闭。

```go
go func() {
	wg.Wait()
	close(results)
}()
```

**带走原则**：close 是所有权动作；多个发送者场景下通常由协调者统一关闭。

**时间控制**：10-12 分钟。

### Case D2-11：并发计数器结果小于预期

**插入位置**：Mutex/RWMutex/atomic 之前。

**前置知识**：goroutine、WaitGroup。

**任务描述**：启动 1000 个 goroutine 做 `count++`。

```go
var count int
var wg sync.WaitGroup

for i := 0; i < 1000; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		count++
	}()
}

wg.Wait()
fmt.Println(count)
```

**预期失败**：结果可能小于 1000；`go run -race` 报告 data race。

**讲师追问**：

- `count++` 是一个原子操作吗？
- 多个 goroutine 同时读、改、写会发生什么？
- Go 有没有工具帮我们检测？

**引入概念**：data race、race detector、Mutex、atomic。

**修复方案**：`sync.Mutex` 或 `atomic.AddInt64`。

**带走原则**：共享变量的复合读写必须同步；`count++` 不是原子操作。

**时间控制**：15 分钟。

### Case D2-12：并发读写 map 直接 fatal

**插入位置**：并发安全与锁，map + goroutine 结合时。

**前置知识**：map、goroutine、WaitGroup。

**任务描述**：多个 goroutine 同时写 map。

```go
m := make(map[int]int)
var wg sync.WaitGroup

for i := 0; i < 8; i++ {
	wg.Add(1)
	go func(id int) {
		defer wg.Done()
		for j := 0; j < 1000; j++ {
			m[j] = id
		}
	}(i)
}

wg.Wait()
```

**预期失败**：fatal error: concurrent map writes。

**讲师追问**：

- 为什么普通变量竞态可能只是结果错，而 map 会直接 fatal？
- Go 的内置 map 是不是并发容器？
- `sync.Map` 是不是所有场景的默认答案？

**引入概念**：map 非并发安全、`RWMutex`、`sync.Map` 适用场景。

**修复方案 1**：`map + sync.Mutex`。

**修复方案 2**：读多写少时 `map + sync.RWMutex`。

**修复方案 3**：特定读多写少、key 相对稳定或不同 goroutine 操作不同 key 时考虑 `sync.Map`。

**带走原则**：裸 map 不能并发读写；常规业务优先 `map + Mutex/RWMutex`，`sync.Map` 不是万能替代。

**时间控制**：15 分钟。

**风险控制**：独立文件或独立命令运行，避免 fatal 影响后续课堂 demo。

### Case D2-13：读多写少为什么用 RWMutex

**插入位置**：Mutex 之后、RWMutex 之前。

**前置知识**：Mutex、map。

**任务描述**：做一个配置读取器，100 个 goroutine 读，1 个 goroutine 偶尔写。第一版全部用 Mutex。

**预期不适**：正确但所有读互相阻塞，表达不出“多个读可以并行”的语义。

**讲师追问**：

- 读操作之间是否互相破坏？
- 写操作和读操作能不能并行？
- 有没有锁能区分读和写？

**引入概念**：`sync.RWMutex`。

**修复方案**：读用 `RLock/RUnlock`，写用 `Lock/Unlock`。

**带走原则**：读多写少且读操作不改变状态时，用 RWMutex 表达读读并行、读写互斥。

**时间控制**：10 分钟。

### Case D2-14：goroutine panic 不能被 main recover

**插入位置**：panic/recover 与 goroutine 结合时。

**前置知识**：panic、recover、goroutine。

**任务描述**：在 main 里 defer recover，再让 goroutine panic。

```go
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()

	go func() {
		panic("worker failed")
	}()

	time.Sleep(time.Second)
}
```

**预期失败**：进程仍然崩溃。

**讲师追问**：

- recover 能跨 goroutine 捕获 panic 吗？
- 如果每个 worker 都可能 panic，保护应该放在哪里？

**引入概念**：panic 只沿当前 goroutine 栈展开。

**修复方案**：在 goroutine 内部 defer recover。

**带走原则**：recover 只能捕获同一个 goroutine 内的 panic。

**时间控制**：10 分钟。

### Case D2-15：context 创建了但 worker 不听

**插入位置**：Context `WithCancel/WithTimeout` 之后。

**前置知识**：goroutine、select、context。

**任务描述**：创建一个 1 秒超时的 context，但 worker 用 `time.Sleep(5 * time.Second)` 模拟任务。

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()

go func() {
	time.Sleep(5 * time.Second)
	fmt.Println("done")
}()

<-ctx.Done()
fmt.Println("timeout")
```

**预期失败或不适**：主流程超时了，但 worker 并没有真的响应取消。

**讲师追问**：

- context 是强制杀 goroutine，还是发取消信号？
- worker 要怎样“听见”这个信号？

**引入概念**：协作式取消。

**修复方案**：

```go
select {
case <-time.After(5 * time.Second):
	fmt.Println("done")
case <-ctx.Done():
	fmt.Println("cancelled:", ctx.Err())
}
```

**带走原则**：context 不是线程终止器；它是取消信号，任务必须主动监听。

**时间控制**：12 分钟。

### Case D2-16：测试里用 `reflect.DeepEqual` 忽略语义

**插入位置**：testing、表格驱动测试之后。

**前置知识**：测试函数、slice/map。

**任务描述**：让学员比较两个结果是否相等。

```go
var a []int = nil
b := []int{}
fmt.Println(reflect.DeepEqual(a, b))
```

**预期失败**：输出 `false`，而很多业务语义里它们都代表“空列表”。

**讲师追问**：

- 测试应该比较结构完全相同，还是业务语义相同？
- 空 slice 和 nil slice 在 JSON、API 返回里可能有什么区别？

**引入概念**：测试断言的语义设计。

**修复方案**：业务上认为都为空时，用 `len(x) == 0`；需要严格区分时写清楚断言。

**带走原则**：测试不是机械比较值，而是表达业务语义。

**时间控制**：8-10 分钟。

## 6. 推荐穿插顺序

### Day 1

| 时间段 | 主知识点 | 推荐 case |
|---|---|---|
| 09:30-10:30 | 环境、包、go run | D1-01 未使用变量与导入 |
| 10:30-11:30 | 变量、作用域、函数 | D1-02、D1-03 |
| 11:30-12:30 | slice、map、string | D1-04、D1-05、D1-06、D1-07 |
| 13:30-14:30 | 指针、nil、值传递 | D1-09 可提前铺垫 |
| 14:30-15:30 | struct、method | D1-08、D1-09 |
| 15:30-16:00 | CLI Task Manager | 复用 D1-05、D1-08、D1-09 作为项目调试点 |

### Day 2

| 时间段 | 主知识点 | 推荐 case |
|---|---|---|
| 09:30-10:30 | interface | D2-01 |
| 10:30-11:30 | error、defer | D2-02、D2-03 |
| 11:30-12:30 | panic/recover | D2-14 |
| 13:30-14:30 | goroutine、WaitGroup、channel | D2-04、D2-05、D2-06、D2-07、D2-08、D2-09、D2-10 |
| 14:30-15:30 | context、锁、race | D2-11、D2-12、D2-13、D2-15 |
| 15:30-16:00 | testing、并发日志分析器 | D2-16，并复盘 D2-07/D2-15 |

实际授课时不建议全部使用。建议每半天选择 2-4 个强 case，其余作为备选或课后材料。

## 7. 强推荐核心 Case

如果课堂时间紧，只保留下面 8 个：

| 优先级 | Case | 原因 |
|---|---|---|
| P0 | D1-04 slice append 共享底层数组 | 能建立 Go 集合语义核心直觉 |
| P0 | D1-05 nil map 可读不可写 | 高频、稳定、直接 |
| P0 | D1-08 map 中 struct 值不可直接改字段 | 任务管理器项目里很容易遇到 |
| P0 | D2-01 typed nil interface | Go interface 最经典语义坑 |
| P0 | D2-07 锁和 slice 手写生产-消费者 | 最适合引出 channel |
| P0 | D2-11 并发计数器 race | 最适合引出 race detector 和锁 |
| P0 | D2-12 并发读写 map fatal | 最适合引出 RWMutex/sync.Map |
| P0 | D2-15 context 创建了但 worker 不听 | 最适合建立协作式取消模型 |

## 8. 讲师执行建议

### 8.1 每个坑都要先让学生预测

不要直接运行。先问：

> 你觉得这段代码会输出什么？会编译失败、panic、deadlock，还是正常运行？

预测之后再运行，记忆会更深。

### 8.2 复盘时先描述现象，不急着给答案

推荐三问：

1. 我们看到了什么现象？
2. 这个现象说明我们之前哪个假设错了？
3. Go 提供了什么机制让我们表达正确意图？

### 8.3 把“坑”翻译成工程原则

不要让学生只背诵陷阱。每个 case 结束时都落成一句原则，例如：

- slice 是视图，不是独立数组。
- map 零值可读不可写。
- channel 的 close 是发送方的所有权动作。
- context 是取消信号，不是强制终止。
- 测试断言要表达业务语义。

### 8.4 控制失败的情绪成本

如果学生连续失败三次，体验会变差。建议每个教学块最多安排一个强失败 case，再配一个轻量预测题。

### 8.5 不要在 D1/D2 过度模拟生产事故

D1/D2 的目标是建立语言和并发心智模型，不是让学生提前背诵复杂工程规则。

例如：

- D1/D2 适合讲 nil map、slice、race、deadlock。
- D3 之后再讲 middleware 顺序、N+1 查询、事务边界、gRPC status code、缓存一致性。

### 8.6 课堂材料组织方式

建议把每个 case 分成两份：

- 学员版：只有任务和初始代码，不给修复方案。
- 讲师版：包含预期失败、追问、修复方案和带走原则。

当前文档是讲师版。后续如果要分发给学员，可以从这里拆出简化版练习单。

## 9. 与现有项目资料的关系

本文件不替代已有文档：

- `docs/module01_basics_lesson_plan.md`：D1 正式授课主线。
- `module02_advanced/instructor/RUNBOOK.md`：D2 正式授课主线。
- `docs/go_darker_corners_supplements.md`：Go 暗角知识点和示例补充。

本文件的作用是把“暗角知识点”组织成课堂可执行的教学剧本，帮助讲师决定什么时候让学生先撞到问题，以及如何用这个问题自然引出下一段知识。
