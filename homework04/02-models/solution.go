package q02models

import "homework04/internal/blog"

// SeedResult groups created records for Q2 model verification.
type SeedResult struct {
	User    blog.User
	Post    blog.Post
	Comment blog.Comment
}

// SeedCoreModels creates user/post/comment records with proper relations.
func SeedCoreModels(app *blog.App, username, email, passwordHash string) (SeedResult, error) {
	u := blog.User{Username: username, Email: email, PasswordHash: passwordHash}
	if err := app.DB.Create(&u).Error; err != nil {
		return SeedResult{}, err
	}

	p := blog.Post{Title: "first post", Content: "hello", UserID: u.ID}
	if err := app.DB.Create(&p).Error; err != nil {
		return SeedResult{}, err
	}

	c := blog.Comment{Content: "nice post", UserID: u.ID, PostID: p.ID}
	if err := app.DB.Create(&c).Error; err != nil {
		return SeedResult{}, err
	}

	return SeedResult{User: u, Post: p, Comment: c}, nil
}
