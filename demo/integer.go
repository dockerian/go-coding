package demo

import "strings"

const (
	// MaxUint represents maximum unsigned integer
	MaxUint uint = ^(uint(0))
	// MaxUint8 represents maximum 8-bit unsigned integer
	MaxUint8 uint8 = ^(uint8(0))
	// MaxUint16 represents maximum 16-bit unsigned integer
	MaxUint16 uint16 = ^(uint16(0))
	// MaxUint32 represents maximum 32-bit unsigned integer
	MaxUint32 uint32 = ^(uint32(0))
	// MaxUint64 represents maximum 64-bit unsigned integer
	MaxUint64 uint64 = ^(uint64(0))

	// MaxInt represents maximum integer
	MaxInt int = int(^uint(0) >> 1) // int(MaxUint >> 1)
	// MinInt represents minimum integer
	MinInt int = ^int(^uint(0) >> 1) // ^MaxInt

	// MaxInt8 represents maximum 8-bit integer
	MaxInt8 int8 = int8(^uint8(0) >> 1) // int(MaxUint8 >> 1)
	// MinInt8 represents minimum 8-bit integer
	MinInt8 int8 = ^int8(^uint8(0) >> 1) // ^MaxInt8

	// MaxInt16 represents maximum 16-bit integer
	MaxInt16 int16 = int16(^uint16(0) >> 1) // int(MaxUint16 >> 1)
	// MinInt16 represents minimum 16-bit integer
	MinInt16 int16 = ^int16(^uint16(0) >> 1) // ^MaxInt16

	// MaxInt32 represents maximum 32-bit integer
	MaxInt32 int32 = int32(^uint32(0) >> 1) // int(MaxUint32 >> 1)
	// MinInt32 represents minimum 32-bit integer
	MinInt32 int32 = ^int32(^uint32(0) >> 1) // ^MaxInt32

	// MaxInt64 represents maximum 64-bit integer
	MaxInt64 int64 = int64(^uint64(0) >> 1) // int64(MaxUint64 >> 1)
	// MinInt64 represents minimum 64-bit integer
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
