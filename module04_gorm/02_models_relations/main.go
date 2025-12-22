package main

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string
	Posts []Post
}

type Tag struct {
	gorm.Model
	Name  string
	Posts []Post `gorm:"many2many:post_tags"`
}

type Post struct {
	gorm.Model
	Title    string
	Content  string
	AuthorID uint
	Author   Author
	Tags     []Tag `gorm:"many2many:post_tags"`
}

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:password@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	}
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Author{}, &Tag{}, &Post{})

	a := Author{Name: "Alice"}
	t1 := Tag{Name: "go"}
	t2 := Tag{Name: "web"}
	p := Post{Title: "Intro", Content: "GORM relations", Author: a, Tags: []Tag{t1, t2}}
	db.Create(&p)

	var out Post
	db.Preload("Author").Preload("Tags").First(&out)
}
