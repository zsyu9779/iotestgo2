# 模块01+02 讲课线性“小抄”（教师讲台备忘录）
 
定位：这一份文档从上往下读/扫一遍，就能把整节课串起来。每个点只保留“讲什么、怎么说、怎么演示、学员做什么、容易踩什么坑、怎么过渡到下一节”。  
覆盖：`module01_basics` + `module02_advanced` 全部小节 + 两个项目。
 
---
 
## 讲前 3 分钟准备（你只看这一段就能开讲）
 
- 打开项目根目录：`/Users/zhangshiyu/class/iotestgo2`
- 统一口令（课堂上固定用这一套，减少临场脑耗）
  - 进入某节：`cd module01_basics/01_hello`（或对应目录）
  - 运行示例：`go run .`
  - 跑测试：`go test -v`
- 课堂节奏提示
  - 每节的结构固定：一句话定位 → 1 个 Demo → 1 个练习 → 1 个坑 → 过渡串词
  - 目标不是“讲完所有细节”，而是“让学员形成脑内地图 + 能动手跑通”
 
---
 
## 开场（1 分钟）
 
你说：
- “今天我们走一条完整主线：从 Go 最小程序结构开始，到能写一个小 CLI 项目，再到并发流水线项目。”
- “Go 的核心优势不在语法花哨，而在：工程化标准统一 + 并发模型清晰 + 标准库够用。”
 
你做：
- 快速展示目录树（不用展开太多）：`module01_basics` → `module02_advanced` → 两个 project
 
过渡串词：
- “先把语法和数据结构打稳（模块01），再讲 Go 的抽象、错误和并发（模块02）。”
 
---
 
## Module01 主线：把 Go 基础一次串起来（按目录顺序）
 
### 01_hello（第 1 个 Go 程序 + 约束感）
 
一句话定位：Go 程序的最小骨架 + Go 对“规范”的强约束（格式、未使用等）。
 
你说（关键句）：
- “`package main` + `func main()` 不是约定，是编译器要求：可执行程序入口。”
- “Go 的理念之一：让团队写出同一种风格的代码。”
 
你做（Demo）：
- `cd module01_basics/01_hello && go run .`
- 提问：把 `package main` 改成别的，会发生什么？（不能运行）
 
学员做（练习）：
- 把输出改成 3 行，包含自己的名字/班级/一句话目标。
 
坑点提醒：
- 大括号 `{` 不能换行（Go 的自动分号插入规则）。
 
过渡串词：
- “能跑起来只是第一步，下一步我们要能正确地表达数据：变量、常量、类型。”
 
---
 
### 02_vars_types（变量/常量/类型推断：Go 的‘稳’）
 
一句话定位：静态强类型 + 语法简洁（类型推断）+ 常量枚举（iota）+ 显式类型转换。
 
你说（关键句）：
- “Go 是静态强类型，但写起来不啰嗦：`:=` 让你像动态语言一样快。”
- “Go 不做隐式类型转换：你写出来，意图就清楚。”
 
你做（Demo）：
- `cd module01_basics/02_vars_types && go run .`
- 指着代码讲 4 个点：
  - `var age int = 30`
  - `name := "Gopher"`（只在函数内）
  - `iota` 三个状态值
  - `float64(i)` 显式转换
 
学员做（练习）：
- 用 `const` + `iota` 定义一组星期枚举（Mon/Tue/…），打印出来。
 
坑点提醒：
- `int32` 溢出会回绕（示例里 `maxInt32 + 1`），强调“别拿溢出做业务逻辑”。
 
过渡串词：
- “有了数据，下一步是让程序做决策和复用：if/for/switch + 函数。”
 
---
 
### 03_control_funcs（流程控制 + 多返回值函数）
 
一句话定位：Go 的控制流非常克制；`for` 是唯一循环；函数多返回值是 Go 错误处理的基础。
 
你说（关键句）：
- “Go 的 `if`/`for` 没有括号，强调可读性。”
- “Go 的 `switch` 默认不贯穿，不用写一堆 break。”
- “多返回值是 Go 的风格：结果 + 状态（以后就是 result + err）。”
 
