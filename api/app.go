// Package apimain :: app.go - main api entrance
package apimain

import (
	"log"
)

// App is the API main entrance
func App() {
	app := NewAppServer()

	log.Fatal(ListenAndServe(app.Server, app.Ctx))
}
