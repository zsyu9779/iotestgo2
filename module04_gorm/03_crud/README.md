# 第 3 课：CRUD 与零值更新

目标：批量 Create、First/Find、Update/Delete，并观察结构体 Updates 跳过零值。Java 对比：Repository save 与显式 Update。

```bash
go run ./module04_gorm/03_crud
go test ./module04_gorm/03_crud/lab/solution
go test -tags=exercise ./module04_gorm/03_crud/lab/starter
```

预期输出包含 `struct_zero=100 map_zero=0`。
