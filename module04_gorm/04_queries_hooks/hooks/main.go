package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"iotestgo/module04_gorm/internal/classroomdb"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        string `gorm:"primaryKey;size:36"`
	Name      string `gorm:"size:80;uniqueIndex"`
	PriceCent int
	CreatedAt time.Time
}

func (Product) TableName() string { return "m04_l04_hook_products" }

func (p *Product) BeforeCreate(*gorm.DB) error {
	p.Name = strings.TrimSpace(p.Name)
	if p.Name == "" {
		return fmt.Errorf("name is required")
	}
	if p.PriceCent < 0 {
		return fmt.Errorf("price cannot be negative")
	}
	p.ID = uuid.NewString()
	return nil
}

func (p *Product) BeforeUpdate(*gorm.DB) error {
	if p.PriceCent < 0 {
		return fmt.Errorf("price cannot be negative")
	}
	return nil
}

func run() error {
	db, err := classroomdb.Open()
	if err != nil {
		return err
	}
	defer classroomdb.Close(db)
	if err := db.AutoMigrate(&Product{}); err != nil {
		return err
	}
	if err := db.Where("name = ?", "M04 Hook Product").Delete(&Product{}).Error; err != nil {
		return err
	}
	p := Product{Name: " M04 Hook Product ", PriceCent: 129900}
	if err := db.Create(&p).Error; err != nil {
		return err
	}
	p.PriceCent = -1
	validationErr := db.Save(&p).Error
	if validationErr == nil {
		return fmt.Errorf("expected update hook validation error")
	}
	fmt.Printf("before_create=id_generated name=%q before_update=blocked_negative\n", p.Name)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
