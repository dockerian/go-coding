// Package api :: markdown.go
package api

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/russross/blackfriday"
	"github.com/urfave/negroni"
)

var htmlTemplate = template.Must(template.New("base").Parse(`
<html>
  <head>
    <title>{{ .Path }}</title>
    <style>
      body {font-family:Helvetica,sans-serif;font-size:12pt;}
    </style>
  </head>
  <body>
    {{ .Body }}
  </body>
</html>
`))

// MarkdownHandler struct is a negroni.Handler
type MarkdownHandler struct {
	dir           string
	indexFile     string
	prefix        string
	staticHandler negroni.Handler
}

// NewMarkdown constructs a MarkdownHandler
func NewMarkdown(prefix, dir, indexFile string) negroni.Handler {
	return &MarkdownHandler{
		dir:       dir,
		indexFile: indexFile,
		prefix:    prefix,
		staticHandler: &negroni.Static{
			Dir:       http.Dir(dir),
			IndexFile: indexFile,
			Prefix:    prefix,
		},
	}
}

// ServeHTTP implements negroni.Handler interface
func (m *MarkdownHandler) ServeHTTP(
	rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	var reqPath = req.URL.Path
	if strings.HasSuffix(reqPath, "/") {
		reqPath = reqPath + m.indexFile
	}

	content, err := ioutil.ReadFile("." + reqPath)
	if err != nil {
		if !strings.HasSuffix(err.Error(), "a directory") {
			var handler = m.staticHandler
			if static, ok := handler.(*negroni.Static); ok {
				if !strings.HasPrefix(reqPath, static.Prefix) {
					log.Printf("[next] %s [%s]\n", reqPath, static.Prefix)
					next(rw, req)
					return
				}
			}
		}
		log.Printf("[fileserver] %v - (%v)\n", "."+req.URL.Path, err)
		http.ServeFile(rw, req, "."+req.URL.Path)
		return
	}

	if !strings.HasSuffix(reqPath, ".md") {
		log.Printf("[static] %s\n", reqPath)
		m.staticHandler.ServeHTTP(rw, req, next)
		return
	}

	log.Printf("[markdown]: %s\n", reqPath)
	md2html := blackfriday.MarkdownCommon(content)

	rw.Header().Set("Content-Type", "text/html")

	htmlTemplate.Execute(rw, struct {
		Body template.HTML
		Path string
	}{
		Body: template.HTML(string(md2html)),
		Path: req.URL.Path,
	})
}
