// +build all ds array test

package ds

import (
	"fmt"
	"testing"
)

// ArrayTestCase struct
type ArrayTestCase struct {
	Data     interface{}
	Expected interface{}
}

// TestArray is a testing function template
func TestArray(t *testing.T) {
	var tests []ArrayTestCase

	for index, test := range tests {
		var val interface{}
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
	}
}
