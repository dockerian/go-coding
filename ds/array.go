package ds

// Insert inserts the value into the slice at the specified index,
// which must be in range and the slice must have room for the new element.
func Insert(slice []interface{}, index int, value interface{}) []interface{} {
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
