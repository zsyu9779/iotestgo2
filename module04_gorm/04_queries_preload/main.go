package main

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string
	Products []Product
}

type Product struct {
	gorm.Model
	Name       string
	CategoryID uint
}

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:password@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	}
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Category{}, &Product{})

	c := Category{Name: "Tech", Products: []Product{{Name: "Laptop"}, {Name: "Phone"}}}
	db.Create(&c)

	var out Category
	db.Preload("Products").First(&out)
}
