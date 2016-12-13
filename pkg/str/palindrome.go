package str

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode"

	u "github.com/dockerian/go-coding/utils"
)

// Palindrome interface
type Palindrome interface {
	GetData() string
	IsPalindrome() bool
}

// PalindromeNumber struct
type PalindromeNumber struct {
	input uint64
}

// PalindromeString struct
type PalindromeString struct {
	input string
}

// GetData returns input data
func (p *PalindromeNumber) GetData() string {
	if p == nil {
		return ""
	}
	return strconv.FormatUint(p.input, 10)
}

// GetData returns input data
func (p *PalindromeString) GetData() string {
	if p == nil {
		return ""
	}
	return p.input
}

// GetSubstring returns the longest palindromic substring
func (p *PalindromeString) GetSubstring() string {
	if p == nil {
		return ""
	}
	return GetPalindromicSubstring(p.input)
}

// GetPalindromicSubstring returns the longest palindromic substring
func GetPalindromicSubstring(str string) string {
	if len(str) <= 0 {
		return ""
	}

	var end = len(str) - 1
	var x, y int

	for k := 0; k < end; k++ {
		var i, j int
		var m, n = k + 2, k + 1
		var continued bool

		if m < end && str[k] == str[m] {
			continued = str[k] == str[n]
			i, j = k, m
		}
		if n < end && str[k] == str[n] {
			continued = true
			i, j = k, n
		}

		if j > 0 {
			if continued {
				for i > 0 && str[i-1] == str[k] {
					i--
				}
				for j < end && str[k] == str[j+1] {
					j++
				}
			}

			for i > 0 && j < end && str[i-1] == str[j+1] {
				i--
				j++
			}

			if j-i > y-x {
				x, y = i, j
			}
		}

		if y-x > 2*(end-k) {
			u.Debug("break: x = %v, y = %v ['%v']; k = %v, end = %v in '%v'\n",
				x, y, str[x:y+1], k, end, str)
			break
		}
	}

	return str[x : y+1]
}

// IsPalindrome checks if an input number is palindrome
func (p *PalindromeNumber) IsPalindrome() bool {
	if p == nil {
		return false
	}
	return IsPalindromeNumber(p.input)
}

// IsPalindrome checks if an input string is palindrome
func (p *PalindromeString) IsPalindrome() bool {
	if p == nil {
		return false
	}
	return IsPalindromeString(p.input)
}

// IsPalindromePhase checks if an input string is palindrome phase
func (p *PalindromeString) IsPalindromePhase() bool {
	if p == nil {
		return false
	}
	return IsPalindromePhase(p.input)
}

// IsPalindrome checks if interface{} is palindrome
func IsPalindrome(input interface{}) (bool, error) {
	v := reflect.ValueOf(input)
	// switch v.Kind() case reflect.int32
	switch input.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		if v.Int() <= 0 {
			return IsPalindromeNumber(uint64(-v.Int())), nil
		}
		return IsPalindromeNumber(v.Uint()), nil
	case string:
		return IsPalindromeString(v.String()), nil
	}
	return false, fmt.Errorf("unsupported type: %v", v.Kind())
}

// IsPalindromeNumber checks if an input number is palindrome
func IsPalindromeNumber(input uint64) bool {
	var number = input
	var numberReversed uint64
	for number != 0 {
		numberReversed = numberReversed*10 + number%10
		number /= 10
	}
	return input == numberReversed
}

// IsPalindromePhase checks if an input string is palindrome phase
func IsPalindromePhase(input string) bool {
	if len(input) == 0 {
		return false
	}
	var runes = []rune{}

	for _, char := range input {
		runes = append(runes, char)
	}
	//u.Debug("input = %v, runes = %+v\n", input, runes)

	var valid = false
	var bound = len(runes) - 1
	for i, j := 0, bound; i <= bound && j >= i; i++ {
		for i <= bound && !u.IsDigitOrLetter(runes[i]) {
			i++
		}
		//u.Debug("runes[%v] = %c, runes[%v] = %c, bound = %v\n", i, runes[i], j, runes[j], bound)
		for j > 1 && !u.IsDigitOrLetter(runes[j]) {
			j--
		}
		//u.Debug("runes[%v] = %c, runes[%v] = %c, bound = %v\n", i, runes[i], j, runes[j], bound)

		if i <= bound && j > i {
			if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
				u.Debug("break: runes[%v] = %c != runes[%v] = %c, input = %v\n",
					i, runes[i], j, runes[j], input)
				return false
			}
			valid = true
		}
		j--
	}
	return valid
}

// IsPalindromeString checks if an input string is palindrome
func IsPalindromeString(input string) bool {
	var size = len(input)
	for i := 0; i < size>>1; i++ {
		if input[i] != input[size-1-i] {
			return false
		}
	}
	return true
}
