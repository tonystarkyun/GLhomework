package q04posts

import (
	"encoding/json"
	"net/http"
	"testing"

	"homework04/internal/testkit"
)

func TestPosts_CRUDAndAuthorPermission(t *testing.T) {
	app := testkit.NewTestApp(t, nil)
	tokenA := testkit.RegisterAndLogin(t, app, "author", "password123", "author@example.com")
	tokenB := testkit.RegisterAndLogin(t, app, "other", "password123", "other@example.com")

	created := CreatePost(app.Router, tokenA, "First", "Content")
	if created.Code != http.StatusCreated {
		t.Fatalf("create post failed: code=%d body=%s", created.Code, string(created.Body))
	}

	var createdBody struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(created.Body, &createdBody); err != nil {
		t.Fatalf("decode create body: %v", err)
	}
	if createdBody.ID == 0 {
		t.Fatalf("created post id is empty")
	}

	list := ListPosts(app.Router)
	if list.Code != http.StatusOK {
		t.Fatalf("list post failed: code=%d body=%s", list.Code, string(list.Body))
	}

	get := GetPost(app.Router, createdBody.ID)
	if get.Code != http.StatusOK {
		t.Fatalf("get post failed: code=%d body=%s", get.Code, string(get.Body))
	}

	forbiddenUpdate := UpdatePost(app.Router, tokenB, createdBody.ID, "hijack", "x")
	if forbiddenUpdate.Code != http.StatusForbidden {
		t.Fatalf("expected 403 for non-author update, got=%d body=%s", forbiddenUpdate.Code, string(forbiddenUpdate.Body))
	}

	okUpdate := UpdatePost(app.Router, tokenA, createdBody.ID, "updated", "updated content")
	if okUpdate.Code != http.StatusOK {
		t.Fatalf("author update failed: code=%d body=%s", okUpdate.Code, string(okUpdate.Body))
	}

	forbiddenDelete := DeletePost(app.Router, tokenB, createdBody.ID)
	if forbiddenDelete.Code != http.StatusForbidden {
		t.Fatalf("expected 403 for non-author delete, got=%d body=%s", forbiddenDelete.Code, string(forbiddenDelete.Body))
	}

	okDelete := DeletePost(app.Router, tokenA, createdBody.ID)
	if okDelete.Code != http.StatusNoContent {
		t.Fatalf("author delete failed: code=%d body=%s", okDelete.Code, string(okDelete.Body))
	}

	notFound := GetPost(app.Router, createdBody.ID)
	if notFound.Code != http.StatusNotFound {
		t.Fatalf("expected 404 after delete, got=%d body=%s", notFound.Code, string(notFound.Body))
	}
}
