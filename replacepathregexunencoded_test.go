package replacepathregexunencoded_test

import (
	"context"
	"github.com/istosSchulz/replacepathregexunencoded"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReplacePathRegexUnencoded(t *testing.T) {
	cfg := replacepathregexunencoded.CreateConfig()
	cfg.Regex = "/api/([^/]+)/websockets/(import/progress)$"
	cfg.Replacement = "/datacenter/$2?tenantId=$1"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := replacepathregexunencoded.New(ctx, next, cfg, "replacePathFromURLRegex-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "wss://api.istos.io/api/138f5014-b17d-4964-81cb-07191ec822eb/websockets/import/progress", nil)
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
