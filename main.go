// Package main :: main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dockerian/go-coding/api"
)

var (
	help = `
	Args:
		api - Start a RESTful API server

		cli - Run CLI commands (TBD)
	`
)

// main function
func main() {
	var arg string
	if len(os.Args) > 1 {
		arg = strings.ToLower(os.Args[1])
	}
	switch arg {
	case "api":
		apimain.App()

	case "cli":
		fallthrough
	default:
		fmt.Println(help)
	}
}
