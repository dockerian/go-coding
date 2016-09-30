// +build all ds math test

package ds

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestLogBaseX tests LogBaseX
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
