// Package apimain :: appRouters.go
package apimain

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dockerian/go-coding/pkg/api"
	"github.com/dockerian/go-coding/pkg/cfg"
)

var (
	routeConfigs = api.RouteConfigs{
		{
			Pattern: "/", Method: "GET", Name: "Index", HandlerFunc: Index,
		},
		{
			Pattern: "/info", Method: "GET", Name: "GetInfo", HandlerFunc: Info,
		},
		{
			Pattern: `/private/{rest:.*}`, Method: "*", Name: "Private",
			Proxy: api.ProxyRoute{
				Prefix: "/private", RedirectURL: privateURL,
			},
		},
		{
			Pattern: `/test/{rest:.*}`, Method: "*", Name: "Test",
			Proxy: api.ProxyRoute{
				Prefix: "/test", RedirectURL: testURL,
			},
		},
		{
			Pattern: `/v1/{rest:.*}`, Method: "*", Name: "Data",
			Proxy: api.ProxyRoute{
				Prefix: "/v1", RedirectURL: dataURL,
				RedirectOnly: true,
			},
		},
		{
			Pattern: `/{rest:.+}`, Method: "*", Name: "NotFound",
			HandlerFunc: NotFound,
		},
	}
)

// Info handles /info path
func Info(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error {
	env := ctx.Env
	data := struct {
		Name        string
		AppAddress  string
		APIVersion  string
		AppVersion  string
		Author      string
		Copyright   string
		Description string
		Doc         string
	}{
		Name:        env.Get("appName"),
		AppAddress:  env.Get("appAddress"),
		APIVersion:  env.Get("apiVersion"),
		AppVersion:  env.Get("appVersion"),
		Description: env.Get("description"),
		Copyright:   env.Get("copyright"),
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
