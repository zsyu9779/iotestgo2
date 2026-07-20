# Module 01 Demo Notes

这份小抄只服务讲师现场 Demo。学员任务、检查点和裁剪决策以 [RUNBOOK](RUNBOOK.md) 为准。以下命令均从仓库根目录运行；Demo 时先问“你预测会发生什么？”，再执行。

## 课前 15 分钟

```bash
go version
go env GOMOD
make module01-verify
make module01-demo-contracts
make module01-teaching-failures
```

预期 `module01-verify` 检查根 Go Module 中的 Module 01 代码，并进入 `teacher/solution` 调用与学员包相同的 `scripts/grade.sh`。它不运行独立 Go Module `student_pack` 中故意保留的 RED。`module01-demo-contracts` 验证正常 Demo 输出；`module01-teaching-failures` 则确认隔离教学 Case 按预期以匹配诊断失败。

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

### 必须看到

- `02_vars_types`：`iota edge cases: 0 1 2 250 250 5 6`、`iota reset: 0 1`，以及 `Zero values: int=0 float=0.0 bool=false string=""`。
- `03_control_funcs`：`switch init: owner has full access`、`Owner branch` 和 `Admin branch reached by fallthrough`；循环体只出现 `loop body: 0` 与 `loop body: 2`。

### 投影提示

1. 在执行前请学员指出 `package main`、`import` 和 `func main`，对比 Java 的 Class 入口。
2. 在 `02_vars_types` 指出 `:=` 只能用在函数体内，整数类型转换必须显式。
3. 在 `03_control_funcs` 先预测 `switch` 是否会贯穿，然后指出 Go 不需要默认 `break`。
4. 把 `calculate` 的两个返回值迁移到 `Grade(score) (string, error)`：失败是接口的显式部分，不是隐藏的异常通道。

| 暂停点 | 先问 | 预期结果 | 准确解释 | 常见误解 | 级别 |
| --- | --- | --- | --- | --- | --- |
| `left, _, right` | `_` 能否打印？ | 只能打印 1 和 3 | `_` 丢弃值，不绑定可读变量 | 把 `_` 当普通变量 | 核心 |
| `const messageLength = len(message)` | `messageLength` 是运行时才计算的吗？ | 3 | `message` 是常量时，`len(message)` 是编译期常量 | 所有 `len` 调用都在运行时计算 | 核心 |
| `consD/consE/consF` | E、F 分别是多少？ | 250、5 | E 复用 `= 250`；iota 未停止，F 使用当前行号 5 | “iota 被打断后重新从 0 开始” | 核心 |
| 第二个 const 块 | resetA 是多少？ | 0 | 每个 const 块独立重置 iota | iota 在整个包连续计数 | 深挖 |
| `UserID` / alias | 哪个需要转换？ | UserID 需要，alias 不需要 | 定义类型与别名语义不同 | 两种 type 写法完全一样 | 深挖 |
| `floatValue := 1.9; int(floatValue)` | 输出 1 还是 2？ | 1 | 变量从浮点转整数会截断小数；`int(1.9)` 会编译失败，因为 1.9 不能精确表示为 int | 数值转换会自动舍入，或 `int(1.9)` 可直接编译 | 核心 |
| if 初始化 | if 外打印哪个 score？ | 外层仍为 50 | 初始化变量只在 if/else 作用域内，外层变量未改变 | 认为短声明修改了外层变量 | 核心 |
| continue / break | 1 和 3 会不会进入循环体？ | 1 被跳过，3 终止循环 | continue 跳过本轮，break 结束当前循环 | 两者都只是跳过本轮 | 核心 |
| 无限 for | 如何安全退出？ | attempts 到 3 时 break | `for {}` 是合法无限循环，需要显式退出条件 | Go 必须写 while | 核心 |
| switch 初始化 | currentRole 在 switch 外可用吗？ | 不可用 | 初始化变量作用域属于 switch | 与普通赋值作用域相同 | 深挖 |
| 多值 case | owner 匹配哪个分支？ | 第一个分支 | 逗号分隔值表示 OR | 一个 case 只能有一个值 | 核心 |
| fallthrough | 下一 case 条件会重算吗？ | 不会 | 无条件进入紧邻 case，业务代码通常不建议使用 | 相当于继续匹配条件 | 深挖 |
| Atoi 失败 | 错误能否忽略？ | err 非 nil | 外部输入解析属于可预期失败，应显式处理 | 转换失败自动得到 0 且安全 | 核心 |

