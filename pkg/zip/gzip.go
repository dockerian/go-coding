// Package zip :: gzip.go - gzip writer and handler
//
// GZipHandler constructs a http handler wrapper to add gzip compression.
// See https://gist.github.com/bryfry/09a650eb8aac0fb76c24
package zip

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
	"sync"
)

// gzipPool contains previously used Writers
var gzipPool = sync.Pool{New: func() interface{} {
	return gzip.NewWriter(nil)
}}

// GZipResponseWriter struct wraps an io.Writer and http.ResponseWriter
type GZipResponseWriter struct {
	io.Writer
	http.ResponseWriter
	done bool
}

// Write implements io.Writer
func (gzw *GZipResponseWriter) Write(bytes []byte) (int, error) {
	if !gzw.done {
		if gzw.Header().Get("Content-Type") == "" {
			gzw.Header().Set("Content-Type", http.DetectContentType(bytes))
		}
		gzw.done = true
	}
	return gzw.Writer.Write(bytes)
}

// GZipHandler wrap a http.Handler to support transparent gzip encoding
func GZipHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Accept-Encoding")
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")

		// get a Writer from the Pool
		gz := gzipPool.Get().(*gzip.Writer)
		// at the end, put the Writer back in to the Pool
		defer gzipPool.Put(gz)

		// reset the writer
		gz.Reset(w)
		defer gz.Close()

		gzrWriter := &GZipResponseWriter{Writer: gz, ResponseWriter: w}
		next.ServeHTTP(gzrWriter, r)
	})
}
