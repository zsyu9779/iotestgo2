# Block 2：Goroutine 与 Channel

## 学习结果

完成后能够启动并等待 goroutine，设计 producer/consumer 通信协议，并正确关闭 channel。

## 时间盒

70 分钟，其中学员动手不少于 40 分钟。

## 前置知识

Block 1、Module 01 的函数和循环。

## Java 对比

Goroutine 类似轻量任务，WaitGroup 可类比 CountDownLatch，Channel 可类比 BlockingQueue，但关闭语义不同。

## 讲师 Demo

```bash
go run ./module02_advanced/blocks/02_goroutines_channels/demo/03_goroutines
go run ./module02_advanced/blocks/02_goroutines_channels/demo/04_channels
```

## 学员任务

实现 `CollectSquares` worker pool，要求处理全部输入、保持结果顺序，并正确处理空输入和 worker 边界。

## 验收命令

```bash
go test -tags=exercise ./module02_advanced/blocks/02_goroutines_channels/lab/starter
```

## 常见错误

- main 提前退出；
- 接收方关闭 channel；
- 向已关闭 channel 写入；
- range channel 前没有关闭发送端。

## 三级提示

1. 先为每个 goroutine 增加 WaitGroup 计数。
2. 让唯一发送方负责 close。
3. 用完成信号测试所有 worker 已退出。

## 复盘问题

- Channel 解决了通信问题，但是否自动解决了所有共享状态问题？
- 无缓冲和有缓冲 channel 分别表达什么协作关系？

## Bonus

观察 nil Channel 动态禁用 case、多个就绪 case 的选择和 `default` 忙等风险。
