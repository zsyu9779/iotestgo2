# Module 01 Exit Quiz（14 题）

用时 5 分钟，不查资料。每题都请根据可观察行为作答，不只选“听起来像 Go”的术语。

## 1. Array 值语义

```go
func change(a [2]int) { a[0] = 9 }
scores := [2]int{60, 80}
change(scores)
```

调用后 `scores` 是：

- A. `[9 80]`
- B. `[60 80]`
- C. `nil`
- D. 不确定

## 2. Slice 共享

```go
scores := []int{60, 80, 90}
middle := scores[1:]
middle[0] = 100
```

调用后 `scores` 是：

- A. `[60 80 90]`
- B. `[60 100 90]`
- C. `[100 80 90]`
- D. 代码 panic

## 3. Append 与底层数组

```go
base := make([]int, 2, 2)
base[0], base[1] = 1, 2
grown := append(base, 3)
grown[0] = 9
```

此时 `base` 是：

- A. `[9 2]`
- B. `[1 2]`
- C. `[1 2 3]`
- D. `[]`

## 4. Map Lookup

```go
m := map[string]int{"done": 0}
a, okA := m["done"]
b, okB := m["missing"]
```

哪个结果正确？

- A. `a=0, okA=true; b=0, okB=false`
- B. `a=0, okA=false; b=nil, okB=false`
- C. 查询 `missing` 会 panic
- D. `a` 与 `b` 的存在性无法区分

## 5. Rune 与 Byte

对 `s := "A你"`，哪个可观察结果正确？

- A. `len(s) == 2`，`range` 迭代 2 次
- B. `len(s) == 4`，`range` 迭代 2 次
- C. `len(s) == 4`，`range` 迭代 4 次且都是完整字符
- D. `len(s)` 取决于操作系统

## 6. 接收者选择

`Student.UpdateScore` 必须修改调用方持有的学生，`Student.Snapshot` 只返回隔离的值副本。最清晰的选择是：

- A. 两者都用值接收者并返回指针
- B. `UpdateScore` 用指针接收者，`Snapshot` 用值接收者并返回值
- C. 两者都必须用指针接收者
- D. 接收者与是否修改无关

## 7. 闭包

```go
func AtLeast(min int) func(int) bool {
	return func(value int) bool { return value >= min }
}
pass := AtLeast(60)
```

`AtLeast` 返回后，`pass(60)` 和 `pass(59)` 分别是：

- A. `true, false`，返回的函数仍可访问 `min`
- B. `false, false`，`min` 已被销毁
- C. `true, true`，闭包不能保存配置
- D. 第一次调用会 panic

## 8. Defer 顺序

```go
record("start")
defer record("end")
record("operation")
```

当这段代码所在函数正常返回时，记录顺序是：

- A. `start, end, operation`
- B. `end, start, operation`
- C. `start, operation, end`
- D. `start, operation`

## 9. 阅读测试失败

测试输出为：

```text
--- FAIL: TestAverage (0.00s)
    scorebook_test.go:91: Average() = 80, want 80.5
```

最佳的下一步是：

- A. 删掉 `TestAverage`，因为 80 已足够接近
- B. 先检查平均值计算是否在整数除法后才转为 `float64`
- C. 一次性重写 Scorebook 所有方法
- D. 把期望值改成 80

## 10. 返回副本

```go
student, _ := book.Find(1)
student.Score = 0
again, _ := book.Find(1)
```

若 `Find` 的契约是返回隔离的 `Student` 值副本，且书中原成绩是 90，则 `again.Score` 应为：

- A. 0
- B. 90
- C. 随机值
- D. 访问时 panic

完成后写下你最不确定的一个题号，作为下次 Code Review 的离场票。

## 11. 变量零值

```go
var n int
var ok bool
var name string
```

哪个结果正确？

- A. `n=1, ok=true, name="null"`
- B. `n=0, ok=false, name=""`
- C. 三个变量都为 `nil`
- D. 代码无法编译

## 12. 多值 Case 与 Fallthrough

```go
role := "owner"
switch role {
case "admin", "owner":
    fmt.Println("full")
case "user":
    fmt.Println("read")
}
```

输出是什么？如果在第一个 case 末尾加上 `fallthrough`，还会发生什么？

- A. 只输出 `full`；加 `fallthrough` 后无条件执行下一个 case
- B. 输出 `full` 和 `read`；加不加 `fallthrough` 都一样
- C. 代码无法编译，因为一个 case 不能有多个值
- D. 输出 `read`，因为 `owner` 会自动贯穿到下一个 case

## 13. String 基础操作

```go
words := strings.Fields("  Go\t语言 ")
```

哪个结果正确？

- A. `[]string{"", "Go", "", "语言", ""}`
- B. `[]string{"Go", "语言"}`
- C. `[]rune{'G', 'o', '语', '言'}`
- D. `nil`

## 14. 学员补充测试

为 `Filter` 增加一个回归测试时，哪个断言最有价值？

- A. 只断言函数没有 panic
- B. 只断言返回 Slice 非 nil
- C. 同时断言筛选结果和原输入中的相对顺序
- D. 修改期望值直到测试通过
