package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORM Hooks 演示：生命周期回调
// Hooks 在 Create/Update/Delete/Query 前后自动执行

type Product struct {
	ID    string `gorm:"primaryKey;size:36"`
	Name  string
	Price float64
	Stock int
}

// BeforeCreate：自动生成 UUID
func (p *Product) BeforeCreate(tx *gorm.DB) error {
	fmt.Printf("  [Hook] BeforeCreate: generating UUID for %s\n", p.Name)
	p.ID = uuid.NewString()
	return nil
}

// AfterCreate：创建后日志
func (p *Product) AfterCreate(tx *gorm.DB) error {
	fmt.Printf("  [Hook] AfterCreate: product %s created with ID=%s\n", p.Name, p.ID)
	return nil
}

// BeforeUpdate：更新前校验
func (p *Product) BeforeUpdate(tx *gorm.DB) error {
	fmt.Printf("  [Hook] BeforeUpdate: validating price for %s\n", p.Name)
	if p.Price < 0 {
		return fmt.Errorf("price cannot be negative")
	}
	return nil
}

// AfterFind：查询后处理
func (p *Product) AfterFind(tx *gorm.DB) error {
	// 示例：价格单位转换（分 → 元）
	p.Price = p.Price / 100
	return nil
}

// SoftDelete 时的 Hook
func (p *Product) BeforeDelete(tx *gorm.DB) error {
	fmt.Printf("  [Hook] BeforeDelete: 准备删除商品 %s\n", p.Name)
	return nil
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
	if err := db.AutoMigrate(&Product{}); err != nil {
		panic("migration failed: " + err.Error())
	}

	fmt.Println("=== GORM Hooks 演示 ===")
	fmt.Println()

	// Create
	fmt.Println("--- Create（触发 BeforeCreate + AfterCreate）---")
	p := Product{Name: "MacBook Pro", Price: 1299900, Stock: 10}
	db.Create(&p)
	fmt.Printf("  商品ID: %s, 价格: %.2f\n\n", p.ID, p.Price)

	// Update
	fmt.Println("--- Update（触发 BeforeUpdate）---")
	db.Model(&p).Update("Price", 1399900)
	fmt.Println()

	// Update with error
	fmt.Println("--- Update with error（校验失败）---")
	err := db.Model(&p).Update("Price", -100).Error
	if err != nil {
		fmt.Printf("  校验失败: %v\n\n", err)
	}

	// Delete
	fmt.Println("--- Delete（触发 BeforeDelete）---")
	db.Delete(&p)
	fmt.Println()

	fmt.Println("=== 可用 Hooks 列表 ===")
	fmt.Println("  BeforeSave / AfterSave")
	fmt.Println("  BeforeCreate / AfterCreate")
	fmt.Println("  BeforeUpdate / AfterUpdate")
	fmt.Println("  BeforeDelete / AfterDelete")
	fmt.Println("  AfterFind")
	fmt.Println()
	fmt.Println("Java 对比：JPA @PrePersist / @PostLoad / @PreUpdate")
	fmt.Println("  GORM 通过方法签名约定（func(*gorm.DB) error）自动注册 Hook")
	fmt.Println("  JPA 通过注解声明，GORM 通过命名约定")
}
