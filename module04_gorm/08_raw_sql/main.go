// 08 原生 SQL 与 SQL Builder
//
// 何时放弃 ORM、何时用原生 SQL、何时用 SQL Builder
package main

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID     uint
	Name   string
	Age    int
	Status int
}

type SumResult struct {
	Status int
	Count  int
	AvgAge float64
}

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:password@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db connection failed: " + err.Error())
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("migration failed: " + err.Error())
	}

	fmt.Println("=== 原生 SQL 与 SQL Builder ===")
	fmt.Println()

	// ===== 1. Raw 查询 =====
	fmt.Println("--- 1. db.Raw() 复杂查询 ---")
	var result User
	db.Raw("SELECT id, name, age FROM users WHERE id = ?", 1).Scan(&result)
	fmt.Println("  Raw query with param")

	var sumResults []SumResult
	db.Raw(`SELECT status, COUNT(*) as count, AVG(age) as avg_age
            FROM users
            WHERE age > ?
            GROUP BY status
            HAVING count > ?
            ORDER BY count DESC`, 18, 2).Scan(&sumResults)
	fmt.Println("  Raw complex SQL with GROUP BY/HAVING/aggregation")

	// ===== 2. Exec 批量更新 =====
	fmt.Println()
	fmt.Println("--- 2. db.Exec() 批量操作 ---")
	db.Exec("UPDATE users SET status = ? WHERE age < ?", 0, 18)
	db.Exec("DELETE FROM users WHERE created_at < ?", "2024-01-01")

	// ===== 3. 命名参数 =====
	fmt.Println()
	fmt.Println("--- 3. 命名参数 ---")
	db.Raw("SELECT * FROM users WHERE name = @name AND age = @age",
		sql.Named("name", "Alice"),
		sql.Named("age", 25),
	).Scan(&result)

	// ===== 4. 返回普通类型 =====
	fmt.Println()
	fmt.Println("--- 4. 标量查询 ---")
	var count int64
	db.Raw("SELECT COUNT(*) FROM users").Scan(&count)

	var names []string
	db.Raw("SELECT name FROM users").Scan(&names)

	// ===== 5. 执行存储过程/DDL =====
	fmt.Println()
	fmt.Println("--- 5. DDL / 存储过程 ---")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_users_name ON users(name)")

	// ===== 决策指南 =====
	fmt.Println()
	fmt.Println("=== 何时用 GORM vs Raw SQL vs SQL Builder ===")
	fmt.Println()
	fmt.Println("用 GORM (ORM)：")
	fmt.Println("  ✓ 简单 CRUD（单表增删改查）")
	fmt.Println("  ✓ 关联查询（Preload 够用时）")
	fmt.Println("  ✓ 事务（普通 ACID 事务）")
	fmt.Println("  ✓ AutoMigrate（原型阶段）")
	fmt.Println()
	fmt.Println("用 db.Raw() (原生 SQL)：")
	fmt.Println("  ✓ 复杂聚合查询（GROUP BY, HAVING, 窗口函数）")
	fmt.Println("  ✓ 子查询、UNION、CTE")
	fmt.Println("  ✓ 存储过程调用")
	fmt.Println("  ✓ 批量 DDL 操作")
	fmt.Println("  ✓ SQL 性能调优（需要精确控制 SQL）")
	fmt.Println()
	fmt.Println("用 SQL Builder（如 squirrel, dbr）：")
	fmt.Println("  ✓ 动态查询条件（需要大量 IF 判断拼 SQL）")
	fmt.Println("  ✓ 复杂的多表 JOIN 条件")
	fmt.Println("  ✓ 既要类型安全又要动态 SQL")
	fmt.Println()
	fmt.Println("Java 对比：")
	fmt.Println("  GORM     = JPA/Hibernate（ORM）")
	fmt.Println("  db.Raw() = JdbcTemplate / Native Query")
	fmt.Println("  Builder  = jOOQ（SQL Builder）")
}