**一级提示：**让学员只预测当前一行输出。

**二级提示：**指出当前操作的是值、地址、底层数组、Map key 或闭包捕获变量中的哪一个。

**三级提示：**运行最小 Case，只解释 actual/want 或编译器第一条诊断，不直接展示 Solution。

### 不得讲错

- `iota` 始终按 ConstSpec 行递增；表达式省略与 `iota` 计数是两个维度。
- String 下标和切片按 byte；`range` 解码 rune，index 仍是 byte offset。
- nil Slice 可以 `append`；nil Map 可以读但不能写。

### 失败演示与救援

```bash
go test -tags=exercise ./module01_basics/blocks/01_go_basics/lab/starter -run '^TestGrade$'
```

起始 Starter 应因返回等级不符合而失败。请学员只读出“测试名、实际值、期望值”；如 8 分钟仍无进展，给一级提示“先把输入分成有效与无效”，不投影 solution。

### 受控编译失败与 panic（按需，8 分钟内）

```bash
make module01-teaching-failures
```

课堂只选择三个核心 Case：同作用域无新变量的 `:=`、Map Struct 字段不可直接赋值、nil Map 写入 panic。逐一先让学员预测诊断或 panic，再展示 [受控失败 Case](../teaching_failures/README.md) 中的正确写法。其余 Case（包级短声明、定义类型赋值、Slice 比较、最后一个 case 的 `fallthrough`、nil 指针）仅按班级情况调用，不逐个占用课堂时间。命令非零且诊断匹配才是通过；路径、权限、依赖或工具链异常不是教学成功。

**一级提示：**让学员只预测当前一条诊断或 panic。

**二级提示：**指出当前操作的是声明位置、可寻址 Map value，还是 nil Map 写入。

**三级提示：**展示正确写法的最小差异，只解释编译器第一条诊断或 panic；不修改失败 fixture，也不展示 Solution。

### Block 1 内容扩展

按班级节奏选择以下镜头，不要删除原有 `01_hello`、`02_vars_types` 和 `03_control_funcs`：

```bash
go run ./module01_basics/blocks/01_go_basics/demo/05_strings_basics
```

**必须看到：**`"A你" bytes=4`；`range over string` 下恰好两次 rune 输出；`Atoi error: true`。

- **核心：**让学员预测 `var int/string/bool` 的零值，说明 nil Slice 可以 `append`，nil Map 只能读不能写。
- **核心：**展示 `strings.TrimSpace`、`Fields`、`ToLower`，再解释 String 下标是 byte、`range` 得到 rune；展示 String/`[]byte` 往返和 `Atoi` 的成功与错误路径。
- **核心：**让学员预测 `case "admin", "owner"`、`switch {}`、`continue/break` 和无限 `for` 的输出；指出 if 初始化变量的作用域。
- **深挖：**实际运行 `fallthrough`，强调它不重新判断下一个 case；说明 switch 初始化变量只属于 switch。
- **可裁：**标签、`goto`、switch 初始化作用域细节和 String 转换成功路径只口头说明或按需演示，不在当天编码。
- **深挖入口：**在 `range` 基础示例后按需打开 [range dark corner](../bonus/dark_corners/range/main.go)，先说明其 Go 1.22+ 语义，再与 Go 1.16 作业环境区分。

## Block 2：Collections（25 分钟）

```bash
go run ./module01_basics/blocks/02_collections/demo/04_arrays_slices
go run ./module01_basics/blocks/02_collections/demo/05_maps_strings
go run ./module01_basics/blocks/02_collections/demo/06_slice_map_edges
go run ./module01_basics/blocks/02_collections/demo/07_string_utf8_edges
```

### 必须看到

