// +build all ds math fibo fibonacci test

package mathEx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFibo tests fibonacci functions
func TestFibo(t *testing.T) {
	for n := 0; n < 100; n++ {
		fib, err1 := Fibo(n)
		seq, err2 := Fibos(n)
		var expected uint64
		if n < len(Fibonacci) {
			expected = Fibonacci[n]
		}
		msg := fmt.Sprintf("expecting Fib(%v) = %v", n, expected)
		t.Logf("Test %2d: %v\n", n, msg)
		if err1 != nil {
			assert.Equal(t, "Fibonacci overflow uint64", err1.Error())
			assert.Equal(t, "Fibonacci overflow uint64", err2.Error())
		} else {
			for i := 0; i <= n; i++ {
				assert.Equal(t, Fibonacci[i], seq[i], msg)
			}
			assert.Equal(t, expected, fib, msg)
		}
	}
}
