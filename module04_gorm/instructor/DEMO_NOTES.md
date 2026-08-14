# Module 04 Demo Notes

每个 Demo 先让学员预测 SQL 和最终状态，再运行命令。不得把“程序没有报错”等同于事务正确，必须查询最终状态。

## 第 1 课

投影固定课堂 DSN、连接打开、`db.DB()` 和三个 Pool 参数。必须看到 `connection=ok`。失败演示：临时把 DSN 端口改错，确认错误能定位到数据库连接。

## 第 2 课

画出 Author→Post、Post↔Tag。运行后检查 `m04_l02_post_tags` 与 `profiles.email`。不得说 AutoMigrate 能安全完成所有生产迁移。

## 第 3 课

在运行前询问 `Updates(Item{Stock: 0})` 是否写入。必须看到 `struct_zero=100 map_zero=0`，解释 GORM 对结构体零值的筛选。

## 第 4 课

依次运行 queries、nplusone、hooks。必须看到 N+1 为 4 次、Preload 为 2 次。强调 Preload 不是 JOIN；Hook 应短小、确定且可测试。

## 第 5 课

先跑成功转账，再观察余额不足与第二步注入失败。必须看到两种失败均回滚，SavePoint 回滚不改变最终余额。

## 第 6 课

将 `Alice' OR 1=1 --` 作为参数传入 Lab。不得现场执行字符串拼接的危险 SQL；只展示其代码形态和参数化修复。

## 第 7 课

先运行离线 sqlmock，再运行带标签的集成测试。指出 Mock 证明“发出了预期 SQL”，MySQL 集成测试证明“真实驱动和数据库行为可用”，两者不能互相替代。

## 现场救援

- 连接拒绝：检查 MySQL 服务和 3306 端口。
- 权限失败：换用具备 `CREATE/ALTER/INDEX/SELECT/INSERT/UPDATE/DELETE` 的课堂账号。
- 表冲突：仅清理 `m04_*` 表后重试。
- 测试 SQL 不匹配：先开启 GORM Logger 查看实际 SQL，不要直接放宽为匹配任意语句。
