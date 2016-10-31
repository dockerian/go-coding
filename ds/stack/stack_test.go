// +build all ds stack test

package stack

import (
	"fmt"
	"testing"
)

// StackTestCase struct
type StackTestCase struct {
	Data     interface{}
	Expected interface{}
}

// TestStack is a testing function template
func TestStack(t *testing.T) {
	var tests []StackTestCase

	for index, test := range tests {
		var val interface{}
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
	}
}