- `04_arrays_slices`：`Original Array: [1 2 3]`、`Original Slice after sub-slice mod: [1 999 3 4 5]`，以及 `copy count=3 dst=[10 20 30 0 0] src=[10 20 30]`。
- `05_maps_strings`：nil Map 读取为 `0`、Map 遍历标注“不保证顺序”，以及 `nested map value: ready`。
- `06_slice_map_edges`：`array after changing range value: [1 2 3]` 与 `shared slice: base=[9 2] view=[9 2 3] len=3 cap=4`；**救援：**先只预测修改的是 range 值副本还是共享的底层数组，再指出当前操作的对象。
- `07_string_utf8_edges`：`text="A你" bytes=4 runes=2` 与 `byte-index=1 rune=你`；**救援：**让学员先标出 byte-index 是 1，再将它与 rune 个数 2 分开解释。

### 投影提示

1. 在修改 `arrCopy[0]` 和 `subSlice[0]` 前暂停，让学员分别预测原值是否改变：Array 赋值复制元素，Slice 描述符副本仍可共享底层数组。
2. 在 `append` 前后同时读 `len` 和 `cap`；只说“可能换底层数组”，不承诺固定扩容倍数。
3. 对 Map 提问：成绩 0 和不存在的学生如何区分？答案必须包含 comma-ok。
4. 对 `"Hello, 世界"` 先收集 `len` 预测，再比较 13 bytes 与 9 runes。

| 暂停点 | 先问 | 预期结果 | 准确解释 | 常见误解 | 级别 |
| --- | --- | --- | --- | --- | --- |
| `copy(destination, source)` | 返回值和剩余两个位置是什么？ | 3，剩余为 0 | 复制数量为 min(len(dst), len(src)) | copy 会自动扩展 dst | 核心 |
| 嵌套 Map | 只 make 外层能否直接写内层？ | 不能，会 panic | 每层 Map 都要初始化 | 外层 make 会递归初始化 | 深挖 |

**一级提示：**让学员只预测当前一行输出。

**二级提示：**指出当前操作的是值、地址、底层数组、Map key 或闭包捕获变量中的哪一个。

**三级提示：**运行最小 Case，只解释 actual/want 或编译器第一条诊断，不直接展示 Solution。

### 不得讲错

- `append` 可能复用也可能更换底层数组；不能承诺固定扩容倍数。
- Map 遍历顺序不是稳定契约。

### 失败演示与救援

```bash
go test -tags=exercise ./module01_basics/blocks/02_collections/lab/starter -run '^TestAnalyze$'
```

把第一个失败字段当作当前任务，不同时修四个计数。如 Map 赋值 panic，询问 `Frequencies` 是 nil 还是已经 `make`；如 byte/rune 混淆，让学员单独打印 `len("你")` 和 `utf8.RuneCountInString("你")`。

### Block 2 内容扩展

- **核心：**用 `06_slice_map_edges` 展示 Array 的 range 值副本、Slice 共享和容量变化、nil Slice/nil Map、comma-ok、Map Struct 写回。
- **核心：**在 `04_arrays_slices` 用 `copy(destination, source)` 说明目标在前、返回复制数量和未覆盖位置保持零值。
- **核心：**用 `07_string_utf8_edges` 展示 `Contains`、`Split`、`Join`、`TrimSpace`、`Fields`、大小写转换和 `range string`。
- **深挖：**在 `05_maps_strings` 说明嵌套 Map 必须逐层初始化；穿插 [Map dark corner](../bonus/dark_corners/map/main.go) 的 Map 顺序、nil Map、不可寻址值；穿插 [String dark corner](../bonus/dark_corners/string/main.go) 的 byte/rune 细节。
- **可裁：**`unsafe.Pointer`、指针运算和复杂 Map 指针场景不进入核心课堂。

## Block 3：Modeling（20 分钟）

```bash
go run ./module01_basics/blocks/03_modeling/demo/06_pointers
go run ./module01_basics/blocks/03_modeling/demo/07_structs_methods
go run ./module01_basics/blocks/03_modeling/demo/08_struct_zero_values
go run ./module01_basics/blocks/03_modeling/demo/09_copy_and_receivers
```

### 必须看到

