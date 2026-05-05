package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// N+1 问题演示：对比"循环查询"和"Preload 预加载"

type CategoryN struct {
	gorm.Model
	Name     string
	Products []ProductN
}

type ProductN struct {
	gorm.Model
	Name       string
	CategoryID uint
	Category   CategoryN
}

func RunNPlusOneDemo() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:password@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), // 开启 SQL 日志查看 SQL 条数
	})
	if err != nil {
		panic("db connection failed: " + err.Error())
	}
	if err := db.AutoMigrate(&CategoryN{}, &ProductN{}); err != nil {
		panic("migration failed: " + err.Error())
	}

	// 造数据：3 个分类，每个 5 个商品
	for i := 1; i <= 3; i++ {
		cat := CategoryN{Name: fmt.Sprintf("Category %d", i)}
		db.Create(&cat)
		for j := 1; j <= 5; j++ {
			db.Create(&ProductN{Name: fmt.Sprintf("Product %d-%d", i, j), CategoryID: cat.ID})
		}
	}

	fmt.Println("=== GORM N+1 问题演示 ===")
	fmt.Println()

	// 方式一：N+1 查询（坏）
	fmt.Println("--- 方式一：循环查（N+1 查询）---")
	db = db.Session(&gorm.Session{}) // 重置
	db.Logger = db.Logger.LogMode(0) // 静默日志
	var categories []CategoryN
	db.Find(&categories) // 1 条查询：SELECT * FROM categories
	for _, cat := range categories {
		var products []ProductN
		db.Where("category_id = ?", cat.ID).Find(&products) // N 条查询！
		cat.Products = products
		fmt.Printf("  Category %s: %d products\n", cat.Name, len(products))
	}
	println("  SQL 条数：1 +", len(categories), "（N 次 — N+1 问题）")

	fmt.Println()
	// 方式二：Preload 预加载（好）
	fmt.Println("--- 方式二：Preload 预加载 ---")
	var categories2 []CategoryN
	db.Preload("Products").Find(&categories2) // 2 条查询！
	for _, cat := range categories2 {
		fmt.Printf("  Category %s: %d products\n", cat.Name, len(cat.Products))
	}
	println("  SQL 条数：2（1 SELECT categories + 1 SELECT products WHERE category_id IN (...)）")
	println()

	fmt.Println("结论：")
	fmt.Println("  - N+1 问题：1 条主查询 + N 条关联查询，总 N+1 条 SQL")
	fmt.Println("  - Preload：GORM 用 IN 子查询批量加载关联数据")
	fmt.Println("  - 用 db.Debug() 查看实际 SQL 条数")
	fmt.Println("  - 生产环境务必避免 N+1！")
	fmt.Println()
	fmt.Println("Java 对比：")
	fmt.Println("  JPA: @EntityGraph + @NamedEntityGraph 或 JOIN FETCH")
	fmt.Println("  GORM: Preload() / Joins()")
}
