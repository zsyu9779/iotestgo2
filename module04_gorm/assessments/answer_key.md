# Module 04 Quiz Answer Key

## Entry

1. `gorm.DB` 提供 ORM 会话/查询能力，`sql.DB` 管理连接池。
2. GORM 通过字段和 Tag 声明映射，没有 JPA 的运行时注解实体模型。
3. 一组操作全部成功或全部回滚。
4. 拼接会破坏语句与数据边界并导致 SQL 注入。
5. `First` 未找到返回 `ErrRecordNotFound`；`Find` 返回空集合且通常无错误。

## Exit

1. GORM 的结构体 Updates 默认忽略零值；使用 Map、Select 或指针。
2. 主查询后对每个父记录再查询一次；Preload 通常是两条 SQL。
3. 第一条必须回滚，最终状态与事务开始前一致。
4. 多对多标签可能仍被其他文章引用，只应清理关联行。
5. sqlmock 验证 SQL 交互契约；集成测试验证真实驱动、方言、DDL 和数据库行为。
6. 它缺少可靠的版本历史、审核、回滚和复杂数据迁移控制。
