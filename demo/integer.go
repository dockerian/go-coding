package demo

import "strings"

const (
	// MaxUint32 represents maximium unsigned integer
	MaxUint32 uint32 = ^(uint32(0))
	// MinUint32 represents minimium unsigned integer
	MinUint32 uint32 = 0
	// MaxUint64 represents maximium 64-bit unsigned integer
	MaxUint64 uint64 = ^(uint64(0))
	// MinUint64 represents minimium unsigned integer
	MinUint64 uint64 = 0
	// MaxInt32 represents maximum integer
	MaxInt32 int32 = int32(^uint32(0) >> 1) // int(MaxUint32 >> 1)
	// MinInt32 represents minimum integer
	MinInt32 int32 = ^int32(^uint32(0) >> 1) // ^MaxInt32
	// MaxInt64 represents 64-bit maximum integer
	MaxInt64 int64 = int64(^uint64(0) >> 1) // int64(MaxUint64 >> 1)
	// MinInt64 represents 64-bit minimum integer
	MinInt64 int64 = ^int64(^uint64(0) >> 1) // ^MaxInt64
)

// Atoi converts string to integer
// Note: 0 for valid decimal number; return max or min on overflow
func Atoi(input string) int64 {
	input = strings.TrimSpace(input)
	var maxint = int64(^uint64(0) >> 1)
	var minint = ^maxint
	var result int64
	var signed int64 = 1

	for i := 0; i < len(input); i++ {
		if i == 0 {
			if input[i] == '-' {
				signed = -1
				continue
			}
			if input[i] == '+' {
				continue
			}
		}
		if input[i] == ',' {
			continue
		}
		if '0' > input[i] || input[i] > '9' {
			break
		}
		bit := int64(input[i] - '0')
		sum := 10*result + bit

		// u.Debug("signed= %v, result= %v, sum= %v, bit= %v\n", signed, result, sum, bit)
		var overflow = signed*result != signed*(sum-bit)/10 ||
			signed == -1 && -sum > -result ||
			signed == 1 && sum < result
		if overflow {
			if signed == 1 {
				return maxint
			}
			return minint
		}

		result = sum
	}

	return signed * result
}

// DecodeInteger decrypts code by length to two integers
func DecodeInteger(code int, length int) (x, y int) {
	x = code / length
	y = code % length
	// fmt.Printf("code(%v), length(%v) => x(%v), y(%v)\n", code, length, x, y)
	return x, y
}

// EncodeIntegers crypts two integers by length
func EncodeIntegers(x, y, length int) int {
	code := x*length + y
	// fmt.Printf("code(%v) <= length(%v), x(%v), y(%v)\n", code, length, x, y)
	return code
}
