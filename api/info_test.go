// +build all app api info

// Package apimain :: info_test.go
package apimain

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dockerian/go-coding/api/v1/client"
	"github.com/dockerian/go-coding/pkg/cfg"
	"github.com/stretchr/testify/assert"
)

// TestGetDbInfo tests func sng.GetDbInfo
func TestGetDbInfo(t *testing.T) {
	ctx := NewAppContext()
	req, _ := http.NewRequest("GET", "/v1/db", nil)
	rec := httptest.NewRecorder()
	err := GetDbInfo(*ctx, rec, req)
	assert.Nil(t, err)

	dbInfo := client.DbSchema{}
	json.NewDecoder(rec.Body).Decode(&dbInfo)
	assert.NotNil(t, dbInfo)
}

// TestGetDbInfoAll tests func sng.GetDbInfoAll
func TestGetDbInfoAll(t *testing.T) {
	ctx := NewAppContext()
	req, _ := http.NewRequest("GET", "/v1/db/all", nil)
	rec := httptest.NewRecorder()
	err := GetDbInfoAll(*ctx, rec, req)
	assert.Nil(t, err)

	dbInfoAll := []client.DbSchema{}
	json.NewDecoder(rec.Body).Decode(&dbInfoAll)
	assert.NotNil(t, dbInfoAll)
}

// TestGetInfo tests func api.GetInfo
func TestGetInfo(t *testing.T) {
	tests := []func(cfg.Context, http.ResponseWriter, *http.Request) error{
		AppIndex,
		Info,
	}
	for _, test := range tests {
		ctx := NewAppContext()
		req, _ := http.NewRequest("GET", "/v1/info", nil)
		rec := httptest.NewRecorder()
		env := ctx.Env
		env.Set("appName", "appName")
		env.Set("appDescription", "app description")
		env.Set("appVersion", "appVersion")

		err := test(*ctx, rec, req)
		assert.Nil(t, err)

		appInfo := client.ApiSchema{}
		json.NewDecoder(rec.Body).Decode(&appInfo)
		assert.Equal(t, "app description", appInfo.Description)
		assert.Equal(t, "appVersion", appInfo.Version)
	}
}

// TestNotDefined tests func api.NotDefined
func TestNotDefined(t *testing.T) {
	ctx := NewAppContext()
	req, _ := http.NewRequest("GET", "/v1/not/implemented/yet", nil)
	rec := httptest.NewRecorder()
	err := NotDefined(*ctx, rec, req)
	assert.Equal(t, http.StatusNotImplemented, rec.Code)
	assert.NotNil(t, err)
}

// TestNotFound tests func api.NotFound
func TestNotFound(t *testing.T) {
	ctx := NewAppContext()
	req, _ := http.NewRequest("GET", "/v1/not/found", nil)
	rec := httptest.NewRecorder()
	err := NotFound(*ctx, rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.NotNil(t, err)
}
