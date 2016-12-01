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

// MaxInt returns maximum int from v ...int
func MaxInt(v ...int) int {
	if len(v) <= 0 {
		return MAXINT
	}
	a := MININT
	for _, i := range v {
		if i > a {
			a = i
		}
	}
	return a
}

// MinInt returns minimum int from v ...int
func MinInt(v ...int) int {
	if len(v) <= 0 {
		return MININT
	}
	a := MAXINT
	for _, i := range v {
		if i < a {
			a = i
		}
	}
	return a
}
