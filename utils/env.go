// +build all utils env

package utils

import (
	"os"
	"strings"
)

// GetEnvron get a map of environment variables
func GetEnvron() map[string]string {
	items := make(map[string]string)
	for _, item := range os.Environ() {
		pair := strings.Split(item, "=")
		key := pair[0]
		val := pair[1]
		items[key] = val
	}
	return items
}

// HasEnv checks if an environment variable is set
func HasEnv(name string, ignoreCase bool) bool {
	for _, item := range os.Environ() {
		pair := strings.Split(item, "=")
		ekey := pair[0]
		if ignoreCase {
			ekey = strings.ToLower(ekey)
			name = strings.ToLower(name)
		}
		if ekey == name {
			return true
		}
	}
	return false
}
