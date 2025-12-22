package service

import (
	"iotestgo/module04_gorm/project_blog_api/internal/model"
	"iotestgo/module04_gorm/project_blog_api/internal/repository"

	"gorm.io/gorm"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) ListPosts() ([]model.Post, error) {
	return s.repo.FindAll()
}

func (s *PostService) CreatePost(title, content string) (*model.Post, error) {
	post := &model.Post{Title: title, Content: content}
	if err := s.repo.Create(post); err != nil {
		return nil, err
	}
	return post, nil
}

// CreatePostWithComment demonstrates a transaction: create post AND a default comment
func (s *PostService) CreatePostWithComment(title, content, commentText string) (*model.Post, error) {
	var newPost *model.Post

	// Use the DB from repo to start a transaction
	err := s.repo.DB().Transaction(func(tx *gorm.DB) error {
		txRepo := s.repo.WithTx(tx)

		// 1. Create Post
		post := &model.Post{Title: title, Content: content}
		if err := txRepo.Create(post); err != nil {
			return err // Rollback
		}

		// 2. Create Comment
		comment := &model.Comment{PostID: post.ID, Content: commentText}
		if err := txRepo.CreateComment(comment); err != nil {
			return err // Rollback
		}

		newPost = post
		return nil // Commit
	})

	if err != nil {
		return nil, err
	}
	return newPost, nil
}
