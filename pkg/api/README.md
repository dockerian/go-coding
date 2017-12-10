# api
--
    import "github.com/dockerian/go-coding/pkg/api"

Package api :: appError.go api.Error interface wraps error with http status.
api.AppError composes error and http status code for http handler without
accessing to header in http.ResponseWriter.

Package api :: appHandler.go api.AppHandler declares an extended http.Handler
with configuration data and error (see api.AppError).

Package api :: appLogger.go - logging handlers

Package api :: appServer.go

Package api :: auth.go - auth handlers

Package api :: formatter.go - api formatters

Package api :: markdown.go - a Markdown handler

Package api :: params.go - http request parameters

Package api :: proxy.go - proxy handler For each Proxy redirect/forward call,
api.ProxyRoute defines a RedirectURL per prefix path. The api.ProxyRoute
implements http.Handler interface so that the struct pointer itself can be
wrapped in a routing configuration; optionally, a Proxy() can construct a
http.Handler with prefix and predefined redirect URL.

Package api :: redirect.go - http redirect handlers

Package api :: route.go

## Usage

#### func  AppLogger

```go
func AppLogger(handler http.Handler, name string) http.Handler
```
AppLogger is a middleware handler returns http.HandlerFunc

#### func  Auth

```go
func Auth(next http.Handler, token string) http.Handler
```
Auth creates a http.Handler wrapper to check api token in request

#### func  BasicAuth

```go
func BasicAuth(next http.Handler, user, password string) http.Handler
```
BasicAuth creates a http.Handler wrapper to check auth in request

#### func  GetJSONEncoder

```go
func GetJSONEncoder(w io.Writer, indent string) *json.Encoder
```
GetJSONEncoder returns JSON by specified indent

#### func  NewLogger

```go
func NewLogger(prefix string) *negroni.Logger
```
NewLogger returns a new negroni.Logger instance

#### func  NewMarkdown

```go
func NewMarkdown(prefix, dir, indexFile string) negroni.Handler
```
NewMarkdown constructs a MarkdownHandler

#### func  NewMiddleware

```go
func NewMiddleware() *negroni.Negroni
```
NewMiddleware returns a negroni middleware

#### func  NewRecovery

```go
func NewRecovery() *negroni.Recovery
```
NewRecovery returns a negroni recovery handler

#### func  NewRouter

```go
func NewRouter(ctx cfg.Context, routeConfigs RouteConfigs) *mux.Router
```
NewRouter returns *mux.Router

#### func  Proxy

```go
func Proxy(prefix, redirectURL string, w http.ResponseWriter, r *http.Request) error
```
Proxy generates new request with prefix path to redirected url

#### func  ProxyHandler

```go
func ProxyHandler(prefix, redirectURL string) http.Handler
```
ProxyHandler constructs an http.Handler by prefix path and redirect URL

#### func  Redirect

```go
func Redirect(prefix, redirectURL string, w http.ResponseWriter, r *http.Request) error
```
Redirect forwards call by prefix path to redirected url

#### func  RedirectHandler

```go
func RedirectHandler(prefix, redirectURL string) http.Handler
```
RedirectHandler constructs an http.Handler by prefix path and redirect URL

#### func  RedirectToHTTPS

```go
func RedirectToHTTPS(next http.Handler) http.Handler
```
RedirectToHTTPS is a middleware handler returns http.HandlerFunc

#### func  WriteJSON

```go
func WriteJSON(w http.ResponseWriter, code int, data interface{})
```
WriteJSON writes status code and response data

#### type AppError

```go
type AppError struct {
	// Err inherits standard error interface
	Err error
	// StatusCode is http status code
	StatusCode int
}
```

AppError represents an error with an associated HTTP status code.

#### func (AppError) Error

```go
func (ae AppError) Error() string
```
Error for AppError to implement the error interface.

#### func (AppError) Status

```go
func (ae AppError) Status() int
```
Status returns HTTP status code.

#### func (AppError) StatusText

```go
func (ae AppError) StatusText() string
```
StatusText returns HTTP status text.

#### type AppHandler

```go
type AppHandler struct {
	cfg.Context
	Handle func(e cfg.Context, w http.ResponseWriter, r *http.Request) error
}
```

AppHandler struct wraps Env and Handle function implementing http.Handler

#### func (AppHandler) ServeHTTP

```go
func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)
```
ServeHTTP implements http.Handler

#### type AppLoggerResponseWriter

```go
type AppLoggerResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}
```

AppLoggerResponseWriter implements http.ResponseWriter

#### func (*AppLoggerResponseWriter) WriteHeader

