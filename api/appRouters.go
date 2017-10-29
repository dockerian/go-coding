// Package apimain :: appRouters.go
package apimain

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dockerian/go-coding/api/api"
	"github.com/dockerian/go-coding/utils"
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
			Pattern: `/private/{rest:[a-zA-Z0-9=\-\/]+}`, Method: "*", Name: "Private",
			Proxy: api.ProxyRoute{
				Prefix: "/private", RedirectURL: privateURL,
			},
		},
		{
			Pattern: `/test/{rest:[a-zA-Z0-9=\-\/]+}`, Method: "*", Name: "Test",
			Proxy: api.ProxyRoute{
				Prefix: "/test", RedirectURL: testURL,
			},
		},
		{
			Pattern: `/v1/{rest:[a-zA-Z0-9=\-\/\s% ]+}`, Method: "*", Name: "Data",
			Proxy: api.ProxyRoute{
				Prefix: "/v1", RedirectURL: dataURL,
			},
		},
	}
)

// Info handles /info path
func Info(env utils.Env, w http.ResponseWriter, r *http.Request) error {
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
func Index(env utils.Env, w http.ResponseWriter, r *http.Request) error {
	Info(env, w, r)
	return nil
}
