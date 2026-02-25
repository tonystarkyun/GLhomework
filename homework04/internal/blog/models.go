package blog

import "time"

// User is the account model.
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"size:64;not null;uniqueIndex" json:"username"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Email        string    `gorm:"size:255;not null;uniqueIndex" json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Post is the blog post model.
type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	User      User      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Comment is the post comment model.
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	User      User      `json:"-"`
	PostID    uint      `gorm:"not null;index" json:"post_id"`
	Post      Post      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PostResponse is the API response shape for posts.
type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CommentResponse is the API response shape for comments.
type CommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	PostID    uint      `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

func toPostResponse(p Post) PostResponse {
	return PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserID,
		Username:  p.User.Username,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func toCommentResponse(c Comment) CommentResponse {
	return CommentResponse{
		ID:        c.ID,
		Content:   c.Content,
		UserID:    c.UserID,
		Username:  c.User.Username,
		PostID:    c.PostID,
		CreatedAt: c.CreatedAt,
	}
}
