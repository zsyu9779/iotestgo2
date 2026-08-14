# Module 04 讲师 Runbook（七节 × 75 分钟）

## 授课边界

- 学员已完成 Gin 模块并能阅读 Java DAO/JPA 代码。
- 每课固定为概念与 Java 对比 15 分钟、Demo 20 分钟、Lab 30 分钟、验收复盘 10 分钟。
- AutoMigrate、固定数据和本机账号仅用于课堂；必须明确生产环境还需要版本迁移、密钥管理、可观测性和备份。

## 课前检查

```bash
go version
make module04-verify
make module04-env-check
make module04-demo-contracts
```

先完成 Entry Quiz。打开本 Runbook、Demo Notes 和 Rubric；准备一个终端运行命令、一个终端执行 MySQL 查询。

## 七节日程

| 课 | 0–15 分钟 | 15–35 分钟 | 35–65 分钟 | 65–75 分钟 |
|---|---|---|---|---|
| 1 | GORM/sql.DB | 连接池 Demo | Pool Lab | Ping 与权限复盘 |
| 2 | Model/Tag | 关系和迁移 | Relation Lab | 软硬删除边界 |
| 3 | CRUD 语义 | 零值 Demo | Update Lab | First/Find 复盘 |
| 4 | 查询计划 | N+1/Hooks | Preload Lab | SQL 条数验收 |
| 5 | ACID | 回滚/SavePoint | Transfer Lab | 最终状态验收 |
| 6 | ORM 边界 | Raw/Exec | 参数化 Lab | 注入风险复盘 |
| 7 | 测试金字塔 | sqlmock/MySQL | Mock Lab | Exit Quiz |

第 4、5、7 课及博客 Lab 不可裁剪。延误时先裁 Java 扩展比较，再裁 Hooks 扩展示例，不裁学员验收。

## 综合 Lab

第 4 课后发放 Starter；第 5 课完成创建事务；第 7 课结束时只演示删除接口验收，学员课后独立完成。教师不得把 Solution 目录复制到学员包。

## 课后验收

```bash
make module04-lab
make module04-integration
make module04-audit
```
