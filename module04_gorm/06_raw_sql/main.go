package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"iotestgo/module04_gorm/internal/classroomdb"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:80;uniqueIndex"`
	Age       int
	Status    int
	CreatedAt time.Time
}

func (User) TableName() string { return "m04_l06_users" }

type Summary struct {
	Status int
	Count  int
	AvgAge float64
}

func run() error {
	db, err := classroomdb.Open()
	if err != nil {
		return err
	}
	defer classroomdb.Close(db)
	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}
	for _, seed := range []User{{Name: "M04 Alice", Age: 25, Status: 1}, {Name: "M04 Bob", Age: 17, Status: 1}, {Name: "M04 Carol", Age: 31, Status: 1}} {
		if err := db.Where("name = ?", seed.Name).Assign(seed).FirstOrCreate(&seed).Error; err != nil {
			return err
		}
	}
	var selected User
	if err := db.Raw("SELECT id, name, age, status, created_at FROM m04_l06_users WHERE name = ?", "M04 Alice").Scan(&selected).Error; err != nil {
		return fmt.Errorf("parameterized Raw: %w", err)
	}
	var summaries []Summary
	if err := db.Raw(`SELECT status, COUNT(*) AS count, AVG(age) AS avg_age
		FROM m04_l06_users WHERE name LIKE ? GROUP BY status ORDER BY status`, "M04%").Scan(&summaries).Error; err != nil {
		return fmt.Errorf("aggregate Raw: %w", err)
	}
	result := db.Exec("UPDATE m04_l06_users SET status = ? WHERE name LIKE ? AND age < ?", 0, "M04%", 18)
	if result.Error != nil {
		return fmt.Errorf("parameterized Exec: %w", result.Error)
	}
	var named User
	if err := db.Raw("SELECT id, name, age, status, created_at FROM m04_l06_users WHERE name = @name", sql.Named("name", "M04 Alice")).Scan(&named).Error; err != nil {
		return fmt.Errorf("named parameter: %w", err)
	}
	if !db.Migrator().HasIndex(&User{}, "idx_m04_l06_status") {
		if err := db.Exec("CREATE INDEX idx_m04_l06_status ON m04_l06_users(status)").Error; err != nil {
			return fmt.Errorf("create controlled index: %w", err)
		}
	}
	fmt.Printf("raw=%s aggregate_groups=%d exec_rows=%d named=%s injection_safe=parameters\n", selected.Name, len(summaries), result.RowsAffected, named.Name)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
