package main

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Owner   string
	Balance int
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
	if err := db.AutoMigrate(&Wallet{}); err != nil {
		panic("migration failed: " + err.Error())
	}
	db.Create(&Wallet{Owner: "Alice", Balance: 100})
	db.Create(&Wallet{Owner: "Bob", Balance: 50})

	fmt.Println("=== GORM 事务 + SavePoint 演示 ===")
	fmt.Println()

	// --- 1. 普通事务 ---
	fmt.Println("--- 1. 普通事务（转账 30） ---")
	err := db.Transaction(func(tx *gorm.DB) error {
		var alice Wallet
		var bob Wallet
		tx.First(&alice, "owner = ?", "Alice")
		tx.First(&bob, "owner = ?", "Bob")
		if alice.Balance < 30 {
			return errors.New("insufficient")
		}
		tx.Model(&alice).Update("Balance", alice.Balance-30)
		tx.Model(&bob).Update("Balance", bob.Balance+30)
		return nil
	})
	if err != nil {
		fmt.Printf("  转账失败: %v\n", err)
	} else {
		fmt.Println("  转账成功")
	}

	// --- 2. 手动事务 + SavePoint（嵌套回滚点）---
	fmt.Println()
	fmt.Println("--- 2. 手动事务 + SavePoint ---")
	tx := db.Begin()

	var alice Wallet
	tx.First(&alice, "owner = ?", "Alice")
	fmt.Printf("  Alice 初始余额: %d\n", alice.Balance)

	// SavePoint A
	tx.SavePoint("sp_a")
	tx.Model(&alice).Update("Balance", alice.Balance-10)
	fmt.Println("  SavePoint A: 扣 10")

	// SavePoint B
	tx.SavePoint("sp_b")
	tx.Model(&alice).Update("Balance", alice.Balance-20)
	fmt.Println("  SavePoint B: 再扣 20")

	// 回滚到 SavePoint B（撤销再扣 20，回到只扣了 10 的状态）
	tx.RollbackTo("sp_b")
	tx.First(&alice, "owner = ?", "Alice")
	fmt.Printf("  RollbackTo sp_b 后余额: %d\n", alice.Balance)

	// 回滚到 SavePoint A（撤销所有操作）
	tx.RollbackTo("sp_a")
	tx.First(&alice, "owner = ?", "Alice")
	fmt.Printf("  RollbackTo sp_a 后余额: %d\n", alice.Balance)

	tx.Commit()
	fmt.Println()

	fmt.Println("--- SavePoint 使用场景 ---")
	fmt.Println("  1. 子流程失败时只回滚子流程，不影响主事务")
	fmt.Println("  2. 复杂业务中的多级回滚")
	fmt.Println("  3. 类似 Java 的 TransactionManager + TransactionDefinition")
	fmt.Println()
	fmt.Println("Java 对比：")
	fmt.Println("  Spring: @Transactional(propagation = Propagation.NESTED)")
	fmt.Println("  GORM: tx.SavePoint() / tx.RollbackTo()")
}
