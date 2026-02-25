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

func TestPostAfterCreateUpdatesUserPostCount(t *testing.T) {
	db := newTestDB(t)

	user := User{Name: "Tom", Email: "tom@example.com"}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("create user failed: %v", err)
	}

	posts := []Post{
		{UserID: user.ID, Title: "p1"},
		{UserID: user.ID, Title: "p2"},
	}
	for _, p := range posts {
		if err := db.Create(&p).Error; err != nil {
			t.Fatalf("create post failed: %v", err)
		}
	}

	var got User
	if err := db.First(&got, user.ID).Error; err != nil {
		t.Fatalf("query user failed: %v", err)
	}
	if got.PostCount != 2 {
		t.Fatalf("expected post_count 2, got %d", got.PostCount)
	}
}

func TestCommentAfterDeleteUpdatesCommentStatus(t *testing.T) {
	db := newTestDB(t)

	user := User{Name: "Jerry", Email: "jerry@example.com"}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("create user failed: %v", err)
	}

	post := Post{UserID: user.ID, Title: "hook post"}
	if err := db.Create(&post).Error; err != nil {
		t.Fatalf("create post failed: %v", err)
	}

	c1 := Comment{PostID: post.ID, Content: "c1"}
	c2 := Comment{PostID: post.ID, Content: "c2"}
	if err := db.Create(&c1).Error; err != nil {
		t.Fatalf("create c1 failed: %v", err)
	}
	if err := db.Create(&c2).Error; err != nil {
		t.Fatalf("create c2 failed: %v", err)
	}

	var gotPost Post
	if err := db.First(&gotPost, post.ID).Error; err != nil {
		t.Fatalf("query post failed: %v", err)
	}
	if gotPost.CommentStatus != "有评论" {
		t.Fatalf("expected status 有评论 after comment created, got %s", gotPost.CommentStatus)
	}

	if err := db.Delete(&c1).Error; err != nil {
		t.Fatalf("delete c1 failed: %v", err)
	}
	if err := db.First(&gotPost, post.ID).Error; err != nil {
		t.Fatalf("query post failed: %v", err)
	}
	if gotPost.CommentStatus != "有评论" {
		t.Fatalf("expected status keeps 有评论 when comments remain, got %s", gotPost.CommentStatus)
	}

	if err := db.Delete(&c2).Error; err != nil {
		t.Fatalf("delete c2 failed: %v", err)
	}
	if err := db.First(&gotPost, post.ID).Error; err != nil {
		t.Fatalf("query post failed: %v", err)
	}
	if gotPost.CommentStatus != "无评论" {
		t.Fatalf("expected status 无评论 after deleting last comment, got %s", gotPost.CommentStatus)
	}
}
