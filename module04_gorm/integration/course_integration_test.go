//go:build integration

package integration

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"iotestgo/module04_gorm/internal/classroomdb"

	"gorm.io/gorm"
)

type Parent struct {
	ID        uint `gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt
	Name      string  `gorm:"size:160;uniqueIndex"`
	Children  []Child `gorm:"foreignKey:ParentID"`
}

func (Parent) TableName() string { return "m04_it_parents" }

type Child struct {
	ID       uint `gorm:"primaryKey"`
	ParentID uint
	Name     string
}

func (Child) TableName() string { return "m04_it_children" }

type MigratedV1 struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (MigratedV1) TableName() string { return "m04_it_migrations" }

type MigratedV2 struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func (MigratedV2) TableName() string { return "m04_it_migrations" }

type Wallet struct {
	ID      uint `gorm:"primaryKey"`
	Owner   string
	Balance int
}

func (Wallet) TableName() string { return "m04_it_wallets" }

func openMySQL(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := classroomdb.Open()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = classroomdb.Close(db) })
	return db
}

func TestModelsMigrationCRUDAndPreload(t *testing.T) {
	db := openMySQL(t)
	if err := db.AutoMigrate(&Parent{}, &Child{}, &MigratedV1{}); err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&MigratedV2{}); err != nil || !db.Migrator().HasColumn(&MigratedV2{}, "Email") {
		t.Fatalf("V2 migration failed: %v", err)
	}
	name := fmt.Sprintf("m04-it-%d", time.Now().UnixNano())
	parent := Parent{Name: name, Children: []Child{{Name: "one"}, {Name: "two"}}}
	if err := db.Create(&parent).Error; err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		db.Unscoped().Where("parent_id = ?", parent.ID).Delete(&Child{})
		db.Unscoped().Delete(&parent)
	})
	var loaded Parent
	if err := db.Preload("Children").First(&loaded, parent.ID).Error; err != nil || len(loaded.Children) != 2 {
		t.Fatalf("preload children=%d err=%v", len(loaded.Children), err)
	}
	if err := db.Model(&loaded).Updates(Parent{Name: ""}).Error; err != nil {
		t.Fatal(err)
	}
	var afterStruct Parent
	db.First(&afterStruct, loaded.ID)
	if afterStruct.Name == "" {
		t.Fatal("struct zero value should have been skipped")
	}
	if err := db.Model(&loaded).Updates(map[string]any{"name": ""}).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Delete(&loaded).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.First(&Parent{}, loaded.ID).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("soft delete lookup error=%v", err)
	}
}

func TestTransactionRollbackSavePointAndRawSQL(t *testing.T) {
	db := openMySQL(t)
	if err := db.AutoMigrate(&Wallet{}); err != nil {
		t.Fatal(err)
	}
	owner := fmt.Sprintf("m04-wallet-%d", time.Now().UnixNano())
	wallet := Wallet{Owner: owner, Balance: 100}
	if err := db.Create(&wallet).Error; err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Unscoped().Delete(&wallet) })
	want := errors.New("rollback")
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&wallet).Update("balance", 70).Error; err != nil {
			return err
		}
		return want
	})
	if !errors.Is(err, want) {
		t.Fatal(err)
	}
	var gotBalance int
	if err := db.Raw("SELECT balance FROM m04_it_wallets WHERE id = ?", wallet.ID).Scan(&gotBalance).Error; err != nil || gotBalance != 100 {
		t.Fatalf("raw balance=%d err=%v", gotBalance, err)
	}
	tx := db.Begin()
	if err := tx.SavePoint("before_change").Error; err != nil {
		t.Fatal(err)
	}
	if err := tx.Model(&wallet).Update("balance", 90).Error; err != nil {
		t.Fatal(err)
	}
	if err := tx.RollbackTo("before_change").Error; err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit().Error; err != nil {
		t.Fatal(err)
	}
}
