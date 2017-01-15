package utils

import (
	"os"
	"strings"
)

// DebugEnv indicates DEBUG = 1|on|true in environment variable (ignoring case)
var DebugEnv = CheckEnvBoolean("DEBUG", true)

// DebugEnvStore keeps a backup of DebugEnv value
var DebugEnvStore = DebugEnv

// CheckEnvBoolean checks if an environment variable is set to non-false/non-zero
func CheckEnvBoolean(name string, ignoreCase bool) bool {
	for _, item := range os.Environ() {
		pair := strings.Split(item, "=")
		ekey := pair[0]
		eVal := strings.ToLower(pair[1])
		if ignoreCase {
			ekey = strings.ToLower(ekey)
			name = strings.ToLower(name)
		}
		if ekey == name {
			return eVal == "true" || eVal == "on" || eVal == "1"
		}
	}
	return false
}

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
