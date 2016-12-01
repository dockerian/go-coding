package utils

import "strings"

// FmtComma formats number with thousands comma
func FmtComma(number string) string {
	parts := strings.Split(number, ".")
	part0 := parts[0]
	szInt := len(part0)
	part1 := ""

	if szInt == 0 || len(parts) > 2 {
		return number
	} else if len(parts) == 2 {
		part1 = "." + parts[1]
	}

	bound := szInt - 1
	bytes := make([]byte, szInt+bound/3)

	for i, j := bound, len(bytes)-1; i >= 0 && j >= 0; i-- {
		isCommaPos := (bound-i)%3 == 0
		if i != bound && isCommaPos {
			bytes[j] = ','
			j--
		}
		bytes[j] = part0[i]
		j--
	}

	return string(bytes) + part1
}
