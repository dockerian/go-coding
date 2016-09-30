package ds

import (
	"math"
)

// LogBaseX caculates base log of x
// Note: Log <base> (x) = Log(x) / Log(base)
func LogBaseX(base int32, x float64) float64 {
	switch base {
	case 2:
		return math.Log2(x)
	case 10:
		return math.Log10(x)
	}
	return math.Log2(x) / math.Log2(float64(base))
}
