# 第 4 课：高级查询、N+1 与 Hooks

目标：Where、Join、条件 Preload、N+1 查询计数和安全生命周期 Hook。Java 对比：JOIN FETCH、EntityGraph 和 JPA Callback。

```bash
go run ./module04_gorm/04_queries_hooks/queries
go run ./module04_gorm/04_queries_hooks/nplusone
go run ./module04_gorm/04_queries_hooks/hooks
go test ./module04_gorm/04_queries_hooks/lab/solution
```

N+1 Demo 应输出 `n_plus_one=4 preload=2`。Hook 只做规范化、ID 生成和校验，不修改字段的业务单位。
