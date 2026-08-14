package main

import (
	"fmt"
	"os"

	"iotestgo/module04_gorm/internal/classroomdb"
)

type environment struct {
	Version      string
	Charset      string
	DatabaseName string
}

func run() error {
	db, err := classroomdb.Open()
	if err != nil {
		return err
	}
	defer classroomdb.Close(db)
	var env environment
	if err := db.Raw("SELECT VERSION() AS version, @@character_set_database AS charset, DATABASE() AS database_name").Scan(&env).Error; err != nil {
		return fmt.Errorf("inspect MySQL: %w", err)
	}
	if err := db.Exec("CREATE TABLE IF NOT EXISTS m04_env_check (id BIGINT PRIMARY KEY)").Error; err != nil {
		return fmt.Errorf("DDL permission check: %w", err)
	}
	if err := db.Exec("DROP TABLE m04_env_check").Error; err != nil {
		return fmt.Errorf("cleanup permission check: %w", err)
	}
	fmt.Printf("mysql=%s database=%s charset=%s ddl=ok\n", env.Version, env.DatabaseName, env.Charset)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
