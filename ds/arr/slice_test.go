// +build all ds array arr slice test

package arr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ArrayTestCase struct
type ArrayTestCase struct {
	Data     interface{}
	Expected interface{}
}

// ReduceTestCase struct
type ReduceTestCase struct {
	Data     []interface{}
	Expected interface{}
	Func     func(interface{}, interface{}) interface{}
}

// TestArray is a testing function template
func TestArray(t *testing.T) {
	var tests []ArrayTestCase

	for index, test := range tests {
		var val interface{}
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.EqualValues(t, 1, 1, msg)
	}
}

// TestReduce is a testing function template
func TestReduce(t *testing.T) {
	var tests = []ReduceTestCase{}
	// 	{[]int{1, 22, 333, 4444, 5555, 1, 5555, 4444, 22}, 333,
	// 		func(a, b int) int {
	// 			return a ^ b
	// 		}},
	// }

	for index, test := range tests {
		var val = Reduce(test.Data, test.Func)
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.EqualValues(t, test.Expected, val, msg)
	}
}
