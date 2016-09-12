// +build all ds linkedlist test

package ds

import (
	"fmt"
	"testing"
)

// LinkedListTestCase struct
type LinkedListTestCase struct {
	Data     interface{}
	Expected interface{}
}

// TestLinkedList is a testing function template
func TestLinkedList(t *testing.T) {
	var tests []LinkedListTestCase

	for index, test := range tests {
		var val interface{}
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
	}
}
