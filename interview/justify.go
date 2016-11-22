// +build all interview map justify string test

package interview

//----------------------------
// file: justify.go
//----------------------------
/*
  In text processing, one approach to making a document appealing to the eye is
  to justify paragraph text. With monospaced fonts in a terminal, as you might
  see in a man page, this is accomplished by inserting spaces between words to
  produce a line of text where the first character of that line begins on the
  left-hand margin and the last printable character is on the right-hand margin.
  For this problem, use a single line of text as input, and justify that text
  into a buffer, where the first character of the line of text is in the first
  spot in the buffer and the last character of text is in the specified slot
  in the buffer. https://play.golang.org/p/pVoCF3EYuz
*/

import "strings"

// justify line to specified length with padding spaces between words
// note: if the input has only one word, padding the spaces at the end;
//       if the input is empty or line is longer than specified length,
//       return the input without processing (no wrapping).
func justify(input string, length int) string {
	line := strings.Trim(input, " ")
	size := len(line)
	diffSize := length - size
	words := strings.Split(line, " ")
	countSpaces := len(words) - 1
	var pad, extraPad int

	if size == 0 || size > length {
		// TODO: wrap line if countSpaces > 0 (and length is wider than single word)
		return line
	}

	if countSpaces > 0 {
		pad = diffSize / countSpaces
		extraPad = diffSize % countSpaces
	}

	buffer := make([]byte, length)
	k := 0

	for i := 0; i < size; i++ {
		if line[i] != ' ' {
			buffer[k] = line[i]
		} else {
			padSize := pad
			if extraPad > 0 {
				padSize++
				extraPad--
			}
			for padSize != 0 {
				buffer[k] = ' '
				padSize--
				k++
			}
			buffer[k] = ' '
		}
		k++
	}

	for k < length {
		buffer[k] = ' '
		k++
	}

	return string(buffer)
}
