// Package utils :: conv.go - extended string formatter functions
package utils

import (
	"strings"
	"unicode"
)

// ToCamel converts a string to camel case format
func ToCamel(in string, keepAllCaps ...bool) string {
	var keepAllCap bool
	if len(keepAllCaps) > 0 {
		keepAllCap = keepAllCaps[0]
	}
	splits := strings.FieldsFunc(in, func(r rune) bool {
		return unicode.IsSpace(r) || r == '_' || r == '-' || r == '.'
	})

	var strValues []string
	for _, str := range splits {
		if keepAllCap && str == strings.ToUpper(str) {
			strValues = append(strValues, str)
		} else {
			strValues = append(strValues, strings.Title(strings.ToLower(str)))
		}
	}

	return strings.Join(strValues, "")
}

// ToSnake converts a string to snake case format with unicode support
// See also https://github.com/serenize/snaker/blob/master/snaker.go
func ToSnake(in string) string {
	runes := []rune(in) // using rune to support unicode
	limit := len(runes) // getting the boundary of runes

	var runesOut []rune
	for idx, char := range runes {
		if idx > 0 && unicode.IsUpper(char) && runes[idx-1] != '_' {
			// for any upper case after the first char, inserting '_' if the
			// char before or after is in lower case
			if idx+1 < limit && unicode.IsLower(runes[idx+1]) || unicode.IsLower(runes[idx-1]) {
				runesOut = append(runesOut, '_')
			}
		}
		runesOut = append(runesOut, unicode.ToLower(char))
	}

	return string(runesOut)
}
