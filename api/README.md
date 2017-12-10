# apimain
--
    import "github.com/dockerian/go-coding/api"

Package apimain :: app.go - main api entrance

Package apimain :: appEnv.go

Package apimain :: appRouters.go

Package apimain :: appServer.go

Package apimain :: root.go

## Usage

```go
var (
	// ListenPort is a configurable http port
	ListenPort = 8181

	// RootRoutes configures root routes
	// optionally read from config or move this to routes.go
	RootRoutes = Routes{
		{
			"/", "GET", rootHandler, "Index",
		},
		{
			"/info", "GET", info.GetInfo, "Info",
		},
	}
)
```

#### func  App

```go
func App()
```
App is the API main entrance

#### func  GetConfig

```go
func GetConfig() *cfg.Config
```
GetConfig returns an application configuration

#### func  Index

```go
func Index(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error
```
Index handles the root of api path

#### func  Info

```go
func Info(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error
```
Info handles /info path

#### func  ListenAndServe

```go
func ListenAndServe(server api.AppServerInterface, ctx *cfg.Context) error
```
ListenAndServe starts a server

#### func  NewAppContext

```go
func NewAppContext() *cfg.Context
```
NewAppContext constructs an cfg.Context for the application

#### func  NewAppEnv

```go
func NewAppEnv() *cfg.Env
```
NewAppEnv constructs an cfg.Env for the application

#### func  NewAppServer

```go
func NewAppServer() *api.AppServer
```
NewAppServer constructs AppServer with - a new *mux.Router - negroni middlewares

    any middleware by negroni.Use() should implements negroni.Handler interface:
       ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
    otherwise, by negroni.UseHandler() should implements http.Handler
    see https://github.com/urfave/negroni#handlers

#### func  Root

```go
func Root()
```
Root is api root entry pointer

#### type Route

```go
type Route struct {
	Pattern string
	Method  string
	Handler http.HandlerFunc
	Name    string
}
```

Route struct encapsulates an http route

#### type Routes

```go
type Routes []Route
```

Routes struct is an array of Route
