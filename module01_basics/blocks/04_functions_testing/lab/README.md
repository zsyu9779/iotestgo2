# 成绩筛选与审计实验

请在 `starter/scores.go` 中完成三个函数，练习把函数当作值传递、用闭包保存配置，以及用 `defer` 保证收尾动作的执行位置。

完成两个已有测试后，运行 `TestFilterPreservesInputOrderForAdditionalCase`，把它当作一次学员自己补充的回归测试：它要求 `Filter` 保持原输入中的相对顺序。Starter 的这个测试初始失败是预期的，不能通过删除测试或修改期望值绕过。

保持以下公开接口不变：

```go
func Filter(values []int, keep func(int) bool) []int
func AtLeast(min int) func(int) bool
func WithAudit(name string, record func(string), operation func())
```

`Filter` 按原顺序遍历 `values`，仅保留使 `keep` 返回 `true` 的值。`AtLeast` 返回一个闭包；该闭包捕获 `min`，并判断传入值是否大于或等于它。因此 `Filter([]int{59, 60, 75, 90}, AtLeast(60))` 应得到 `[]int{60, 75, 90}`。

`WithAudit` 必须先记录 `"start:" + name`，再执行 `operation`，最后记录 `"end:" + name`。请使用 `defer` 安排结束记录。对于名称 `average`，可观察到的事件顺序必须是：

```text
start:average
operation
end:average
```

从仓库根目录一次只运行一个具名测试。先关注函数值与闭包：

```bash
go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter -run '^TestFilterWithClosure$'
```

实现并通过它之后，再单独运行 `defer` 测试：

```bash
go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter -run '^TestWithAuditRecordsEndAfterOperation$'
```

两个具名测试都通过后，运行整个练习包，确认没有回归：

```bash
go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter
```

每一步都遵循 RED、GREEN、REFACTOR：先读失败信息，再写刚好能使当前测试通过的代码，最后在保持测试为绿色的前提下整理命名或重复。

## 三级提示

1. 一级：`Filter` 遍历每个值并调用传入的 `keep`；只有结果为 `true` 时才追加到结果 Slice。
2. 二级：`AtLeast` 的返回值本身是一个函数，该函数仍然可以使用外层参数 `min`。
3. 三级：`WithAudit` 先调用 `record("start:" + name)`，随后注册结束记录的 `defer`，最后调用 `operation()`。
