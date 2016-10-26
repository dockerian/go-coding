package utils

import (
	"encoding/json"
	"fmt"
)

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

// ToJSON function returns pretty-printed JSON for a struct.
func ToJSON(t interface{}) string {
	json.MarshalIndent(t, "", "  ")
	jsonBytes, err := json.MarshalIndent(t, "", "  ")

	if err != nil {
		return fmt.Sprintf("{Error: \"%v\"}", err.Error())
	}

	return string(jsonBytes)
}
