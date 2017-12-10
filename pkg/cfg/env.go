// Package cfg :: env.go
package cfg

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
	//TODO [jzhu]: considering to return a string representation by
	//e.g. fmt.Sprintf("%+v", val)
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
