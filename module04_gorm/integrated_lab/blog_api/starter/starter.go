package starter

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var ErrNotImplemented = errors.New("not implemented")

type Post struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string
	Content   string
	// TODO: add Comments and Tags relationships.
}

func (Post) TableName() string { return "m04_blog_posts" }

type Comment struct {
	ID        uint           `gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	PostID    uint
	Content   string
}

func (Comment) TableName() string { return "m04_blog_comments" }

type Tag struct {
	ID   uint
	Name string
}

func (Tag) TableName() string { return "m04_blog_tags" }

func NormalizeTags([]string) []string {
	// TODO: trim, lowercase, drop empty values and de-duplicate.
	return nil
}

type TransactionRunner interface {
	WithinTransaction(func() error) error
}

func CreatePostWithComment(TransactionRunner, string, string, string, []string) error {
	// TODO: put post, tag association and comment creation in one transaction.
	return ErrNotImplemented
}

func DeletePost(TransactionRunner, uint) error {
	// TODO: clear associations and soft-delete comments/post in one transaction.
	return ErrNotImplemented
}