你做（Demo）：
- `cd module01_basics/03_control_funcs && go run .`
- 指着讲：
  - `if score := 85; score >= 60 { ... }`
  - `switch role { ... }`（提 `fallthrough`）
  - `for i := 0; ...` + “while-like for”
  - `calculate` 返回 `(sum, diff)`
 
学员做（练习）：
- 写一个 `max(a,b int) int`，再写一个 `minmax(a,b int) (min,max int)`。
 
坑点提醒：
- 讲“匿名函数/闭包”先不要深入（后面 09_advanced_functions 再系统讲）。
 
过渡串词：
- “控制流会了，接下来进入 Go 最常用的两类容器：数组/切片（Go 的‘动态数组’）。”
 
---
 
### 04_arrays_slices（数组 vs 切片：值语义 vs 引用语义）
 
一句话定位：数组是定长值类型；切片是动态视图（指针+len+cap），`append` 可能触发扩容。
 
你说（关键句）：
- “数组赋值是拷贝；切片更像引用，但本质是一个小结构体指向底层数组。”
- “`len` 是当前长度，`cap` 是到底层数组末尾的容量。”
 
你做（Demo）：
- `cd module01_basics/04_arrays_slices && go run .`
- 讲三个瞬间：
  - `arrCopy := arr` 改了 copy 不影响原数组
  - `subSlice[0] = 999` 会影响原 slice（共享底层数组）
  - `make([]int, 0, 5)` 预分配容量
 
学员做（练习）：
- 创建一个切片，连续 append 10 次，每次打印 `len/cap`，观察 cap 增长规律。
 
坑点提醒（必讲）：
- “`append` 触发扩容后，切片会指向新数组：此时你以为还在改原来的，其实已经分家。”
 
过渡串词：
- “切片解决顺序集合；下一节 Map 解决键值查找，再顺便把字符串的 Unicode 坑讲掉。”
 
---
 
### 05_maps_strings（map + string/rune：两大高频坑）
 
一句话定位：map 是哈希表；读写要用 `make` 初始化；字符串是字节序列，中文要用 rune。
 
你说（关键句）：
- “map 的零值是 nil，nil map 只能读不能写，写会 panic。”
- “`len(str)` 是字节数，不是字符数；`[]rune(str)` 才是按 Unicode 字符看。”
 
你做（Demo）：
- `cd module01_basics/05_maps_strings && go run .`
- 讲：
  - `val, ok := scores["Alice"]` 的 comma-ok
  - `delete(scores, "Bob")`
  - `len("Hello, 世界")` vs `len([]rune(...))`
 
学员做（练习）：
- WordCount：统计一句英文里每个单词出现次数（提示用 `strings.Fields`）。
 
坑点提醒：
- map 遍历无序：输出顺序每次可能不同，别写依赖顺序的逻辑/测试。
 
过渡串词：
- “map 和 slice 都是引用味道很重的类型，下一节我们把‘引用’讲透：指针。”
 
---
 
### 06_pointers（指针：你必须会但不需要恐惧）
 
一句话定位：Go 只有值传递；用指针是为了“修改外部变量/避免大拷贝”；Go 指针不能运算，更安全。
 
你说（关键句）：
- “Go 只有 pass-by-value：传指针也只是把地址这个值拷贝过去。”
- “能用值就用值，需要修改就传指针。”
 
你做（Demo）：
- `cd module01_basics/06_pointers && go run .`
- 讲：
  - `ptr := &x` / `*ptr` 读写
  - nil 指针解引用会 panic（演示只说不执行）
  - `modifyValue` vs `modifyPointer`
 
学员做（练习）：
- 写 `swap(a,b *int)`，交换两个值（提示：用临时变量）。
 
坑点提醒：
- “不要为了显得高级而到处用指针：可读性和并发安全会变差。”
 
过渡串词：
- “有了指针，我们就能解释 Go 方法接收器的选择：值接收器 vs 指针接收器。”
 
---
 
### 07_structs_methods（struct + method：Go 的面向对象‘够用版’）
 
一句话定位：struct 负责数据；method 负责行为；组合（embedding）替代继承；指针接收器是常态。
 
你说（关键句）：
- “Go 没有 class 继承链；更鼓励组合：把一个 struct 嵌入到另一个 struct。”
- “方法接收器决定：你是在改副本，还是在改原对象。”
 
