package main

import (
	"errors"
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
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Wallet{})
	db.Create(&Wallet{Owner: "Alice", Balance: 100})
	db.Create(&Wallet{Owner: "Bob", Balance: 50})

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
		panic(err)
	}
}
