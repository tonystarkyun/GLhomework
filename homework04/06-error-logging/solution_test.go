package q06errorlogging

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	"homework04/internal/testkit"
)

func TestErrorHandling_UniformShape(t *testing.T) {
	app := testkit.NewTestApp(t, nil)

	badID := testkit.JSONRequest(t, app.Router, http.MethodGet, "/posts/not-number", nil, "")
	if badID.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got=%d body=%s", badID.Code, badID.Body.String())
	}

	payload, err := ParseErrorBody(badID.Body.Bytes())
	if err != nil {
		t.Fatalf("parse error body: %v", err)
	}
	if payload.Error == "" {
		t.Fatalf("expected non-empty error message")
	}
}

func TestErrorHandling_And_Logging(t *testing.T) {
	var logs bytes.Buffer
	app := testkit.NewTestApp(t, &logs)
	token := testkit.RegisterAndLogin(t, app, "logger", "password123", "logger@example.com")

	invalidTokenResp := testkit.JSONRequest(t, app.Router, http.MethodPost, "/posts", map[string]string{
		"title":   "x",
		"content": "y",
	}, "bad-token")
	if invalidTokenResp.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for invalid token, got=%d body=%s", invalidTokenResp.Code, invalidTokenResp.Body.String())
	}

	notFoundResp := testkit.JSONRequest(t, app.Router, http.MethodGet, "/posts/99999", nil, "")
	if notFoundResp.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for missing post, got=%d body=%s", notFoundResp.Code, notFoundResp.Body.String())
	}

	_ = token // token proved successful auth path setup

	allLogs := logs.String()
	if !strings.Contains(allLogs, "error status=401") {
		t.Fatalf("expected unauthorized error log, got=%s", allLogs)
	}
	if !strings.Contains(allLogs, "error status=404") {
		t.Fatalf("expected not found error log, got=%s", allLogs)
	}
	if !strings.Contains(allLogs, "request method=") {
		t.Fatalf("expected request logs, got=%s", allLogs)
	}
}
