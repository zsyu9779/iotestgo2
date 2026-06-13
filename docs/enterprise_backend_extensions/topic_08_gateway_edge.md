# Topic 08: API Gateway 与边缘层职责

## 适合插入位置

Module 05 grpc-gateway 或 Module 06 网关部署之后。

## 核心问题

Gateway 不是业务逻辑垃圾桶。它适合处理跨服务统一问题，不适合塞满订单、用户、商品规则。

## Gateway 适合做

- TLS 终止
- 鉴权前置
- 路由
- 限流
- 灰度
- 请求 ID 注入
- 基础审计

## Gateway 不适合做

- 订单价格计算
- 库存扣减
- 用户状态变更
- 数据库事务

## 练习

让学生把 10 个功能分到 Gateway / API Service / RPC Service 三层。
