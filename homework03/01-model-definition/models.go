package main

import (
	"time"

	"gorm.io/gorm"
)

// User 与 Post 是一对多关系。
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:120;uniqueIndex;not null"`
	Posts     []Post `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Post 与 Comment 是一对多关系。
type Post struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`
	Title     string    `gorm:"size:200;not null"`
	Content   string    `gorm:"type:text"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	PostID    uint   `gorm:"index;not null"`
	Content   string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func AutoMigrateBlog(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Post{}, &Comment{})
}
