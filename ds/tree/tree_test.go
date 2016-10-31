// +build all ds tree test

package tree

import (
	"fmt"
	"testing"
)

// TreeTestCase struct
type TreeTestCase struct {
	Data     interface{}
	Expected interface{}
}

// TestTree is a testing function template
func TestTree(t *testing.T) {
	var tests []TreeTestCase

	for index, test := range tests {
		var val interface{}
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
	}
}
