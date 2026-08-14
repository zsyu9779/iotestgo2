package main

import (
	"fmt"
	"os"

	"iotestgo/module04_gorm/internal/classroomdb"
)

type Item struct {
	ID    uint   `gorm:"primaryKey"`
	SKU   string `gorm:"size:32;uniqueIndex"`
	Name  string `gorm:"size:80"`
	Stock int
	Price int
}

func (Item) TableName() string { return "m04_l03_items" }

func run() error {
	db, err := classroomdb.Open()
	if err != nil {
		return err
	}
	defer classroomdb.Close(db)
	if err := db.AutoMigrate(&Item{}); err != nil {
		return fmt.Errorf("migrate items: %w", err)
	}
	if err := db.Where("sku LIKE ?", "M04-L03-%").Delete(&Item{}).Error; err != nil {
		return fmt.Errorf("clean lesson rows: %w", err)
	}
	items := []Item{{SKU: "M04-L03-BOOK", Name: "Book", Stock: 10, Price: 50}, {SKU: "M04-L03-PEN", Name: "Pen", Stock: 100, Price: 5}}
	if err := db.Create(&items).Error; err != nil {
		return fmt.Errorf("batch create: %w", err)
	}
	var first Item
	if err := db.First(&first, "sku = ?", "M04-L03-PEN").Error; err != nil {
		return fmt.Errorf("First: %w", err)
	}
	var found []Item
	if err := db.Find(&found, "sku LIKE ?", "M04-L03-%").Error; err != nil {
		return fmt.Errorf("Find: %w", err)
	}
	if err := db.Model(&first).Updates(Item{Stock: 0}).Error; err != nil {
		return fmt.Errorf("struct update: %w", err)
	}
	var afterStruct Item
	if err := db.First(&afterStruct, first.ID).Error; err != nil {
		return err
	}
	if err := db.Model(&first).Updates(map[string]any{"stock": 0}).Error; err != nil {
		return fmt.Errorf("map zero-value update: %w", err)
	}
	var afterMap Item
	if err := db.First(&afterMap, first.ID).Error; err != nil {
		return err
	}
	if err := db.Delete(&first).Error; err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	fmt.Printf("batch=%d first=%s struct_zero=%d map_zero=%d delete=ok\n", len(found), first.SKU, afterStruct.Stock, afterMap.Stock)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
