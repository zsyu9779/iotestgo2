package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name  string
	Stock int
}

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:password@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	}
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Item{})

	db.Create(&Item{Name: "Book", Stock: 10})
	db.Create(&Item{Name: "Pen", Stock: 100})

	var pen Item
	db.First(&pen, "name = ?", "Pen")
	db.Model(&pen).Update("Stock", 80)

	var items []Item
	db.Find(&items)
	fmt.Println("Items count:", len(items))

	db.Delete(&pen)
}
