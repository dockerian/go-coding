package utils

import (
	"bufio"
)

// Readlines returns lines (without the ending \n) from buffered reader,
// additionally returns error from buffered reader if there is any.
func Readlines(r *bufio.Reader) ([]string, error) {
	var (
		err      error
		isPrefix bool
		line     []byte
		lines    []string
	)
	lines = make([]string, 0)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		lines = append(lines, string(line))
	}
	return lines, err
}
