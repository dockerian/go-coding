// +build all demo integer test

package demo

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AtoiTestData struct
type AtoiTestData struct {
	inputs string
	result int64
}

type ConstTestData struct {
	data     interface{}
	expected interface{}
}

// DecodeIntegerTestData struct
type DecodeIntegerTestData struct {
	X      int
	Y      int
	Length int
	Code   int
}

// TestAtoi tests Atoi function
// Note: integer range
//       int16: -32,768 to 32,767
//       int32: -2,147,483,648 to 2,147,483,647
//       int64: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
// Note: maximium int64 + 1 => minimium int64 (negative)
//       minimium int64 - 1 => maximium int64 (positive)
func TestAtoi(t *testing.T) {
	testData := []AtoiTestData{
		{"", 0},
		{"-1", -1},
		{"12345", 12345},
		{"-1234", -1234},
		{"  -0012a42", -12},
		{"    10522545459", 10522545459},
		{"a", 0},
		{"2147483648", 2147483648},
		{"-9,223,372,036,854,775,808", -9223372036854775808},
		{"9223372036854775807", 9223372036854775807},
		{"9223372036854775808", 9223372036854775807},
	}

	for index, test := range testData {
		var val = Atoi(test.inputs)
		var msg = fmt.Sprintf("expecting '%v' to '%v'", test.inputs, test.result)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.result, val, msg)
	}
}

// TestDecodeEncodeIntegers tests decoding and encoding integers
func TestDecodeEncodeIntegers(t *testing.T) {
	testData := []DecodeIntegerTestData{
		{78, 11, 33, 2585},
		{116, 4, 7, 816},
		{128, 4, 5, 644},
	}

	for index, v := range testData {
		t.Logf("\nTest(%v):\n", index)
		msg1 := fmt.Sprintf("code(%v) <= length(%v), x(%v), y(%v)",
			v.Code, v.Length, v.X, v.Y)
		t.Logf("\tExpected: %v\n", msg1)
		code := EncodeIntegers(v.X, v.Y, v.Length)
		assert.Equal(t, v.Code, code, msg1)

		msg2 := fmt.Sprintf("code(%v), length(%v) => x(%v), y(%v)",
			v.Code, v.Length, v.X, v.Y)
		t.Logf("\tExpected: %v\n", msg2)
		x, y := DecodeInteger(v.Code, v.Length)
		assert.Equal(t, v.X, x, msg2)
		assert.Equal(t, v.Y, y, msg2)
	}
}

// TestIntegerConstants tests integer constants
// See https://en.wikipedia.org/wiki/Two%27s_complement
func TestIntegerConstants(t *testing.T) {
	// minInt32 represents minimum integer
	var minInt32 int32 = -MaxInt32 - 1
	var minInt64 int64 = -MaxInt64 - 1
	var testData = []ConstTestData{
		{^MaxInt32, minInt32},
		{^MaxInt64, minInt64},
		{MaxInt32, math.MaxInt32},
		{MaxInt64, math.MaxInt64},
		{MinInt32, math.MinInt32},
		{MinInt64, math.MinInt64},
	}
	for index, test := range testData {
		var msg = fmt.Sprintf("expecting %v == %v", test.data, test.expected)
		t.Logf("Test %2d: %v\n", index, msg)
		assert.EqualValues(t, test.expected, test.data, msg)
	}
}
