package str

import "strconv"

// Str type is an alias to string
type Str string

// Compress returns a compressed string to omit repeats
func (s *Str) Compress() string {
	siz := len(*s)
	runes := make([]rune, 0, siz)
	var count = 1
	var hasEscape bool
	var prev rune

	repeats := func(num int) { // define a closure func
		str := strconv.Itoa(num)
		runes = append(runes, '*')
		for _, c := range str {
			runes = append(runes, c)
		}
		runes = append(runes, '\\')
	}

	for _, r := range *s {
		if r == prev {
			count++
		} else {
			if count > 1 {
				repeats(count)
				count = 1
			}
			if r == '*' || r == '\\' {
				hasEscape = true
				runes = append(runes, '\\')
			}
			prev = r
			runes = append(runes, r)
			if !hasEscape && len(runes) > siz {
				return string(*s)
			}
		}
	}
	if count > 1 {
		repeats(count)
	}
	if !hasEscape && len(runes) > siz {
		return string(*s)
	}
	return string(runes)
}

// Decompress returns a decompressed string
func (s *Str) Decompress() string {
	siz := len(*s)
	result := make([]rune, 0, siz*2)
	var counts []rune
	var escape, start bool
	var prev rune
	for _, r := range *s {
		if start {
			if '0' <= r && r <= '9' {
				counts = append(counts, r)
			} else {
				xn, _ := strconv.Atoi(string(counts))
				for i := 1; i < xn; i++ {
					result = append(result, prev)
				}
				counts = make([]rune, 0)
				start = false
			}
			continue
		}
		if r == '*' && !escape {
			start = true
		} else {
			if r != '\\' || escape {
				escape = false
				result = append(result, r)
				prev = r
			} else {
				escape = !start
			}
		}
	}
	return string(result)
}
