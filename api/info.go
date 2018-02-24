// Package apimain :: info.go
package apimain

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dockerian/go-coding/api/db"
	"github.com/dockerian/go-coding/pkg/api"
	"github.com/dockerian/go-coding/pkg/cfg"
)

// APIInfo struct
type APIInfo struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"desc,omitempty"`
	Copyright   string `json:"copyright,omitempty"`
	Author      string `json:"author,omitempty"`
	APIURL      string `json:"api_url,omitempty"`
	APIInfoURL  string `json:"api_info_url,omitempty"`
	APIVersion  string `json:"api_version,omitempty"`
	Version     string `json:"version,omitempty"`
}

var (
	apiInfo = APIInfo{
		Name:        "Go API",
		Description: "Go RESTful API Example",
		Copyright:   "(C) 2016 Dockerian",
		Author:      "Dockerian Seattle",
		APIURL:      "/api/v1",
		APIInfoURL:  "/api/info",
		APIVersion:  "v1",
		Version:     "0.0.1",
	}
)

// AppIndex handles the root of api path
func AppIndex(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error {
	return Info(ctx, w, r)
}

// GetDbInfo handles /info/db path
func GetDbInfo(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error {
	dbInfo := db.SchemaInfo(ctx)
	log.Printf("[DbSchema] %+v\n", dbInfo)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	encoder := api.GetJSONEncoder(w, "  ")
	encoder.Encode(dbInfo)
	return nil
}

// GetDbInfoAll handles /info/db/all path
func GetDbInfoAll(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error {
	dbInfoAll := db.SchemaInfoAll(ctx)
	log.Printf("[DbSchema] %+v\n", dbInfoAll)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	encoder := api.GetJSONEncoder(w, "  ")
	encoder.Encode(dbInfoAll)
	return nil
}

// GetInfo is api/info handler
func GetInfo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(apiInfo); err != nil {
		panic(err)
	}
}

// Info handles /info path
func Info(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error {
	env := ctx.Env
	data := struct {
		Name        string `json:"name"`
		AppAddress  string `json:"appURL"`
		APIVersion  string `json:"apiVersion"`
		AppVersion  string `json:"version"`
		Author      string `json:"author"`
		Copyright   string `json:"copyright"`
		Description string `json:"description"`
		Doc         string `json:"doc"`
	}{
		Name:        env.Get("appName"),
		AppAddress:  env.Get("appAddress"),
		APIVersion:  env.Get("apiVersion"),
		AppVersion:  env.Get("appVersion"),
		Description: env.Get("appDescription"),
		Copyright:   env.Get("appCopyright"),
		Doc:         env.Get("docLocation"),
	}
	fmt.Printf("info: %+v\n", data)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
	return nil
}

// Index handles the root of api path
func Index(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error {
	Info(ctx, w, r)
	return nil
}

// NotDefined handles any unimplemented path
func NotDefined(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error {
	return api.WriteError(w, http.StatusNotImplemented, "API gateway not implemented")
}

// NotFound handles /{rest} path
func NotFound(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error {
	return api.WriteError(w, http.StatusNotFound, "API gateway request not found")
}
