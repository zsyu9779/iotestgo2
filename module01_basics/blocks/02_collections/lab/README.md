# 中英文文本统计实验

请在 `starter/textstats.go` 中实现 `Analyze(text string) Stats`，返回以下统计结果：

- `Bytes`：文本的 UTF-8 字节数，使用 `len(text)` 计算。
- `Runes`：文本的 rune 数，使用 `utf8.RuneCountInString` 计算。
- `Words`：由 `strings.Fields` 得到的单词数。
- `Frequencies`：每个单词出现的次数。

分词规则是按空白分隔：连续空白不会产生空单词。词频统计先对每个完整单词使用 `strings.ToLower`，因此英文单词不区分大小写。标点不删除，也不作为分隔符；例如 `Go, go` 的两个词项是 `"go,"` 和 `"go"`。

空文本的三个计数都应为 0，`Frequencies` 应为空 map。

从仓库根目录运行验收测试：

```bash
go test -tags=exercise ./module01_basics/blocks/02_collections/lab/starter
```

初始测试失败是正常现象。实现后重复运行，直到全部通过。

## 三级提示

1. 一级：先用 `strings.Fields` 得到单词切片，字节数、rune 数和单词数可以分别计算。
2. 二级：创建 `map[string]int`，遍历单词切片，把小写形式作为 key。
3. 三级：每看到一个单词，就将 `frequencies[strings.ToLower(word)]` 加一；返回时填入 `Stats` 的四个字段。
