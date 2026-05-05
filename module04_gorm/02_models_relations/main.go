package main

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ========== 软删除演示 ==========

type Article struct {
	gorm.Model
	Title     string
	Content   string
	DeletedAt gorm.DeletedAt // 逻辑删除字段
	// 删除时 GORM 自动设置 DeletedAt = 当前时间
	// 查询时 GORM 自动过滤 DeletedAt IS NOT NULL 的记录
}

type Author struct {
	gorm.Model
	Name    string
	Posts   []Post
	DeletedAt gorm.DeletedAt
}

type Tag struct {
	gorm.Model
	Name    string
	Posts   []Post `gorm:"many2many:post_tags"`
	DeletedAt gorm.DeletedAt
}

type Post struct {
	gorm.Model
	Title     string
	Content   string
	AuthorID  uint
	Author    Author
	Tags      []Tag `gorm:"many2many:post_tags"`
	DeletedAt gorm.DeletedAt
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
	if err := db.AutoMigrate(&Author{}, &Tag{}, &Post{}, &Article{}); err != nil {
		panic("migration failed: " + err.Error())
	}

	// --- 软删除演示 ---
	article := Article{Title: "Hello", Content: "GORM soft delete"}
	db.Create(&article)

	// 删除：只设置 DeletedAt，不真正删除
	db.Delete(&article)

	// 查不到（被自动过滤）
	var result Article
	err = db.First(&result, article.ID).Error
	println("After delete, found:", err == nil) // false

	// 使用 Unscoped() 可以查到已删除记录
	db.Unscoped().First(&result, article.ID)
	println("With Unscoped, found:", result.Title) // "Hello"

	// 硬删除（真正删除）
	db.Unscoped().Delete(&result)

	println()
	println("--- 逻辑删除说明 ---")
	println("  gorm.DeletedAt 字段实现软删除")
	println("  db.Delete() → 设置 DeletedAt = now()")
	println("  db.First() → 自动加 AND deleted_at IS NULL")
	println("  db.Unscoped() → 跳过自动 filter")
	println()
	println("Java 对比：@SQLDelete + @Where 注解")
	println("  Spring Data JPA 通过注解实现，GORM 通过字段类型声明实现")

	// 原有关系演示（保留）
	a := Author{Name: "Alice"}
	t1 := Tag{Name: "go"}
	t2 := Tag{Name: "web"}
	p := Post{Title: "Intro", Content: "GORM relations", Author: a, Tags: []Tag{t1, t2}}
	db.Create(&p)

	var out Post
	db.Preload("Author").Preload("Tags").First(&out)
}
