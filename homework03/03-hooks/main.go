package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("blog_q3.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("open database failed: %v", err)
	}

	if err := AutoMigrateBlog(db); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	user := User{Name: "HookUser", Email: "hook-user@example.com"}
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("create user failed: %v", err)
	}

	post := Post{UserID: user.ID, Title: "Hook Post", Content: "demo"}
	if err := db.Create(&post).Error; err != nil {
		log.Fatalf("create post failed: %v", err)
	}

	var refreshedUser User
	if err := db.First(&refreshedUser, user.ID).Error; err != nil {
		log.Fatalf("query user failed: %v", err)
	}
	fmt.Printf("Q3 hook1: user post_count = %d\n", refreshedUser.PostCount)

	comment := Comment{PostID: post.ID, Content: "first comment"}
	if err := db.Create(&comment).Error; err != nil {
		log.Fatalf("create comment failed: %v", err)
	}
	if err := db.Delete(&comment).Error; err != nil {
		log.Fatalf("delete comment failed: %v", err)
	}

	var refreshedPost Post
	if err := db.First(&refreshedPost, post.ID).Error; err != nil {
		log.Fatalf("query post failed: %v", err)
	}
	fmt.Printf("Q3 hook2: post comment_status = %s\n", refreshedPost.CommentStatus)
}
