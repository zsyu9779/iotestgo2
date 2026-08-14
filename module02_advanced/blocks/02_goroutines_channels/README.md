# Block 2：Goroutine 与 Channel

## 学习结果

完成后能够解释 goroutine 与无缓冲 channel 的协作方式，追踪并动态组合 producer/filter 流水线，并识别这种教学实现的生命周期和性能边界。

## 时间盒

70 分钟课堂精讲，不布置 Starter 或课后作业。

## 前置知识

Block 1、Module 01 的函数和循环。

## Java 对比

Goroutine 类似轻量任务，WaitGroup 可类比 CountDownLatch，Channel 可类比 BlockingQueue，但关闭语义不同。

## 讲师 Demo

```bash
go run ./module02_advanced/blocks/02_goroutines_channels/demo/03_goroutines
go run ./module02_advanced/blocks/02_goroutines_channels/demo/04_channels
```

## 精讲主线

1. `GenerateNatural` 启动一个 goroutine，按需产生 `2, 3, 4, ...`。
2. 第一次从当前数据流取出的数字一定是素数。
3. `PrimeFilter` 为该素数启动一个过滤 goroutine，并返回新的数据流。
4. `ch = PrimeFilter(ch, prime)` 每执行一次，流水线就增长一级。

```text
GenerateNatural -> Filter(2) -> Filter(3) -> Filter(5) -> ... -> next prime
```

无缓冲 channel 让下游消费速度反向约束上游，形成背压。这个示例展示的是 CSP 风格的动态并发流水线，不是欧拉线性筛，也不是面向大规模计算的高性能筛法。

## 常见错误

- 把它误称为欧拉线性筛；
- 忽略每发现一个素数就新增一个 goroutine 和 channel；
- 误以为 goroutine 串成流水线后一定会获得并行加速；
- 在长期运行的服务中照搬无限数据流，却没有取消和退出协议。

## 复盘问题

- 为什么当前数据流的第一个数字一定是素数？
- 数字 `49` 会经过哪些过滤器，最终在哪里被丢弃？
- 无缓冲 channel 如何阻止生成器无限占用内存？
- 为什么这个实现适合讲并发组合，却不适合计算超大范围素数？

## Bonus

将无限生成器和过滤器增加 `context.Context` 取消协议，观察每一级 goroutine 如何退出。
