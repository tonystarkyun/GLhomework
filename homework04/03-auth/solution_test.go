package q03auth

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"homework04/internal/blog"
	"homework04/internal/testkit"
)

func TestAuth_RegisterLoginAndJWT(t *testing.T) {
	app := testkit.NewTestApp(t, nil)

	reg := RegisterUser(app.Router, "bob", "password123", "bob@example.com")
	if reg.StatusCode != http.StatusCreated {
		t.Fatalf("register status: got=%d body=%s", reg.StatusCode, string(reg.Body))
	}

	var stored blog.User
	if err := app.DB.Where("username = ?", "bob").First(&stored).Error; err != nil {
		t.Fatalf("load user: %v", err)
	}
	if stored.PasswordHash == "password123" || !strings.HasPrefix(stored.PasswordHash, "$2") {
		t.Fatalf("password is not bcrypt hashed: %s", stored.PasswordHash)
	}

	login := LoginUser(app.Router, "bob", "password123")
	if login.StatusCode != http.StatusOK {
		t.Fatalf("login status: got=%d body=%s", login.StatusCode, string(login.Body))
	}
	var loginResp struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(login.Body, &loginResp); err != nil {
		t.Fatalf("decode login body: %v", err)
	}
	if loginResp.Token == "" {
		t.Fatalf("expected token in login response")
	}

	unauthPost := CreatePostWithToken(app.Router, "", "t", "c")
	if unauthPost.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected unauthorized without token, got=%d body=%s", unauthPost.StatusCode, string(unauthPost.Body))
	}

	authPost := CreatePostWithToken(app.Router, loginResp.Token, "t", "c")
	if authPost.StatusCode != http.StatusCreated {
		t.Fatalf("expected post create success with token, got=%d body=%s", authPost.StatusCode, string(authPost.Body))
	}
}

func TestAuth_LoginWrongPassword(t *testing.T) {
	app := testkit.NewTestApp(t, nil)
	RegisterUser(app.Router, "carol", "password123", "carol@example.com")

	bad := LoginUser(app.Router, "carol", "wrong")
	if bad.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected 401 for wrong password, got=%d body=%s", bad.StatusCode, string(bad.Body))
	}
}
