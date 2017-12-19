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

// TestInsertIntoSlice
func TestInsertIntoSlice(t *testing.T) {
	slice := make([]interface{}, 3, 4)
	for i := range slice {
		t.Logf("slice[%d] = %v\n", i, i)
		slice[i] = i
	}
	slice = InsertIntoSlice(slice, 2, 12)
	slice = InsertIntoSlice(slice, 3, 22)
	t.Logf("result: %+v\n", slice)
	assert.Equal(t, 5, cap(slice))
	assert.Equal(t, 5, len(slice))
	assert.Equal(t, 12, slice[2])
	assert.Equal(t, 22, slice[3])
}

// TestMaps
func TestMaps(t *testing.T) {
	slice := []interface{}{0, 1, 2, 3, 4, 5}
	slice = Maps(slice, func(i interface{}) interface{} {
		if num, ok := i.(int); ok {
			return num * num
		}
		return i
	})
	assert.Equal(t, []interface{}{0, 1, 4, 9, 16, 25}, slice)
}

// TestReduce is a testing function template
func TestReduce(t *testing.T) {
	var tests = []ReduceTestCase{
		{
			[]interface{}{1, 22, 333, 4444, 5555, 1, 5555, 4444, 22}, 20377,
			func(a, b interface{}) interface{} {
				if num1, ok := a.(int); ok {
					if num2, ok := b.(int); ok {
						return num1 + num2
					}
				}
				return b
			},
		},
	}

	for index, test := range tests {
		var val = Reduce(test.Data, test.Func)
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.EqualValues(t, test.Expected, val, msg)
	}
}

// TestReverse
func TestReverse(t *testing.T) {
	slice := []interface{}{0, 1, 2, 3, 4, 5, 6}
	Reverse(slice)
	assert.Equal(t, []interface{}{6, 5, 4, 3, 2, 1, 0}, slice)
}

// TestShift
func TestShift(t *testing.T) {
	slice := []interface{}{0, 1, 2, 3, 4, 5}
	shift, result := Shift(slice)
	assert.Equal(t, []interface{}{1, 2, 3, 4, 5}, result)
	assert.Equal(t, 0, shift)
}
