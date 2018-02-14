// +build all common pkg api markdown

// Package api :: markdown_test.go
package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	_ "path/filepath"
	"testing"

	"path/filepath"

	"github.com/stretchr/testify/assert"
)

// TestMarkdownHandler tests func api.MarkdownHandler
func TestMarkdownHandler(t *testing.T) {
	pwd, _ := os.Getwd()

	t.Logf("pwd: %+v", pwd)

	tests := []struct {
		prefix     string
		reqPath    string
		docPath    string
		httpStatus int
	}{
		{"/doc", "/", pwd, http.StatusOK},
		{"/doc", "/foo", ".", http.StatusOK},
		{"/doc", "/foo", ".foo/test", http.StatusOK},
		{"/doc", "/test.md", "./test/md", http.StatusOK},
		{"/doc", "/test", pwd, http.StatusMovedPermanently},
		{"/doc", "/markdown.go", ".", http.StatusOK},
	}

	// creating a "test" dir
	testPath := filepath.Join(pwd, "test")
	os.MkdirAll(testPath, 0777)

	for idx, test := range tests {
		req, err := http.NewRequest("GET", test.reqPath, nil)
		if err != nil {
			t.Fatal(err)
		}

		next := func(w http.ResponseWriter, r *http.Request) {}
		appHandler := NewMarkdown(test.prefix, test.docPath, "README.md")
		rr := httptest.NewRecorder()
		appHandler.ServeHTTP(rr, req, http.HandlerFunc(next))
		msg := fmt.Sprintf("%s [%s] - rr: %+v", test.reqPath, test.docPath, rr.Code)
		t.Logf("Test %2d: %s", idx+1, msg)
		assert.Equal(t, test.httpStatus, rr.Code, msg)
		assert.False(t, rr.Flushed)
	}

	// deleting the "test" dir
	os.RemoveAll(testPath)
}
