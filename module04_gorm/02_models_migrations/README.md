# 第 2 课：模型、关系、软删除与迁移

目标：掌握标签、一对多、多对多、软/硬删除和 V1→V2 `AutoMigrate`。Java 对比：JPA Entity、`@ManyToMany` 与 Flyway。

```bash
go run ./module04_gorm/02_models_migrations
go test ./module04_gorm/02_models_migrations/lab/solution
go test -tags=exercise ./module04_gorm/02_models_migrations/lab/starter
```

预期输出包含 `relations=ok`、`migration=email`。生产环境应使用版本化迁移工具，不能把 AutoMigrate 当作完整 Flyway 替代品。
