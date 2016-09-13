package ds

// Insert inserts the value into the slice at the specified index,
// which must be in range and the slice must have room for the new element.
func Insert(slice []interface{}, index int, value interface{}) []interface{} {
	// Grow the slice by one element.
	slice = slice[0 : len(slice)+1]
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(slice[index+1:], slice[index:])
	// Store the new value.
	slice[index] = value
	// Return the result.
	return slice
}
