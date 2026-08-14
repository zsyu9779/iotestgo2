# 第 1 课：连接、配置与连接池

目标：解释 `gorm.Open` 与 `sql.DB` 的职责，配置连接池并用 Ping 验证环境。Java 对比：DataSource/HikariCP。

```bash
make module04-env-check
go run ./module04_gorm/01_setup
go test ./module04_gorm/01_setup/lab/solution
go test -tags=exercise ./module04_gorm/01_setup/lab/starter
```

预期 Demo 包含 `connection=ok max_open=10`。常见失败：代码内的课堂 DSN 与本机 MySQL 不匹配、数据库不存在或账号没有建表权限。
