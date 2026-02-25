package main

import (
	"fmt"
	"strings"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", strings.ReplaceAll(t.Name(), "/", "_"))
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite failed: %v", err)
	}

	if err := AutoMigrateBlog(db); err != nil {
		t.Fatalf("auto migrate failed: %v", err)
	}

	return db
}

func seedForQueryTest(t *testing.T, db *gorm.DB) (uint, uint, uint) {
	t.Helper()

	user1 := User{Name: "U1", Email: "u1@example.com"}
	user2 := User{Name: "U2", Email: "u2@example.com"}
	if err := db.Create(&user1).Error; err != nil {
		t.Fatalf("create user1 failed: %v", err)
	}
	if err := db.Create(&user2).Error; err != nil {
		t.Fatalf("create user2 failed: %v", err)
	}

	p1 := Post{UserID: user1.ID, Title: "p1"}
	p2 := Post{UserID: user1.ID, Title: "p2"}
	p3 := Post{UserID: user2.ID, Title: "p3"}
	if err := db.Create(&p1).Error; err != nil {
		t.Fatalf("create p1 failed: %v", err)
	}
	if err := db.Create(&p2).Error; err != nil {
		t.Fatalf("create p2 failed: %v", err)
	}
	if err := db.Create(&p3).Error; err != nil {
		t.Fatalf("create p3 failed: %v", err)
	}

	comments := []Comment{
		{PostID: p1.ID, Content: "p1-c1"},
		{PostID: p1.ID, Content: "p1-c2"},
		{PostID: p2.ID, Content: "p2-c1"},
		{PostID: p3.ID, Content: "p3-c1"},
		{PostID: p3.ID, Content: "p3-c2"},
		{PostID: p3.ID, Content: "p3-c3"},
	}
	if err := db.Create(&comments).Error; err != nil {
		t.Fatalf("create comments failed: %v", err)
	}

	return user1.ID, p1.ID, p3.ID
}

func TestGetUserPostsWithComments(t *testing.T) {
	db := newTestDB(t)
	userID, _, _ := seedForQueryTest(t, db)

	posts, err := GetUserPostsWithComments(db, userID)
	if err != nil {
		t.Fatalf("query failed: %v", err)
	}

	if len(posts) != 2 {
		t.Fatalf("expected 2 posts, got %d", len(posts))
	}

	if len(posts[0].Comments) != 2 {
		t.Fatalf("expected first post has 2 comments, got %d", len(posts[0].Comments))
	}
	if len(posts[1].Comments) != 1 {
		t.Fatalf("expected second post has 1 comment, got %d", len(posts[1].Comments))
	}
}

func TestGetMostCommentedPost(t *testing.T) {
	db := newTestDB(t)
	_, _, expectedPostID := seedForQueryTest(t, db)

	got, err := GetMostCommentedPost(db)
	if err != nil {
		t.Fatalf("query failed: %v", err)
	}
	if got == nil {
		t.Fatalf("expected non-nil result")
	}
	if got.Post.ID != expectedPostID {
		t.Fatalf("expected post id %d, got %d", expectedPostID, got.Post.ID)
	}
	if got.CommentCount != 3 {
		t.Fatalf("expected comment count 3, got %d", got.CommentCount)
	}
}

func TestGetMostCommentedPostWhenNoComments(t *testing.T) {
	db := newTestDB(t)

	user := User{Name: "NoCommentUser", Email: "no-comment@example.com"}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("create user failed: %v", err)
	}
	post := Post{UserID: user.ID, Title: "empty"}
	if err := db.Create(&post).Error; err != nil {
		t.Fatalf("create post failed: %v", err)
	}

	got, err := GetMostCommentedPost(db)
	if err != nil {
		t.Fatalf("query failed: %v", err)
	}
	if got != nil {
		t.Fatalf("expected nil result when no comments")
	}
}
