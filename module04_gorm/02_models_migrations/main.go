package main

import (
	"fmt"
	"os"
	"time"

	"iotestgo/module04_gorm/internal/classroomdb"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"size:80;uniqueIndex"`
	Posts     []Post
}

func (Author) TableName() string { return "m04_l02_authors" }

type Tag struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:40;uniqueIndex"`
	Posts []Post `gorm:"many2many:m04_l02_post_tags"`
}

func (Tag) TableName() string { return "m04_l02_tags" }

type Post struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         `gorm:"size:160"`
	AuthorID  uint
	Author    Author
	Tags      []Tag `gorm:"many2many:m04_l02_post_tags"`
}

func (Post) TableName() string { return "m04_l02_posts" }

type ProfileV1 struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:80"`
}

func (ProfileV1) TableName() string { return "m04_l02_profiles" }

type ProfileV2 struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:80"`
	Email string `gorm:"size:160"`
}

func (ProfileV2) TableName() string { return "m04_l02_profiles" }

func run() error {
	db, err := classroomdb.Open()
	if err != nil {
		return err
	}
	defer classroomdb.Close(db)
	if err := db.AutoMigrate(&Author{}, &Tag{}, &Post{}, &ProfileV1{}); err != nil {
		return fmt.Errorf("migrate v1 models: %w", err)
	}
	if err := db.AutoMigrate(&ProfileV2{}); err != nil {
		return fmt.Errorf("migrate profile v2: %w", err)
	}
	if !db.Migrator().HasColumn(&ProfileV2{}, "Email") {
		return fmt.Errorf("migration contract failed: email column missing")
	}

	var author Author
	if err := db.Where("name = ?", "M04 Alice").FirstOrCreate(&author, Author{Name: "M04 Alice"}).Error; err != nil {
		return fmt.Errorf("seed author: %w", err)
	}
	tags := make([]Tag, 0, 2)
	for _, name := range []string{"go", "gorm"} {
		var tag Tag
		if err := db.Where("name = ?", name).FirstOrCreate(&tag, Tag{Name: name}).Error; err != nil {
			return fmt.Errorf("seed tag %s: %w", name, err)
		}
		tags = append(tags, tag)
	}
	post := Post{Title: fmt.Sprintf("Relations %d", time.Now().UnixNano()), AuthorID: author.ID, Tags: tags}
	if err := db.Create(&post).Error; err != nil {
		return fmt.Errorf("create related post: %w", err)
	}
	var loaded Post
	if err := db.Preload("Author").Preload("Tags").First(&loaded, post.ID).Error; err != nil {
		return fmt.Errorf("preload relations: %w", err)
	}
	if err := db.Delete(&loaded).Error; err != nil {
		return fmt.Errorf("soft delete post: %w", err)
	}
	var softDeleted Post
	if err := db.Unscoped().First(&softDeleted, loaded.ID).Error; err != nil {
		return fmt.Errorf("find soft-deleted post: %w", err)
	}
	if err := db.Unscoped().Model(&softDeleted).Association("Tags").Clear(); err != nil {
		return fmt.Errorf("clear tag associations before hard delete: %w", err)
	}
	if err := db.Unscoped().Delete(&softDeleted).Error; err != nil {
		return fmt.Errorf("hard delete post: %w", err)
	}
	fmt.Printf("relations=ok author=%s tags=%d migration=email soft_delete=ok hard_delete=ok\n", loaded.Author.Name, len(loaded.Tags))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
