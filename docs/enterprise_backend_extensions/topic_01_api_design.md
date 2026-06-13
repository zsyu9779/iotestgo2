# Topic 01: API 设计与兼容性

## 适合插入位置

Module 03 Gin API 设计之后。

## 要解决的问题

学生通常会把接口理解为“能返回 JSON 就行”。企业项目更关心接口是否可演进、可排查、可被前端/客户端稳定使用。

## 核心概念

- 路由版本：`/api/v1/users`
- 统一错误响应：`code/message/request_id`
- 幂等性：重复请求是否安全
- 兼容性：新增字段通常安全，删除/改名字段通常危险
- 分页：`page/page_size` 或 `cursor/limit`

## 课堂 Demo

基于 `module03_web_gin/06_api_design/main.go`，把一个随意接口改成版本化接口：

```json
{
  "code": "OK",
  "message": "success",
  "request_id": "req-20260614-001",
  "data": {
    "user_id": 1,
    "username": "gopher"
  }
}
```

## 练习

让学生设计 `GET /api/v1/posts` 的响应，要求包含：

- 分页信息
- 文章列表
- 请求追踪 ID
- 空列表时的响应

## 讨论题

为什么“直接返回数据库模型”在小作业里很方便，但在长期维护中容易出问题？
