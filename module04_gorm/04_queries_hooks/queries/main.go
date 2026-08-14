package main

import (
	"fmt"
	"os"

	"iotestgo/module04_gorm/internal/classroomdb"
)

type Category struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"size:80;uniqueIndex"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}

func (Category) TableName() string { return "m04_l04_categories" }

type Product struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:80;uniqueIndex"`
	Price      int
	Active     bool
	CategoryID uint
}

func (Product) TableName() string { return "m04_l04_products" }

type JoinedProduct struct {
	ProductName  string
	CategoryName string
}

func run() error {
	db, err := classroomdb.Open()
	if err != nil {
		return err
	}
	defer classroomdb.Close(db)
	if err := db.AutoMigrate(&Category{}, &Product{}); err != nil {
		return err
	}
	var category Category
	if err := db.Where("name = ?", "M04 Tech").FirstOrCreate(&category, Category{Name: "M04 Tech"}).Error; err != nil {
		return err
	}
	seeds := []Product{{Name: "M04 Laptop", Price: 9000, Active: true, CategoryID: category.ID}, {Name: "M04 Phone", Price: 5000, Active: false, CategoryID: category.ID}}
	for i := range seeds {
		if err := db.Where("name = ?", seeds[i].Name).Assign(seeds[i]).FirstOrCreate(&seeds[i]).Error; err != nil {
			return err
		}
	}
	var categories []Category
	if err := db.Preload("Products", "active = ?", true).Where("name LIKE ?", "M04%").Find(&categories).Error; err != nil {
		return fmt.Errorf("conditional preload: %w", err)
	}
	var joined []JoinedProduct
	err = db.Table("m04_l04_products AS p").
		Select("p.name AS product_name, c.name AS category_name").
		Joins("JOIN m04_l04_categories AS c ON c.id = p.category_id").
		Where("p.price >= ?", 5000).Order("p.id").Scan(&joined).Error
	if err != nil {
		return fmt.Errorf("join query: %w", err)
	}
	active := 0
	if len(categories) > 0 {
		active = len(categories[0].Products)
	}
	fmt.Printf("where=ok preload_active=%d joins=%d\n", active, len(joined))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
