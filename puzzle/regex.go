package puzzle

import (
	u "github.com/dockerian/go-coding/utils"
)

// BasicRegex struct
type BasicRegex struct {
	input string
	regex string
}

// NewBasicRegex returns a pointer to a new BasicRegex object
func NewBasicRegex(input, regex string) *BasicRegex {
	return &BasicRegex{input: input, regex: regex}
}

// IsMatch tests if the input matches regex
func (br *BasicRegex) IsMatch() bool {
	input := br.input
	regex := br.regex
	return isMatchBasicRegexDP(&input, &regex)
}

// isMatchBasicRegexDP uses dynamic programming
// Note: dp[i][j] value depends on the following conditions
// 1, If p[j] == s[i] || p[j] == '.'
//          dp[i][j] = dp[i-1][j-1] // if matched previously
// 2, If p[j] == '*' (two sub conditions)
//    1) p[j-1] != s[i]
//          dp[i][j] = dp[i][j-2]  // 'a*' only counts as empty
//    2) p[j-1] == s[i] || p[j-1] == '.'
//          dp[i][j] = dp[i-1][j]  // 'a*' counts as multiple 'a'
//       or dp[i][j] = dp[i][j-1]  // 'a*' counts as single 'a'
//       or dp[i][j] = dp[i][j-2]  // 'a*' counts as empty
func isMatchBasicRegexDP(input, regex *string) bool {
	if input == nil || regex == nil {
		return false
	}
	s, p, sLen, pLen := *input, *regex, len(*input), len(*regex)

	if pLen == 0 {
		return sLen == 0
	}

	dp := make([][]bool, sLen+1)
	for i := range dp {
		dp[i] = make([]bool, pLen+1)
	}

	// u.Debug("\ninput: s = '%v', p = '%v'\n", s, p)
	dp[0][0] = true
	for j := 2; j <= pLen; j++ {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			}
			if j >= 2 && p[j-1] == '*' {
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i][j-1] || dp[i-1][j] || dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}

	return dp[sLen][pLen]
}

// isMatchBasicRegexNR uses no recursive call
func isMatchBasicRegexNR(input, regex *string) bool {
	var prevMatch byte = 0xFF
	s, p := *input, *regex
	sLen, pLen := len(s), len(p)
	j, k := 0, 0

	u.Debug("\ninput: s = '%v', p = '%v'\n", s, p)

	for pLen-k > 0 {
		u.Debug("  ---: s[j] = '%v'[%v], p[k] = '%v'[%v], prevMatch = %c\n",
			u.GetSliceAtIndex(s, j), j, u.GetSliceAtIndex(p, k), k, prevMatch)
		if sLen-j == 0 {
			u.Debug("  -*-: s[j] = '%v'[%v], p[k] = '%v'[%v], prevMatch = %c\n",
				u.GetSliceAtIndex(s, j), j, u.GetSliceAtIndex(p, k), k, prevMatch)
			if pLen-k == 1 || p[k+1] != '*' && p[k+1] != '.' {
				return p[k] == prevMatch || p[k] == '.'
			}
			if p[k+1] == '.' {
				return false
			}
			k += 2
		} else if pLen-k < 2 || p[k+1] != '*' {
			if sLen-j == 0 || s[j] != p[k] && p[k] != '.' {
				// the first chars do not match
				return false
			}
			// both moving to next char
			j++
			k++
			prevMatch = 0xFF
		} else {
			// repeat until matched all chars with *
			for sLen-j > 0 && (s[j] == p[k] || s[j] == prevMatch || p[k] == '.') {
				prevMatch = s[j]
				j++
			}
			// check the rest
			k += 2
		}
	}

	return sLen-j == 0
}

func isMatchBasicRegexAtFirst(input, regex *string) bool {
	s := *input
	p := *regex
	if len(p) > 0 {
		return len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	}
	return len(s) == 0
}

func isMatchBasicRegexRecursive(input, regex *string) bool {
	if input == nil || regex == nil {
		return false
	}
	s, p := *input, *regex

	if len(p) == 0 {
		return len(s) == 0
	}

	// check the first char if it matches regex char without following *
	if len(p) < 2 || p[1] != '*' {
		if isMatchBasicRegexAtFirst(&s, &p) {
			s1 := u.ShiftSlice(s, 1)
			p1 := u.ShiftSlice(p, 1)
			// both moving to next char and call recursively
			return isMatchBasicRegexRecursive(&s1, &p1)
		}
		// the first chars do not match
		return false
	}

	p2 := u.ShiftSlice(p, 2)
	if isMatchBasicRegexRecursive(&s, &p2) {
		return true
	}

	// repeat until matched all chars with *
	for isMatchBasicRegexAtFirst(&s, &p) {
		s = u.ShiftSlice(s, 1)
		// check the rest
		if isMatchBasicRegexRecursive(&s, &p2) {
			return true
		}
	}

	return false
}

func isMatchBasicRegexSlice(input, regex *string) bool {
	s, p := *input, *regex
	var prevMatch byte = 0xFF

	// u.Debug("\ninput: s = '%v', p = '%v'\n", s, p)

	for len(p) > 0 {
		if len(s) == 0 {
			u.Debug(" zero: s = '%v', p = '%v', prevMatch = '%c'\n", s, p, prevMatch)
			if len(p) == 1 || p[1] != '*' && p[1] != '.' {
				return p[0] == prevMatch || p[0] == '.'
			}
			if p[1] == '.' {
				return false
			}
			p = u.ShiftSlice(p, 2)
		} else if len(p) < 2 || p[1] != '*' {
			// check the first char if it matches regex char without following *
			// u.Debug("first: s = '%v', p = '%v'\n", s, p)
			if len(s) == 0 || s[0] != p[0] && s[0] != prevMatch && p[0] != '.' {
				// the first chars do not match
				return false
			}
			// both moving to next char
			s = u.ShiftSlice(s, 1)
			p = u.ShiftSlice(p, 1)
			prevMatch = 0xFF
		} else {
			// u.Debug("match: s = '%v', p = '%v'\n", s, p)
			// repeat until matched all chars with *
			for len(s) > 0 && (s[0] == p[0] || s[0] == prevMatch || p[0] == '.') {
				prevMatch = s[0]
				s = s[1:]
			}
			// check the rest
			p = u.ShiftSlice(p, 2)
		}
	}

	// u.Debug("  end: s = '%v', p = '%v'\n", s, p)
	return len(s) == 0
}
