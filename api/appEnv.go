// Package apimain :: appEnv.go
package apimain

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dockerian/go-coding/utils"
)

const (
	defaultDataURL    = "http://localhost:5001"
	defaultPrivateURL = "http://localhost:5003"
	defaultTestURL    = "http://echo.jsontest.com"
)

var (
	// config is application-wide configuration model
	config = GetConfig()
	// dataURL is the private endpoint
	dataURL = config.Get("api.data_url", defaultDataURL)
	// privateURL is the private endpoint
	privateURL = config.Get("api.private_url", defaultPrivateURL)
	// testURL is the test endpoint
	testURL = config.Get("api.test_url", defaultTestURL)
)

// GetConfig returns an application configuration
func GetConfig() *utils.Config {
	pwd, _ := os.Getwd()
	return utils.GetConfig(filepath.Join(pwd, "config.yaml"))
}

// NewAppEnv constructs an utils.Env for the application
func NewAppEnv() utils.Env {
	pwd, _ := os.Getwd()
	configHost := config.Get("api.host")
	configPort := config.Get("api.port", "8001")
	configAddress := fmt.Sprintf("%s:%s", configHost, configPort)
	appName := config.Get("api.name", "Go API")
	appVersion := config.Get("api.version", "1.0.0")

	env := utils.Env{
		"appName":     appName,
		"author":      "Dockerian Seattle",
		"copyright":   "(C) 2016 Dockerian",
		"description": "Go RESTful API Example",
		"apiURL":      "/api/v1",
		"apiInfoURL":  "/api/info",
		"apiVersion":  "v1",
		"appVersion":  appVersion,
		"appAddress":  configAddress,
		"appHost":     configHost,
		"appPort":     configPort,
		"appLocation": pwd,
		"docLocation": filepath.Join(pwd, "doc"),
		"docIndex":    "README.md",
		"dataURL":     dataURL,
		"privateURL":  privateURL,
		"testURL":     testURL,
	}

	return env
}