你做（Demo）：
- `cd module01_basics/07_structs_methods && go run .`
- 指着讲：
  - `func (u User) String() string`（值接收器：读）
  - `func (u *User) UpdateName(...)`（指针接收器：改）
  - `type Admin struct { User; Level int }`（embedding）
 
学员做（练习）：
- 写 `type Rectangle struct { W,H int }` + `Area()` + `Scale(k int)`（Scale 要用指针接收器）。
 
坑点提醒：
- “embedding 不是继承：它是字段提升（field promotion）。概念要说清楚。”
 
过渡串词：
- “现在我们用这些能力写一个‘真的数据结构’：链表 + 栈。”
 
---
 
### 08_data_structures（链表 + 栈：把指针与 struct 绑在一起）
 
一句话定位：链表是指针驱动的数据结构；通过它强化 nil 判断、指针移动、方法封装。
 
你说（关键句）：
- “链表题不是为了写链表，是为了训练：指针的移动和边界处理。”
- “Go 写数据结构靠组合：Node + Next 指针 + List 维护 Head/Size。”
 
你做（Demo）：
- `cd module01_basics/08_data_structures && go run .`
- 讲 Add/Remove 的关键分支：
  - 空链表
  - 删除头结点
  - 删除中间节点
 
学员做（练习）：
- 在 LinkedList 上加一个 `Contains(val int) bool` 或 `Reverse()`（可选）。
 
坑点提醒：
- “处理链表：先画图再写代码，别靠脑补。”
 
过渡串词：
- “Go 的表达能力不仅是数据结构，还有函数：下一节把函数当变量/当返回值来用。”
 
---
 
### 09_advanced_functions（函数一等公民：闭包/高阶函数/defer/柯里化）
 
一句话定位：函数可以存、可以传、可以返回；闭包保留环境；defer 管资源；函数式技巧在 Go 中主要用于‘配置/中间件/组合’。
 
你说（关键句）：
- “闭包 = 函数 + 它捕获的外部变量环境。”
- “defer 的价值在于：把‘收尾动作’放在‘拿到资源’的旁边，减少遗漏。”
- “Go 不是函数式语言，但函数式模式在工程里非常常见：中间件、选项模式。”
 
你做（Demo）：
- `cd module01_basics/09_advanced_functions && go run .`
- 课堂抓 5 个镜头讲：
  1) 函数变量：`type MathFunc func(int,int) int`（函数类型）
  2) 闭包计数器：`createCounter()`（每个 counter 有自己的 count）
  3) map/filter/reduce：泛型版高阶函数（强调“模式”，不纠结泛型细节）
  4) defer：LIFO + 捕获参数值（defer 参数求值时机）
  5) 柯里化/组合：`curryAdd`、`compose`、`pipeline`
 
加餐（强烈推荐，用测试文件讲工程模式）：
- 中间件/装饰器（柯里化的真实用途）：
  - 路径：`module01_basics/09_advanced_functions/curry_best_practice_test.go`
  - 课堂演示：`go test -v -run TestMiddleware`
- Builder vs Functional Options（Go 社区常用配置模式）：
  - 路径：`module01_basics/09_advanced_functions/patterns_comparison_test.go`
  - 课堂演示：`go test -v -run TestFunctionalOptionsPattern`
 
学员做（练习）：
- 写一个 `WithPrefix(prefix string) func(string) string`，把输入字符串加前缀；再组合两个装饰器。
 
坑点提醒：
- “闭包+循环变量：如果你在 goroutine 里用循环变量，一定要把变量当参数传进去（后面并发课再强调）。”
 
过渡串词：
- “基础内容到这里够写一个小项目了。现在我们把所有东西压缩成一个 CLI 任务管理器。”
 
---
 
### project_task_manager（小项目：把基础能力落成‘可跑’的软件）
 
一句话定位：用 struct+slice+方法+输入输出，做一个可交互的 CLI；顺带引出测试和工程意识。
 
你说（关键句）：
- “项目不是写 UI，是训练：拆结构、封装方法、处理输入、写测试。”
- “先让它能用，再让它好用（可测试、可扩展、可维护）。"
 
你做（Demo）：
- `cd module01_basics/project_task_manager && go run .`
- 按顺序操作一遍（你做，学员看）：
  - Add Task → List → Complete → List → Delete → Exit
 
