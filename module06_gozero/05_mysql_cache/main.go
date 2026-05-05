// 05 数据库与缓存：sqlx 包装器 + Model 生成 + Cache-Aside 缓存一致性
package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== 05 数据库与缓存 ===")
	fmt.Println()

	fmt.Println("--- go-zero sqlx 包装器 ---")
	fmt.Println("  sqlx 是 go-zero 对 database/sql 的增强：")
	fmt.Println("  1. 自动处理连接池（MaxOpenConns / MaxIdleConns）")
	fmt.Println("  2. 结构化错误处理（ErrNotFound）")
	fmt.Println("  3. 支持事务、Prepare、NamedExec")
	fmt.Println("  4. 与 Model 生成工具集成")
	fmt.Println()

	fmt.Println("--- Model 生成 ---")
	fmt.Println("  从 DDL 生成 Model 代码：")
	fmt.Println("  goctl model mysql ddl -src user.sql -dir internal/model -c")
	fmt.Println("  # -c 表示生成带缓存的 Model")
	fmt.Println()
	fmt.Println("  生成的文件：")
	fmt.Println("  - userModel.go:   数据库操作（Insert/Update/Delete/FindOne）")
	fmt.Println("  - userModel_gen.go: 自动生成的基础 CRUD（勿手动修改）")
	fmt.Println("  - vars.go:        表名常量、错误定义")
	fmt.Println()

	fmt.Println("--- 用户表 DDL 示例 (user.sql) ---")
	fmt.Println(`  CREATE TABLE user (
      id BIGINT AUTO_INCREMENT PRIMARY KEY,
      username VARCHAR(64) NOT NULL UNIQUE,
      password VARCHAR(128) NOT NULL,
      email VARCHAR(128) DEFAULT '',
      status TINYINT DEFAULT 1,
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
      updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`)
	fmt.Println()

	fmt.Println("--- 带缓存的 Model 使用 ---")
	fmt.Println(`  // 自动缓存模式（生成代码中已实现）：
  user, err := userModel.FindOne(ctx, userId)
  // 内部流程：
  // 1. 先查 Redis 缓存: cache:user:id:{userId}
  // 2. 缓存命中 → 直接返回
  // 3. 缓存未命中 → 查 MySQL → 写回 Redis → 返回
  // 4. Update/Delete 时自动删除/更新 Redis 缓存`)

	fmt.Println()
	fmt.Println("--- 缓存一致性模式（Cache-Aside） ---")
	fmt.Println("  读取流程：")
	fmt.Println("  1. 读缓存 → 命中返回")
	fmt.Println("  2. 未命中 → 读数据库 → 写缓存 → 返回")
	fmt.Println()
	fmt.Println("  写入流程：")
	fmt.Println("  1. 写数据库")
	fmt.Println("  2. 删除缓存（或更新缓存）")
	fmt.Println("  go-zero 默认使用 \"先删缓存，再写数据库\" 或 \"先写数据库，再删缓存\"")
	fmt.Println()

	fmt.Println("--- 缓存穿透/击穿/雪崩防护 ---")
	fmt.Println("  穿透（查不存在的数据）：单飞（singleflight）+ 布隆过滤器")
	fmt.Println("  击穿（热点 key 过期）：singleflight 合并并发请求")
	fmt.Println("  雪崩（大量 key 同时过期）：TTL + 随机偏移")
	fmt.Println()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: MyBatis + RedisTemplate / Redisson")
	fmt.Println("  go-zero: sqlx + 内置缓存层（自动 Cache-Aside）")
	fmt.Println("  go-zero 的缓存集成更自动化，不需要手动写缓存逻辑")

	_ = fmt.Sprint
}
