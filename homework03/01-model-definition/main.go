package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("blog_q1.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("open database failed: %v", err)
	}

	if err := AutoMigrateBlog(db); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	fmt.Println("Q1 done: tables for users/posts/comments are created.")
}
