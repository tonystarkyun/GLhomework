package main

import "gorm.io/gorm"

// GetUserPostsWithComments 查询某个用户的所有文章及其评论信息。
func GetUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Where("user_id = ?", userID).
		Preload("Comments").
		Order("id ASC").
		Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

type MostCommentedPost struct {
	Post         Post
	CommentCount int64
}

// GetMostCommentedPost 查询评论数量最多的文章。
func GetMostCommentedPost(db *gorm.DB) (*MostCommentedPost, error) {
	var row struct {
		PostID       uint
		CommentCount int64
	}

	err := db.Model(&Comment{}).
		Select("post_id, COUNT(*) AS comment_count").
		Group("post_id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&row).Error
	if err != nil {
		return nil, err
	}

	if row.PostID == 0 {
		return nil, nil
	}

	var post Post
	if err := db.Preload("Comments").First(&post, row.PostID).Error; err != nil {
		return nil, err
	}

	return &MostCommentedPost{Post: post, CommentCount: row.CommentCount}, nil
}
