package interview

import (
	u "github.com/dockerian/go-coding/utils"
)

// FindMostOccurrences func
// For a string of length N, figure out the number of occurrences of
// the most frequent substring of length L in this string. Assuming:
//      2 <= N && N <= 100000
//      2 <= K && K <= L
//      the number of distinct characters must not exceed M <= 26
//      the string contains only lower-case letters (a-z)
func FindMostOccurrences(s string, substrLen int) (string, int) {
	hash := make(map[string]int)
	size := len(s)

	if substrLen >= size {
		if len(s) > 0 {
			return s, 1
		}
		return "", 0
	}

	for i := 0; i <= size; i++ {
		for j := i + 2; j-i <= substrLen && j <= size; j++ {
			// u.Debug("build: %s (%d, %d)\n", s[i:j], i, j)
			str := s[i:j]
			if count, ok := hash[str]; ok {
				hash[str] = count + 1
			} else {
				hash[str] = 1
			}
		}
	}
	u.Debug("substr hash: %+v\n", hash)

	result, counts := "", 0
	for str, count := range hash {
		if count > counts || count >= counts && len(str) > len(result) {
			counts = count
			result = str
		}
	}
	return result, counts
}
