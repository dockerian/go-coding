// +build all demo integer test

package demo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// DecodeIntegerTestData struct
type DecodeIntegerTestData struct {
	X      int
	Y      int
	Length int
	Code   int
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
