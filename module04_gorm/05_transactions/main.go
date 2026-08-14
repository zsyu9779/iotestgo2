package main

import (
	"errors"
	"fmt"
	"os"

	"iotestgo/module04_gorm/internal/classroomdb"

	"gorm.io/gorm"
)

var errInjected = errors.New("injected second-step failure")

type Wallet struct {
	ID      uint   `gorm:"primaryKey"`
	Owner   string `gorm:"size:80;uniqueIndex"`
	Balance int
}

func (Wallet) TableName() string { return "m04_l05_wallets" }

func transfer(db *gorm.DB, from, to string, amount int, failSecondStep bool) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var source, target Wallet
		if err := tx.Where("owner = ?", from).First(&source).Error; err != nil {
			return err
		}
		if err := tx.Where("owner = ?", to).First(&target).Error; err != nil {
			return err
		}
		if amount <= 0 || source.Balance < amount {
			return fmt.Errorf("insufficient balance")
		}
		if err := tx.Model(&source).Update("balance", source.Balance-amount).Error; err != nil {
			return err
		}
		if failSecondStep {
			return errInjected
		}
		return tx.Model(&target).Update("balance", target.Balance+amount).Error
	})
}

func balance(db *gorm.DB, owner string) (int, error) {
	var wallet Wallet
	if err := db.Where("owner = ?", owner).First(&wallet).Error; err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}

func run() error {
	db, err := classroomdb.Open()
	if err != nil {
		return err
	}
	defer classroomdb.Close(db)
	if err := db.AutoMigrate(&Wallet{}); err != nil {
		return err
	}
	for _, seed := range []Wallet{{Owner: "M04 Alice", Balance: 100}, {Owner: "M04 Bob", Balance: 50}} {
		if err := db.Where("owner = ?", seed.Owner).Assign(Wallet{Balance: seed.Balance}).FirstOrCreate(&seed).Error; err != nil {
			return err
		}
	}
	if err := transfer(db, "M04 Alice", "M04 Bob", 30, false); err != nil {
		return fmt.Errorf("successful transfer: %w", err)
	}
	if err := transfer(db, "M04 Alice", "M04 Bob", 1000, false); err == nil {
		return fmt.Errorf("expected insufficient balance rollback")
	}
	before, err := balance(db, "M04 Alice")
	if err != nil {
		return err
	}
	if err := transfer(db, "M04 Alice", "M04 Bob", 10, true); !errors.Is(err, errInjected) {
		return fmt.Errorf("expected injected rollback, got %v", err)
	}
	after, err := balance(db, "M04 Alice")
	if err != nil {
		return err
	}
	if before != after {
		return fmt.Errorf("rollback changed source balance: before=%d after=%d", before, after)
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.SavePoint("before_bonus").Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&Wallet{}).Where("owner = ?", "M04 Alice").Update("balance", gorm.Expr("balance + ?", 5)).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.RollbackTo("before_bonus").Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	fmt.Printf("success=ok insufficient=rolled_back second_step=rolled_back savepoint=ok alice=%d\n", after)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
