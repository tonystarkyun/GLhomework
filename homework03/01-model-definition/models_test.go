package main

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite failed: %v", err)
	}

	if err := AutoMigrateBlog(db); err != nil {
		t.Fatalf("auto migrate failed: %v", err)
	}

	return db
}

func TestAutoMigrateCreatesTables(t *testing.T) {
	db := newTestDB(t)

	if !db.Migrator().HasTable(&User{}) {
		t.Fatalf("users table was not created")
	}
	if !db.Migrator().HasTable(&Post{}) {
		t.Fatalf("posts table was not created")
	}
	if !db.Migrator().HasTable(&Comment{}) {
		t.Fatalf("comments table was not created")
	}
}

func TestRelationsWorkWithPreload(t *testing.T) {
	db := newTestDB(t)

	user := User{Name: "Alice", Email: "alice@example.com"}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("create user failed: %v", err)
	}

	post := Post{UserID: user.ID, Title: "GORM 入门", Content: "content"}
	if err := db.Create(&post).Error; err != nil {
		t.Fatalf("create post failed: %v", err)
	}

	comment := Comment{PostID: post.ID, Content: "写得很好"}
	if err := db.Create(&comment).Error; err != nil {
		t.Fatalf("create comment failed: %v", err)
	}

	var got User
	if err := db.Preload("Posts.Comments").First(&got, user.ID).Error; err != nil {
		t.Fatalf("query user failed: %v", err)
	}

	if len(got.Posts) != 1 {
		t.Fatalf("expected 1 post, got %d", len(got.Posts))
	}
	if len(got.Posts[0].Comments) != 1 {
		t.Fatalf("expected 1 comment, got %d", len(got.Posts[0].Comments))
	}
}
