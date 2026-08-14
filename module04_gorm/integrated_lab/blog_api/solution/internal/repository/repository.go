package repository

import (
	"context"
	"fmt"

	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	WithinTransaction(context.Context, func(PostRepository) error) error
	CreatePost(context.Context, *model.Post) error
	FindOrCreateTag(context.Context, string) (*model.Tag, error)
	ReplaceTags(context.Context, *model.Post, []model.Tag) error
	CreateComment(context.Context, *model.Comment) error
	ListPosts(context.Context) ([]model.Post, error)
	FindPost(context.Context, uint) (*model.Post, error)
	DeleteComments(context.Context, uint) error
	ClearTags(context.Context, *model.Post) error
	DeletePost(context.Context, *model.Post) error
}

type repository struct{ db *gorm.DB }

func New(db *gorm.DB) PostRepository { return &repository{db: db} }

func (r *repository) WithinTransaction(ctx context.Context, fn func(PostRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error { return fn(&repository{db: tx}) })
}

func (r *repository) CreatePost(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *repository) FindOrCreateTag(ctx context.Context, name string) (*model.Tag, error) {
	tag := model.Tag{Name: name}
	if err := r.db.WithContext(ctx).Where("name = ?", name).FirstOrCreate(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *repository) ReplaceTags(ctx context.Context, post *model.Post, tags []model.Tag) error {
	return r.db.WithContext(ctx).Model(post).Association("Tags").Replace(tags)
}

func (r *repository) CreateComment(ctx context.Context, comment *model.Comment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

func (r *repository) ListPosts(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post
	err := r.db.WithContext(ctx).Preload("Comments").Preload("Tags").Order("id DESC").Find(&posts).Error
	return posts, err
}

func (r *repository) FindPost(ctx context.Context, id uint) (*model.Post, error) {
	var post model.Post
	if err := r.db.WithContext(ctx).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *repository) DeleteComments(ctx context.Context, postID uint) error {
	return r.db.WithContext(ctx).Where("post_id = ?", postID).Delete(&model.Comment{}).Error
}

func (r *repository) ClearTags(ctx context.Context, post *model.Post) error {
	if err := r.db.WithContext(ctx).Model(post).Association("Tags").Clear(); err != nil {
		return fmt.Errorf("clear tags: %w", err)
	}
	return nil
}

func (r *repository) DeletePost(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Delete(post).Error
}
