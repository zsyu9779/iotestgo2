# Module 04：GORM 七课教师备课教案

适用对象：已完成 Gin 模块、具备 Java/SQL 基础的本科生。总课时为七节，每节 75 分钟；博客项目作为贯穿式 Lab 和课后作业，不另占第八课。

## 固定课堂结构

每节依次进行概念与 Java 对比 15 分钟、可观察 Demo 20 分钟、Starter Lab 30 分钟、自动验收与复盘 10 分钟。所有命令从仓库根目录运行。

| 课 | 主题 | Demo | 学员验收 |
|---|---|---|---|
| 1 | GORM、sql.DB、连接池 | `go run ./module04_gorm/01_setup` | `go test -tags=exercise ./module04_gorm/01_setup/lab/starter` |
| 2 | 模型、关系、软删除、AutoMigrate | `go run ./module04_gorm/02_models_migrations` | `go test -tags=exercise ./module04_gorm/02_models_migrations/lab/starter` |
| 3 | CRUD 与零值更新 | `go run ./module04_gorm/03_crud` | `go test -tags=exercise ./module04_gorm/03_crud/lab/starter` |
| 4 | Where、Join、Preload、N+1、Hooks | 见课程 README 的三个 Demo | `go test -tags=exercise ./module04_gorm/04_queries_hooks/lab/starter` |
| 5 | 事务、回滚、SavePoint | `go run ./module04_gorm/05_transactions` | `go test -tags=exercise ./module04_gorm/05_transactions/lab/starter` |
| 6 | Raw SQL、Exec 与参数化 | `go run ./module04_gorm/06_raw_sql` | `go test -tags=exercise ./module04_gorm/06_raw_sql/lab/starter` |
| 7 | sqlmock 与 MySQL 集成测试 | `go test -v ./module04_gorm/07_testing_mysql` | `go test -tags=exercise ./module04_gorm/07_testing_mysql/lab/starter` |

## 课前与课后

课前运行 `make module04-verify module04-env-check`，数据库必须是本机 MySQL 8.0+ 并配置 `MYSQL_DSN`。课后运行 `make module04-audit`。逐课讲法、不可讲错内容和救援步骤以 [Module 04 Runbook](../module04_gorm/instructor/RUNBOOK.md) 和 [Demo Notes](../module04_gorm/instructor/DEMO_NOTES.md) 为准。

## 综合博客项目

从第 4 课开始使用 [Starter/Solution](../module04_gorm/integrated_lab/blog_api/README.md)。验收重点是文章—标签多对多、文章—评论一对多、无 N+1 列表、创建事务、删除事务、共享标签保留和 HTTP 状态码。分页、认证、缓存和生产迁移不在本模块核心范围。
