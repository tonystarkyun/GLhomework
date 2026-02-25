package q04posts

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
)

// Result wraps HTTP response status/body for post APIs.
type Result struct {
	Code int
	Body []byte
}

func requestJSON(handler http.Handler, method, path string, payload any, token string) Result {
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

func CreatePost(handler http.Handler, token, title, content string) Result {
	return requestJSON(handler, http.MethodPost, "/posts", map[string]string{"title": title, "content": content}, token)
}

func ListPosts(handler http.Handler) Result {
	return requestJSON(handler, http.MethodGet, "/posts", nil, "")
}

func GetPost(handler http.Handler, id int) Result {
	return requestJSON(handler, http.MethodGet, "/posts/"+itoa(id), nil, "")
}

func UpdatePost(handler http.Handler, token string, id int, title, content string) Result {
	return requestJSON(handler, http.MethodPut, "/posts/"+itoa(id), map[string]string{"title": title, "content": content}, token)
}

func DeletePost(handler http.Handler, token string, id int) Result {
	return requestJSON(handler, http.MethodDelete, "/posts/"+itoa(id), nil, token)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	buf := [20]byte{}
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}
