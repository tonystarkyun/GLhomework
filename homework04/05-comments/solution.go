package q05comments

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
)

// Result wraps comment API responses.
type Result struct {
	Code int
	Body []byte
}

func do(handler http.Handler, method, path string, payload any, token string) Result {
	var body io.Reader
	if payload != nil {
		b, _ := json.Marshal(payload)
		body = bytes.NewReader(b)
	}

	req := httptest.NewRequest(method, path, body)
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return Result{Code: w.Code, Body: w.Body.Bytes()}
}

// CreateComment creates a comment on a post by ID.
func CreateComment(handler http.Handler, token string, postID int, content string) Result {
	return do(handler, http.MethodPost, "/posts/"+strconv.Itoa(postID)+"/comments", map[string]string{"content": content}, token)
}

// ListComments lists comments on a post by ID.
func ListComments(handler http.Handler, postID int) Result {
	return do(handler, http.MethodGet, "/posts/"+strconv.Itoa(postID)+"/comments", nil, "")
}
