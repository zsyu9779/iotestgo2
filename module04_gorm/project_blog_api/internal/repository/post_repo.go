package repository

import (
	"iotestgo/module04_gorm/project_blog_api/internal/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
	CreateComment(comment *model.Comment) error
	FindAll() ([]model.Post, error)
	Delete(id uint) error
	// Transaction support
	WithTx(tx *gorm.DB) PostRepository
	DB() *gorm.DB
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) DB() *gorm.DB {
	return r.db
}

func (r *postRepository) WithTx(tx *gorm.DB) PostRepository {
	return &postRepository{db: tx}
}

func (r *postRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) CreateComment(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *postRepository) FindAll() ([]model.Post, error) {
	var posts []model.Post
	err := r.db.Preload("Comments").Order("id desc").Find(&posts).Error
	return posts, err
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&model.Post{}, id).Error
}
