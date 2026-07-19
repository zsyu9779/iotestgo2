# Module 01 内容增量合并设计

## 目标

在不破坏当前 `module01_basics` 教学结构、Lab、Scorebook、Task Manager 和验收流程的前提下，将早期 `/Users/zhangshiyu/iotestgo` 中较完整的 Go 语言素材增量合并进当前 Module 01，并把 `module01_basics/bonus/dark_corners` 中的基础语义陷阱穿插到核心课程。

本次工作的重点是补足语言内容的广度和语义深度，不重做课程架构，也不要求课堂一次讲完所有新增材料。Runbook 会明确核心、可裁剪和进阶层次，由讲师根据班级节奏选择。

## 不变的主线

以下内容保留为当前 Module 01 的权威主线，不被旧项目整体替换：

- 四个 Block 的顺序、学习目标和时间盒。
- 现有 Starter、公开练习测试、Solution 和显式 `exercise` build tag。
- Scorebook 综合 Lab 的接口、检查点和返回副本/不变量要求。
- Task Manager 学员包、教师答案、验收脚本和发布边界。
- 当前 Module 01 的唯一学员入口和唯一讲师入口。

新增内容优先通过新 Demo、补充段落、复盘问题、讲师提示和评估题接入，不删除当前主线文件中的有效示例。

## 素材合并原则

### 1. 保留与增量

当前 Demo 继续承担最小入门镜头；早期项目内容按主题拆成额外 Demo 或语义补充。重复主题不做二选一，而是形成由浅入深的序列：

```text
当前最小示例 → 早期项目扩展语法 → dark corner 反例 → Lab 迁移
```

### 2. 课堂内容分层

每个主题按三层组织：

- 核心：当天应理解并能解释的基本语义。
- 深挖：课堂有余力时穿插的边界行为，尤其是 `dark_corners`。
- 后置：保留在 Bonus 或后续 Module 的独立专题。

“核心”描述完整语言语义的最低覆盖，不等同于每个班级都必须在当天讲完所有示例。

### 3. 教学清洗

旧项目只作为内容来源，不直接复制其过时或不严谨的注释。合并时必须修正：

- 定义新类型与类型别名的区别。
- Embedding 是组合和方法提升，不等同于 Java 继承。
- 指针参数、指针接收者和“Go 只有值传递”的关系。
- String 的 byte/rune、`range` index 和 UTF-8 切片边界。
- Go 版本相关的循环变量语义。
- 过时 API、错误处理方式和容易误导初学者的类比。

`unsafe.Pointer`、Cgo、反射、网络、文件、并发和接口继续按现有 Module 规划处理，不因旧项目素材丰富而提前塞入 Module 01 核心。

## Block 内容映射

### Block 1：Go Basics

保留当前入口、变量、控制流、函数和显式错误返回主线，增加：

- 变量声明形式、零值、`:=`、空白标识符。
- 常量、`iota`、类型转换、定义新类型与类型别名的区别。
- 常用基本类型、整数溢出边界和位运算的轻量示例。
- `if` 初始化语句与作用域。
- `for` 的无限循环、条件循环、三段式循环、`break` 和 `continue`。
- 普通 `switch`、多值 `case`、`default`、无表达式 `switch`。
- `fallthrough` 的真实行为、限制和“不建议在业务代码中使用”的结论。
- 标签控制嵌套循环；`goto` 只作为可识别但不推荐的扩展。
- 多返回值、命名返回值、变参函数和基础字符串/整数转换。

`mygolang` 中复杂的交互输入和 `goto` 场景只提炼语义，不作为课堂主实验。

### Block 2：Collections

保留当前 Array、Slice、Map、UTF-8 和 TextStats 主线，增加：

- Array 的值复制、`range` 值副本和按索引修改。
- `[...]T` 长度推断。
- Slice 的 `len`、`cap`、子 Slice 共享、`append` 可能换底层数组。
- nil Slice、`copy` 和 Slice 传参语义。
- Map 的初始化、nil Map 能读不能写、comma-ok、delete 和遍历无序。
- Map 中 Struct 值取出修改再写回，与 Map 指针值的差别。
- String 拼接、比较、下标、byte 切片以及 `strings` 常用函数。
- `range string` 的 byte offset/rune 行为与 `[]rune` 的适用场景。

