package utils

import (
	"fmt"
)

// Debugln prints logging message (with new line) if DEBUG is set
func Debugln(a ...interface{}) {
	if DebugEnv {
		fmt.Println(a...)
	}
}

// Debug prints logging message if DEBUG is set
func Debug(format string, v ...interface{}) {
	if DebugEnv {
		fmt.Printf(format, v...)
	}
}
