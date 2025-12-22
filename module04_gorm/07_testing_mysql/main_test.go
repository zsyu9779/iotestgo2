package main

import (
    "os"
    "testing"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Note struct {
    gorm.Model
    Title string
}

func TestMySQLCreateSkipIfNoDSN(t *testing.T) {
    dsn := os.Getenv("MYSQL_DSN")
    if dsn == "" {
        t.Skip("MYSQL_DSN not set; skipping MySQL integration test")
    }
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        t.Fatalf("open: %v", err)
    }
    if err := db.AutoMigrate(&Note{}); err != nil {
        t.Fatalf("migrate: %v", err)
    }
    if err := db.Create(&Note{Title: "hello"}).Error; err != nil {
        t.Fatalf("create: %v", err)
    }
    var out Note
    if err := db.First(&out, "title = ?", "hello").Error; err != nil {
        t.Fatalf("read: %v", err)
    }
}