- `06_pointers`：`After value pass: 5`、`After pointer pass: 100`、`pointer field sugar: 3`。
- `09_copy_and_receivers`：`value receiver mutation keeps: Alice` 与 `pointer receiver: 1:Bob=80`。
- `07_structs_methods`：`Updated: John Doe` 与 `Admin Name: Admin, Level: 1`；**救援：**先问更新发生在同一个对象还是副本，再把提升字段解释为组合提供的访问便利，不等同于继承。
- `08_struct_zero_values`：`struct zero value: main.User{ID:0, Name:"", Enabled:false, Address:main.Address{City:""}}` 与 `after pointer parameter: shared value`；**救援：**让学员分别圈出 Struct 零值字段和指针参数修改的对象，再回到“传入的是地址值的副本”。

### 投影提示

1. 在 `modifyValue` 和 `modifyPointer` 前分别预测 `val`，用输出强调 Go 只有值传递；传入指针时复制的值是地址。
2. 把 `User.String` 标为只读观察，把 `UpdateName` 标为修改同一对象，让学员选接收者。
3. 将 Java 类的构造器对比为普通 `New(...)` 函数，强调校验成功后才创建有效状态。
4. Embedding 只做“组合可以提升字段和方法”的一句展示，不说它等于继承。

| 暂停点 | 先问 | 预期结果 | 准确解释 | 常见误解 | 级别 |
| --- | --- | --- | --- | --- | --- |
| `counter.Value` / `(*counter).Value` | 是否修改同一字段？ | 最终为 3 | Go 自动解引用字段访问 | `p.field` 会复制对象 | 深挖 |
| `RenameCopy` | student.Name 会变吗？ | 仍为 Alice | 值接收者得到 Struct 副本 | 方法天然拥有 Java this 引用语义 | 核心 |
| `Rename` | 为什么变为 Bob？ | 指针接收者修改同一对象 | 传入的仍是值，只是该值为地址 | Go 存在“引用传递” | 核心 |

**一级提示：**让学员只预测当前一行输出。

**二级提示：**指出当前操作的是值、地址、底层数组、Map key 或闭包捕获变量中的哪一个。

**三级提示：**运行最小 Case，只解释 actual/want 或编译器第一条诊断，不直接展示 Solution。

### 不得讲错

- Go 所有参数传递都是值传递；传指针时复制的是地址值。
- 值接收者操作副本；指针接收者用于修改同一对象，但仍是值传递。

### 失败演示与救援

```bash
go test -tags=exercise ./module01_basics/blocks/03_modeling/lab/starter -run '^TestStudentLifecycle$'
```

当修改丢失时，让学员圈出接收者是 `Student` 还是 `*Student`；当无效修改污染状态时，让其标出第一次赋值和第一次校验的顺序。救援只给“先校验，后赋值”，不代写方法。

### Block 3 内容扩展

- **核心：**用 `08_struct_zero_values` 观察 Struct 零值、复合字面量、值复制、指针参数和 nil 指针检查。
- **核心：**用 `09_copy_and_receivers` 对比 `RenameCopy` 的副本修改、指针接收者、快照隔离和 Embedding 的字段提升。
- **深挖：**用 `06_pointers` 说明 `counter.Value` 是 `(*counter).Value` 的自动解引用语法糖；强调 Embedding 是组合，不是继承；讲解指针参数和指针接收者都遵守值传递，但复制的值不同。
- **可裁：**自动解引用语法糖、匿名字段的复杂方法集、指针嵌入和反射不进入核心课堂。

## Block 4：Functions & Testing（15 分钟）

核心 Demo 保留函数值、闭包和 `defer`，并增加变参、命名函数类型、闭包状态和 defer 求值；柯里化、函数组合和并发继续移入可裁 Bonus。

```bash
go run ./module01_basics/blocks/04_functions_testing/demo/09_advanced_functions
go run ./module01_basics/blocks/04_functions_testing/demo/10_function_forms
go run ./module01_basics/blocks/04_functions_testing/demo/11_defer_edges
go test ./module01_basics/blocks/04_functions_testing/lab/solution -run '^TestFilterWithClosure$'
```

### 必须看到

- `10_function_forms`：`closure state: 11 12` 与 `independent closure state: 11`，证明两个 counter 状态独立。
- `11_defer_edges`：LIFO 段先输出 `defer second registered` 再输出 `defer first registered`，普通参数读取 `1`，闭包读取 `2`。
- `09_advanced_functions`：`75 passed: true`、`filtered: [60 75]`，以及 `start` 后才出现 `end`；**救援：**先让学员只预测谓词对 75 的结果和过滤后的顺序，再指出函数值是参数、`defer` 的 `end` 在当前函数返回时执行。

