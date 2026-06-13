package service

import (
	"regexp"
	"testing"

	"iotestgo/module04_gorm/project_blog_api/internal/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreatePostWithComment(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock: %v", err)
	}
	defer sqlDB.Close()

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("open gorm db: %v", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `posts` (`created_at`,`updated_at`,`deleted_at`,`title`,`content`) VALUES (?,?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "gRPC 入门", "streaming rpc").
		WillReturnResult(sqlmock.NewResult(42, 1))
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `comments` (`created_at`,`updated_at`,`deleted_at`,`post_id`,`content`) VALUES (?,?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, uint(42), "first comment").
		WillReturnResult(sqlmock.NewResult(7, 1))
	mock.ExpectCommit()

	svc := NewPostService(repository.NewPostRepository(db))
	post, err := svc.CreatePostWithComment("gRPC 入门", "streaming rpc", "first comment")
	if err != nil {
		t.Fatalf("CreatePostWithComment returned error: %v", err)
	}
	if post.ID != 42 {
		t.Fatalf("expected post ID 42, got %d", post.ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}
