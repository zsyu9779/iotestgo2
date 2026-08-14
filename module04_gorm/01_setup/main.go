package main

import (
	"fmt"
	"os"
	"time"

	"iotestgo/module04_gorm/internal/classroomdb"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Code  string `gorm:"size:32;uniqueIndex"`
	Price uint
}

func (Product) TableName() string { return "m04_l01_products" }

func run() error {
	db, err := gorm.Open(mysql.Open(classroomdb.DSN), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("open MySQL: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get connection pool: %w", err)
	}
	defer sqlDB.Close()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("ping MySQL: %w", err)
	}
	if err := db.AutoMigrate(&Product{}); err != nil {
		return fmt.Errorf("migrate lesson table: %w", err)
	}
	p := Product{Code: "M04-L01-A001", Price: 100}
	if err := db.Where("code = ?", p.Code).Assign(Product{Price: p.Price}).FirstOrCreate(&p).Error; err != nil {
		return fmt.Errorf("seed product: %w", err)
	}
	fmt.Printf("connection=ok max_open=10 product=%s price=%d\n", p.Code, p.Price)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
