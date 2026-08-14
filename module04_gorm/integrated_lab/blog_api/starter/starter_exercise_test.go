//go:build exercise

package starter

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

type transactionSpy struct{ called bool }

func (s *transactionSpy) WithinTransaction(fn func() error) error {
	s.called = true
	return fn()
}

func TestNormalizeTagsExercise(t *testing.T) {
	want := []string{"go", "gorm"}
	if got := NormalizeTags([]string{" Go ", "go", "", "GORM"}); !reflect.DeepEqual(got, want) {
		t.Fatalf("NormalizeTags()=%v, want %v", got, want)
	}
}

func TestModelRelationsExercise(t *testing.T) {
	typeOfPost := reflect.TypeOf(Post{})
	comments, ok := typeOfPost.FieldByName("Comments")
	if !ok || !strings.Contains(comments.Tag.Get("gorm"), "foreignKey") {
		t.Fatal("Post.Comments must declare the one-to-many relation")
	}
	tags, ok := typeOfPost.FieldByName("Tags")
	if !ok || !strings.Contains(tags.Tag.Get("gorm"), "many2many:m04_blog_post_tags") {
		t.Fatal("Post.Tags must declare the m04_blog_post_tags relation")
	}
}

func TestCreateUsesTransactionExercise(t *testing.T) {
	spy := &transactionSpy{}
	err := CreatePostWithComment(spy, "Title", "Content", "Comment", []string{"go"})
	if errors.Is(err, ErrNotImplemented) || !spy.called {
		t.Fatalf("implement transaction: called=%v err=%v", spy.called, err)
	}
}

func TestDeleteUsesTransactionExercise(t *testing.T) {
	spy := &transactionSpy{}
	err := DeletePost(spy, 1)
	if errors.Is(err, ErrNotImplemented) || !spy.called {
		t.Fatalf("implement delete transaction: called=%v err=%v", spy.called, err)
	}
}
