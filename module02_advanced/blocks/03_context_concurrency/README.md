# Block 3：Context 与并发安全

## 学习结果

完成后能够让任务响应取消、使用 Mutex/RWMutex/Atomic 保护共享状态，并用 race detector 验证实现。

## 时间盒

65 分钟，其中学员动手不少于 35 分钟。

## 前置知识

Block 2、channel 和 WaitGroup。

## Java 对比

Context 不是线程本地变量，而是请求生命周期树；Mutex 和 RWMutex 可类比 Java 锁，但 Go 更强调显式传递。

## 讲师 Demo

```bash
go run ./module02_advanced/blocks/03_context_concurrency/demo/05_context
go run ./module02_advanced/blocks/03_context_concurrency/demo/06_concurrency_safety
```

## 学员任务

修复一个无锁计数器，使其通过 `go test -race`；再给一个长任务增加可取消等待。

## 验收命令

```bash
go test -tags=exercise -race ./module02_advanced/blocks/03_context_concurrency/lab/starter
```

## 常见错误

- 只取消 Context，不等待 goroutine 退出；
- 在锁外读取共享字段；
- 复制包含 Mutex 的结构体；
- 用 RWMutex 代替所有 Mutex。

## 三级提示

1. 找出所有共享读写点。
2. 用 defer 保证解锁。
3. 用 `go test -race` 而不是肉眼判断正确性。

## 复盘问题

- 取消信号如何穿过多个函数层级？
- 为什么先写正确再谈锁粒度和性能？

## Bonus

比较 `sync.Once`、sync.Map 与 map+Mutex 的场景边界。
