package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:120;uniqueIndex;not null"`
	PostCount int    `gorm:"not null;default:0"`
	Posts     []Post `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"index;not null"`
	Title         string    `gorm:"size:200;not null"`
	Content       string    `gorm:"type:text"`
	CommentStatus string    `gorm:"size:20;not null;default:无评论"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
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

// AfterCreate: 文章创建成功后，用户文章数量 +1。
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// AfterCreate: 评论创建后，文章状态设为“有评论”。
func (c *Comment) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&Post{}).
		Where("id = ?", c.PostID).
		Update("comment_status", "有评论").Error
}

// AfterDelete: 删除评论后若无剩余评论，文章状态设为“无评论”。
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var cnt int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&cnt).Error; err != nil {
		return err
	}

	if cnt == 0 {
		return tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}

	return nil
}
