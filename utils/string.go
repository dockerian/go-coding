// Package utils :: string.go
package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GetSliceAtIndex returns indexed one-byte slice, or empty string
func GetSliceAtIndex(input string, index int) string {
	if 0 > index || index >= len(input) {
		return ""
	}
	return input[index : index+1]
}

// ReplaceProxyURL searches prefix in url and replaces with proxyURL
func ReplaceProxyURL(url, prefix, proxyURL string) string {
	if prefix == "" || url == "" {
		return url
	}
	prefix = strings.TrimRight(prefix, "/")
	pos := strings.Index(url, prefix)
	posNext := pos + len(prefix)
	okToReplace := pos >= 0 && len(url) > posNext && url[posNext] == '/' || strings.HasSuffix(url, prefix)

	if okToReplace {
		substURL := proxyURL + url[posNext:]
		return substURL
	}

	return url
}

// ShiftSlice returns slice by shift index
func ShiftSlice(input string, shift int) string {
	if len(input) > shift {
		return input[shift:]
	}
	return ""
}

// StringIn check if an input is in an array of strings; optional to ignore case
func StringIn(stringInput string, stringList []string, options ...bool) bool {
	if stringInput == "" || len(stringList) == 0 {
		return false
	}

	ignoreCase := false
	if len(options) > 0 {
		ignoreCase = options[0]
	}
	if ignoreCase {
		stringInput = strings.ToLower(stringInput)
	}
	for _, str := range stringList {
		if str != "" {
			if ignoreCase {
				str = strings.ToLower(str)
			}
			if stringInput == str {
				return true
			}
		}
	}
	return false
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
