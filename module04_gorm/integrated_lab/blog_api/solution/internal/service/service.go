package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/model"
	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/repository"

	"gorm.io/gorm"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("post not found")
)

type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CommentResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

type PostResponse struct {
	ID        uint              `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	Tags      []TagResponse     `json:"tags"`
	Comments  []CommentResponse `json:"comments"`
	CreatedAt time.Time         `json:"created_at"`
}

type PostService interface {
	CreatePost(context.Context, string, string, []string) (*PostResponse, error)
	CreatePostWithComment(context.Context, string, string, string, []string) (*PostResponse, error)
	ListPosts(context.Context) ([]PostResponse, error)
	DeletePost(context.Context, uint) error
}

type service struct{ repo repository.PostRepository }

func New(repo repository.PostRepository) PostService { return &service{repo: repo} }

func normalizeText(value string) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", ErrInvalidInput
	}
	return value, nil
}

func NormalizeTags(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.ToLower(strings.TrimSpace(value))
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func (s *service) create(ctx context.Context, title, content string, tags []string, comment *string) (*PostResponse, error) {
	var err error
	if title, err = normalizeText(title); err != nil {
		return nil, fmt.Errorf("title: %w", err)
	}
	if content, err = normalizeText(content); err != nil {
		return nil, fmt.Errorf("content: %w", err)
	}
	if comment != nil {
		normalized, normalizeErr := normalizeText(*comment)
		if normalizeErr != nil {
			return nil, fmt.Errorf("comment: %w", normalizeErr)
		}
		comment = &normalized
	}
	post := &model.Post{Title: title, Content: content}
	err = s.repo.WithinTransaction(ctx, func(txRepo repository.PostRepository) error {
		if err := txRepo.CreatePost(ctx, post); err != nil {
			return err
		}
		for _, name := range NormalizeTags(tags) {
			tag, err := txRepo.FindOrCreateTag(ctx, name)
			if err != nil {
				return err
			}
			post.Tags = append(post.Tags, *tag)
		}
		if err := txRepo.ReplaceTags(ctx, post, post.Tags); err != nil {
			return err
		}
		if comment != nil {
			created := model.Comment{PostID: post.ID, Content: *comment}
			if err := txRepo.CreateComment(ctx, &created); err != nil {
				return err
			}
			post.Comments = append(post.Comments, created)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("create post transaction: %w", err)
	}
	response := toResponse(*post)
	return &response, nil
}

func (s *service) CreatePost(ctx context.Context, title, content string, tags []string) (*PostResponse, error) {
	return s.create(ctx, title, content, tags, nil)
}

func (s *service) CreatePostWithComment(ctx context.Context, title, content, comment string, tags []string) (*PostResponse, error) {
	return s.create(ctx, title, content, tags, &comment)
}

func (s *service) ListPosts(ctx context.Context) ([]PostResponse, error) {
	posts, err := s.repo.ListPosts(ctx)
	if err != nil {
		return nil, fmt.Errorf("list posts: %w", err)
	}
	responses := make([]PostResponse, 0, len(posts))
	for _, post := range posts {
		responses = append(responses, toResponse(post))
	}
	return responses, nil
}

func (s *service) DeletePost(ctx context.Context, id uint) error {
	err := s.repo.WithinTransaction(ctx, func(txRepo repository.PostRepository) error {
		post, err := txRepo.FindPost(ctx, id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		if err != nil {
			return err
		}
		if err := txRepo.DeleteComments(ctx, id); err != nil {
			return err
		}
		if err := txRepo.ClearTags(ctx, post); err != nil {
			return err
		}
		return txRepo.DeletePost(ctx, post)
	})
	if err != nil {
		return fmt.Errorf("delete post transaction: %w", err)
	}
	return nil
}

func toResponse(post model.Post) PostResponse {
	response := PostResponse{ID: post.ID, Title: post.Title, Content: post.Content, CreatedAt: post.CreatedAt, Tags: []TagResponse{}, Comments: []CommentResponse{}}
	for _, tag := range post.Tags {
		response.Tags = append(response.Tags, TagResponse{ID: tag.ID, Name: tag.Name})
	}
	for _, comment := range post.Comments {
		response.Comments = append(response.Comments, CommentResponse{ID: comment.ID, Content: comment.Content})
	}
	return response
}
