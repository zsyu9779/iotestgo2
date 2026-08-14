# Module 04：数据持久化与 GORM

Module 04 面向具备 Java 基础的本科生，由七节 75 分钟课程和一个博客综合作业组成。核心边界是数据建模、查询和事务；固定数据与本机 MySQL 用于教学，不代表生产部署方案。

## 学习结果

完成后，学员能够配置 GORM 与连接池，设计一对多/多对多模型，正确处理零值更新，发现并修复 N+1，使用事务保证一致性，编写参数化 Raw SQL，并区分 sqlmock 单测与真实 MySQL 集成测试。

## 课前环境

要求 Go 1.25.x 和 MySQL 8.0+。课堂代码固定连接本机 MySQL：

```text
root:password@tcp(127.0.0.1:3306)/gorm_demo
```

先创建对应数据库：

```sql
CREATE DATABASE gorm_demo CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

无需导出环境变量，直接运行 `make module04-env-check`。如本机账号或密码不同，修改 `internal/classroomdb/db.go` 中的 `DSN`。所有课堂表以 `m04_` 开头；如需清理，只能删除这些表，不要清空整个数据库。

## 七节课程

| 课 | 主题 | Demo |
|---|---|---|
| 1 | 连接、配置与连接池 | `go run ./module04_gorm/01_setup` |
| 2 | 模型、关系、软删除与迁移 | `go run ./module04_gorm/02_models_migrations` |
| 3 | CRUD 与零值更新 | `go run ./module04_gorm/03_crud` |
| 4 | 查询、N+1 与 Hooks | 见该课 README 的三个命令 |
| 5 | 事务、回滚与 SavePoint | `go run ./module04_gorm/05_transactions` |
| 6 | Raw SQL 与 ORM 选型 | `go run ./module04_gorm/06_raw_sql` |
| 7 | MySQL 与 sqlmock 测试 | `go test -v ./module04_gorm/07_testing_mysql` |

每课目录均包含 README、Demo、`lab/starter` 与 `lab/solution`。Starter 初始 RED 需显式运行 `-tags=exercise`。

## 综合作业与讲师入口

- [博客综合 Lab](integrated_lab/blog_api/README.md)
- [讲师 Runbook](instructor/RUNBOOK.md)
- [Demo Notes](instructor/DEMO_NOTES.md)
- [评分 Rubric](instructor/RUBRIC.md)
- [Entry Quiz](assessments/entry_quiz.md) / [Exit Quiz](assessments/exit_quiz.md)

## 验收

```bash
make module04-verify          # 离线：格式、Vet、单测、构建
make module04-lab             # 教师 Solution
make module04-env-check       # 本机 MySQL
make module04-demo-contracts  # 实际运行七课 Demo
make module04-integration     # 真实 MySQL 集成测试
make module04-audit           # 完整授课审计
```

常见故障：`Unknown database` 表示尚未建库；`Access denied` 表示代码内的账号密码不匹配或账号缺少连接、DDL 权限；重复键或列冲突通常说明使用了旧版无前缀表。
