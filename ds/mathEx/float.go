package mathEx

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const (
	// Precision set precision for math functions
	Precision = 0.00000000000001
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

// NearEqual compares a and b in range of precision
func NearEqual(a, b, precision float64) bool {
	diff := math.Abs(a - b)
	return diff < precision
}

// RandRangeFloat64 returns a pseudo-random float64 in [min,max)
func RandRangeFloat64(min, max float64) float64 {
	if max < min {
		max, min = min, max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float64()*(max-min) + min
}

// RandRangeInt64 returns a pseudo-random int64 in [min,max)
func RandRangeInt64(min, max int64) int64 {
	if max < min {
		max, min = min, max
	}
	if max-min == 0 {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min) + min
}

// Sqrt calculates square root of x (to demonstrate float point comparison)
func Sqrt(x float64) float64 {
	switch {
	case x == 0 || math.IsNaN(x) || math.IsInf(x, 1):
		return x
	case x < 0:
		return math.NaN()
	}

	const BEGIN, PRECISION = 1.0, Precision
	var result = BEGIN

	for {
		temp := result - ((result*result - x) / (2 * result))
		if diff := math.Abs(result - temp); diff < PRECISION {
			break
		} else {
			result = temp
		}
	}

	resultString := fmt.Sprintf("%.20f", result)
	result, _ = strconv.ParseFloat(resultString, 20)

	return result
}
