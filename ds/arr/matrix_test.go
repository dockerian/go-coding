// +build all ds array test

package arr

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ArrayTestCase struct
type CreateMatrixTestCase struct {
	Rows        int
	Columns     int
	Expected    interface{}
	Empty       interface{}
	ElementType reflect.Type
	HasError    bool
}

// TestCreateMatrix tests CreateMatrix
func TestCreateMatrix(t *testing.T) {
	matrix0 := [][]int{
		{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0},
	}
	matrix1 := [][]string{
		{"", "", "", "", ""},
		{"", "", "", "", ""},
	}
	tests := []CreateMatrixTestCase{
		{3, 4, matrix0, make([][]int, 0), reflect.TypeOf(int(0)), false},
		{2, 5, matrix1, make([][]string, 0), reflect.TypeOf(""), false},
	}

	for index, test := range tests {
		val, err := CreateMatrix(test.ElementType, test.Rows, test.Columns)
		msg := fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v.1: %v\n", index+1, msg)
		assert.Equal(t, test.HasError, err != nil, msg)
		assert.EqualValues(t, test.Expected, val, msg)
	}
}
