// Package cfg :: config.go
// Get project settings from os env or specified config.yaml
package cfg

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/dockerian/go-coding/pkg/str"
	"gopkg.in/yaml.v2"
)

var (
	// _configParser is a ConfigParserFunc
	_configParser ConfigParserFunc = yaml.Unmarshal
	// _reader is a ConfigReaderFunc
	_configReader ConfigReaderFunc = ioutil.ReadFile
	// _decryptFunc is a decrypt function to return decrypted key value
	_decryptFunc DecryptFunc = DecryptKeyTextByKMS
	// configs use config file path mapping to a Config struct
	configs = map[string]Config{"": {}}
	// syncMtx is a private lock used by GetConfig
	syncMtx = &sync.Mutex{}
)

// Config represents a flattened settings per config file
type Config struct {
	configFile string
	settings   map[string]string
}

// ConfigParserFunc is a generic parser function
type ConfigParserFunc func([]byte, interface{}) error

// ConfigReaderFunc is a generic reader function
type ConfigReaderFunc func(string) ([]byte, error)

// DecryptFunc is a generic decrypt function
type DecryptFunc func(string, string) string

/*******************************************************************************
 * Config methods
 *******************************************************************************
 */

// Get gets string value of the key in os.Environ() or Config.settings
// or using defaultValues[0] if provided; otherwise return ""
func (c Config) Get(key string, defaultValues ...string) string {
	keyValue := ""
	if len(defaultValues) > 0 {
		keyValue = defaultValues[0]
	}
	keyUpper := strings.ToUpper(strings.Replace(key, ".", "_", -1))
	if enVal := os.Getenv(keyUpper); enVal != "" {
		keyValue = enVal
	} else if settingValue, okay := c.settings[key]; okay {
		keyValue = settingValue
	}
	if keyValue != "" {
		return DecryptKeyTextByKMS(key, keyValue)
	}
	return keyValue
}

// GetBool gets boolean value of the key, or defaultValues[0], or false
func (c Config) GetBool(key string, defaultValues ...bool) bool {
	defValue := len(defaultValues) > 0 && defaultValues[0]
	keyValue := c.Get(key)

	// check special boolean values: enabled|disabled, on|off, yes|no, or 1|0
	strLower := strings.TrimSpace(strings.ToLower(keyValue))
	if str.StringIn(strLower, []string{"1", "yes", "true", "on", "enabled"}) {
		return true
	}

	// check normal boolean values: true|false, 1|0
	if v, err := strconv.ParseBool(keyValue); err == nil {
		return v
	}

	return defValue
}

// GetInt32 gets int32 value of the key, or defaultValues[0], or 0
func (c Config) GetInt32(key string, defaultValues ...int32) int32 {
	var defaultValue int32
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	vInt64 := c.GetInt64(key, int64(defaultValue))

	if math.MinInt32 <= vInt64 && vInt64 <= math.MaxInt32 {
		return int32(vInt64)
	}
	return defaultValue
}

// GetInt64 gets int64 value of the key, or defaultValues[0], or 0
func (c Config) GetInt64(key string, defaultValues ...int64) int64 {
	if v, err := strconv.ParseInt(c.Get(key), 10, 64); err == nil {
		return v
	}
	if len(defaultValues) > 0 {
		return defaultValues[0]
	}
	return 0
}

// GetUint32 gets uint32 value of the key, or defaultValues[0], or 0
func (c Config) GetUint32(key string, defaultValues ...uint32) uint32 {
	var defaultValue uint32
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	vUint64 := c.GetUint64(key, uint64(defaultValue))

	// use default value on overflow
	if vUint64 > math.MaxUint32 {
		return defaultValue
	}
	return uint32(vUint64)
}

// GetUint64 gets uint64 value of the key, or defaultValues[0], or 0
func (c Config) GetUint64(key string, defaultValues ...uint64) uint64 {
	if v, err := strconv.ParseUint(c.Get(key), 10, 64); err == nil {
		return v
	}
	if len(defaultValues) > 0 {
		return defaultValues[0]
	}
	return 0
}

/*******************************************************************************
 * Config helper functions
 *******************************************************************************
 */

// Flatten builds a flattened key/value string pairs map
func Flatten(prefix string, value interface{}, kvmap map[string]string) {
	if value == nil {
		return
	}
	if list, ok := value.([]interface{}); ok {
		// flatten the value as a list
		for idx, val := range list {
			nextPrefix := fmt.Sprintf("%s.%d", prefix, idx)
			Flatten(nextPrefix, val, kvmap)
		}
	} else if dict, ok := value.(map[interface{}]interface{}); ok {
		// flattern value as a dictionary
		for key, val := range dict {
			nextPrefix := fmt.Sprintf("%s.%s", prefix, key)
			Flatten(nextPrefix, val, kvmap)
		}
	} else {
		kvmap[prefix] = fmt.Sprintf("%v", value)
	}
}

// FlattenConfig loads a config file (.yaml) to flattened key/value map
func FlattenConfig(file string) map[string]string {
	keyvalmap := map[string]string{}
	stringmap := map[string]interface{}{}

	if data, err := _configReader(file); err == nil {
		err = _configParser(data, stringmap)
		if err == nil {
			for key, val := range stringmap {
				if val != nil {
					Flatten(key, val, keyvalmap)
				}
			}
		}
	}
	return keyvalmap
}

// GetConfig gets a singleton instance of Config
func GetConfig(file string) *Config {
	syncMtx.Lock()
	defer syncMtx.Unlock()

	if fullpath, errf := filepath.Abs(file); errf == nil {
		if config, okay := configs[fullpath]; okay {
			return &config
		}
		if newConfigRef := newConfig(fullpath); newConfigRef != nil {
			configs[fullpath] = *newConfigRef
			return newConfigRef
		}
	}
	if config, ok := configs[""]; ok {
		return &config
	}
	return nil
}

/*******************************************************************************
 * Config constructor (private)
 *******************************************************************************
 */

// newConfig constructs a Config struct per file; otherwise, return nil
func newConfig(file string) *Config {
	config := Config{}
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		if fullpath, err := filepath.Abs(file); err == nil {
			config.configFile = fullpath
			config.settings = FlattenConfig(fullpath)
			return &config
		}
	}
	return nil
}
