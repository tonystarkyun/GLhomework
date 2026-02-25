package testkit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"homework04/internal/blog"
)

// NewTestApp creates an isolated in-memory app for tests.
func NewTestApp(t *testing.T, writer io.Writer) *blog.App {
	t.Helper()
	if writer == nil {
		writer = io.Discard
	}
	secret := "test-secret"
	dsn := fmt.Sprintf("file:blog_test_%d?mode=memory&cache=shared", time.Now().UnixNano())

	app, err := blog.NewSQLiteApp(blog.Config{
		DSN:       dsn,
		JWTSecret: secret,
		Logger:    log.New(writer, "", 0),
	})
	if err != nil {
		t.Fatalf("new app: %v", err)
	}
	return app
}

// JSONRequest issues an HTTP request with optional JSON body and bearer token.
func JSONRequest(t *testing.T, handler http.Handler, method, path string, payload any, token string) *httptest.ResponseRecorder {
	t.Helper()

	var body io.Reader
	if payload != nil {
		buf, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("marshal payload: %v", err)
		}
		body = bytes.NewReader(buf)
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
	return w
}

// DecodeBody decodes JSON body into out.
func DecodeBody(t *testing.T, w *httptest.ResponseRecorder, out any) {
	t.Helper()
	if err := json.Unmarshal(w.Body.Bytes(), out); err != nil {
		t.Fatalf("decode body: %v, body=%s", err, w.Body.String())
	}
}

// RegisterAndLogin returns a valid token.
func RegisterAndLogin(t *testing.T, app *blog.App, username, password, email string) string {
	t.Helper()

	reg := JSONRequest(t, app.Router, http.MethodPost, "/register", map[string]string{
		"username": username,
		"password": password,
		"email":    email,
	}, "")
	if reg.Code != http.StatusCreated {
		t.Fatalf("register failed: code=%d body=%s", reg.Code, reg.Body.String())
	}

	login := JSONRequest(t, app.Router, http.MethodPost, "/login", map[string]string{
		"username": username,
		"password": password,
	}, "")
	if login.Code != http.StatusOK {
		t.Fatalf("login failed: code=%d body=%s", login.Code, login.Body.String())
	}

	var resp struct {
		Token string `json:"token"`
	}
	DecodeBody(t, login, &resp)
	if resp.Token == "" {
		t.Fatalf("empty token in response")
	}
	return resp.Token
}
