// +build all utils integer

package demo

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