你说（边做边讲结构）：
- `Task`：ID/Title/Completed
- `TaskManager`：`tasks []*Task` + `nextID`
- 核心方法：Add/List/Complete/Delete
- `bufio.Scanner`：循环读取用户输入，switch 分发命令
 
学员做（练习）：
- 必做：新增一个菜单项 `6. List Completed`（只列已完成任务）。
- 可选：把 tasks 持久化到本地文件（JSON）。（提前埋伏笔：后面会讲 file I/O）
 
坑点提醒：
- `Scanner` 默认 token 长度有限，真实项目需注意（这里只提醒即可）。
- 删除 slice 元素用 `append(slice[:i], slice[i+1:]...)`，别写错边界。
 
过渡串词：
- “模块01你们已经能写出完整的小程序。模块02开始讲：抽象（接口）+ 边界（错误）+ 并发（goroutine/channel/context）+ 工程化（测试/IO/运行时）。”
 
---
 
## Module02 主线：抽象 + 错误 + 并发 + 工程化（按目录顺序）
 
### 01_interfaces（隐式实现：Go 抽象能力的核心）
 
一句话定位：接口不是“继承体系”，而是“行为契约”；Go 通过隐式实现实现低耦合。
 
你说（关键句）：
- “在 Go 里，接口的意义是：把‘用什么’和‘怎么做’分开。”
- “实现接口不需要声明：只要方法集匹配，就自然实现。”
 
你做（Demo）：
- `cd module02_advanced/01_interfaces && go run .`
- 讲三段：
  1) `Animal` + `Dog/Cat` + `MakeSound`（多态）
  2) `interface{}`（any）是‘未知类型容器’
  3) type assertion + type switch
 
学员做（练习）：
- 定义 `type Shape interface { Area() float64 }`，实现 `Rect` 与 `Circle`，写一个 `TotalArea([]Shape)`。
 
坑点提醒（经典必讲）：
- 接口的 nil 陷阱：`var p *Dog=nil; var a Animal=p; a==nil?`（提示：接口内部含 type 信息）
 
过渡串词：
- “接口让我们抽象得更优雅，但工程里更要命的是：失败怎么表达？下一节讲 Go 的错误处理。”
 
---
 
### 02_errors_defer（Errors are values + defer/recover 边界）
 
一句话定位：Go 不鼓励用异常控制流程；错误是返回值；panic 只用于真正不可恢复；recover 只在边界兜底。
 
你说（关键句）：
- “错误不是例外，是一种数据：你要显式地处理它。”
- “panic 是程序‘不该发生’的事；recover 只在最外层做兜底，别到处 catch。”
 
你做（Demo）：
- `cd module02_advanced/02_errors_defer && go run .`
- 讲：
  - defer LIFO（两条 defer 的顺序）
  - 自定义错误 `MyError` 实现 `Error()`，`errors.As` 抽出具体类型
  - `safeFunction` 里 panic，被 defer recover 捕获后 main 继续执行
 
学员做（练习）：
- 写 `func ParsePort(s string) (int, error)`：非法就返回 error，不要 panic。
 
坑点提醒：
- “在库代码里不要随便 panic；在 main/handler 边界可以 recover 兜底并记录日志。”
 
过渡串词：
- “错误处理讲清楚之后，我们开始讲 Go 的王牌：并发。先从 goroutine 的生命周期开始。”
 
---
 
### 03_goroutines（并发起点：go 关键字 + WaitGroup）
 
一句话定位：goroutine 是轻量协程；main 退出会杀死所有 goroutine；用 WaitGroup 管生命周期。
 
你说（关键句）：
- “并发是结构，是否并行要看 CPU 和调度（GOMAXPROCS）。"
- “生命周期管理是第一要务：不要让 goroutine 变成幽灵。”
 
你做（Demo）：
- `cd module02_advanced/03_goroutines && go run .`
- 讲：
  - `runtime.NumCPU()` 输出
  - `go func(){...}()` 启动 goroutine
  - `WaitGroup`：Add/Done/Wait（类比 Java CountDownLatch）
 
学员做（练习）：
- 启动 5 个 worker，每个睡眠不同时间后打印“done”，保证 main 等到全部结束。
 
