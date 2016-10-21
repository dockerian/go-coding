// +build all ds str test

package str

import (
	"reflect"
	"runtime"
	"testing"
)

var (
	benchMarkRepeatSize    = 999
	benchMarkSubTestString = "- This !$ a te$t: ……知心幾見曾來往，水隔山遙望眼枯……"
	benchMarkTestString    = getReverseTestString()
)

func getReverseTestString() string {
	return generateTestString(benchMarkSubTestString, benchMarkRepeatSize)
}

// BenchmarkReverse benchmarks on func reverse
func BenchmarkReverse(b *testing.B) {
	b.Logf("Benchmark testing for reserve string functions\n")
	for _, exec := range funcReverses {
		path := runtime.FuncForPC(reflect.ValueOf(exec.Func).Pointer()).Name()
		name := reflect.TypeOf(exec.Func).Name()
		// TODO: check why the name is empty
		b.Logf("- %v [%v]\n", path, name)
		b.Run(exec.Name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				exec.Func(benchMarkTestString)
			}
		})
	}
}
