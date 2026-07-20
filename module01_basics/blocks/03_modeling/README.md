# Block 3：Modeling

## 学习结果

完成本区块后，你能够：

- 使用指针读取和修改同一个值，并说明值传递与指针传递的差异。
- 使用 Struct 把相关字段建模为一个领域对象，并通过构造函数维护对象不变量。
- 为 Struct 定义方法，根据是否需要修改原对象选择指针接收者或值接收者。
- 使用哨兵错误表达可预期的校验失败，并返回与原对象隔离的值快照。

## 时间盒：65 分钟

- 讲师 Demo：20 分钟
- 学员结对实现：35 分钟
- 测试与复盘：10 分钟

## 前置知识

完成 Block 1 和 Block 2，能够阅读 Go 的函数、多返回值和 `error`，并能运行测试。你应当理解 Java 的类、对象引用、构造方法和实例方法，但不要求熟悉 Go 指针。

## Java 对比

| Java | Go |
| --- | --- |
| 对象变量通常保存引用 | Struct 变量默认保存值；`*Student` 才是指向该值的指针 |
| `new Student(...)` 调用构造方法 | 常用普通函数 `New(...)` 校验后返回 `*Student` |
| 字段与方法都定义在 Class 内 | Struct 声明字段，方法通过接收者与类型关联 |
| 实例方法通常能修改当前对象 | 值接收者操作副本；需要修改原值时使用指针接收者 |
| 常用异常报告构造参数非法 | Go 常返回约定的 `error`，调用方显式处理 |
| 防御性复制常返回新对象 | 小型 Struct 可以直接按值返回快照副本 |

Go 指针支持解引用和共享修改，但不支持指针算术。对 Struct 指针，`p.Field` 是 `(*p).Field` 的字段访问语法糖，两种写法操作同一字段，不会复制 Struct。方法调用时 Go 在可寻址的值与指针之间提供便利转换；接收者的选择仍然表达重要语义：这个方法是观察值，还是修改同一个对象。值接收者拿到 Struct 副本，所以即使在方法内给字段赋值，调用方的原值也不会改变。

## 讲师 Demo

按顺序运行并讲解：

```bash
go run ./module01_basics/blocks/03_modeling/demo/06_pointers
go run ./module01_basics/blocks/03_modeling/demo/07_structs_methods
go run ./module01_basics/blocks/03_modeling/demo/08_struct_zero_values
go run ./module01_basics/blocks/03_modeling/demo/09_copy_and_receivers
```

第一个 Demo 对比值传递与指针传递，观察解引用、共享修改、nil 指针，以及 `counter.Value` 与 `(*counter).Value` 操作同一字段。第二个 Demo 展示 Struct、值接收者、指针接收者和嵌入，先观察 `RenameCopy` 不会改变 `student.Name`，再观察 `Rename` 为什么能够修改原对象。

新增 Demo 补充 Struct 零值、复合字面量、值复制、指针参数、Snapshot 隔离和组合语义。注意：指针参数与指针接收者是两个相关但不同的概念；Embedding 提升字段和方法，但不建立 Java 意义上的继承关系。

## 学员任务

进入 `lab/starter`，实现学生模型：`New` 负责建立满足不变量的 `Student`，`Rename` 和 `UpdateScore` 修改原对象，`Snapshot` 返回按值复制的当前状态。姓名需要去除首尾空白，ID 必须为正数，成绩必须位于 0–100（含两端）。完整接口、错误约定与接收者说明见 [lab/README.md](lab/README.md)。

先运行测试观察失败，再围绕当前失败补充最小实现。完成后解释为什么两个修改方法使用指针接收者，而快照方法使用值接收者。

## 验收命令

```bash
go test -tags=exercise ./module01_basics/blocks/03_modeling/lab/starter
```

所有测试通过，且能说明三种方法的接收者选择，即完成基础任务。

## 常见错误

- 直接保存原始姓名，忘记先用 `strings.TrimSpace` 去除首尾空白。
- 将 0 或 100 错误地排除在有效成绩之外。
- 用值接收者实现 `Rename` 或 `UpdateScore`，方法内部改变的只是副本。
- 校验失败后仍然修改字段，破坏原对象已有的有效状态。
- 返回临时创建的文本错误，而不是约定的 `ErrInvalidID`、`ErrInvalidName` 或 `ErrInvalidScore`。
- 把 `Snapshot` 改成返回指针，使调用方能够通过快照修改原对象。

## 三级提示

1. `New` 先检查 `id > 0`，再对姓名使用 `strings.TrimSpace`，最后检查 `0 <= score && score <= 100`。
2. 修改方法先校验新值；只有校验成功后，才通过 `s.Name` 或 `s.Score` 修改指针指向的对象。
3. `Snapshot` 的接收者和返回值都是 `Student`；直接 `return s` 就会返回一份值副本。

## 复盘问题

- 为什么 `Rename` 和 `UpdateScore` 必须修改调用者持有的同一个 `Student`？
- 为什么 `Snapshot` 使用值接收者并返回 `Student`，而不是返回 `*Student`？
- Struct 的零值为什么通常可以直接使用？哪些字段或嵌套对象仍需要额外初始化？
- 指针参数和指针接收者分别解决什么问题？它们为什么仍然符合 Go 的值传递规则？
- Embedding 为什么是组合而不是继承？被提升的字段访问是否改变了对象之间的类型关系？
- `New` 集中维护不变量，相比调用方直接写 Struct 字段有什么好处？
- Java 对象引用与 Go 的 Struct 值、Struct 指针分别有哪些相似和不同？

## Bonus

先增加表驱动测试，再检查所有无效边界：ID 为 0 和负数、纯空白姓名、成绩 -1 和 101。确认每次失败都返回对应哨兵错误，并且失败的修改不会改变对象原来的有效状态。
