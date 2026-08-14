package main

import (
	"fmt"
	"log"
	"os"

	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/handler"
	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/model"
	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/repository"
	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/service"
	"iotestgo/module04_gorm/internal/classroomdb"
)

func run() error {
	db, err := classroomdb.Open()
	if err != nil {
		return err
	}
	defer classroomdb.Close(db)
	if err := db.AutoMigrate(&model.Post{}, &model.Comment{}, &model.Tag{}); err != nil {
		return fmt.Errorf("migrate blog tables: %w", err)
	}
	log.Println("Module 04 blog API listening on :8091")
	return handler.Router(service.New(repository.New(db))).Run(":8091")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
