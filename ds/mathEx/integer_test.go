// +build all ds math integer test

package mathEx

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ParseInt64TestCase struct {
	Data    string
	Integer int64
	IsValid bool
}

// ReverseInt64TestCase struct
type ReverseInt64TestCase struct {
	Number   int64
	Reversed int64
}

// TestParseInt64 tests ParseInt64 function
// Note: int64 range: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
func TestParseInt64(t *testing.T) {
	var tests = []ParseInt64TestCase{
		{"", 0, true},
		{"-9223372036854775808", math.MinInt64, true},
		{"9223372036854775808", 922337203685477580, false},
		{"9223372036854775807", math.MaxInt64, true},
		{"  +31415926", 31415926, true},
		{"123 - 456", 123, true},
		{"-2147483648", -2147483648, true},
		{"3.1415926", 3, false},
		{"0xFFFF", 0, false},
	}

	for index, test := range tests {
		v, err := ParseInt64(test.Data)
		var msg = fmt.Sprintf("expecting %20d <= [%5v] '%s'", test.Integer, test.IsValid, test.Data)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, err == nil, test.IsValid, msg)
		if err == nil {
			assert.Equal(t, test.Integer, v, msg)
		}
	}
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
