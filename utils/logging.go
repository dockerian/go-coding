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

// Debugln prints logging message (with new line) if DEBUG is set
func Debugln(a ...interface{}) {
	if DebugEnv {
		fmt.Println(a...)
	}
}

// DebugReset reset to debug mode by environment setting
func DebugReset() {
	DebugEnv = DebugEnvStore
}

// DebugOff turns off debug
func DebugOff() {
	DebugEnv = false
}

// DebugOn turns on debug
func DebugOn() {
	DebugEnv = true
}
