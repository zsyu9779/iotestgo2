# Topic 10: 性能分析与容量意识

## 适合插入位置

Module 02 benchmark 或 Module 06 结束后。

## 核心问题

性能优化不是猜。先测量，再定位，再优化。

## 工具顺序

1. Benchmark：函数级性能
2. pprof：CPU/内存热点
3. 压测：接口级吞吐和延迟
4. 指标：线上持续观察

## 练习

比较字符串拼接：

```bash
go test -bench=. ./module01_basics/05_maps_strings
```

让学生解释为什么 `strings.Builder` 通常比循环 `+=` 更适合大量拼接。
