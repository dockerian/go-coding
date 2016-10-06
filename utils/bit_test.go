// +build all utils bit test

package utils

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type BitCountOneTestCase struct {
	xInt64 int64
	counts int
}

func TestBitCountOne(t *testing.T) {
	tests := []BitCountOneTestCase{
		{int64(0x0), 0},
		{int64(0x1), 1},
		{int64(0x7), 3},
		{int64(0x777), 9},
		{int64(0x1234567890), 15},
		{int64(1234567890), 12},
		{math.MaxInt64, 63},
		{math.MinInt64, 1},
	}

	for index, test := range tests {
		var uuu = uint64(test.xInt64)
		var bbb = ToBinaryString(uuu, ",")
		var hex = ToHexString(uuu, "|")
		var num = BitCountOneUint64(uuu)
		var val = BitCountOne(test.xInt64)
		var msg = fmt.Sprintf("expecting %2d (1)s in %s (%s)", val, hex, bbb)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.counts, val, msg)
		assert.Equal(t, num, val, msg)
	}
}

// TestBitSumIntegers tests BitSumIntegers func
func TestBitSumIntegers(t *testing.T) {
	// Create and seed the generator.
	// Typically to use a non-fixed seed, e.g. time.Now().UnixNano()
	// Using a fixed seed will produce the same output on every run
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 1; i < 1000; i++ {
		a := r.Int()
		b := r.Int()
		c := BitSumInt(a, b)
		var msg = fmt.Sprintf("expecting %d + %d == %d", a, b, c)
		// t.Logf("Test %3d: %v\n", i, msg)
		assert.Equal(t, a+b, c, msg)

		x := r.Int63()
		y := r.Int63()
		s := BitSumInt64(x, y)
		assert.Equal(t, x+y, s)
	}
}
