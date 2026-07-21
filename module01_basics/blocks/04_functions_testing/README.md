# Block 4：Functions and Testing

## 学习结果

完成本区块后，你能够：

- 区分首字母大写的公有函数和首字母小写的私有函数，并阅读完整的函数声明结构。
- 使用普通多值返回、命名返回值、裸 `return`、`error` 返回和 `_` 丢弃不需要的返回值。
- 编写变参函数，并用 `values...` 把 Slice 展开为变参参数。
- 把函数赋给变量、作为参数传入另一个函数，并解释函数值如何替代单一方法策略。
- 声明命名函数类型，理解它与未命名 `func(...) ...` 类型的区别；认识函数值取地址和 `new` 得到函数类型指针的合法但少见写法。
- 编写返回函数的高阶函数，并使用闭包捕获创建时的配置。
- 使用 `defer` 安排函数返回前必须执行的收尾动作，并判断可观察事件的顺序。
- 使用 Go 的 `testing` 包阅读失败信息，通过 `-run` 一次聚焦一个具名测试。

## 时间盒：45 分钟

- 讲师 Demo：15 分钟
- 学员结对实现：20 分钟
- 测试与复盘：10 分钟

## 前置知识

完成 Block 1–3，能够阅读函数、Slice、Struct 和基础 Go 测试。你应当理解 Java 的方法、Lambda、`Predicate<Integer>` 与 `try/finally`，但不要求熟悉 Go 闭包或 `defer`。

## Java 对比

| Java | Go |
| --- | --- |
| Lambda 常匹配 `Predicate<T>` 等函数式接口 | 函数值直接拥有 `func(...) ...` 类型，无需额外接口 |
| 方法通常依附于 Class | 普通函数可以独立存在，也能赋值和传递 |
| Lambda 可以捕获外层变量 | 闭包可以捕获创建它的词法作用域中的变量 |
| `try/finally` 显式包围需要清理的代码 | `defer` 在当前函数返回前执行已注册的调用 |
| JUnit 用注解标记测试 | Go 测试由 `TestXxx(t *testing.T)` 的命名和签名识别 |
| 常按 Class 或方法筛选测试 | `go test` 测试包，`-run` 用正则筛选测试名 |

函数值使行为可以像数据一样被组合。闭包保留的是外层变量的访问能力，而不只是“带一个隐藏字段的方法对象”。`defer` 在声明处表达收尾意图，但调用会延迟到外围函数返回前执行；多个 `defer` 按后进先出的顺序运行。

## 讲师 Demo

先运行核心的函数值、闭包与 `defer` 演示；完整函数模式已移到 Bonus，可在核心时间盒之外运行和测试：

```bash
go run ./module01_basics/blocks/04_functions_testing/demo/09_advanced_functions
go run ./module01_basics/blocks/04_functions_testing/demo/10_function_forms
go run ./module01_basics/blocks/04_functions_testing/demo/11_defer_edges
go run ./module01_basics/bonus/function_patterns
go test ./module01_basics/bonus/function_patterns
```

Demo 现在完整覆盖函数声明、返回值、变参、函数类型、高阶函数、闭包状态和 `defer` 的 LIFO/参数求值语义；柯里化、函数组合和并发继续移入可裁 Bonus。Bonus 测试用于识别 Arrange、Act、Assert；可用 `go test ./module01_basics/bonus/function_patterns -run TestURLBuilder` 单独执行一个测试，说明测试名为何也是开发反馈接口的一部分。

## 函数与测试基础

- 函数名首字母大写表示可被其他包访问，首字母小写表示包内私有；这不是 Java 的 `public/private` 关键字，而是 Go 的导出规则。
- 函数签名由名称、参数列表和返回列表组成；`func f(a, b int) (int, error)` 合并了参数类型，返回值可以是多个不同类型。
- 命名返回值会先得到零值，裸 `return` 返回当前结果；短函数可读，复杂函数中应优先使用显式返回值避免隐藏数据流。
- 多值返回常与 `error` 一起使用；不需要某个结果时用 `_`，不能声明后完全不使用局部变量。
- 变参参数在函数体内表现为 Slice；调用已有 Slice 时使用 `values...`，不能把 Slice 直接当成一个普通单值参数。
- 函数可以作为变量、参数和返回值；可以定义命名函数类型。
- 命名函数类型是新类型，虽然底层签名相同，也需要显式使用该类型；函数值可以取地址，但通常直接传函数值更清楚。
- 变参函数接收一个 Slice，调用方可以用 `values...` 展开已有 Slice。
- 闭包可以捕获外层变量；每次调用外层函数都会创建独立的捕获状态。
- 多个 `defer` 按后进先出执行；普通参数在注册时求值，闭包体在执行时读取变量。
- 测试文件使用 `_test.go`，测试函数使用 `TestXxx(t *testing.T)`；`t.Run` 适合组织子场景，表格驱动测试适合覆盖边界。

