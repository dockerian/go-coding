// +build all ds integer test

package ds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ReverseInt64TestCase struct
type ReverseInt64TestCase struct {
	Number   int64
	Reversed int64
}

// TestReverseInt64 tests ReverseInt function
func TestReverseInt64(t *testing.T) {
	var tests = []ReverseInt64TestCase{
		{121, 121},
		{12345, 54321},
		{-1234567, -7654321},
		{31415926, 62951413},
		{-666, -666},
		{777, 777},
		{9112000, 2119},
		{500, 5},
		{0, 0},
	}

	for index, test := range tests {
		var val = ReverseInt64(test.Number)
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Reversed)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.Reversed, val, msg)
	}
}
