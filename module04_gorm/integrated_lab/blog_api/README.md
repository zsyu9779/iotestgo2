# 博客数据层综合 Lab

本 Lab 同时作为 Module 04 的课后作业。课堂完成模型、预加载和创建事务，课后完成删除事务与 HTTP 验收。

## 检查点

1. 为 `Post`、`Comment`、`Tag` 建立一对多和多对多关系。
2. 列表查询用两次 `Preload` 返回评论和标签。
3. 创建文章、规范化标签和首条评论必须位于同一事务。
4. 删除时软删除评论与文章、清除关联表，但保留共享标签。
5. HTTP 接口实现 201、204、400、404 和 500 状态码。

## Starter

```bash
go test -tags=exercise ./module04_gorm/integrated_lab/blog_api/starter/...
```

初始失败是预期 RED。依次实现 `NormalizeTags`、模型关系、事务回调和状态码。

## 教师 Solution

```bash
go test ./module04_gorm/integrated_lab/blog_api/solution/...
make run-blog-api
```

测试接口：

```bash
curl -X POST http://localhost:8091/posts \
  -H 'Content-Type: application/json' \
  -d '{"title":"GORM","content":"relations","tags":["go","gorm"]}'
curl http://localhost:8091/posts
```