坑点提醒（口头预告，下一节更系统）：
- goroutine 之间不要靠共享变量乱读写，先学 channel 通信。
 
过渡串词：
- “goroutine 负责并发执行，那它们怎么协作？Go 的答案是 channel。”
 
---
 
### 04_channels（通道：通信 + 协调 + 关闭协议）
 
一句话定位：channel 是并发通信原语；无缓冲同步、有缓冲解耦；close 是广播式‘不再发送’信号；select 处理多路输入。
 
你说（关键句）：
- “不要通过共享内存通信；通过通信来共享内存。”
- “close 的语义是：发送方宣布‘我不会再发了’，不是接收方的权力。”
 
你做（Demo）：
- `cd module02_advanced/04_channels && go run .`
- 讲：
  - `make(chan int, 2)` 缓冲通道
  - producer 发送 + close
  - receiver 用 `for val := range ch` 消费直到关闭
  - `select` + `time.After` 超时兜底
 
学员做（练习）：
- 写一个 producer 发 10 个数字，2 个 consumer 并发处理（打印平方），最后 main 等待完成。
 
坑点提醒（必讲）：
- 向已关闭的 channel 写会 panic；从关闭 channel 读会返回零值 + ok=false（本例用 range 更直观）。
 
过渡串词：
- “有了 channel，系统能跑起来；但真实系统还需要‘控制它什么时候停’：下一节 context。”
 
---
 
### 05_context（取消/超时：把‘停下来’也工程化）
 
一句话定位：context 用于控制请求生命周期（取消、超时、截止时间）和传递请求级数据；它是并发系统的刹车系统。
 
你说（关键句）：
- “context 像一棵树：父节点取消，子节点都得停。”
- “context 作为参数从上往下传，不要塞进 struct 当全局变量用。”
 
你做（Demo）：
- `cd module02_advanced/05_context && go run .`
- 讲两个镜头：
  - WithCancel：goroutine 监听 Done 退出
  - WithTimeout：超时后 ctx.Err() 返回 deadline exceeded
 
学员做（练习）：
- 模拟一个任务循环打印“Working…”，2 秒后 cancel，保证 goroutine 能退出并打印“canceled”。
 
坑点提醒：
- `WithTimeout` 一定要 `defer cancel()`（释放 timer 资源）。
 
过渡串词：
- “通信（channel）+ 刹车（context）够了，但还有一类场景必须共享状态：下一节讲并发安全与锁。”
 
---
 
### 06_concurrency_safety（锁/原子：共享状态的正确姿势）
 
一句话定位：并发的敌人是 data race；Mutex 保证临界区；Atomic 适合简单计数；`-race` 是第一诊断工具。
 
你说（关键句）：
- “没有锁的并发读写不是‘偶尔错’，是‘必然错’，只是你没撞到而已。”
- “先写正确，再谈性能；性能优化的前提是可测量。”
 
你做（Demo）：
- `cd module02_advanced/06_concurrency_safety && go run .`
- 讲：
  - SafeCounter 的 `Inc/Value` 用 mutex 包起来（配合 defer Unlock）
  - `atomic.AddUint64` 的使用场景
- 口头提示：可以用 `go run -race .` 看竞态检测（不强求课堂一定跑）
 
学员做（练习）：
- 把一个“并发加 1000 次”的错误计数器修好：用 Mutex 或 Atomic。
 
坑点提醒：
- 锁的粒度：锁太大影响吞吐，锁太小容易死锁/复杂；教学里强调“先正确”。
 
过渡串词：
- “并发代码最怕‘以为对了’，所以我们要学会用测试证明它对：下一节 testing。”
 
---
 
### 07_testing（测试：把正确性写进代码里）
 
一句话定位：Go 的测试内建且简单；表格驱动是 Go 社区最常见写法；测试让重构不怕。
 
你说（关键句）：
- “测试不是最后补的，是开发的一部分；尤其在并发/边界条件场景。”
 
你做（Demo）：
- `cd module02_advanced/07_testing && go test -v`
- 指着讲：
  - `TestAdd` 断言
  - `TestAddTable` 表格驱动测试
 