`dark_corners/map` 和 `dark_corners/string` 在这一 Block 中作为语义陷阱穿插；不把所有深层示例都变成独立 Lab。

### Block 3：Modeling

保留当前 Student 模型、不变量、方法和 Snapshot 主线，增加：

- Struct 零值、字段初始化、复合字面量和嵌套 Struct。
- Struct 值传递与指针传递。
- 值接收者与指针接收者的选择及方法调用的自动取址/解引用便利。
- 匿名字段和 Embedding 的组合语义，明确排除“继承”类比。
- 返回副本、浅拷贝和可变字段隔离边界。
- nil 指针的可观察行为和安全检查。

Embedding 不作为 Block 3 的主要 Lab 要求，除非讲师选择深挖。

### Block 4：Functions & Testing

保留当前函数值、闭包、`defer` 和测试反馈主线，增加：

- 多返回值和错误返回的不同组合。
- 命名返回值的可读性边界。
- 变参函数以及变参 Slice 展开。
- 函数类型、函数作为参数、函数作为返回值。
- 闭包捕获状态和多个闭包之间的状态隔离。
- `defer` 的 LIFO 顺序、参数立即求值、闭包延迟求值。
- `_test.go`、`TestXxx`、`testing.T`、`t.Run`、表格驱动测试和 `-run`。
- 至少一个由学员补写的边界测试，而不只是实现已有测试。

`panic/recover` 不在本 Block 展开，只在讲师说明中标记为 Module 02 的错误处理专题。

## Dark Corners 的核心穿插点

原有 `module01_basics/bonus/dark_corners` 文件先保留，避免破坏现有路径；核心入口通过 Block README、Demo Notes 和 Runbook 链接到它们，必要时新增课堂友好的短 Demo。

- `dark_corners/range`：放在 Block 1/Block 2 的 `for`、`range` 和循环变量之后；标明 Go 1.22+ 循环变量语义，避免与 Go 1.16 作业要求混淆。
- `dark_corners/map`：放在 Block 2 Map 初始化、遍历和结构体值之后。
- `dark_corners/string`：放在 Block 2 String byte/rune 和字符串遍历之后。

每个 dark corner 必须配套四句话：现象、原因、推荐写法、版本或适用边界。它们用于预测输出、短讨论或讲师演示，不强制扩展课堂验收面。

## 评估与文档同步

扩充内容时同步更新：

- Block README 的学习结果、Demo 命令、常见错误和复盘问题。
- `instructor/DEMO_NOTES.md` 的预测题、讲解顺序和裁剪建议。
- `instructor/RUNBOOK.md` 的核心/深挖/可裁内容标记。
- Entry/Exit Quiz 对新增核心语义的覆盖，尤其是零值、String、`range`、多值 `case`、`fallthrough`、Map nil 和测试基础。
- `bonus/README.md`，说明 dark corners 已由核心入口引用，其他 Bonus 仍由讲师选择。

现有 Lab 的公开接口和行为契约不因新增内容改变；新增测试只验证新增示例或学习目标，不把实现细节强行加入原有 Lab。

## 排除项

本次不做：

- 重写四 Block 或更换 Scorebook/Task Manager 领域模型。
- 把早期项目所有目录直接搬进 `module01_basics`。
- 将接口、反射、并发、Cgo、文件、网络、泛型和复杂函数式模式提前纳入核心。
- 用新增内容挤掉现有 Starter 实操、Scorebook 检查点或 Task Manager 作业启动；新增内容由讲师按 Runbook 裁剪。

## 验收标准

完成后应满足：

1. 当前 Module 01 原有 Lab、Scorebook 和 Task Manager 教师答案仍可通过完整验收。
2. 每个新增核心主题至少有一个可运行 Demo、一个讲师复盘点和一个学员可观察结果。
3. `dark_corners` 在 Block 文档和 Runbook 中有明确穿插位置，不再只出现在 Bonus 总览。
4. 新增内容不依赖 Module 02 的接口、并发或反射知识。
5. 所有 Go 示例经过 `gofmt`、`go vet` 和对应测试/运行命令验证。
6. 文档不再把 String 基础操作、变量零值、`range` 基础和 Map nil 语义当作仅供课后阅读的内容。
