package utils

import (
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

// GetTestName returns the name of the test function from testing.T.
func GetTestName(t *testing.T) string {
	v := reflect.ValueOf(t)
	testName := v.Elem().FieldByName("name")
	return testName.String()
}

// GetTestNameByCaller returns the name of the test function from the call stack.
func GetTestNameByCaller() string {
	// the one and only identifier after a package specifier
	var testNameRegexp = regexp.MustCompile(`\.(Test[\p{L}_\p{N}]*)$`)

	programCounters := make([]uintptr, 32)
	n := runtime.Callers(0, programCounters)

	for i := 0; i < n; i++ {
		fname := runtime.FuncForPC(programCounters[i]).Name()
		match := testNameRegexp.FindStringSubmatch(fname)
		if match == nil {
			continue
		}
		return match[1]
	}

	return "Unknown"
}
