# Block 1：Go Basics

## 学习结果

完成本区块后，你能够：

- 运行一个最小的 Go 程序，并说明 `package main`、`import` 与 `func main` 的作用。
- 使用变量、常量、基本类型和短变量声明。
- 使用 `if`、`switch`、`for` 与函数组织简单的业务规则。
- 用 `(value, error)` 显式表达失败，并通过测试检查边界条件。

## 时间盒：45 分钟

- 讲师 Demo：15 分钟
- 学员结对实现：20 分钟
- 测试与复盘：10 分钟

## 前置知识

你应当能够阅读基础 Java 代码，理解变量、条件分支、循环、方法和异常。不要求已有 Go 经验。

## Java 对比

| Java | Go |
| --- | --- |
| 类中的 `public static void main` | `package main` 中的 `func main` |
| 变量类型通常写在变量名前 | 类型写在变量名后，也可由 `:=` 推断 |
| `while`、传统 `for`、增强 `for` | 统一使用 `for` |
| 用异常传递失败 | 常用额外的 `error` 返回值显式传递失败 |
| `switch` 默认贯穿，常需 `break` | `switch` 默认不贯穿 |

Go 允许忽略表达式周围的圆括号，但要求分支体使用花括号。未使用的局部变量和导入会导致编译失败。

## 讲师 Demo

按顺序运行并讲解：

```bash
go run ./module01_basics/blocks/01_go_basics/demo/01_hello
go run ./module01_basics/blocks/01_go_basics/demo/02_vars_types
go run ./module01_basics/blocks/01_go_basics/demo/03_control_funcs
go run ./module01_basics/blocks/01_go_basics/demo/05_strings_basics
```

重点观察程序入口、类型推断、变量零值、String 基础操作、`for`、`range`、多值 `case`、`fallthrough`、`switch {}`、函数参数和多返回值。

## 语言语义补充

### 零值

Go 声明变量但没有提供初始值时，会自动使用该类型的零值：整数为 `0`，浮点数为 `0`，布尔值为 `false`，字符串为 `""`，Struct 的每个字段也使用各自的零值。

复合类型还要区分“可直接使用”和“必须初始化”：nil Slice 可以读取长度并 `append`，nil Map 可以读取但不能写入。运行 `02_vars_types` 观察这些结果。

### 变量与常量语义

- `_` 是空白标识符，用于显式丢弃值，它不是可读取变量。
- const 块中省略表达式会复用上一行完整 ConstSpec；`iota` 仍按声明行递增。
- 每个新的 const 块都会让 `iota` 从 0 重新开始。
- `type UserID int` 定义新类型；`type UserID = int` 声明别名。
- 浮点数转整数会截去小数部分。

### 控制流边界

- `for` 是 Go 唯一的循环关键字，可写成三段式、条件式或无限循环。
- `break` 结束当前循环，`continue` 跳过当前迭代。
- 一个 `case` 可以写多个值，例如 `case "admin", "owner"`。
- Go 的 `switch` 默认不会自动贯穿；`fallthrough` 必须显式写出，并且无条件进入紧邻的下一个 `case`。它在业务代码中通常不推荐使用。
- `switch {}` 是按顺序判断条件的简洁写法，适合成绩等级等互斥条件。

### String 基础

String 是不可变的 UTF-8 字节序列。拼接、比较、`strings.Fields`、`strings.TrimSpace` 和大小写转换适合日常文本处理；下标和切片按 byte 工作，按字符遍历应使用 `range`，需要字符 Slice 时再转换为 `[]rune`。运行 `05_strings_basics` 对比 byte index 与 rune index。

### 核心语义深挖

- [Range 变量与闭包陷阱](../../bonus/dark_corners/range/main.go)：放在基本 `for`/`range` 之后，先讲正常遍历，再讲取地址和闭包捕获的版本差异。

## 学员任务

进入 `lab/starter`，实现 `Grade(score int) (string, error)`。有效成绩应返回 A、B、C、D 或 F；无效成绩应返回 `ErrScoreOutOfRange`。完整规则见 [lab/README.md](lab/README.md)。

先运行测试观察失败，再只写足以通过当前失败的代码。每次修改后重新运行测试。

## 验收命令

```bash
go test -tags=exercise ./module01_basics/blocks/01_go_basics/lab/starter
```

所有测试通过即完成基础任务。

## 常见错误

- 将 90、80、70、60 等边界归入了较低等级。
- 忘记先验证 `score` 是否位于 0–100，导致无效成绩也得到字母等级。
- 返回了一个文本错误，而不是约定的 `ErrScoreOutOfRange`。
- 从低分到高分判断时条件过宽，前面的分支提前匹配。
- 把 Java 的异常或 `switch` 贯穿习惯直接套到 Go。

## 三级提示

1. 先写有效区间检查；区间之外返回空字符串和约定错误。
2. 等级边界是 90、80、70、60，按从高到低的顺序思考判断条件。
3. 处理完无效输入后，用一个无表达式的 `switch` 或等价的 `if` 链返回等级。

## 复盘问题

- 为什么 Go 的函数签名把 `error` 与正常结果并列返回？
- `var i int`、`var s string` 和 `var m map[string]int` 的零值分别是什么？它们哪些可以直接写入？
- `case "admin", "owner"` 与 `fallthrough` 分别表达什么？为什么业务代码通常不需要 `fallthrough`？
- 为什么 `range` 处理 String 时 index 是 byte offset，而不是字符序号？
- `Grade(90)`、`Grade(60)`、`Grade(0)` 与 `Grade(101)` 分别应该走哪条路径？
- 如果新增 `S` 等级，哪些测试应当先改变？

## Bonus

以下内容不计入本 Block 的 45 分钟核心时间盒，仅在课后或进度提前时选做。

在不改变 `Grade` 接口的前提下，补充表驱动测试，覆盖每个等级的上下边界以及 0 和 100。先写测试并确认它能暴露遗漏，再调整实现。
