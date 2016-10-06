package utils

/*
 * bit.go - some collection of bitwise operations
 * see more @
 *	- https://en.wikipedia.org/wiki/Bitwise_operation
 *  - https://en.wikipedia.org/wiki/Bitwise_operations_in_C
 *  - http://www.cprogramming.com/tutorial/bitwise_operators.html
 *  - https://discuss.leetcode.com/topic/50315/
 */

import (
	"fmt"
)

var (
	binaryFormat = "%s%s%s%s%s%s%s%s"
	// BinaryString is an array of all 4-bit binary representation
	BinaryString = map[string][]string{
		"b": []string{
			"0000", "0001", "0010", "0011",
			"0100", "0101", "0110", "0111",
			"1000", "1001", "1010", "1011",
			"1100", "1101", "1110", "1111",
		},
		"x": []string{
			"0", "1", "2", "3", "4", "5", "6", "7",
			"8", "9", "A", "B", "C", "D", "E", "F",
		},
	}
)

// BitAllOne returns integer with all bits are 1
func BitAllOne() int64 {
	return ^0
}

// BitCheck checks on nth bit of x
func BitCheck(x int64, n uint8) bool {
	return (x & 1 << n) != 0
}

// BitClear sets 0 on nth bit of x
func BitClear(x int64, n uint8) int64 {
	return x & ^(1 << n)
}

// BitCountOne returns number of 1 in x (aka Hamming weight)
func BitCountOne(x int64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

// BitCountOneUint64 returns number of 1 in x (aka Hamming weight)
func BitCountOneUint64(x uint64) int {
	var count int
	var mask uint64 = 1
	for i := 0; i < 64; i++ {
		if mask&x != 0 {
			count++
		}
		mask = mask << 1
	}
	return count
}

// BitIntersection applies bitwise AND (&) operator on a and b (interaction)
func BitIntersection(a, b int64) int64 {
	return a & b
}

// BitInverse returns inverted x
func BitInverse(x int64) int64 {
	return ^x
}

// BitIsPowerOf2 checks if the number is power of 2
func BitIsPowerOf2(number int64) bool {
	return number > 0 && 0 == number&(number-1)
}

// BitIsPowerOf4 checks if the number is power of 4
func BitIsPowerOf4(number int64) bool {
	return number > 0 && 0 == number&(number-1) && 0 == number&0x5555555555555555
}

// BitNegativeInt returns negative number of x (0 - x)
func BitNegativeInt(x int64) int64 {
	var y int64 = 1
	return BitSumInt64(^x, y)
}

// BitSet sets 1 on nth bit of x
func BitSet(x int64, n uint8) int64 {
	return x | 1<<n
}

// BitString converts uint64 x to zero-padding bits representation
func BitString(x uint64, key, delimiter string) string {
	var mask uint64 = 0xF
	var sFmt = binaryFormat + binaryFormat
	var xbin = key == "b"
	var xpad = delimiter != ""
	var xmap []string
	var xval string
	var okay bool

	if xpad {
		sFmt += sFmt
	}
	if xmap, okay = BinaryString[key]; !okay {
		return ""
	}

	Debug("\nx= %d [0x%x / %b], xmap= %+v\n", x, x, x, xmap)

	bnum := len(xmap)
	data := make([]interface{}, 0, bnum*2)

	for n := 1; n <= bnum; n++ {
		ndx := int(mask & (x >> uint(4*(bnum-n))))
		Debug("n= %2d, shift= %2d, ndx= %2d [%b]\n", n, 4*(bnum-n), ndx, ndx)
		data = append(data, xmap[ndx])
		if xpad {
			if n < bnum && (xbin || n%2 == 0) {
				data = append(data, delimiter)
			} else {
				data = append(data, "")
			}
		}
	}

	Debug("%+v\n", data)
	xval = fmt.Sprintf(sFmt, data...)

	return xval
}

// BitSubstraction applies A & ~B
func BitSubstraction(a, b int64) int64 {
	return a & ^b
}

// BitSumInt64 calculates sum of two integers without using arithmetic operators
func BitSumInt64(x, y int64) int64 {
	if y != 0 {
		return BitSumInt64(x^y, (x&y)<<1)
	}
	return x
}

// BitSumInt calculates sum of two integers without using arithmetic operators
func BitSumInt(x, y int) int {
	// Iterate till there is no carry
	for y != 0 {
		// carry now contains common set bits of x and y
		carry := x & y
		// XOR on bits of x and y where at least one of the bits is not set
		x = x ^ y
		// carry is shifted by one so that adding it to x gives the required sum
		y = carry << 1
	}
	return x
}

// BitUnion applies bitwise OR (|) operator on a and b (union)
func BitUnion(a, b int64) int64 {
	return a | b
}

// ToBinaryString converts uint64 x to zero-padding binary representation
func ToBinaryString(x uint64, delimiter string) string {
	return BitString(x, "b", delimiter)
}

// ToHexString converts uint64 x to zero-padding hexidecimal representation
func ToHexString(x uint64, delimiter string) string {
	return BitString(x, "x", delimiter)
}
