package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         `gorm:"size:160;not null"`
	Content   string         `gorm:"type:text;not null"`
	Comments  []Comment      `gorm:"foreignKey:PostID"`
	Tags      []Tag          `gorm:"many2many:m04_blog_post_tags"`
}

func (Post) TableName() string { return "m04_blog_posts" }

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	PostID    uint           `gorm:"index;not null"`
	Content   string         `gorm:"type:text;not null"`
}

func (Comment) TableName() string { return "m04_blog_comments" }

type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:40;uniqueIndex;not null"`
}

func (Tag) TableName() string { return "m04_blog_tags" }
