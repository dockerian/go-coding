package utils

// GetSliceAtIndex returns indexed one-byte slice, or empty string
func GetSliceAtIndex(input string, index int) string {
	if 0 > index || index >= len(input) {
		return ""
	}
	return input[index : index+1]
}

// ShiftSlice returns slice by shift index
func ShiftSlice(input string, shift int) string {
	if len(input) > shift {
		return input[shift:]
	}
	return ""
}
