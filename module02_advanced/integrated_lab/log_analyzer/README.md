# 并发日志分析器综合实验

本实验组合接口、错误、Goroutine、Channel、Context、测试和 Benchmark。不要复制任一 Block 的 Solution。

## 时间盒：40 分钟

1. 0–5 分钟：实现同步 `CountErrors`。
2. 5–15 分钟：使用固定数量 worker 并明确 Channel 关闭方。
3. 15–22 分钟：拒绝小于 1 的 worker 数量。
4. 22–30 分钟：让所有阻塞接收和发送响应 Context。
5. 30–35 分钟：补空输入、预取消和处理中取消测试。
6. 35–40 分钟：运行 Benchmark 与 race detector。

Starter 验收：

```bash
go test -tags=exercise -race ./module02_advanced/integrated_lab/log_analyzer/starter
```

教师 Solution：

```bash
go test -race ./module02_advanced/integrated_lab/log_analyzer/solution
```

只有日志生产者关闭输入 Channel，只有协调者在全部 worker 退出后关闭结果 Channel。Context 取消不能替代等待 goroutine 退出。
