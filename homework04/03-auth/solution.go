package q03auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// APIResult captures HTTP status and response body.
type APIResult struct {
	StatusCode int
	Body       []byte
}

func doJSON(router http.Handler, method, path string, payload any, token string) APIResult {
	var body io.Reader
	if payload != nil {
		buf, _ := json.Marshal(payload)
		body = bytes.NewReader(buf)
	}

	req, _ := http.NewRequest(method, path, body)
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	w := &responseRecorder{header: make(http.Header)}
	router.ServeHTTP(w, req)
	return APIResult{StatusCode: w.code, Body: w.body.Bytes()}
}

// RegisterUser calls the register endpoint.
func RegisterUser(router http.Handler, username, password, email string) APIResult {
	return doJSON(router, http.MethodPost, "/register", map[string]string{
		"username": username,
		"password": password,
		"email":    email,
	}, "")
}

// LoginUser calls the login endpoint and returns response with JWT payload.
func LoginUser(router http.Handler, username, password string) APIResult {
	return doJSON(router, http.MethodPost, "/login", map[string]string{
		"username": username,
		"password": password,
	}, "")
}

// CreatePostWithToken calls protected create-post API.
func CreatePostWithToken(router http.Handler, token, title, content string) APIResult {
	return doJSON(router, http.MethodPost, "/posts", map[string]string{
		"title":   title,
		"content": content,
	}, token)
}

type responseRecorder struct {
	header http.Header
	body   bytes.Buffer
	code   int
}

func (r *responseRecorder) Header() http.Header {
	return r.header
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	if r.code == 0 {
		r.code = http.StatusOK
	}
	return r.body.Write(b)
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.code = statusCode
}
