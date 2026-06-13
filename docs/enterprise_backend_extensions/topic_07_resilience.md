# Topic 07: 韧性设计：超时、重试、熔断、限流

## 适合插入位置

Module 05 metadata/interceptor 或 Module 06 go-zero resilience 之后。

## 核心问题

分布式系统一定会部分失败。韧性设计不是让失败消失，而是让失败被限制、被观察、被恢复。

## 四个概念

- Timeout：不要无限等。
- Retry：只对可重试错误重试，必须有次数上限。
- Circuit Breaker：下游持续失败时快速失败。
- Rate Limit：保护系统不被流量打穿。

## 练习

给 API -> RPC 调用设计策略：

- 超时：3 秒
- 重试：最多 2 次，只重试 `Unavailable`
- 熔断：连续失败率超过阈值后打开
- 限流：每秒最多 100 请求
