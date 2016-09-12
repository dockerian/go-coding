package utils

import (
	"fmt"
)

// Debug prints logging message if DEBUG is set
func Debug(format string, v ...interface{}) {
	if DebugEnv {
		fmt.Printf(format, v...)
	}
}
