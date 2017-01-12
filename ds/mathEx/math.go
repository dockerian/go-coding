package mathEx

const (
	// MAXUINT represents maximum uint
	MAXUINT = ^uint(0)
	// MAXINT represents maximum int
	MAXINT = int(^uint(0) >> 1)
	// MININT represents minimum int
	MININT = ^int(^uint(0) >> 1)

	// MAXUINT8 represents maximum 8-bit unsigned integer
	MAXUINT8 uint8 = ^(uint8(0))
	// MAXUINT16 represents maximum 16-bit unsigned integer
	MAXUINT16 uint16 = ^(uint16(0))
	// MAXUINT32 represents maximum 32-bit unsigned integer
	MAXUINT32 uint32 = ^(uint32(0))
	// MAXUINT64 represents maximum 64-bit unsigned integer
	MAXUINT64 uint64 = ^(uint64(0))

	// MAXINT8 represents maximum 8-bit integer
	MAXINT8 int8 = int8(^uint8(0) >> 1) // int(MAXUINT8 >> 1)
	// MININT8 represents minimum 8-bit integer
	MININT8 int8 = ^int8(^uint8(0) >> 1) // ^MAXINT8

	// MAXINT16 represents maximum 16-bit integer
	MAXINT16 int16 = int16(^uint16(0) >> 1) // int(MAXUINT16 >> 1)
	// MININT16 represents minimum 16-bit integer
	MININT16 int16 = ^int16(^uint16(0) >> 1) // ^MAXINT16

	// MAXINT32 represents maximum 32-bit integer
	MAXINT32 int32 = int32(^uint32(0) >> 1) // int(MAXUINT32 >> 1)
	// MININT32 represents minimum 32-bit integer
	MININT32 int32 = ^int32(^uint32(0) >> 1) // ^MAXINT32

	// MAXINT64 represents maximum 64-bit integer
	MAXINT64 int64 = int64(^uint64(0) >> 1) // int64(MAXUINT64 >> 1)
	// MININT64 represents minimum 64-bit integer
	MININT64 int64 = ^int64(^uint64(0) >> 1) // ^MAXINT64
)

// IsSquareRootInteger returns true if square root of num is integer
func IsSquareRootInteger(num int64) bool {
	if num < 0 || num == 0 || num == 1 {
		return num >= 0
	}
	var i int64 = 1
	var half = num / int64(2)
	for ; i < half; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}

// MaxAndMin returns maximum and minimum integers from v ...int
func MaxAndMin(v ...int) (int, int) {
	max, min := MAXINT, MININT
	if len(v) > 0 {
		max, min = min, max
		for _, i := range v {
			if i > max {
				max = i
			}
			if i < min {
				min = i
			}
		}
	}
	return max, min
}

// MaxAndMinInt16 returns maximum and minimum 16-bit integers from v ...int
func MaxAndMinInt16(v ...int16) (int16, int16) {
	max, min := MAXINT16, MININT16
	if len(v) > 0 {
		max, min = min, max
		for _, i := range v {
			if i > max {
				max = i
			}
			if i < min {
				min = i
			}
		}
	}
	return max, min
}

// MaxAndMinInt32 returns maximum and minimum 32-bit integers from v ...int
func MaxAndMinInt32(v ...int32) (int32, int32) {
	max, min := MAXINT32, MININT32
	if len(v) > 0 {
		max, min = min, max
		for _, i := range v {
			if i > max {
				max = i
			}
			if i < min {
				min = i
			}
		}
	}
	return max, min
}

// MaxAndMinInt64 returns maximum and minimum 64-bit integers from v ...int
func MaxAndMinInt64(v ...int64) (int64, int64) {
	max, min := MAXINT64, MININT64
	if len(v) > 0 {
		max, min = min, max
		for _, i := range v {
			if i > max {
				max = i
			}
			if i < min {
				min = i
			}
		}
	}
	return max, min
}

// MaxAndMinInt8 returns maximum and minimum 8-bit integers from v ...int
func MaxAndMinInt8(v ...int8) (int8, int8) {
	max, min := MAXINT8, MININT8
	if len(v) > 0 {
		max, min = min, max
		for _, i := range v {
			if i > max {
				max = i
			}
			if i < min {
				min = i
			}
		}
	}
	return max, min
}

// MaxInt returns maximum int from v ...int
func MaxInt(v ...int) int {
	a, _ := MaxAndMin(v...)
	return a
}

// MinInt returns minimum int from v ...int
func MinInt(v ...int) int {
	_, a := MaxAndMin(v...)
	return a
}
