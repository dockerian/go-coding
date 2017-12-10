// Package zip :: zip.go - zip extensions
package zip

import (
	"archive/zip"
	"errors"
	"io"
	"time"
)

// Source defines a zip reader struct with name and reader
type Source struct {
	io.Reader
	Name string
	Size int
}

// CreateZip copies from multiple sources to a writer
func CreateZip(sources []*Source, w io.Writer) error {
	if w == nil || len(sources) == 0 {
		return errors.New("undefined writer and readers")
	}

	// note: zip.NewWriter does not create a writer implements io.Writer
	container := zip.NewWriter(w)
	// defer container.Close()

	for _, source := range sources {
		if source != nil && source.Reader != nil {
			header := &zip.FileHeader{
				Method: zip.Deflate,
				Name:   source.Name,
			}
			header.SetModTime(time.Now().UTC())

			zw, err := container.CreateHeader(header) // construct an io.Writer

			if err != nil {
				return err
			}

			if _, err := io.Copy(zw, source.Reader); err != nil {
				return err
			}
		}
	}

	err := container.Close()

	return err
}
