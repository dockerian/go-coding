// +build all utils config

// Package utils :: config_test.go
package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_configFile                        = "config_test.yaml"
	_configReaderMock ConfigReaderFunc = func(file string) ([]byte, error) {
		f, _ := os.Open(_configFile)
		defer f.Close()
		return ioutil.ReadAll(f)
	}
	_configGetTests = map[string]string{
		"some.non.exist.var":   "default value",
		"more.environment.var": "another test for default value",
		"TEST_CASE.1":          "test case 1 name",
		"123_TEST_NAME":        "digital leading env var",
		"env var with space":   "environment variable name with space?",
	}
	_configGetBoolTests = []ConfigGetBoolTestCase{
		{keyValue: "False", expected: false},
		{keyValue: "FALSE", expected: false},
		{keyValue: "false", expected: false},
		{keyValue: "disabled", expected: false},
		{keyValue: "DISABLED", expected: false},
		{keyValue: "OFF", expected: false},
		{keyValue: "off", expected: false},
		{keyValue: "No", expected: false},
		{keyValue: "NO", expected: false},
		{keyValue: "-5", expected: false},
		{"1111", true, true},
		{"55555", false, false},
		{"Allowed", true, true},
		{keyValue: "accepted", expected: false},
		{keyValue: "1000", expected: false},
		{keyValue: "0", expected: false},
		{keyValue: "1", expected: true},
		{keyValue: "Enabled", expected: true},
		{keyValue: "enabled", expected: true},
		{keyValue: "On", expected: true},
		{keyValue: "ON", expected: true},
		{keyValue: "True", expected: true},
		{keyValue: "TRUE", expected: true},
		{keyValue: "true", expected: true},
		{keyValue: "YES", expected: true},
		{keyValue: "yes", expected: true},
	}
	_configGetInt32Tests = []ConfigGetInt32TestCase{
		{"ABCD", 1234, 1234},
		{keyValue: "-32768", expected: -32768},
		{keyValue: "-2147483649", expected: 0},
		{keyValue: "-2147483648", expected: -2147483648},
		{keyValue: "2147483647", expected: 2147483647},
		{keyValue: "2147483648", expected: 0},
		{keyValue: "-1", expected: -1},
		{keyValue: "0", expected: 0},
		{keyValue: "0xFF", expected: 0},
		{keyValue: "abcdef", expected: 0},
	}
	_configGetInt64Tests = []ConfigGetInt64TestCase{
		{"0xFFFF", 645321, 645321},
		{keyValue: "0567", expected: 567},
		{keyValue: "0xABCDEF", expected: 0},
		{"INFINITE", 9223372036854775807, 9223372036854775807},
		{keyValue: "-9223372036854775809", expected: 0},
		{keyValue: "-9223372036854775808", expected: -9223372036854775808},
		{keyValue: "9223372036854775807", expected: 9223372036854775807},
		{keyValue: "9223372036854775808", expected: 0},
		{keyValue: "-1", expected: -1},
		{keyValue: "0", expected: 0},
	}
	_configGetUint32Tests = []ConfigGetUint32TestCase{
		{"0xFFFF", 3210, 3210},
		{keyValue: "-32768", expected: 0},
		{keyValue: "-2147483649", expected: 0},
		{keyValue: "2147483647", expected: 2147483647},
		{keyValue: "2147483648", expected: 2147483648},
		{keyValue: "4294967295", expected: 4294967295},
		{keyValue: "4294967296", expected: 0},
		{keyValue: "-1", expected: 0},
		{keyValue: "0", expected: 0},
	}
	_configGetUint64Tests = []ConfigGetUint64TestCase{
		{"0xFFFF", 64, 64},
		{keyValue: "18446744073709551616", expected: 0},
		{keyValue: "18446744073709551615", expected: 18446744073709551615},
		{keyValue: "-1", expected: 0},
		{keyValue: "0", expected: 0},
	}
	_configTests = []ConfigTestCase{
		{"go.math.maxInt8", "127"},
		{"go.math.minInt8", "-128"},
		{"go.math.maxUint8", "255"},
		{"go.math.maxInt16", "32767"},
		{"go.math.minInt16", "-32768"},
		{"go.math.maxUint16", "65535"},
		{"go.math.maxInt32", "2147483647"},
		{"go.math.minInt32", "-2147483648"},
		{"go.math.maxUint32", "4294967295"},
		{"go.math.maxInt64", "9223372036854775807"},
		{"go.math.minInt64", "-9223372036854775808"},
		{"go.math.maxUint64", "18446744073709551615"},
		{"goobar.api_key", "SOCKEYE_API_KEY"},
		{"goobar.api_url", "https://sockeye/api"},
		{"goobar.bool_tests.0", "false"},
		{"goobar.bool_tests.1", "true"},
		{"goobar.bool_tests.2", "false"},
		{"goobar.bool_tests.3", "true"},
		{"goobar.bool_tests.4", "false"},
		{"goobar.bool_tests.5", "true"},
		{"goobar.bool_tests.6", "disabled"},
		{"goobar.bool_tests.7", "ENABLED"},
		{"goobar.bool_tests.8", "Disabled"},
		{"goobar.bool_tests.9", "enabled"},
		{"goobar.bool_tests.10", "TurnOff"},
		{"goobar.bool_tests.11", "true"},
		{"goobar.bool_tests.12", "false"},
		{"goobar.bool_tests.13", "true"},
		{"goobar.bool_tests.14", "false"},
		{"goobar.bool_tests.15", "true"},
		{"goobar.bool_tests.16", "0"},
		{"goobar.bool_tests.17", "1"},
		{"goobar.data.threat_class", "ThreatClass"},
		{"goobar.data.threat_family", "ThreatFamily"},
		{"goobar.data.threat_source", "ThreatSource"},
		{"xxdb.debug", "true"},
		{"xxdb.host", "127.0.0.1"},
		{"xxdb.port", "3306"},
		{"xxdb.username", "admin"},
		{"xxdb.password", "pass"},
	}
)

