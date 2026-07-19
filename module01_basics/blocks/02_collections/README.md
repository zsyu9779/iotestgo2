# Block 2：Collections

## 学习结果

完成本区块后，你能够：

- 说明 Array 的值语义与 Slice 共享底层数组带来的行为差异。
- 使用 `len`、`cap` 和 `append`，并识别 Slice 扩容可能改变底层数组。
- 初始化和读写 Map，使用 comma-ok 区分“不存在”与“零值”，并理解 Map 迭代无序。
- 区分 String 的 byte 数与 rune 数，正确统计中英文文本和词频。

## 时间盒：75 分钟

- 讲师 Demo：25 分钟
- 学员结对实现：40 分钟
- 测试与复盘：10 分钟

## 前置知识

完成 Block 1，能阅读 Go 的变量、`for`、`range` 和函数。你应当理解 Java 的数组、`List`、`Map` 和 `String`，但不要假设 Go 集合与它们拥有相同语义。

## Java 对比

| Java | Go |
| --- | --- |
| 数组是引用类型 | Array 是值，赋值或传参会复制全部元素 |
| `ArrayList` 隐藏容量管理 | Slice 用 `len` 和 `cap` 暴露长度与容量，`append` 可能扩容 |
| `HashMap.get` 常与 `containsKey` 配合 | Map 查询可用 `value, ok := m[key]` 同时得到值和存在性 |
| 不应依赖 `HashMap` 迭代顺序 | Go Map 迭代顺序不保证 |
| `String.length()` 计算 UTF-16 code unit | `len(string)` 计算 UTF-8 字节，rune 计数需单独处理 |

Slice 不是“数组引用”的简单别名；它是描述底层数组一段区域的值。Map 读取不存在的 key 会得到值类型的零值，需要 comma-ok 才能判断 key 是否真实存在。

## 讲师 Demo

按顺序运行并讲解：

```bash
go run ./module01_basics/blocks/02_collections/demo/04_arrays_slices
go run ./module01_basics/blocks/02_collections/demo/05_maps_strings
go run ./module01_basics/blocks/02_collections/demo/06_slice_map_edges
go run ./module01_basics/blocks/02_collections/demo/07_string_utf8_edges
```

第一个 Demo 对比 Array 复制、Slice 分片和共享修改，并观察 `append` 前后的 `len` 与 `cap`。第二个 Demo 展示 Map 初始化、comma-ok、`delete`、String 不可变性以及 byte 与 rune 的差异。

新增 Demo 补充早期项目中的基础边界：Array 的 `range` 值副本、Slice 在容量充足或不足时的底层数组行为、nil Slice 与 nil Map 的区别、Map 中 Struct 值的写回、Map 遍历无序，以及 String 常用操作和 `range string` 的 byte offset/rune 关系。

## 核心语义陷阱

- [Map 边界演示](../../bonus/dark_corners/map/main.go)：nil Map 能读不能写、Map 迭代顺序不保证、Map 中 Struct 值需要取出修改再写回，或改用指针值。
- [String/UTF-8 边界演示](../../bonus/dark_corners/string/main.go)：byte、rune、字符串遍历和字符串长度的差异。

这两份材料现在是 Block 2 的课堂深挖入口，不再只是课后 Bonus。讲师至少应从每份材料中选择一个现象让学员预测输出；Go 版本相关的行为要先标明版本。

## 学员任务

进入 `lab/starter`，实现 `textstats.Analyze(text string) Stats`。结果需包含 UTF-8 字节数、rune 数、按空白分隔的单词数，以及不区分英文大小写的词频。完整规则见 [lab/README.md](lab/README.md)。

先运行测试观察失败，再只写足以通过当前失败的代码。每次修改后重新运行测试。

## 验收命令

```bash
go test -tags=exercise ./module01_basics/blocks/02_collections/lab/starter
```

所有测试通过即完成基础任务。

## 常见错误

- 把 `len(text)` 当成“字符数”，导致中文文本统计错误。
- 使用普通空格切分，没有正确处理连续空白或换行。
- 忘记将单词转为小写，把 `Go` 和 `go` 计成两个 key。
- 自行删除标点，与实验约定的空白分词规则不一致。
- 向 nil Map 写入，或将 Map 的无序迭代误认为稳定输出。
- 忽略 Slice 共享底层数组，误以为子 Slice 修改不会影响原 Slice。

## 三级提示

1. 用 `strings.Fields` 把文本变成单词 Slice；先确定四个返回字段各自的数据来源。
2. 为词频创建 `map[string]int`，遍历单词，用 `strings.ToLower` 得到每次更新的 key。
3. `Bytes` 来自 `len(text)`，`Runes` 来自 `utf8.RuneCountInString(text)`，`Words` 来自单词 Slice 长度，`Frequencies` 来自词频 Map。

## 复盘问题

- 为什么 `"Go go 你好"` 的 byte 数和 rune 数不同？
- 为什么从 Map 中读到 0 不能证明 key 存在？
- nil Slice 和 nil Map 都是零值，为什么前者可以 `append`，后者不能直接写入？
- 为什么 Map 中的 Struct 值需要“取出、修改、写回”，而 Map 中的指针值可以直接改字段？
- 什么情况下修改一个子 Slice 会影响原 Slice？`append` 后这个答案为什么可能改变？
- 如果需要按词频高低稳定输出，为什么不能直接依赖 Map 迭代？
- `range string` 返回的 index 为什么是 byte offset，而不是字符序号？

## Bonus

先增加测试，再扩展分析结果：记录出现次数最多的单词。当多个单词次数相同时，明确一个稳定的决胜规则，不要依赖 Map 迭代顺序。
