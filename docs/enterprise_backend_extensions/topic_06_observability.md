# Topic 06: 可观测性：日志、指标、链路追踪

## 适合插入位置

Module 06 `07_observability` 之后。

## 核心问题

上线后的系统不能靠猜。可观测性回答三类问题：

- 发生了什么：日志
- 现在健康吗：指标
- 慢在哪里：链路追踪

## 三大支柱

| 类型 | 适合回答 | 例子 |
|---|---|---|
| Logs | 单次事件细节 | 某个订单创建失败的错误 |
| Metrics | 趋势和告警 | QPS、P95、错误率 |
| Traces | 跨服务路径 | API -> RPC -> DB 花了多久 |

## 练习

给 `order-api` 的一次请求设计 5 个日志字段：

- request_id
- user_id
- order_id
- latency_ms
- error_code
