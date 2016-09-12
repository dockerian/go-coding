package ds

// ReverseInt64 reverses a decimal integer
func ReverseInt64(number int64) int64 {
	var reversed int64
	for number != 0 {
		reversed = reversed*10 + number%10
		number /= 10
	}

	return reversed
}
