package str

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// Reverse a string
func Reverse(s string) string {
	return reverseFunc(s)
}

// reverse a string
// see http://golangcookbook.com/chapters/strings/reverse/
func reverseRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; {
		runes[i], runes[j] = runes[j], runes[i]
		i, j = i+1, j-1
	}
	return string(runes)
}

// reverseByCopy reverses a string by built-in copy func
func reverseByCopy(s string) string {
	bytes := []byte(s)
	cs := make([]byte, len(bytes))
	bn := 0
	for len(bytes) > 0 {
		r, size := utf8.DecodeLastRune(bytes)
		d := make([]byte, size)
		_ = utf8.EncodeRune(d, r)
		bn += copy(cs[bn:], d)
		bytes = bytes[:len(bytes)-size]
	}
	return string(cs)
}

// reverseByRunes reverses a string by using Runes type
func reverseByRunes(s string) string {
	runes := Runes(s)
	return string(runes.Reverse())
}

// reverseBySprintf reverses a string by using fmt.Sprintf
func reverseBySprintf(s string) string {
	var rs string
	bytes := []byte(s)
	for len(bytes) > 0 {
		r, size := utf8.DecodeLastRune(bytes)
		rs += fmt.Sprintf("%c", r)
		bytes = bytes[:len(bytes)-size]
	}
	return rs
}

// reverseFunc reverses a string by defer func
func reverseFunc(s string) (rs string) {
	for _, v := range s {
		defer func(r rune) { rs += string(r) }(v)
	}
	return
}

// reverseGrapheme reverses a unicode string by using regexp
func reverseGrapheme(s string) string {
	regex := regexp.MustCompile("\\PM\\pM*|.")
	slice := regex.FindAllString(s, -1)
	sz := len(slice)
	rs := make([]string, sz)

	for i := 0; i < sz; i++ {
		rs[i] = slice[sz-1-i]
	}

	return strings.Join(rs, "")
}

// reverseGraphemeUnicode reverses a unicode string
func reverseGraphemeUnicode(s string) string {
	runes := []rune("")
	checked := false
	rs := ""

	for _, c := range s {
		if !unicode.Is(unicode.M, c) {
			if len(runes) > 0 {
				rs = string(runes) + rs
			}
			runes = runes[:0]
			runes = append(runes, c)

			if checked == false {
				checked = true
			}
		} else if checked == false {
			rs = string(append([]rune(""), c)) + rs
		} else {
			runes = append(runes, c)
		}
	}

	return string(runes) + rs
}

// reverseNorm reverses a unicode string by using unicode/norm
// preserves sequences of combining characters intact, w/ invalid UTF-8 input
func reverseNorm(s string) string {
	bound := make([]int, 0, len(s)+1)

	var iter norm.Iter
	iter.InitString(norm.NFD, s)
	bound = append(bound, 0)
	for !iter.Done() {
		iter.Next()
		bound = append(bound, iter.Pos())
	}
	bound = append(bound, len(s))
	out := make([]byte, 0, len(s))
	for i := len(bound) - 2; i >= 0; i-- {
		out = append(out, s[bound[i]:bound[i+1]]...)
	}
	return string(out)
}

// reverseUTF8 using utf8 decoding/encoding functions
func reverseUTF8(s string) string {
	sz := len(s)
	rs := make([]byte, sz)
	for start := 0; start < sz; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(rs[sz-start:], r)
	}
	return string(rs)
}
