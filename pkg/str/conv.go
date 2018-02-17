// Package str :: conv.go - extended string formatter functions
package str

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

var (
	// translateNumberTo maps a translate function by language.
	translateNumberTo = map[string]TranslateFunc{
		"default": FormatNumber,
		"en":      FromNumber,
	}
)

// TranslateFunc defines a type of function to translate number (uint64) to string.
type TranslateFunc func(uint64) string

// FormatNumber returns a comma delimited decimal string
func FormatNumber(number uint64) string {
	if number == 0 {
		return "0"
	}

	var str string
	for num, exp := number, 18; exp >= 0 && num >= 0; {
		// 10 to the power of exp (10^exp) by every 1,000 (exp % 3 == 0)
		pow := uint64(math.Pow(float64(10), float64(exp)))
		if division := num / pow; str != "" || division > 0 {
			if str != "" {
				str += fmt.Sprintf("%03d,", division)
			} else {
				str += fmt.Sprintf("%d,", division)
			}
		}
		num %= pow
		exp -= 3
	}
	return strings.Trim(str, ",")
}

// FromNumber returns an English words representation for a number.
// ex. 1024 => "one thousand twenty four"
func FromNumber(number uint64) string {
	if number == 0 {
		return "zero"
	}

	// toPluralUnits returns unit string with postfix 's' if x > 1
	toPluralUnits := func(x uint64, unit string) string {
		if x > 1 {
			return unit + "s"
		}
		return unit
	}
	// toWord converts a num (<1000) to words
	toWord := func(num uint64) string {
		var str string
		var mis = []string{
			"zero",
			"one", "two", "three", "four", "five",
			"six", "seven", "eight", "nine", "ten",
			"eleven", "twleve", "thirteen", "fourteen", "fifteen",
			"sixteen", "seventeen", "eighteen", "nineteen",
		}
		var nty = []string{
			"zero", "ten", "twenty", "thirty", "forty",
			"fifty", "sixty", "seventy", "eighty", "ninety",
		}

		if num >= 100 {
			x := num / 100 // between [0, 10]
			str += mis[x] + toPluralUnits(x, " hundred")
			num %= 100
		}
		if num > 19 { // num < 100
			x := num / 10 // between [1, 9]
			str += " " + nty[x]
			if reminder := num % 10; reminder > 0 {
				str += " " + mis[reminder]
			}
		} else if num > 0 { // between [1, 19]
			str += " " + mis[num]
		}
		return strings.Trim(str, " ")
	}

	var scales = []string{
		"zero", "thousand", "million", "billion", "trillion",
		"quadrillion", "quintillion", "sextillion", "septillion",
	}
	var str string

	// initialize calculated num and exponent
	for num, exp := number, 18; exp >= 0 && num > 0; {
		// 10 to the power of exp (10^exp) by every 1,000 (exp % 3 == 0)
		pow := uint64(math.Pow(float64(10), float64(exp)))
		if division := num / pow; division > 0 {
			str += " " + toWord(division)
			if x := exp / 3; x > 0 {
				str += toPluralUnits(division, " "+scales[x])
			}
		}
		num %= pow
		exp -= 3
	}
	return strings.Trim(str, " ")
}

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

// TranslateNumber translates a number to string by specific function.
func TranslateNumber(number uint64, xFunc TranslateFunc) string {
	if xFunc == nil {
		return FormatNumber(number)
	}
	return xFunc(number)
}

// TranslateTo returns a string representation of number by specific language.
func TranslateTo(lang string, number uint64) string {
	if xFunc, ok := translateNumberTo[lang]; ok {
		return xFunc(number)
	}
	return FormatNumber(number)
}
