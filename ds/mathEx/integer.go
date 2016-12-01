package mathEx

import (
	"fmt"
	"math"

	u "github.com/dockerian/go-coding/utils"
)

// Compare returns 0 if a == b; 1 if a > b; or -1 if a < b
func Compare(a, b uint64) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

// EqualSign checks if a and b are both same signed (positive or negative)
func EqualSign(a, b int64) bool {
	return a >= 0 && b >= 0 || a < 0 && b < 0
}

// MultiplyUint64 gets muliplication of two integers
func MultiplyUint64(a, b uint64) (uint64, error) {
	m := a * b
	if m/a != b {
		return 0, fmt.Errorf("overflow: muliplication of %v, %v", a, b)
	}
	return m, nil
}

// ParseInt64 converts a string to integer, assuming base-10
func ParseInt64(s string) (int64, error) {
	var init = true
	var icut int64 = math.MaxInt64/10 + 1
	var imod int64 = math.MaxInt64 % 10
	var ival int64
	var sign int64 = 1
	var size = len(s)
	var perr error

	for i := 0; i < size; i++ {
		if s[i] == '+' {
			init = false
			continue
		} else if s[i] == '-' {
			init = false
			sign = -1
			continue
		}
		if s[i] == ' ' || s[i] == '+' || s[i] == '-' {
			if init {
				continue
			}
			break
		}
		if '0' <= s[i] && s[i] <= '9' {
			digit := int64(s[i] - '0')
			n := ival*10 + digit
			if ival >= icut ||
				(ival == icut-1) &&
					(sign == -1 && digit > imod+1 || sign == 1 && digit > imod) {
				perr = fmt.Errorf("Overflow: %s [range: %d, %d]",
					s, math.MinInt64, math.MaxInt64)
				return ival, perr
			}
			init = false
			ival = n
		} else {
			perr = fmt.Errorf("Invalid digit: '%c' in '%s'", s[i], s)
			return ival, perr
		}
	}

	return sign * ival, perr
}

// ReverseInt64 reverses a decimal integer
func ReverseInt64(number int64) int64 {
	var result int64
	for number != 0 {
		var unit = number % 10
		var test = result*10 + unit
		// u.Debug("result= %v, test= %v, unit= %v\n", result, test, unit)
		// Note: maximum int64 + 1 => minimium int64 (negative)
		//       minimium int64 - 1 => maximum int64 (positive)
		if (test-unit)/10 != result || result != 0 && !EqualSign(test, result) {
			return 0
		}
		// u.Debug("result= %v, test= %v, number= %v\n", result, test, number)
		result = test
		number /= 10
	}

	u.Debug("number= %v, result= %v\n\n", number, result)
	return result
}

// SumUint64 gets sum of two integers
func SumUint64(a, b uint64) (uint64, error) {
	s := a + b
	if s-a != b {
		return 0, fmt.Errorf("overflow: sum of %v, %v", a, b)
	}
	return s, nil
}

// SwapInt swaps two integers
func SwapInt(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return b, a
}
