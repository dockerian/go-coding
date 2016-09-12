package puzzle

import (
	"fmt"
	"reflect"
	"strconv"
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
	return strconv.FormatUint(p.input, 10)
}

// GetData returns input data
func (p *PalindromeString) GetData() string {
	return p.input
}

// IsPalindrome checks if an input number is palindrome
func (p *PalindromeNumber) IsPalindrome() bool {
	return isPalindromeNumber(p.input)
}

// IsPalindrome checks if an input string is palindrome
func (p *PalindromeString) IsPalindrome() bool {
	return isPalindromeString(p.input)
}

func isPalindrome(input interface{}) (bool, error) {
	v := reflect.ValueOf(input)
	// switch v.Kind() case reflect.int32
	switch input.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		if v.Int() <= 0 {
			return isPalindromeNumber(uint64(-v.Int())), nil
		}
		return isPalindromeNumber(v.Uint()), nil
	case string:
		return isPalindromeString(v.String()), nil
	}
	return false, fmt.Errorf("unsupported type: %v", v.Kind())
}

// isPalindromeNumber checks if an input number is palindrome
func isPalindromeNumber(input uint64) bool {
	var number = input
	var numberReversed uint64
	for number != 0 {
		numberReversed = numberReversed*10 + number%10
		number /= 10
	}
	return input == numberReversed
}

// isPalindromeString checks if an input string is palindrome
func isPalindromeString(input string) bool {
	var size = len(input)
	for i := 0; i < size>>1; i++ {
		if input[i] != input[size-1-i] {
			return false
		}
	}
	return true
}
