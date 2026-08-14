//go:build integration

package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/model"
	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/repository"
	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/service"
	"iotestgo/module04_gorm/internal/classroomdb"

	"gorm.io/gorm"
)

func TestBlogLifecycleMySQL(t *testing.T) {
	db, err := classroomdb.Open()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = classroomdb.Close(db) })
	if err := db.AutoMigrate(&model.Post{}, &model.Comment{}, &model.Tag{}); err != nil {
		t.Fatal(err)
	}
	suffix := time.Now().UnixNano()
	titleOne := fmt.Sprintf("Integration Post One %d", suffix)
	titleTwo := fmt.Sprintf("Integration Post Two %d", suffix)
	sharedTag := fmt.Sprintf("shared-%d", suffix)
	t.Cleanup(func() {
		var postIDs []uint
		db.Unscoped().Model(&model.Post{}).Where("title IN ?", []string{titleOne, titleTwo}).Pluck("id", &postIDs)
		if len(postIDs) > 0 {
			db.Exec("DELETE FROM m04_blog_post_tags WHERE post_id IN ?", postIDs)
		}
		db.Unscoped().Where("title IN ?", []string{titleOne, titleTwo}).Delete(&model.Post{})
		db.Unscoped().Where("content = ?", fmt.Sprintf("Comment %d", suffix)).Delete(&model.Comment{})
		db.Unscoped().Where("name IN ?", []string{sharedTag, fmt.Sprintf("go-%d", suffix)}).Delete(&model.Tag{})
	})
	ctx := context.Background()
	svc := service.New(repository.New(db))
	first, err := svc.CreatePostWithComment(ctx, titleOne, "Content", fmt.Sprintf("Comment %d", suffix), []string{sharedTag, fmt.Sprintf("go-%d", suffix)})
	if err != nil {
		t.Fatal(err)
	}
	second, err := svc.CreatePost(ctx, titleTwo, "Content", []string{sharedTag})
	if err != nil {
		t.Fatal(err)
	}
	queryCount := 0
	if err := db.Callback().Query().Before("gorm:query").Register("m04_blog:count_queries", func(*gorm.DB) { queryCount++ }); err != nil {
		t.Fatal(err)
	}
	posts, err := svc.ListPosts(ctx)
	if err != nil || len(posts) < 2 {
		t.Fatalf("posts=%d err=%v", len(posts), err)
	}
	if queryCount != 4 {
		t.Fatalf("ListPosts executed %d queries, want 4 (posts, comments, join rows, tags)", queryCount)
	}
	if err := svc.DeletePost(ctx, first.ID); err != nil {
		t.Fatal(err)
	}
	var sharedCount int64
	if err := db.Model(&model.Tag{}).Where("name = ?", sharedTag).Count(&sharedCount).Error; err != nil || sharedCount != 1 {
		t.Fatalf("shared tag count=%d err=%v", sharedCount, err)
	}
	if err := svc.DeletePost(ctx, second.ID); err != nil {
		t.Fatal(err)
	}
}
