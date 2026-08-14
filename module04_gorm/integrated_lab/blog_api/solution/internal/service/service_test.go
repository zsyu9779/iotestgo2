package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/model"
	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/repository"

	"gorm.io/gorm"
)

type fakeRepo struct {
	posts        []model.Post
	steps        []string
	failAt       string
	transaction  bool
	nextID       uint
	findNotFound bool
}

func (f *fakeRepo) fail(step string) error {
	f.steps = append(f.steps, step)
	if f.failAt == step {
		return errors.New("injected " + step)
	}
	return nil
}

func (f *fakeRepo) WithinTransaction(_ context.Context, fn func(repository.PostRepository) error) error {
	f.transaction = true
	if err := f.fail("begin"); err != nil {
		return err
	}
	return fn(f)
}
func (f *fakeRepo) CreatePost(_ context.Context, post *model.Post) error {
	if err := f.fail("create-post"); err != nil {
		return err
	}
	f.nextID++
	post.ID = f.nextID
	return nil
}
func (f *fakeRepo) FindOrCreateTag(_ context.Context, name string) (*model.Tag, error) {
	if err := f.fail("tag-" + name); err != nil {
		return nil, err
	}
	f.nextID++
	return &model.Tag{ID: f.nextID, Name: name}, nil
}
func (f *fakeRepo) ReplaceTags(context.Context, *model.Post, []model.Tag) error {
	return f.fail("replace-tags")
}
func (f *fakeRepo) CreateComment(_ context.Context, comment *model.Comment) error {
	if err := f.fail("create-comment"); err != nil {
		return err
	}
	f.nextID++
	comment.ID = f.nextID
	return nil
}
func (f *fakeRepo) ListPosts(context.Context) ([]model.Post, error) { return f.posts, f.fail("list") }
func (f *fakeRepo) FindPost(context.Context, uint) (*model.Post, error) {
	if f.findNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	if err := f.fail("find-post"); err != nil {
		return nil, err
	}
	return &model.Post{ID: 1}, nil
}
func (f *fakeRepo) DeleteComments(context.Context, uint) error    { return f.fail("delete-comments") }
func (f *fakeRepo) ClearTags(context.Context, *model.Post) error  { return f.fail("clear-tags") }
func (f *fakeRepo) DeletePost(context.Context, *model.Post) error { return f.fail("delete-post") }

func TestNormalizeTags(t *testing.T) {
	got := NormalizeTags([]string{" Go ", "go", "", " GORM ", "gorm"})
	want := []string{"go", "gorm"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("NormalizeTags() = %v, want %v", got, want)
	}
}

func TestCreateWithCommentUsesTransaction(t *testing.T) {
	repo := &fakeRepo{}
	created, err := New(repo).CreatePostWithComment(context.Background(), " Title ", " Content ", " First ", []string{"Go", " go "})
	if err != nil {
		t.Fatal(err)
	}
	if !repo.transaction || created.Title != "Title" || len(created.Tags) != 1 || len(created.Comments) != 1 {
		t.Fatalf("unexpected result: transaction=%v created=%#v", repo.transaction, created)
	}
}

func TestCreateSecondStepFailureReturnsError(t *testing.T) {
	repo := &fakeRepo{failAt: "create-comment"}
	_, err := New(repo).CreatePostWithComment(context.Background(), "Title", "Content", "Comment", nil)
	if err == nil || !repo.transaction {
		t.Fatalf("expected transactional failure, got %v", err)
	}
}

func TestDeleteOrderAndNotFound(t *testing.T) {
	repo := &fakeRepo{}
	if err := New(repo).DeletePost(context.Background(), 1); err != nil {
		t.Fatal(err)
	}
	want := []string{"begin", "find-post", "delete-comments", "clear-tags", "delete-post"}
	if !reflect.DeepEqual(repo.steps, want) {
		t.Fatalf("delete steps = %v, want %v", repo.steps, want)
	}
	repo = &fakeRepo{findNotFound: true}
	if err := New(repo).DeletePost(context.Background(), 99); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestRejectsWhitespace(t *testing.T) {
	_, err := New(&fakeRepo{}).CreatePost(context.Background(), "   ", "content", nil)
	if !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}
