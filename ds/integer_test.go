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
// Note: integer range
//       int16: -32,768 to 32,767
//       int32: -2,147,483,648 to 2,147,483,647
//       int64: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
// Note: maximium int64 + 1 => minimium int64 (negative)
//       minimium int64 - 1 => maximium int64 (positive)
func TestReverseInt64(t *testing.T) {
	var tests = []ReverseInt64TestCase{
		{121, 121},
		{12345, 54321},
		{-1234567, -7654321},
		{31415926, 62951413},
		{0x7FFFFFFFFFFFFFFF, 7085774586302733229},
		{7085774586302733229, 9223372036854775807},
		{8085774586302733229, 0},
		{-8085774586302733229, -9223372036854775808},
		{-8085774586302733230, -323372036854775808},
		{-666, -666},
		{777, 777},
		{9112000, 2119},
		{500, 5},
		{0, 0},
	}

	for index, test := range tests {
		var val = ReverseInt64(test.Number)
		var msg = fmt.Sprintf("expecting %v => %v", test.Number, test.Reversed)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.Reversed, val, msg)
	}
}
