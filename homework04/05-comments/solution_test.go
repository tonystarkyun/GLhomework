package q05comments

import (
	"encoding/json"
	"net/http"
	"testing"

	"homework04/internal/testkit"
)

func TestComments_CreateAndList(t *testing.T) {
	app := testkit.NewTestApp(t, nil)
	authorToken := testkit.RegisterAndLogin(t, app, "author", "password123", "author@example.com")
	commenterToken := testkit.RegisterAndLogin(t, app, "commenter", "password123", "commenter@example.com")

	postResp := testkit.JSONRequest(t, app.Router, http.MethodPost, "/posts", map[string]string{
		"title":   "post",
		"content": "body",
	}, authorToken)
	if postResp.Code != http.StatusCreated {
		t.Fatalf("create post failed: code=%d body=%s", postResp.Code, postResp.Body.String())
	}

	var post struct {
		ID int `json:"id"`
	}
	testkit.DecodeBody(t, postResp, &post)

	createdComment := CreateComment(app.Router, commenterToken, post.ID, "nice article")
	if createdComment.Code != http.StatusCreated {
		t.Fatalf("create comment failed: code=%d body=%s", createdComment.Code, string(createdComment.Body))
	}

	list := ListComments(app.Router, post.ID)
	if list.Code != http.StatusOK {
		t.Fatalf("list comments failed: code=%d body=%s", list.Code, string(list.Body))
	}

	var comments []struct {
		Content  string `json:"content"`
		Username string `json:"username"`
	}
	if err := json.Unmarshal(list.Body, &comments); err != nil {
		t.Fatalf("decode comments: %v", err)
	}
	if len(comments) != 1 {
		t.Fatalf("expected 1 comment, got=%d", len(comments))
	}
	if comments[0].Content != "nice article" || comments[0].Username != "commenter" {
		t.Fatalf("unexpected comment payload: %+v", comments[0])
	}
}

func TestComments_AuthAndNotFound(t *testing.T) {
	app := testkit.NewTestApp(t, nil)
	token := testkit.RegisterAndLogin(t, app, "user1", "password123", "u1@example.com")

	unauth := CreateComment(app.Router, "", 1, "x")
	if unauth.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 without token, got=%d body=%s", unauth.Code, string(unauth.Body))
	}

	notFound := CreateComment(app.Router, token, 9999, "x")
	if notFound.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for missing post, got=%d body=%s", notFound.Code, string(notFound.Body))
	}
}