学员做（练习）：
- 给 task_manager 里的某个行为补 1 个测试用例（例如删除不存在的 ID）。
 
坑点提醒：
- 测试输出不稳定通常是并发/随机数/时间相关；需要可控依赖（后面项目会提到）。
 
过渡串词：
- “测试能保证逻辑正确，但工程还要和系统打交道：命令、信号、文件。先从 OS 交互开始。”
 
---
 
### 08_os_interaction（执行命令 + 监听信号：优雅退出的基本功）
 
一句话定位：用 `os/exec` 运行外部命令；用 `os/signal` 捕获 Ctrl+C/SIGTERM；结合 context 做优雅退出。
 
你说（关键句）：
- “服务进程必须能优雅退出：收到信号 → 通知取消 → 清理资源 → 退出。”
 
你做（Demo）：
- `cd module02_advanced/08_os_interaction && go run .`
- 讲：
  - exec.Command + Output
  - 管道：给 grep 写 stdin、读 stdout
  - signal.Notify + ctx timeout（10 秒自动退出，或 Ctrl+C）
 
学员做（练习）：
- 把 timeout 从 10 秒改成 3 秒；再试一次 Ctrl+C 退出（观察输出）。
 
坑点提醒：
- 外部命令依赖系统环境（Windows/macOS/Linux 可能不同），课堂强调“理念与 API”，不纠结命令差异。
 
过渡串词：
- “下一节讲文件 I/O：很多小工具/服务离不开读写文件。”
 
---
 
### 09_file_io（文件读写：从低级到缓冲，再到 Seek）
 
一句话定位：掌握 `os.Open/Create`、`io.Reader/Writer`、`bufio`、`io.Copy`；理解 flush/close/seek 的边界。
 
你说（关键句）：
- “所有 IO 都是接口：`Reader/Writer` 让代码可组合、可测试。”
- “写文件最常见的坑不是写不出来，是没 flush / 没 close / 权限不对。”
 
你做（Demo）：
- `cd module02_advanced/09_file_io && go run .`
- 讲流程：
  - 创建临时目录 `test_files`（程序退出会清理）
  - `basicWrite` 写入
  - `basicRead` 分块读取并输出
  - 手写 copy vs `io.Copy` + bufio
  - `Seek(7)` 后读到 `World`
 
学员做（练习）：
- 写一个 `copyfile(src,dst)`：优先用 `io.Copy`，再比较手写 buffer 版本。
 
坑点提醒：
- `defer f.Close()` 的位置：拿到资源立刻 defer，避免遗漏。
 
过渡串词：
- “到这里你们已经能写很多工具了。接下来讲一个‘高级但谨慎用’的能力：反射。”
 
---
 
### 10_reflection（反射：框架为什么能‘自动’）
 
一句话定位：反射能在运行时检查类型/字段/方法并修改值；但成本高、可读性差，除非必要不要用。
 
你说（关键句）：
- “你不用天天写反射，但你必须知道 ORM/JSON/DI 为什么能工作。”
- “反射的关键是：Type/Value；改值必须是可设置的（指针 + CanSet）。"
 
你做（Demo）：
- `cd module02_advanced/10_reflection && go run .`
- 讲三个镜头：
  - `inspectStruct`：TypeOf/ValueOf，遍历字段
  - `modifyValue(&x, 200)`：为什么必须传指针
  - `callMethod(s, "Study")`：按名字调用方法
 
学员做（练习）：
- 给 `Student` 增加一个字段，观察 `inspectStruct` 输出变化。
 
坑点提醒：
- 反射能做的，泛型/接口/代码生成往往更清晰；反射优先用于“通用框架层”。
 
过渡串词：
- “最后补上运行时知识：CPU 使用、调度让出、调用栈信息。你写并发/性能问题时会用到。”
 
---
 
### 11_runtime_control（运行时：你对调度器有最小认知即可）
 
一句话定位：知道 `NumCPU/GOMAXPROCS/Gosched/Caller` 这些工具，就能解释很多并发现象并做最基本的诊断。
 
你说（关键句）：
- “并发性能问题，先看 CPU 利用，再看调度，再看锁竞争，再看 IO。”
 
