# 第 5 课：事务、回滚与 SavePoint

目标：闭包事务、失败回滚、手动事务和 SavePoint。Java 对比：`@Transactional` 与 NESTED propagation。

```bash
go run ./module04_gorm/05_transactions
go test ./module04_gorm/05_transactions/lab/solution
go test -tags=exercise ./module04_gorm/05_transactions/lab/starter
```

预期输出包含 `second_step=rolled_back savepoint=ok`。课堂必须验证最终余额，不能只观察返回值。
