# Module 01: Go 语言基础 - 教师备课教案

**适用对象**: Go 语言零基础或转语言学员  
**总课时**: 预计 6-8 小时 (含练习)  
**教学目标**: 掌握 Go 语言核心语法，理解其独特的类型系统和内存模型，能够编写简单的命令行工具。

---

## 目录
1. [第 1 课: Hello World与环境初探 (01_hello)](#第-1-课-hello-world与环境初探-01_hello)
2. [第 2 课: 变量与基本类型 (02_vars_types)](#第-2-课-变量与基本类型-02_vars_types)
3. [第 3 课: 流程控制与基础函数 (03_control_funcs)](#第-3-课-流程控制与基础函数-03_control_funcs)
4. [第 4 课: 数组与切片详解 (04_arrays_slices)](#第-4-课-数组与切片详解-04_arrays_slices)
5. [第 5 课: Map与字符串处理 (05_maps_strings)](#第-5-课-map与字符串处理-05_maps_strings)
6. [第 6 课: 指针的本质 (06_pointers)](#第-6-课-指针的本质-06_pointers)
7. [第 7 课: 结构体与方法 (07_structs_methods)](#第-7-课-结构体与方法-07_structs_methods)
8. [第 8 课: 数据结构综合应用 (08_data_structures)](#第-8-课-数据结构综合应用-08_data_structures)
9. [第 9 课: 高级函数特性 (09_advanced_functions)](#第-9-课-高级函数特性-09_advanced_functions)
10. [项目实战: 任务管理器 (project_task_manager)](#项目实战-任务管理器-project_task_manager)

---

## 第 1 课: Hello World与环境初探 (01_hello)

**源码路径**: `module01_basics/01_hello/main.go`

### 1. 教学目标
- 理解 Go 程序的最小结构 (`package main`, `func main`).
- 掌握 `go run` 和 `go build` 的区别.
- 了解 Go 语言的简洁性哲学 (无分号, 强制格式化).

### 2. 理论讲解 (Lecture Notes)
- **Package 概念**: Go 代码组织在包中. `package main` 告诉编译器这是一个可执行程序, 而不是库.
- **Import**: 导入标准库 `fmt` (Format). 对比 Python `import` 或 Java `import`.
- **Main 函数**: 程序的入口. 注意: 无参数, 无返回值 (不同于 C/Java 的 args).
- **格式化**: 介绍 `gofmt`. Go 只有一种标准代码风格.

### 3. 代码演示脚本 (Teacher's Script)
- **展示代码**:
  ```go
  package main
  import "fmt"
  func main() {
      fmt.Println("Hello, World!")
  }
  ```
- **提问**: "如果把 `package main` 改成 `package hello`, 还能运行吗?"
  - *答案*: 不能. `go run` 需要 main 包.
- **提问**: "为什么行尾没有分号?"
  - *答案*: 编译器自动插入. 实际上是有的, 只是为了简洁省略了.

### 4. 实操指南 (Hands-on)
1. **运行程序**: `go run main.go`
2. **编译程序**: `go build -o hello` 然后 `./hello`
3. **尝试错误**: 
   - 尝试删除 `import "fmt"` 但保留 `fmt.Println` -> 编译错误 (Go 不允许未使用或未导入).
   - 尝试左大括号 `{` 换行 -> 语法错误 (Go 强制大括号不换行).

---

## 第 2 课: 变量与基本类型 (02_vars_types)

**源码路径**: `module01_basics/02_vars_types/main.go`

### 1. 教学目标
- 掌握三种变量声明方式 (`var`, `:=`, 块声明).
- 理解零值 (Zero Value) 概念.
- 熟悉基本类型 (`int`, `string`, `bool`, `float`).

### 2. 理论讲解
- **静态类型 vs 动态类型**: Go 是静态强类型, 但有类型推断 (`:=`).
- **零值机制**: 强调 Go 变量声明即初始化 (0, "", false, nil), 不会出现 "未初始化内存" 错误.
- **常量**: `const`. 介绍 `iota` (如果代码中有涉及).

### 3. 代码演示脚本
- **变量声明**: 
  - `var i int = 10` (标准)
  - `var i = 10` (类型推断)
  - `i := 10` (短变量声明, 仅限函数内)
- **对比 Java/C++**: 强调 Go 的类型在变量名之后.
- **类型转换**: Go 不支持隐式转换 (如 `int` 转 `float`), 必须显式强转 `T(v)`.

### 4. 实操指南
- **练习**: 声明一个未初始化的 `string`, 打印它, 观察输出 (空行).
- **陷阱**: 尝试在函数外使用 `:=` -> 报错.

---

## 第 3 课: 流程控制与基础函数 (03_control_funcs)

**源码路径**: `module01_basics/03_control_funcs/main.go`

### 1. 教学目标
- 掌握 `if` (不带括号), `for` (唯一的循环), `switch` (默认 break).
- 理解函数的定义, 多返回值特性.

### 2. 理论讲解
- **If**: 条件表达式不需要括号, 左大括号必须同行. 支持初始化语句 `if err := foo(); err != nil`.
- **For**: Go 没有 `while`. `for` 既是 `for` 又是 `while`.
- **Switch**: 默认不需要 `break`. `fallthrough` 关键字.
- **函数**: 一等公民. 多返回值是 Go 处理错误的基石.

### 3. 代码演示脚本
- **For 循环变体**:
  ```go
  // 标准
  for i := 0; i < 10; i++ {}
  // While 模式
  for i < 10 {}
  // 死循环
  for {}
  ```
- **多返回值**: 演示 `func div(a, b int) (int, int)` 返回商和余数.

### 4. 实操指南
- **任务**: 编写一个 FizzBuzz 程序 (1-100, 3的倍数打印Fizz, 5的倍数Buzz).
- **调试**: 故意写一个死循环, 然后用 `Ctrl+C` 终止.

---

## 第 4 课: 数组与切片详解 (04_arrays_slices)

**源码路径**: `module01_basics/04_arrays_slices/main.go`

### 1. 教学目标
- **区分数组与切片**: 数组是值类型且定长, 切片是引用类型且动态.
- 掌握 `make`, `append`, `copy`, 切片截取 `[:]`.
- 理解切片的底层原理 (Pointer, Len, Cap).

### 2. 理论讲解
- **数组**: `[3]int` 和 `[4]int` 是不同类型. 数组赋值是拷贝!
- **切片 (Slice)**: 它是数组的"视图" (Window). 
- **Append**: 自动扩容机制. 当 Cap 不足时, 通常翻倍 (或 1.25 倍).

### 3. 代码演示脚本
- **修改切片影响原数组**:
  ```go
  arr := [5]int{1, 2, 3, 4, 5}
  s := arr[1:3] // [2, 3]
  s[0] = 100
  // 问: arr 变成了什么? -> [1, 100, 3, 4, 5]
  ```
- **Append 陷阱**: 如果 `append` 导致扩容, 切片会指向新数组, 修改不再影响原数组. 这是一个常见的面试/实战考点.

### 4. 实操指南
- **练习**: 创建一个 len=3, cap=5 的切片. 连续 append 3 个元素, 观察 len 和 cap 的变化.

---

## 第 5 课: Map与字符串处理 (05_maps_strings)

**源码路径**: `module01_basics/05_maps_strings/main.go`

### 1. 教学目标
- 掌握 Map 的声明 (`make` vs `var`), 读写, 删除 (`delete`).
- 理解 Map 的无序性.
- 字符串: 不可变性, `rune` 处理中文.

### 2. 理论讲解
- **Map**: 哈希表实现. 必须用 `make` 初始化才能写入, 否则 panic (`nil map`).
- **Comma-ok 模式**: `val, ok := m["key"]` 判断键是否存在.
- **String**: 底层是 `[]byte`. 遍历字符串时, `range` 会自动按 rune (Unicode 字符) 迭代.

### 3. 代码演示脚本
- **Map 查找**:
  ```go
  m := map[string]int{"a": 1}
  if v, ok := m["b"]; ok { ... } else { fmt.Println("Key not found") }
  ```
- **字符串长度**: `len("你好")` 是 6 (UTF-8 字节数), 不是 2. 要获取字符数需转为 `[]rune`.

### 4. 实操指南
- **任务**: 统计一段英文文本中每个单词出现的频率 (Word Count).
- **调试**: 尝试向一个 nil map 写入数据, 触发 panic.

---

## 第 6 课: 指针的本质 (06_pointers)

**源码路径**: `module01_basics/06_pointers/main.go`

### 1. 教学目标
- 理解什么是指针 (内存地址).
- 掌握 `&` (取地址) 和 `*` (解引用).
- 理解 Go 只有"值传递" (Pass by Value). 指针传递本质上是拷贝了地址值.

### 2. 理论讲解
- **Go vs C**: Go 指针不能运算 (no pointer arithmetic), 安全但灵活.
- **应用场景**: 修改函数参数的值; 避免大结构体拷贝.

### 3. 代码演示脚本
- **值传递 vs 指针传递**:
  ```go
  func modifyVal(n int) { n = 100 }
  func modifyPtr(n *int) { *n = 100 }
  ```
- **Nil 指针**: 解引用 nil 指针会 panic. 强调判空的重要性.

### 4. 实操指南
- **练习**: 编写一个 swap 函数, 交换两个 int 变量的值.

---

## 第 7 课: 结构体与方法 (07_structs_methods)

**源码路径**: `module01_basics/07_structs_methods/main.go`

### 1. 教学目标
- 定义 `struct`.
- 方法 (Method) 与函数 (Function) 的区别: 接收器 (Receiver).
- 值接收器 vs 指针接收器.

### 2. 理论讲解
- **类比**: Struct 类似 Java 的 Class (但没有继承).
- **接收器选择**: 
  - 如果要修改状态 -> 必须用指针接收器.
  - 如果结构体很大 -> 建议用指针接收器 (避免拷贝).
  - 否则 -> 值接收器 (并发安全, 语义清晰).

### 3. 代码演示脚本
- **定义**:
  ```go
  type User struct { Name string; Age int }
  func (u *User) Grow() { u.Age++ }
  ```
- **嵌入 (Embedding)**: 简单的组合复用 (Go 不提倡深层继承).

### 4. 实操指南
- **任务**: 定义一个 `Rectangle` 结构体, 实现 `Area()` 和 `Perimeter()` 方法.

---

## 第 8 课: 数据结构综合应用 (08_data_structures)

**源码路径**: `module01_basics/08_data_structures/main.go`

### 1. 教学目标
- 综合运用 Struct, Slice, Map 建模复杂数据.
- JSON 序列化与 Tag (如果代码中有涉及).

### 2. 理论讲解
- **复杂类型**: Slice of Structs, Map of Structs.
- **构造函数模式**: Go 没有构造函数, 通常用 `NewUser()` 这样的工厂函数.

### 3. 代码演示脚本
- 分析 `main.go` 中的自定义数据结构, 讲解如何通过组合基础类型构建业务模型.

---

## 第 9 课: 高级函数特性 (09_advanced_functions)

**源码路径**: `module01_basics/09_advanced_functions/main.go`

### 1. 教学目标
- 理解函数作为一等公民 (变量, 参数, 返回值).
- 掌握匿名函数与闭包 (Closure).
- 了解高阶函数 (Map/Filter/Reduce) 与柯里化.

### 2. 理论讲解
- **闭包**: 函数 + 其引用的外部环境. 即使外部函数返回了, 闭包依然持有变量引用.
- **Defer**: 延迟执行. 即使 panic 也会执行. 常用于 `f.Close()`, `mu.Unlock()`.
- **函数式编程**: Go 不是函数式语言, 但支持部分特性.

### 3. 代码演示脚本 (基于源码)
- **闭包计数器**:
  ```go
  func createCounter() func() int {
      count := 0
      return func() int { count++; return count }
  }
  ```
  - *提问*: "如果调用两次 `createCounter()`, 它们共享 count 吗?" -> *答案*: 不, 每次调用创建新的闭包环境.
- **Defer 执行顺序**: 栈 (LIFO). 演示 `defer fmt.Println` 在 `return` 之后执行 (但在返回值赋值之后).

### 4. 实操指南
- **任务**: 编写一个累加生成器.
- **高级任务**: 使用 defer 实现一个简单的执行时间统计 (`defer timeTrack(time.Now())`).

---

## 项目实战: 任务管理器 (project_task_manager)

**源码路径**: `module01_basics/project_task_manager`

### 1. 项目概述
- 一个基于命令行的任务管理工具 (CLI Task Manager).
- 功能: 添加任务, 列出任务, 完成任务, 删除任务.

### 2. 代码架构解析
- **Model**: `Task` 结构体 (ID, Title, Done).
- **Store**: 使用 `[]Task` 或 `map[int]Task` 存储内存数据.
- **Manager**: 封装增删改查逻辑.
- **Main**: 处理命令行参数 (`os.Args` 或 `flag` 包) 和用户交互.

### 3. 教学步骤
1. **需求分析**: 定义 Task 结构体.
2. **实现核心逻辑**: 编写 TaskManager 的 Add/List/Complete 方法.
3. **编写测试**: 运行 `go test` 验证逻辑.
4. **CLI 交互**: 在 main 中解析命令.

### 4. 扩展挑战 (Homework)
- **持久化**: 目前数据在内存中, 重启即失. 请尝试将数据保存到 `tasks.json` 文件中 (使用 `encoding/json`).
- **过滤**: 增加 `list --done` 只显示已完成的任务.

