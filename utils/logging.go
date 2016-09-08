// +build all utils logging

package utils

import (
	"fmt"
)

// Debug prints logging message if DEBUG is set
func Debug(format string, v ...interface{}) {
	if HasEnv("DEBUG", true) {
		fmt.Printf(format, v...)
	}
}
