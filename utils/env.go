// Package utils :: env.go - extended os env functions
package utils

import (
	"os"
	"strings"
)

// Env struct stores application-wide configuration
type Env map[string]interface{}

// Delete removes a key and the mapping value
func (env Env) Delete(key string) {
	delete(env, key)
}

// Get returns string for the mapping value by the key
func (env Env) Get(key string) string {
	if val := env.GetValue(key); val != nil {
		if strValue, ok := val.(string); ok {
			return strValue
		}
	}
	return ""
}

// GetInt returns int for the mapping value by the key
func (env Env) GetInt(key string) int {
	if val := env.GetValue(key); val != nil {
		if intValue, ok := val.(int); ok {
			return intValue
		}
	}
	return 0
}

// GetValue returns the mapping value by the key
func (env Env) GetValue(key string) interface{} {
	if val, ok := env[key]; ok {
		return val
	}
	return nil
}

// Set overwrite the mapping value by the key
func (env Env) Set(key string, value interface{}) {
	env[key] = value
}

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
