package q01projectinit

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"homework04/internal/blog"
	"homework04/internal/testkit"
)

func TestBuildApp_BootstrapAndMigrate(t *testing.T) {
	dsn := fmt.Sprintf("file:q1_%d?mode=memory&cache=shared", time.Now().UnixNano())
	app, err := BuildApp(dsn)
	if err != nil {
		t.Fatalf("build app: %v", err)
	}

	if !app.DB.Migrator().HasTable(&blog.User{}) {
		t.Fatalf("users table was not migrated")
	}
	if !app.DB.Migrator().HasTable(&blog.Post{}) {
		t.Fatalf("posts table was not migrated")
	}
	if !app.DB.Migrator().HasTable(&blog.Comment{}) {
		t.Fatalf("comments table was not migrated")
	}
}

func TestBuildApp_RoutesAreMounted(t *testing.T) {
	app := testkit.NewTestApp(t, nil)
	w := testkit.JSONRequest(t, app.Router, http.MethodPost, "/register", map[string]string{}, "")
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected register route to exist and validate payload, got code=%d body=%s", w.Code, w.Body.String())
	}
}
