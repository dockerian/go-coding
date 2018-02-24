// Package apimain :: appRouters.go
package apimain

import (
	"github.com/dockerian/go-coding/pkg/api"
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
