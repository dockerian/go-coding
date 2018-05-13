// +build all puzzle str readlink

// Package puzzle :: readlink.go
package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestReadlink tests Readlink func.
func TestReadlink(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{"", "."},
		{"/", "/"},
		{"////", "/"},
		{".", "."},
		{"..", ".."},
		{".//", "."},
		{"../", ".."},
		{"一級/二級/subdir", "一級/二級/subdir"},
		{"./a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u/v/w/x/y/z", "a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u/v/w/x/y/z"},
		{"a/b", "a/b"},
		{"a/b/../..", "."},
		{"a/b/../../..", ".."},
		{"a/b/../../../..", "../.."},
		{"/a/b/c/..", "/a/b"},
		{"/a/b/../..", "/"},
		{"/a/b/c/.././../d", "/a/d"},
		{"/a/../../up", "/up"},
		{"/a/../../..", "/"},
	}
	for index, v := range tests {
		var result = Readlink(v.Input)
		var msg = fmt.Sprintf("expecting '%v' => '%v'", v.Input, v.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, v.Expected, result, msg)
	}
}
