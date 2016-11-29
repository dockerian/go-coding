package demo

// demonstrate io.Pipe to avoid unnecessary buffering and memory allocations
// see https://medium.com/stupid-gopher-tricks/streaming-data-in-go-without-buffering-3285ddd2a1e5#.nrnhmld57

import (
	"encoding/json"
	"io"
	"net/http"
)

func pipeStream(v interface{}) (*http.Response, error) {
	// Set up the pipe to write data directly into the Reader.
	pr, pw := io.Pipe()

	// Write JSON-encoded data to the Writer end of the pipe.
	// Write in a separate concurrent goroutine, and remember
	// to Close the PipeWriter, to signal to the paired PipeReader
	// that we’re done writing.
	go func() {
		err := json.NewEncoder(pw).Encode(&v)
		pw.CloseWithError(err)
	}()

	// Send the HTTP request. Whatever is read from the Reader
	// will be sent in the request body.
	// As data is written to the Writer, it will be available
	// to read from the Reader.
	resp, err := http.Post("example.com", "application/json", pr)

	return resp, err
}
