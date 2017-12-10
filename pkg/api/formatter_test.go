// +build all common pkg api formatter

// Package api :: formatter_test.go
package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetJSONEncoder tests GetJSONEncoder function
func TestGetJSONEncoder(t *testing.T) {
	buffwrt := bytes.NewBufferString("")
	encoder := GetJSONEncoder(buffwrt, "    ")

	author := "Jason Zhu <jason.zhuyx@gmail.com>"
	dbVersion := "1.x.x"
	script := "deploy-1.x.x.sql"
	data := struct {
		Author    string `json:"author"`
		DbVersion string `json:"dbVersion"`
		Script    string `json:"script"`
	}{
		Author:    author,
		DbVersion: dbVersion,
		Script:    script,
	}

	expected := `{
    "author": "Jason Zhu <jason.zhuyx@gmail.com>",
    "dbVersion": "1.x.x",
    "script": "deploy-1.x.x.sql"
}
`
	encoder.Encode(&data)
	assert.Equal(t, expected, buffwrt.String())
}

// TestWriteError tests func api.WriteError
func TestWriteError(t *testing.T) {
	reqInfo := "GET /app/error/path"
	msg := fmt.Sprintf("%s - %s", reqInfo, "404 not found")
	rec := httptest.NewRecorder()
	code := http.StatusNotFound
	err := WriteError(rec, code, msg)
	assert.Equal(t, code, rec.Code)
	assert.NotNil(t, err)
}

// TestWriteJSON tests func api.WriteJSON
func TestWriteJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	code := http.StatusOK
	data := map[string]string{"test": "pass"}
	WriteJSON(rec, code, data)
	assert.Equal(t, code, rec.Code)
}
