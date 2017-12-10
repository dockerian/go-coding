// +build all common pkg cfg env

// Package cfg :: env_test.go
package cfg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_cfg = Config{}
	_env = Env{}
)

// TestEnv tests Env methods
func TestEnv(t *testing.T) {
	tests := []struct {
		key string
		val interface{}
	}{
		{"alphabet", "abcdefghijklmnopqrstuvwxyz"},
		{"cfg", _cfg},
		{"some key", "某些结果"},
		{"变量Ａ", "ＧＯ语言的系统设置 (Chinese)"},
		{"t", "test"},
	}

	for idx, test := range tests {
		_env.Delete(test.key)
		result1 := _env.Get(test.key)
		reason1 := fmt.Sprintf("_env[%s] => '' (actual: %s)", test.key, result1)
		t.Logf("Test %2d-Get(): %s\n", idx+1, reason1)
		assert.Equal(t, "", result1, reason1)

		_env.Set(test.key, test.val)

		result2 := _env.GetValue(test.key)
		reason2 := fmt.Sprintf("_env[%s] => %+v (actual: %+v)", test.key, test.val, result2)
		t.Logf("Test %2d-GetValue(): %s\n", idx+1, reason2)
		assert.Equal(t, test.val, result2, reason2)

		textVal := fmt.Sprintf("%+v", test.val)
		_env.Set(test.key, textVal)

		result3 := _env.Get(test.key)
		reason3 := fmt.Sprintf("_env[%s] => '%s' (actual: %s)", test.key, test.val, result3)
		t.Logf("Test %2d-GetValue(): %s\n", idx+1, reason2)
		assert.Equal(t, textVal, result3, reason3)
	}
}

// TestEnvGetInt tests Env method GetInt()
func TestEnvGetInt(t *testing.T) {
	tests := []struct {
		key string
		val interface{}
	}{
		{"neg1", -1},
		{"neg314", -314},
		{"piDigit", 31415926},
		{"test", 93457},
		{"zero", 0},
	}

	for idx, test := range tests {
		_env.Delete(test.key)
		result1 := _env.GetInt(test.key)
		reason1 := fmt.Sprintf("_env[%s] => 0 (actual: %d)", test.key, result1)
		t.Logf("Test %2d-GetInt(): %s\n", idx+1, reason1)
		assert.Equal(t, 0, result1, reason1)

		_env.Set(test.key, test.val)

		result2 := _env.GetValue(test.key)
		reason2 := fmt.Sprintf("_env[%s] => %+v (actual: %v)", test.key, test.val, result2)
		t.Logf("Test %2d-GetValue(): %s\n", idx+1, reason2)
		assert.Equal(t, test.val, result2, reason2)

		result3 := _env.GetInt(test.key)
		reason3 := fmt.Sprintf("_env[%s] => %d (actual: %d)", test.key, test.val, result3)
		t.Logf("Test %2d-GetInt(): %s\n", idx+1, reason2)
		assert.Equal(t, test.val, result3, reason3)
	}
}
