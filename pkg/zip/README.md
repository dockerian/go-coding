# zip
--
    import "github.com/dockerian/go-coding/pkg/zip"

Package zip :: gzip.go - gzip writer and handler

GZipHandler constructs a http handler wrapper to add gzip compression. See
https://gist.github.com/bryfry/09a650eb8aac0fb76c24

Package zip :: zip.go - zip extensions

## Usage

#### func  CreateZip

```go
func CreateZip(sources []*Source, w io.Writer) error
```
CreateZip copies from multiple sources to a writer

#### func  GZipHandler

```go
func GZipHandler(next http.Handler) http.Handler
```
GZipHandler wrap a http.Handler to support transparent gzip encoding

#### type GZipResponseWriter

```go
type GZipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}
```

GZipResponseWriter struct wraps an io.Writer and http.ResponseWriter

#### func (*GZipResponseWriter) Write

```go
func (gzw *GZipResponseWriter) Write(bytes []byte) (int, error)
```
Write implements io.Writer

#### type Source

```go
type Source struct {
	io.Reader
	Name string
	Size int
}
```

Source defines a zip reader struct with name and reader