### 投影提示

1. 将 `func(int) bool` 直接与 Java `Predicate<Integer>` 对比：Go 的函数值不需要额外函数式接口。
2. 让学员预测 `atLeast(60)` 返回的闭包为何能继续访问 `min`，再观察 `passed(75)`。
3. 在 `defer` 注册处指向“现在声明意图，函数返回前执行”，不延伸 panic/recover。
4. 用 `-run` 说明测试名是反馈接口；失败时先读测试名，再读 actual/want，最后回到最小行为。

| 暂停点 | 先问 | 预期结果 | 准确解释 | 常见误解 | 级别 |
| --- | --- | --- | --- | --- | --- |
| 两个 counter | 第二个 counter 会接着第一个计数吗？ | 不会，独立输出 11 | 每次调用工厂函数各自捕获一份状态 | 所有闭包共享同一个局部变量 | 核心 |
| defer LIFO | 哪个 defer 先执行？ | 后注册的先执行 | defer 栈按 LIFO 执行 | defer 按书写顺序执行 | 核心 |
| 参数与闭包 | 为什么一个读 1、一个读 2？ | 普通参数 1，闭包 2 | 普通参数在注册时求值；闭包体在执行时读取 | 两者都在 defer 执行时才取值 | 核心 |

**一级提示：**让学员只预测当前一行输出。

**二级提示：**指出当前操作的是值、地址、底层数组、Map key 或闭包捕获变量中的哪一个。

**三级提示：**运行最小 Case，只解释 actual/want 或编译器第一条诊断，不直接展示 Solution。

### 不得讲错

- defer 普通参数在注册时求值；闭包体读取发生在执行时。

### 失败演示与救援

```bash
go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter -run '^TestFilterWithClosure$'
go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter -run '^TestWithAuditRecordsEndAfterOperation$'
```

始终一次只运行一个具名测试。Filter 有多余零值时，检查是否先创建了 `len(values)` 长度再 `append`；audit 顺序错误时，让学员在纸上写出三个可观察事件，再移动 `defer`。

### Block 4 内容扩展

- **核心：**用 `10_function_forms` 展示命名函数类型、变参、Slice 展开、高阶函数和闭包状态隔离。
- **核心：**用 `11_defer_edges` 展示 LIFO、普通参数立即求值和闭包延迟读取。
- **核心：**让学员运行并解释 `scores_additional_exercise_test.go`，把“补一个测试”纳入测试反馈循环。
- **可裁：**柯里化、函数组合、Functional Options 和并发继续留在 Bonus；`panic/recover` 放到 Module 02。

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
(cd module01_basics/homework/task_manager/student_pack && make grade)
```

预期 gofmt 和 Vet 通过，测试在 `Add` 的 `ErrNotImplemented` 行为处失败。该 RED 证明 Starter 和公开测试已对接，不需要现场“修好”。如因权限、依赖或编译失败，按 [Troubleshooting](../homework/task_manager/teacher/TROUBLESHOOTING.md) 处理环境问题。

讲师答案的绿色对照只在课前或课后运行：

```bash
make module01-homework-solution
```

不向学员投影 `teacher/solution`。

## Exit Quiz 诊断（5 分钟）

课堂计时内只收核心题 1–14。15–17 是提前完成者或下一次课 Code Review 的选做诊断，不占用 Task Manager 启动、`make grade` 或分支交付说明的时间。

| 题号 | 正确答案 | 回收误解的讲师语言 |
| --- | --- | --- |
| 15（选做） | A | 选 B 时追问“省略的 ConstSpec 复用了哪一整行表达式？”；选 C 时重申 iota 在同一个 const 块内不重置。 |
| 16（选做） | B | 让学员先指出 `copy(dst, src)` 中谁是目标，再计算 `min(len(dst), len(src))`；不要把返回值误当成目标长度。 |
| 17（选做） | C | 选 A 时让学员区分 Map index 的 value 与可寻址变量；选 B 时先分类这是编译失败还是运行时 panic，再看编译器第一条诊断。 |
