package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("blog_q2.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("open database failed: %v", err)
	}

	if err := AutoMigrateBlog(db); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	seedUserID, err := seedData(db)
	if err != nil {
		log.Fatalf("seed data failed: %v", err)
	}

	posts, err := GetUserPostsWithComments(db, seedUserID)
	if err != nil {
		log.Fatalf("query user posts failed: %v", err)
	}
	fmt.Printf("Q2 query1: user %d has %d posts\n", seedUserID, len(posts))

	most, err := GetMostCommentedPost(db)
	if err != nil {
		log.Fatalf("query most commented post failed: %v", err)
	}
	if most == nil {
		fmt.Println("Q2 query2: no comments in database")
		return
	}

	fmt.Printf("Q2 query2: post %d (%s) has %d comments\n", most.Post.ID, most.Post.Title, most.CommentCount)
}

func seedData(db *gorm.DB) (uint, error) {
	user1 := User{Name: "Alice", Email: "alice-q2@example.com"}
	user2 := User{Name: "Bob", Email: "bob-q2@example.com"}
	if err := db.Create(&user1).Error; err != nil {
		return 0, err
	}
	if err := db.Create(&user2).Error; err != nil {
		return 0, err
	}

	post1 := Post{UserID: user1.ID, Title: "Alice-Post-1", Content: "..."}
	post2 := Post{UserID: user1.ID, Title: "Alice-Post-2", Content: "..."}
	post3 := Post{UserID: user2.ID, Title: "Bob-Post-1", Content: "..."}
	if err := db.Create(&post1).Error; err != nil {
		return 0, err
	}
	if err := db.Create(&post2).Error; err != nil {
		return 0, err
	}
	if err := db.Create(&post3).Error; err != nil {
		return 0, err
	}

	comments := []Comment{
		{PostID: post1.ID, Content: "c1"},
		{PostID: post1.ID, Content: "c2"},
		{PostID: post2.ID, Content: "c3"},
		{PostID: post3.ID, Content: "c4"},
		{PostID: post3.ID, Content: "c5"},
		{PostID: post3.ID, Content: "c6"},
	}
	if err := db.Create(&comments).Error; err != nil {
		return 0, err
	}

	return user1.ID, nil
}