// ConfigGetFunc defines the function signature of Config getters
type ConfigGetFunc func(string, ...interface{}) interface{}

// ConfigGetFuncTestCase struct
type ConfigGetFuncTestCase struct {
	keyValue string
	expected interface{}
	vDefault interface{}
}

// ConfigTestCase struct
type ConfigTestCase struct {
	key string
	val string
}

// ConfigGetBoolTestCase struct
type ConfigGetBoolTestCase struct {
	keyValue string
	expected bool
	vDefault bool
}

// ConfigGetInt32TestCase struct
type ConfigGetInt32TestCase struct {
	keyValue string
	expected int32
	vDefault int32
}

// ConfigGetInt64TestCase struct
type ConfigGetInt64TestCase struct {
	keyValue string
	expected int64
	vDefault int64
}

// ConfigGetUint32TestCase struct
type ConfigGetUint32TestCase struct {
	keyValue string
	expected uint32
	vDefault uint32
}

// ConfigGetUint64TestCase struct
type ConfigGetUint64TestCase struct {
	keyValue string
	expected uint64
	vDefault uint64
}

// doMockDecrypt mocks _decryptFunc
func doMockDecrypt(key, value string) string {
	return value
}

// doTestConfigGetFunc tests Config getter functions
func doTestConfigGetFunc(t *testing.T, testKey string, testCases []ConfigGetFuncTestCase, getFunc ConfigGetFunc) {
	config := &Config{testKey, make(map[string]string)}

	for index, test := range testCases {
		config.settings[testKey] = test.keyValue
		result := getFunc(testKey, test.vDefault)
		msg := fmt.Sprintf("config[%v]: %v (default: %v) --> %v (actual: %v)",
			testKey, test.keyValue, test.vDefault, test.expected, result)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, result, msg)
	}
}

// doTestCleanup clean up for all tests
func doTestCleanup() {
	_decryptFunc = DecryptKeyTextByKMS
}

// doTestSetup prepares for all tests
func doTestSetup() {
	_decryptFunc = doMockDecrypt
}

// TestMain runs each test with setup and shutdown
func TestMain(m *testing.M) {
	doTestSetup()
	defer doTestCleanup()
	code := m.Run()
	os.Exit(code)
}

// TestCongigGet tests func Config.Get
func TestConfigGet(t *testing.T) {
	config := Config{}
	tindex := 1
	for key, val := range _configGetTests {
		t.Logf("Test %2d: %v => '' or default value: '%v'\n", tindex, key, val)
		os.Unsetenv(key)
		result1 := config.Get(key)
		reason1 := fmt.Sprintf("unset env: %s --> '' (actual: %s)\nos.Environ:%s\n",
			key, result1, os.Environ())
		assert.Equal(t, "", result1, reason1)

		result2 := config.Get(key, val)
		reason2 := fmt.Sprintf("unset env: %s --> default value: %s (actual: %s)",
			key, val, result2)
		assert.Equal(t, val, result2, reason2)
		tindex++
	}
}

