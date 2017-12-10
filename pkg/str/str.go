// Package str :: str.go - extended string functions
package str

import (
	"strings"
)

// Append concatenates byte slices
func Append(slice, data []byte) []byte {
	lenData := len(data)
	lenSlice := len(slice)
	if lenSlice+lenData > cap(slice) { // reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (lenSlice+lenData)*2)
		// copy data, optionally use `bytes.Copy()`
		for idx, item := range slice {
			newSlice[idx] = item
		}
		slice = newSlice
	}
	slice = slice[0 : lenSlice+lenData]
	for idx, item := range data {
		slice[lenSlice+idx] = item
	}
	return slice
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
