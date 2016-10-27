package arr

import (
	"fmt"
	"reflect"
)

// CreateMatrix makes an array of array by rows * columns
func CreateMatrix(elementType reflect.Type, rows, columns int) (interface{}, error) {
	if rows <= 0 || columns <= 0 {
		return nil, fmt.Errorf("rows (%v) and columns (%v) must be larger than 0", rows, columns)
	}

	size := columns * rows
	slices := reflect.MakeSlice(reflect.SliceOf(elementType), size, size)
	matrix := reflect.MakeSlice(reflect.SliceOf(slices.Type()), rows, rows)

	for i := 0; i < rows; i++ {
		index := i * columns
		// u.Debug("i= %v: slice @ index= %v, end=%v\n", i, index, index+columns)
		matrix.Index(i).Set(slices.Slice(index, index+columns))
	}

	return matrix.Interface(), nil
}