// TestConfigGetBool tests func Config.GetBool
func TestConfigGetBool(t *testing.T) {
	mapkey := "bool"
	config := &Config{mapkey, make(map[string]string)}

	for index, test := range _configGetBoolTests {
		config.settings[mapkey] = test.keyValue
		result := config.GetBool(mapkey, test.vDefault)
		msg := fmt.Sprintf("config[%s]: %s (default: %t) --> %t (actual: %t)",
			mapkey, test.keyValue, test.vDefault, test.expected, result)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, result, msg)
	}
}

// TestConfigGetInt32 tests func Config.GetInt32
func TestConfigGetInt32(t *testing.T) {
	mapkey := "int32"
	config := &Config{mapkey, make(map[string]string)}

	for index, test := range _configGetInt32Tests {
		config.settings[mapkey] = test.keyValue
		result := config.GetInt32(mapkey, test.vDefault)
		msg := fmt.Sprintf("config[%s]: %s (default: %d) --> %d (actual: %d)",
			mapkey, test.keyValue, test.vDefault, test.expected, result)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, result, msg)
	}
}

// TestConfigGetInt64 tests func Config.GetInt64
func TestConfigGetInt64(t *testing.T) {
	mapkey := "int64"
	config := &Config{mapkey, make(map[string]string)}

	for index, test := range _configGetInt64Tests {
		config.settings[mapkey] = test.keyValue
		result := config.GetInt64(mapkey, test.vDefault)
		msg := fmt.Sprintf("config[%s]: %s (default: %v) --> %v (actual: %v)",
			mapkey, test.keyValue, test.vDefault, test.expected, result)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, result, msg)
	}

	config.settings[mapkey] = "FFFF"
	result := config.GetInt64(mapkey) // no default value
	assert.Equal(t, int64(0), result)
}

// TestConfigGetUint32 tests func Config.GetUint32
func TestConfigGetUint32(t *testing.T) {
	mapkey := "uint32"
	config := &Config{mapkey, make(map[string]string)}

	for index, test := range _configGetUint32Tests {
		config.settings[mapkey] = test.keyValue
		result := config.GetUint32(mapkey, test.vDefault)
		msg := fmt.Sprintf("config[%s]: %s (default: %d) --> %d (actual: %d)",
			mapkey, test.keyValue, test.vDefault, test.expected, result)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, result, msg)
	}
}

// TestConfigGetUint64 tests func Config.GetUint32
func TestConfigGetUint64(t *testing.T) {
	mapkey := "uint64"
	config := &Config{mapkey, make(map[string]string)}

	for index, test := range _configGetUint64Tests {
		config.settings[mapkey] = test.keyValue
		result := config.GetUint64(mapkey, test.vDefault)
		msg := fmt.Sprintf("config[%s]: %s (default: %v) --> %v (actual: %v)",
			mapkey, test.keyValue, test.vDefault, test.expected, result)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, result, msg)
	}

	config.settings[mapkey] = "ABCD"
	result := config.GetUint64(mapkey) // no default value
	assert.Equal(t, uint64(0), result)
}

// TestGetConfig tests func Config.GetConfig
func TestGetConfig(t *testing.T) {
	config := GetConfig(_configFile)

	for index, test := range _configTests {
		result := config.Get(test.key)
		msg := fmt.Sprintf("config[%s] == %s", test.key, test.val)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, result, test.val, msg)
	}

	anotherConfig := GetConfig(_configFile)
	assert.Equal(t, anotherConfig, config) // should be the same

	envKey := "NON_EXIST_ENVIRONMENT_KEY__"
	envVal := "This should not be set"
	os.Setenv(envKey, envVal)
	result := config.Get(envKey)
	assert.Equal(t, envVal, result)

	t.Logf("Testing: nil config\n")
	emptyConfig := GetConfig("../../../non-exist/path/to/config")
	assert.Equal(t, &Config{}, emptyConfig)
	assert.NotNil(t, emptyConfig)

	configsCopy := configs
	configs = map[string]Config{}
	nilConfig := GetConfig("/does not exist config.yaml")
	assert.Nil(t, nilConfig)
	configs = configsCopy

	newConfig := newConfig("../../../non-exist/path/to/config")
	assert.Nil(t, newConfig)
}
