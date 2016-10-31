// +build all ds queue test

package queue

import (
	"fmt"
	"testing"
)

// QueueTestCase struct
type QueueTestCase struct {
	Data     interface{}
	Expected interface{}
}

// TestQueue is a testing function template
func TestQueue(t *testing.T) {
	var tests []QueueTestCase

	for index, test := range tests {
		var val interface{}
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
	}
}
