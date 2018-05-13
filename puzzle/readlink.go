// Package puzzle :: readlink.go
package puzzle

import (
	"strings"
)

// Readlink normalizes a path string to remove extra '.' and '..'.
func Readlink(linkPath string) string {
	var normalizedPaths []string
	input := strings.Trim(linkPath, " ")
	paths := strings.Split(input, "/")

	pos := -1
	startWithRoot, prefix := false, ""
	if len(input) > 0 && input[0] == '/' {
		startWithRoot, prefix, paths = true, "/", paths[1:]
	}
	for _, path := range paths {
		if path == ".." {
			if pos >= 0 {
				normalizedPaths = normalizedPaths[0:pos]
			} else if !startWithRoot {
				normalizedPaths = append(normalizedPaths, "..")
			}
			pos--
		} else if path != "." && path != "" {
			normalizedPaths = append(normalizedPaths, path)
			pos++
		}
	}

	if len(normalizedPaths) > 0 {
		return prefix + strings.Join(normalizedPaths, "/")
	} else if startWithRoot {
		return prefix
	}

	return "."
}
