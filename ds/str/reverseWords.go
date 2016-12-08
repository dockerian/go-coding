package str

import (
	"bytes"
)

// ReverseWords reverses a string by words
func ReverseWords(s string) string {
	strLen := len(s)
	rwList := make([]string, 0, strLen/4)

	var k int
	var endOfWord = true
	// fmt.Printf("s = '%s'\n", s)
	for i, ch := range s {
		if ch == ' ' {
			if !endOfWord {
				rwList = append(rwList, s[k:i])
				endOfWord = true
			}
		} else if endOfWord {
			endOfWord = false
			k = i
		}
	}

	if k > 0 && !endOfWord {
		rwList = append(rwList, s[k:strLen])
	}

	// note: alternatively using strings.Split
	// rwList = strings.Split(s, " ")

	var buffer bytes.Buffer
	for j := len(rwList) - 1; j >= 0; j-- {
		buffer.WriteString(rwList[j])
		if j > 0 {
			buffer.WriteString(" ")
		}
	}

	// note: alternatively using strings.Join
	// return strings.Join(rwList, " ")
	return buffer.String()
}
