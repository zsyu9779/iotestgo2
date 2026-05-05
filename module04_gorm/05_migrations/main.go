package main

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
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

	// ===== 1. AutoMigrate（自动迁移）=====
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("migration failed: " + err.Error())
	}

	// ===== 2. 手动迁移：AddColumn / ModifyColumn =====
	// 场景：User 表已存在，需要新增 Age 字段
	type UserV2 struct {
		User // 嵌入 V1 字段
		Age  int
	}

	// 手动 MigrateColumn - 使用 GORM 的 Migrator 接口
	// 这些操作也可以通过 AutoMigrate 完成，但手动控制更精确
	if !db.Migrator().HasColumn(&User{}, "Age") {
		db.Migrator().AddColumn(&User{}, "Age")
	}

	// 修改列类型（需要数据库兼容）
	// db.Migrator().AlterColumn(&User{}, "Age")

	// 重命名列
	// db.Migrator().RenameColumn(&User{}, "old_name", "new_name")

	// 创建索引
	// db.Migrator().CreateIndex(&User{}, "idx_name")

	// 删除索引
	// db.Migrator().DropIndex(&User{}, "idx_name")

	println("=== 迁移操作完成 ===")
	println("  AutoMigrate + 手动 AddColumn 已执行")
	println()
	println("  手动迁移 vs AutoMigrate：")
	println("    AutoMigrate：自动同步 model 与表结构（不删除列、不改类型）")
	println("    Migrator：精确控制每个 DDL 操作")
	println()
	println("  Java 对比：Flyway / Liquibase 版本化迁移")
	println("    GORM 的 AutoMigrate 是声明式迁移")
	println("    Flyway 是命令式迁移（每个变更写一个 SQL 文件）")
}

