# Module 01 Demo Notes

这份小抄只服务讲师现场 Demo。学员任务、检查点和裁剪决策以 [RUNBOOK](RUNBOOK.md) 为准。以下命令均从仓库根目录运行；Demo 时先问“你预测会发生什么？”，再执行。

## 课前 15 分钟

```bash
go version
go env GOMOD
make module01-verify
```

预期 `module01-verify` 检查根 Go Module 中的 Module 01 代码，并进入 `teacher/solution` 调用与学员包相同的 `scripts/grade.sh`。它不运行独立 Go Module `student_pack` 中故意保留的 RED。

打开下列页面，但不在学员投影画面中展示 `solution/` 或 [answer key](../assessments/answer_key.md)：

- [Block 1 Lab](../blocks/01_go_basics/lab/README.md)
- [Block 2 Lab](../blocks/02_collections/lab/README.md)
- [Block 3 Lab](../blocks/03_modeling/lab/README.md)
- [Block 4 Lab](../blocks/04_functions_testing/lab/README.md)
- [Scorebook Lab](../integrated_lab/scorebook/README.md)
- [Task Manager 学员说明](../homework/task_manager/student_pack/README.md)

## Block 1：Go Basics（15 分钟）

```bash
go run ./module01_basics/blocks/01_go_basics/demo/01_hello
go run ./module01_basics/blocks/01_go_basics/demo/02_vars_types
go run ./module01_basics/blocks/01_go_basics/demo/03_control_funcs
```

### 投影提示

1. 在执行前请学员指出 `package main`、`import` 和 `func main`，对比 Java 的 Class 入口。
2. 在 `02_vars_types` 指出 `:=` 只能用在函数体内，整数类型转换必须显式。
3. 在 `03_control_funcs` 先预测 `switch` 是否会贯穿，然后指出 Go 不需要默认 `break`。
4. 把 `calculate` 的两个返回值迁移到 `Grade(score) (string, error)`：失败是接口的显式部分，不是隐藏的异常通道。

### 失败演示与救援

```bash
go test -tags=exercise ./module01_basics/blocks/01_go_basics/lab/starter -run '^TestGrade$'
```

起始 Starter 应因返回等级不符合而失败。请学员只读出“测试名、实际值、期望值”；如 8 分钟仍无进展，给一级提示“先把输入分成有效与无效”，不投影 solution。

## Block 2：Collections（25 分钟）

```bash
go run ./module01_basics/blocks/02_collections/demo/04_arrays_slices
go run ./module01_basics/blocks/02_collections/demo/05_maps_strings
```

### 投影提示

1. 在修改 `arrCopy[0]` 和 `subSlice[0]` 前暂停，让学员分别预测原值是否改变：Array 赋值复制元素，Slice 描述符副本仍可共享底层数组。
2. 在 `append` 前后同时读 `len` 和 `cap`；只说“可能换底层数组”，不承诺固定扩容倍数。
3. 对 Map 提问：成绩 0 和不存在的学生如何区分？答案必须包含 comma-ok。
4. 对 `"Hello, 世界"` 先收集 `len` 预测，再比较 13 bytes 与 9 runes。

### 失败演示与救援

```bash
go test -tags=exercise ./module01_basics/blocks/02_collections/lab/starter -run '^TestAnalyze$'
```

把第一个失败字段当作当前任务，不同时修四个计数。如 Map 赋值 panic，询问 `Frequencies` 是 nil 还是已经 `make`；如 byte/rune 混淆，让学员单独打印 `len("你")` 和 `utf8.RuneCountInString("你")`。

## Block 3：Modeling（20 分钟）

```bash
go run ./module01_basics/blocks/03_modeling/demo/06_pointers
go run ./module01_basics/blocks/03_modeling/demo/07_structs_methods
```

### 投影提示

1. 在 `modifyValue` 和 `modifyPointer` 前分别预测 `val`，用输出强调 Go 只有值传递；传入指针时复制的值是地址。
2. 把 `User.String` 标为只读观察，把 `UpdateName` 标为修改同一对象，让学员选接收者。
3. 将 Java 类的构造器对比为普通 `New(...)` 函数，强调校验成功后才创建有效状态。
4. Embedding 只做“组合可以提升字段和方法”的一句展示，不说它等于继承。

### 失败演示与救援

```bash
go test -tags=exercise ./module01_basics/blocks/03_modeling/lab/starter -run '^TestStudentLifecycle$'
```

当修改丢失时，让学员圈出接收者是 `Student` 还是 `*Student`；当无效修改污染状态时，让其标出第一次赋值和第一次校验的顺序。救援只给“先校验，后赋值”，不代写方法。

## Block 4：Functions & Testing（15 分钟）

主 Demo 输出很长，投影时只聚焦函数值、闭包和 `defer`；柯里化、函数组合和并发为可裁 Bonus。

```bash
go run ./module01_basics/blocks/04_functions_testing/demo/09_advanced_functions
go test ./module01_basics/blocks/04_functions_testing/demo/09_advanced_functions -run '^TestURLBuilder$'
```

### 投影提示

1. 将 `func(int) bool` 直接与 Java `Predicate<Integer>` 对比：Go 的函数值不需要额外函数式接口。
2. 调用同一 `counter` 三次，让学员解释为何么 `createCounter` 返回后 `count` 仍可访问。
3. 在 `defer` 注册处指向“现在声明意图，函数返回前执行”，不延伸 panic/recover。
4. 用 `-run` 说明测试名是反馈接口；失败时先读测试名，再读 actual/want，最后回到最小行为。

### 失败演示与救援

```bash
go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter -run '^TestFilterWithClosure$'
go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter -run '^TestWithAuditRecordsEndAfterOperation$'
```

始终一次只运行一个具名测试。Filter 有多余零值时，检查是否先创建了 `len(values)` 长度再 `append`；audit 顺序错误时，让学员在纸上写出三个可观察事件，再移动 `defer`。

## Scorebook 综合 Lab（40 分钟）

不做新 Demo，只投影 [Scorebook Lab](../integrated_lab/scorebook/README.md) 的当前检查点命令。在 10、20、30 分钟时说一句“保留当前输出，转到下一检查点”；40 分钟必须停止。

救援顺序固定为：

1. 请学员读出第一个失败测试。
2. 请其读出 actual 和 want，并指出哪个不变量被破坏。
3. 给 Lab README 的一级提示；仍阻塞才给二级。
4. 不打开 solution，不越过当前检查点预写后续方法。

## Task Manager 启动（最后 20 分钟内）

先声明“这是一周作业，现在不写实现”，再运行：

```bash
cd module01_basics/homework/task_manager/student_pack
make grade
```

预期 gofmt 和 Vet 通过，测试在 `Add` 的 `ErrNotImplemented` 行为处失败。该 RED 证明 Starter 和公开测试已对接，不需要现场“修好”。如因权限、依赖或编译失败，按 [Troubleshooting](../homework/task_manager/teacher/TROUBLESHOOTING.md) 处理环境问题。

讲师答案的绿色对照只在课前或课后运行：

```bash
make module01-homework-solution
```

不向学员投影 `teacher/solution`。
