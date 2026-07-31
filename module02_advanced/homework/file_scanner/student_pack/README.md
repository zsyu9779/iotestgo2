# Module 02 Homework：可取消的并发文件扫描器

实现 `scanner.Scan`，使用固定数量 worker 扫描目录中的普通文件，统计包含指定文本的行数。

```go
type Result struct {
    Files         int
    MatchingLines int
}

func Scan(ctx context.Context, root, needle string, workers int) (Result, error)
```

要求：

- `root` 必须是存在的目录，`needle` 不能为空，`workers` 必须大于 0。
- 使用 `filepath.WalkDir` 发现文件，使用固定数量 goroutine 读取文件。
- 每个阻塞发送和接收都必须能够响应 Context 取消。
- 文件遍历或读取失败时返回包含文件路径的错误。
- CLI 支持 `-root`、`-text`、`-workers` 和 `-timeout`。
- 只使用 Go 标准库。

验收：

```bash
make grade
```

Starter 初始会在测试阶段因 `ErrNotImplemented` 失败，这是预期 RED。不要删除或弱化公开测试。
