package q02models

import (
	"testing"

	"homework04/internal/blog"
	"homework04/internal/testkit"
)

func TestSeedCoreModels_Relations(t *testing.T) {
	app := testkit.NewTestApp(t, nil)
	seed, err := SeedCoreModels(app, "alice", "alice@example.com", "hashed-password")
	if err != nil {
		t.Fatalf("seed models: %v", err)
	}

	if seed.Post.UserID != seed.User.ID {
		t.Fatalf("post.user_id mismatch: got=%d want=%d", seed.Post.UserID, seed.User.ID)
	}
	if seed.Comment.UserID != seed.User.ID || seed.Comment.PostID != seed.Post.ID {
		t.Fatalf("comment relation mismatch")
	}

	var got blog.Comment
	if err := app.DB.Preload("User").Preload("Post").First(&got, seed.Comment.ID).Error; err != nil {
		t.Fatalf("query with preload: %v", err)
	}
	if got.User.Username != "alice" {
		t.Fatalf("expected preloaded user alice, got=%s", got.User.Username)
	}
	if got.Post.Title != "first post" {
		t.Fatalf("expected preloaded post title first post, got=%s", got.Post.Title)
	}
}

func TestSeedCoreModels_UniqueUserConstraints(t *testing.T) {
	app := testkit.NewTestApp(t, nil)
	_, err := SeedCoreModels(app, "same", "same@example.com", "hash")
	if err != nil {
		t.Fatalf("first seed failed: %v", err)
	}
	_, err = SeedCoreModels(app, "same", "same@example.com", "hash")
	if err == nil {
		t.Fatalf("expected duplicate username/email to fail")
	}
}
