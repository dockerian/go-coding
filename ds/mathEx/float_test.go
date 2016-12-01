// +build all ds math test

package mathEx

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// RandRangeFloat64TestCase struct
type RandRangeFloat64TestCase struct {
	min, max float64
}

// RandRangeInt64TestCase
type RandRangeInt64TestCase struct {
	min, max int64
}

// TestLogBaseX tests LogBaseX functions
func TestLogBaseX(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		b := r.Int31n(100) + 2
		x := r.Float64() * 100
		if math.Abs(x) < 0.01 {
			x = float64(r.Int31n(100) + 50)
		}
		result := LogBaseX(b, x)
		target := math.Pow(float64(b), result)
		diff := math.Abs(target - x)
		test := diff < 1.0e-13
		msg1 := fmt.Sprintf("Diff = math.Abs(%v - %v) == %v", result, x, diff)
		msg2 := fmt.Sprintf("\t\t:  math.Pow(%v, %v) == %v", b, result, target)
		msg3 := fmt.Sprintf("\t\t:  LogBaseX(%v, %v) == %v", b, x, result)
		message := fmt.Sprintf("%v\n%v\n%v", msg1, msg2, msg3)
		t.Logf("Test %2d: %v\n\n", i, message)
		assert.True(t, test, message)
	}
}

// TestRandRange tests both RandRangeFloat64 and RandRangeInt64 functions
func TestRandRange(t *testing.T) {
	tests := []RandRangeFloat64TestCase{
		{0.0, 0.0},
		{0.0, 0.1},
		{2.0, 100.5},
		{-3.3, 10.1},
		{-3.0, -10.5},
		{99.9, -10.7},
		{99.9, 99.9},
		{123.456789, 98.7654321},
		{777.777, 0},
	}

	for index, test := range tests {
		times := 100
		max, maxF, maxN := test.max, test.max, int64(test.max)
		min, minF, minN := test.min, test.min, int64(test.min)

		if min > max {
			max, min = min, max
		}
		var rangeF = fmt.Sprintf("range [%v, %v)", min, max)
		var rangeN = fmt.Sprintf("range [%v, %v)", int64(min), int64(max))

		if max == min {
			times = 3
		}

		for n := 0; n < times; n++ {
			var resultF = RandRangeFloat64(minF, maxF)
			var msgF = fmt.Sprintf("expecting %v in %v", resultF, rangeF)
			t.Logf("Test %v.%2d [F]: %v\n", index+1, n, msgF)
			if max != min {
				assert.True(t, min <= resultF && resultF < max, msgF)
			} else {
				assert.Equal(t, min, resultF, msgF)
			}

			var resultN = RandRangeInt64(minN, maxN)
			var msgN = fmt.Sprintf("expecting %v in %v", resultN, rangeN)
			t.Logf("Test %v.%2d [N]: %v\n", index+1, n, msgN)
			if int64(max) != int64(min) {
				assert.True(t, int64(min) <= resultN && resultN < int64(max), msgN)
			} else {
				assert.Equal(t, int64(min), resultN, msgN)
			}
		}
		t.Log("")
	}
}

// TestSqrt tests Sqrt function
func TestSqrt(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		rnd := r.Float64() * 100
		exp := math.Sqrt(rnd)
		val := Sqrt(rnd)
		msg := fmt.Sprintf("expecting Sqrt(%v): %v ~= %v", rnd, val, exp)
		t.Logf("Test %2d: %v\n", i+1, msg)
		assert.True(t, NearEqual(exp, val, Precision), msg)
	}
}
