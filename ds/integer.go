package ds

import (
	"fmt"
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

// MultiplyUint64 gets muliplication of two integers
func MultiplyUint64(a, b uint64) (uint64, error) {
	m := a * b
	if m/a != b {
		return 0, fmt.Errorf("overflow: muliplication of %v, %v", a, b)
	}
	return m, nil
}

// SumUint64 gets sum of two integers
func SumUint64(a, b uint64) (uint64, error) {
	s := a + b
	if s-a != b {
		return 0, fmt.Errorf("overflow: sum of %v, %v", a, b)
	}
	return s, nil
}

// ReverseInt64 reverses a decimal integer
func ReverseInt64(number int64) int64 {
	var reversed int64
	for number != 0 {
		reversed = reversed*10 + number%10
		number /= 10
	}

	return reversed
}
