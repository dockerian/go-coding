// +build all ds string test

package ds

import (
	"fmt"
	"testing"
)

// StringTestCase struct
type StringTestCase struct {
	Data     interface{}
	Expected interface{}
}

// TestString is a testing function template
func TestString(t *testing.T) {
	var tests []StringTestCase

	for index, test := range tests {
		var val interface{}
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
	}
}