课堂上至少补写一个测试。`lab/starter/scores_additional_exercise_test.go` 提供了一个额外的输入顺序场景；Starter 保持预期 RED，学生应先读失败，再完成 `Filter`。

## 学员任务

进入 `lab/starter`，实现 `Filter`、`AtLeast` 和 `WithAudit`。`Filter` 接收函数值决定保留哪些成绩；`AtLeast` 返回捕获最低分数的闭包；`WithAudit` 使用 `defer` 确保结束事件发生在操作之后。完整契约和逐测试命令见 [lab/README.md](lab/README.md)。

先用 `-run` 一次运行一个具名测试：通过闭包筛选测试后，再转向审计顺序测试。最后运行整个练习包检查回归。

## 验收命令

```bash
go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter
```

演示自身也有逐知识点测试：

```bash
go test ./module01_basics/blocks/04_functions_testing/demo/10_function_forms
go test ./module01_basics/blocks/04_functions_testing/demo/11_defer_edges
```

所有测试通过，且能解释 `AtLeast` 捕获了什么、`WithAudit` 的三个事件为何按约定排列，即完成基础任务。

## 常见错误

- 在 `Filter` 中硬编码最低分，而没有调用传入的 `keep` 函数。
- 预先分配长度为 `len(values)` 的结果 Slice 后继续 `append`，留下多余零值。
- 在 `AtLeast` 返回的闭包中使用 `>`，错误排除恰好等于 `min` 的成绩。
- 让 `AtLeast` 总返回相同结果，没有使用闭包捕获的 `min`。
- 在 `WithAudit` 中先执行 `operation`，导致开始事件出现得太晚。
- 直接记录结束事件或过早调用它，没有使用 `defer` 表达返回前收尾。
- 每次都运行全部测试，导致反馈中混入尚未开始实现的另一个行为。
- 把首字母大小写只当作命名风格，忽略它决定跨包可见性。
- 把命名返回值、裸 `return`、多值返回和 `_` 当成同一个概念；它们分别涉及返回变量、返回语句、返回列表和丢弃值。
- 把变参参数误认为数组，或忘记调用已有 Slice 时的 `...`。
- 把命名函数类型误认为普通类型别名，或把函数值的指针误讲成 Go 有独立的函数指针语法。
- 测试只验证“能通过”而不验证顺序、边界或输入不变性；先写一个具名失败，再运行最小测试。

## 三级提示

1. 先写一个空结果 Slice，遍历 `values`，把每个 `value` 交给 `keep` 判断。
2. `AtLeast` 内部直接返回 `func(value int) bool`；这个匿名函数可以读取参数 `min`。
3. 在开始记录之后写 `defer record("end:" + name)`，再调用 `operation()`；退出 `WithAudit` 前会触发结束记录。

## 复盘问题

- Go 的 `func(int) bool` 与 Java 的 `Predicate<Integer>` 在表达和调用上有哪些差异？
- `AtLeast(60)` 返回后，参数 `min` 为什么仍然可用？
- 为什么 `defer record("end:" + name)` 写在 `operation()` 之前，结束事件却发生在操作之后？
- 普通 `defer fmt.Println(value)` 与 `defer func() { fmt.Println(value) }()` 的读取时机有什么区别？
- 为什么给 `Filter` 写测试时要同时检查结果内容和原输入顺序？
- `PublicFunctionDemo` 与 `addText` 的首字母差异会影响什么？
- `addTextNamed` 的裸 `return` 返回了哪些变量？什么时候显式返回更容易读？
- 为什么 `sum(values...)` 合法，而 `sum(values)` 不合法？
- 命名函数类型 `Combiner` 与 `func(int, int) int` 有什么关系和区别？为什么函数值指针是合法但通常不推荐的写法？
- 单独运行具名测试相比每次运行整个包，如何缩短 RED、GREEN、REFACTOR 的反馈循环？

## Bonus

以下内容不计入本 Block 的 45 分钟核心时间盒，仅在课后或进度提前时选做。

先增加一个失败测试，再实现 `Between(min, max int) func(int) bool`。明确边界是否包含在范围内，并将返回的闭包传给 `Filter`；不要修改 `Filter` 的接口。
