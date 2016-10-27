package arr

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

// Maps func
func Maps(arr []interface{}, f func(interface{}) interface{}) []interface{} {
	for i, v := range arr {
		arr[i] = f(v)
	}
	return arr
}

// Reduce func
func Reduce(arr []interface{}, f func(prev, curr interface{}) interface{}) interface{} {
	var r interface{}
	for _, v := range arr {
		r = f(r, v)
	}
	return r
}

// Reverse func re-arranges array items reversely
func Reverse(slice []interface{}) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shift func to get the first item in slice and shifted the rest
func Shift(slice []interface{}) (interface{}, []interface{}) {
	elm := slice[0]
	cpy := make([]interface{}, len(slice)-1)
	copy(cpy, slice[1:])
	return elm, cpy
}
