# Module 01 Entry Quiz（8 题）

用时 5 分钟，不查资料，每题选一个最佳答案。本 Quiz 用于识别 Java-to-Go 的起点假设，不计作业分。

## 1. 程序入口

要让 `go run` 直接运行一个最小命令行程序，哪个结构是必需的？

- A. 任意 package 中的 `public static void main`
- B. `package main` 中的 `func main()`
- C. 名为 `Main` 的 Struct 和 `Run` 方法
- D. 带 `@EntryPoint` 注解的函数

## 2. 局部变量

下列哪一项最符合 Go 的局部变量行为？

- A. 必须总是先写类型，再写变量名
- B. `name := "Gopher"` 可以在函数体内声明并推断类型
- C. 声明后不使用的局部变量不影响编译
- D. `:=` 可用于 package 级声明

## 3. 循环

如果要把 Java `while (ready) { ... }` 迁移到 Go，应优先写：

- A. `while ready { ... }`
- B. `for ready { ... }`
- C. `loop ready { ... }`
- D. `do { ... } while ready`

## 4. Switch

Go `switch` 的一个 case 命中后，没有显式 `fallthrough` 时默认会：

- A. 继续执行后面所有 case
- B. 在当前 case 结束后离开 `switch`
- C. 要求写 `break`，否则编译失败
- D. 抛出异常

## 5. 可预期失败

一个成绩函数需要报告输入超出 0–100，哪个 Go 接口最直接表达这个契约？

- A. `func Grade(score int) string`，失败时返回 `""`
- B. `func Grade(score int) string`，失败时 panic
- C. `func Grade(score int) (string, error)`
- D. 定义 `GradeException` 类

## 6. Array 赋值

```go
a := [2]int{1, 2}
b := a
b[0] = 9
```

执行后 `a` 是什么？

- A. `[9 2]`
- B. `[1 2]`
- C. `nil`
- D. 代码无法编译

## 7. Map 查询

`scores` 是 `map[string]int`。要区分“Alice 存在且成绩为 0”与“Alice 不存在”，应使用：

- A. `scores["Alice"] == nil`
- B. `scores.containsKey("Alice")`
- C. `score, ok := scores["Alice"]`
- D. 遍历 Map 并假设固定顺序

## 8. String 长度

UTF-8 字符串 `"Go你好"` 中有 4 个 rune。`len("Go你好")` 返回：

- A. 4，因为有 4 个字符
- B. 8，因为 `len(string)` 计算 UTF-8 byte
- C. 6，因为中文每字两个 byte
- D. 编译器决定，结果不固定

完成后只提交题号与选项，不在讲解前打开答案文件。
