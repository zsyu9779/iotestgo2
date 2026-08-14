# 第 6 课：Raw SQL、Exec 与选型

目标：参数化查询、命名参数、聚合、批量更新和索引；判断 ORM 与 Raw SQL 的边界。Java 对比：JdbcTemplate/Native Query。

```bash
go run ./module04_gorm/06_raw_sql
go test ./module04_gorm/06_raw_sql/lab/solution
go test -tags=exercise ./module04_gorm/06_raw_sql/lab/starter
```

所有外部输入必须作为参数传递，禁止通过字符串拼接构造 SQL。