你做（Demo）：
- `cd module02_advanced/11_runtime_control && go run .`
- 讲：
  - NumCPU 和 GOMAXPROCS
  - Gosched 让出时间片（两个 goroutine 交替输出）
  - Caller 打印调用信息（定位问题思路）
 
学员做（练习）：
- 把两个 goroutine 的循环次数改大一点，观察输出交替情况。
 
过渡串词：
- “标准库里还有很多‘即插即用’的工程工具：正则、JSON、base64、hash、time。下一节快速扫一遍把它们串起来。”
 
---
 
### 12_stdlib_utils（标准库速览：工程实用技能包）
 
一句话定位：用最少的心智负担掌握最常用标准库：regexp/json/base64/crypto/time。
 
你说（关键句）：
- “Go 标准库很强：很多项目只靠标准库就能跑起来。”
- “时间格式化 layout 必背：`2006-01-02 15:04:05`。”
 
你做（Demo）：
- `cd module02_advanced/12_stdlib_utils && go run .`
- 讲：
  - regexp：邮箱提取 + 版本号分组
  - json tag：`omitempty`、`-`（密码字段不输出）
  - base64 encode/decode
  - sha1：只示范 hash API（顺便提醒生产里密码别用 sha1）
  - time：format/parse/add/unix
 
学员做（练习）：
- 写一个结构体 `Config`，加 json tag，序列化后检查字段命名是否正确。
 
过渡串词：
- “现在把并发三件套（goroutine + channel + context）落地到一个流水线项目：日志分析器。”
 
---
 
## 项目实战：project_log_analyzer（并发流水线：从可跑到可测）
 
一句话定位：用 pipeline 模式模拟处理海量日志：生成 → 多 worker 处理 → 汇总；顺带讲关闭通道、WaitGroup、benchmark。
 
你说（关键句）：
- “并发项目最核心的不是开 goroutine，而是：如何收尾（close/Wait/cancel）。”
- “打印日志会毁掉 benchmark：性能测试要尽量纯净。”
 
你做（Demo）：
- `cd module02_advanced/project_log_analyzer && go run .`
- 指着讲结构（按执行顺序）：
  1) `RunPipeline(numProcessors, logCount)` 总入口
  2) `LogGenerator(ctx, logsCh, count)`：生产者，结束时 `close(out)`
  3) `LogProcessor(...)`：worker pool，从 `logsCh` range，筛 ERROR 发到 `errorsCh`
  4) collector：range `errorsCh` 计数
  5) wg 等 processors → `close(errorsCh)` → collectorWg 等收尾
 
学员做（练习）：
- 必做：把 ERROR 的判定改成“Level == WARN 或 ERROR 都统计”，观察结果变化。
- 可选：把 `numProcessors` 改成 1/5/10，感受吞吐差异（先感性，再讲 benchmark）。
 
加餐（benchmark）：
- `go test -bench .`（路径：`module02_advanced/project_log_analyzer/benchmark_test.go`）
- 讲：基准测试里 b.N 循环；对比不同 worker 数量的吞吐。
 
坑点提醒：
- 多 goroutine 下不要在热路径 fmt.Println（会被锁+IO 放大），需要日志采样/异步。
 
---
 
## 收尾（3 分钟：把课“扣住”）
 
你说（收束金句：你只要背这 10 句就能串起全课）：
1. “Go 程序的最小单元：package main + func main。”  
2. “Go 静态强类型，但类型推断让你写得快。”  
3. “Go 不做隐式类型转换：写出来，意图更清楚。”  
4. “数组是值类型，切片是指向底层数组的视图。”  
5. “map 的零值是 nil：能读不能写。”  
6. “Go 只有值传递：传指针只是把地址值拷贝过去。”  
7. “方法接收器决定你改的是副本还是原对象。”  
8. “错误是返回值；panic 只用于真正不可恢复。”  
9. “并发先管生命周期：WaitGroup / channel close / context cancel。”  
10. “先写正确，再谈性能；用测试和 benchmark 说话。”  
 
你布置作业（从易到难）：
- A：给 task_manager 加一个新命令（过滤/搜索/导出）。  
- B：把 task_manager 的数据落盘到 JSON 文件（用 file IO + json tag）。  
- C：给 log_analyzer 加统计维度（按 Level 分桶 / TopN），并写 benchmark 对比不同 worker 数量。  
 
