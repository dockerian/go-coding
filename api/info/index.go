package info

import (
	"encoding/json"
	"net/http"
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

// GetInfo is api/info handler
func GetInfo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(apiInfo); err != nil {
		panic(err)
	}
}
