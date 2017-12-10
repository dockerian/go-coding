// Package cfg :: context.go
package cfg

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

// Context struct wraps Env, http.Cookie, gorilla Session, and Context
type Context struct {
	context.Context
	Cookie  *http.Cookie
	Session *sessions.Session
	Env     *Env
}

// Value implements context.Context
func (ctx *Context) Value(key interface{}) interface{} {
	if key == "Env" && ctx.Env != nil {
		return ctx.Env
	}
	if ctx.Context != nil {
		return ctx.Context.Value(key)
	}
	return nil
}
