package plugindemo_test

import (
	"context"
	"github.com/traefik/plugindemo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReplacePathFromUrlRegex(t *testing.T) {
	cfg := plugindemo.CreateConfig()
	cfg.Regex = "^/api/([^/]+)/websockets/(import/progress)$"
	cfg.Replacement = "/datacenter/$2?tenantId=$1"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := plugindemo.New(ctx, next, cfg, "replacePathFromURLRegex-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.planning.cloud/api/75cf3229-9763-47c5-9789-b2e3c7fa8051/websockets/import/progress", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertPath(t, req, "/a/uri")

}

func assertPath(t *testing.T, req *http.Request, expected string) {
	t.Helper()

	var currentPath string
	if req.URL.RawPath == "" {
		currentPath = req.URL.Path
	} else {
		currentPath = req.URL.RawPath
	}

	if currentPath != expected {
		t.Errorf("invalid RawPath value: %s, expected: %s", currentPath, expected)
	}
}