```go
func (alw *AppLoggerResponseWriter) WriteHeader(code int)
```
WriteHeader implements http.ResponseWriter interface

#### type AppServer

```go
type AppServer struct {
	Ctx     *cfg.Context       // app context
	Doc     negroni.Handler    // negroni static handler
	Handler *negroni.Negroni   // negroni handler
	Logger  *negroni.Logger    // negroni logger
	Router  *mux.Router        // mux router
	Server  AppServerInterface // http.Server implements ListenAndServe
}
```

AppServer represents partial http.Server

#### type AppServerInterface

```go
type AppServerInterface interface {
	ListenAndServe() error
	ListenAndServeTLS(string, string) error
}
```

AppServerInterface represents partial http.Server interface

#### type BodyParams

```go
type BodyParams map[string]interface{}
```

BodyParams struct contains an JSON key/value map object

#### type Error

```go
type Error interface {
	error
	Status() int
}
```

Error represents a handler error to provide Status() and embed the built-in
error interface.

#### func  WriteAppError

```go
func WriteAppError(w http.ResponseWriter, apiError Error) Error
```
WriteAppError writes status code and returns Error

#### func  WriteError

```go
func WriteError(w http.ResponseWriter, code int, message string) Error
```
WriteError writes status code and returns Error

#### type HandlerFunc

```go
type HandlerFunc func(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error
```

HandlerFunc represents a handler function with cfg.Context

#### type MarkdownHandler

```go
type MarkdownHandler struct {
}
```

MarkdownHandler struct is a negroni.Handler

#### func (*MarkdownHandler) ServeHTTP

```go
func (m *MarkdownHandler) ServeHTTP(
	rw http.ResponseWriter, req *http.Request, next http.HandlerFunc)
```
ServeHTTP implements negroni.Handler interface

#### type Params

```go
type Params struct {
	Form url.Values
	Body []byte
	Path string
	Post BodyParams
	Vars map[string]string
}
```

Params struct contains key/value pairs from URL path, request body, and query
string

#### func  NewParams

```go
func NewParams(r *http.Request) *Params
```
NewParams returns pointer to an instance of Params struct with parsed
http.Request

#### func (*Params) GetBody

```go
func (params *Params) GetBody(key string) interface{}
```
GetBody method returns pointer to body param by key name

#### func (*Params) GetDateRange

```go
func (params *Params) GetDateRange(key string) ([]time.Time, error)
```
GetDateRange method returns a date range by the key name

#### func (*Params) GetDateValues

```go
func (params *Params) GetDateValues(key string) ([]time.Time, error)
```
GetDateValues method returns sorted date values by the key name

#### func (*Params) GetInt

```go
func (params *Params) GetInt(key string, defaultValues ...int) (int, error)
```
GetInt method returns int value by the key name or the second parameter as
default value

#### func (*Params) GetIntByRange

```go
func (params *Params) GetIntByRange(key string, rangeValues ...int) int
```
GetIntByRange method returns int value by the key name and within the range of
rangeValues parameters

#### func (*Params) GetNextPageURL

```go
func (params *Params) GetNextPageURL(pgOffsetKey string, pgOffset int) string
```
GetNextPageURL returns next page URL per current page offset

#### func (*Params) GetValue

```go
func (params *Params) GetValue(key string, defaultValues ...string) string
```
GetValue method returns the value string by the key name or the second parameter
as default value

#### func (*Params) GetValues

```go
func (params *Params) GetValues(key string) []string
```
GetValues method returns values by the key name

#### func (*Params) HasKey

```go
func (params *Params) HasKey(key string) bool
```
HasKey returns true if the params has the key; otherwise, return false

#### type ProxyClient

```go
type ProxyClient interface {
	Do(*http.Request) (*http.Response, error)
}
```

ProxyClient interface

#### type ProxyRoute

```go
type ProxyRoute struct {
	Prefix      string
	RedirectURL string
}
```

ProxyRoute struct defines a redirecting URL based on pattern

#### func (*ProxyRoute) ServeHTTP

```go
func (p *ProxyRoute) ServeHTTP(w http.ResponseWriter, r *http.Request)
```
ServeHTTP implements http.Handler

#### type RouteConfig

```go
type RouteConfig struct {
	Pattern     string
	Method      string
	Name        string
	HandlerFunc func(e cfg.Context, w http.ResponseWriter, r *http.Request) error
	Proxy       ProxyRoute
}
```

RouteConfig struct

#### type RouteConfigs

```go
type RouteConfigs []RouteConfig
```

RouteConfigs is a list of RouteConfig struct
