package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	gorm.Model
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
}
