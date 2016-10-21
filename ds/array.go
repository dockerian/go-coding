package ds

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

// InsertIntoSlice inserts the value into the slice at the specified index,
// which must be in range and the slice must have room for the new element.
func InsertIntoSlice(slice []interface{}, index int, value interface{}) []interface{} {
	var newSlice []interface{}

	// Grow the slice by one element.
	if cap(slice) == len(slice) {
		newSlice = make([]interface{}, len(slice)+1)
		copy(slice[0:index], newSlice[0:index])
	} else {
		newSlice = slice[0 : len(slice)+1]
	}

	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(slice[index+1:], newSlice[index:])
	// Store the new value.
	newSlice[index] = value
	// Return the result.
	return newSlice
}

// ReverseArray re-arranges array items reversely
func ReverseArray(slice []interface{}) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
