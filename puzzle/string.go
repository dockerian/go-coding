package puzzle

import (
	"fmt"

	u "github.com/dockerian/go-coding/utils"
)

// CheckMatchedPair determines if a given pair, e.g. parenthesis or bracket,
// in a string all have valid matches
func CheckMatchedPair(s, begin, close string) (bool, error) {
	matchCount := 0
	beginCount, closeCount := 0, 0
	beginLen, closeLen, sz := len(begin), len(close), len(s)
	var okay = true
	var err error

	for i := 0; i < sz; i++ {
		if i+beginLen <= sz && s[i:i+beginLen] == begin {
			beginCount++
		}
		if i+closeLen <= sz && s[i:i+closeLen] == close {
			if beginCount-matchCount >= 1 {
				matchCount++
			}
			closeCount++
		}
	}

	okay = beginCount == matchCount && beginCount == closeCount

	u.Debug("begin/close: '%s' [%d] and '%s' [%d] in '%s' [%d] - matched %v\n", begin, beginCount, close, closeCount, s, sz, matchCount)

	if !okay {
		msg := fmt.Sprintf("begin/close: '%s' [%d] and '%s' [%d]",
			begin, beginCount, close, closeCount)
		if beginCount == closeCount {
			err = fmt.Errorf("Paired but misplaced %s", msg)
		} else {
			err = fmt.Errorf("Unproperly paired %s", msg)
		}
	}

	return okay, err
}

// GetAllCases returns all variations of upper and lower cases
func GetAllCases(s string) []string {
	var size = len(s)
	var maxL = int64(1) << uint(size)
	var diff = byte('a' - 'A')
	results := make([]string, 0, maxL)
	var getAll func([]byte, int)

	getAll = func(bytes []byte, len int) {
		pos := len - 1
		if pos <= 0 {
			results = append(results, string(bytes))
		}
		getAll(bytes, pos)
		bytes[pos] = bytes[pos] + diff
		getAll(bytes, pos)
	}

	bytes := []byte(s)
	getAll(bytes, size)

	return results
}

// GetLongestSubstringLength solves the following problem:
// Given a string, find the longest non-repeating substring length.
// Note: assuming all input are ASCII characters
// Tags: hash table, map, two pointers, string
func GetLongestSubstringLength(input string) (int, string) {
	var indexL, maxLeng, substrL, substrR int
	var lookup = make(map[byte]int)
	var length = len(input)

	for indexR := 0; indexR < len(input); indexR++ {
		currByte := input[indexR]
		if i, ok := lookup[currByte]; ok && indexL <= i {
			u.Debug("current input[%v : %v] = %v\n", indexL, indexR, input[indexL:indexR])
			u.Debug("    max input[%v : %v] = %v\n", substrL, substrR, input[substrL:substrR])
			if indexR-indexL >= maxLeng {
				maxLeng = indexR - indexL
				substrL = indexL
				substrR = indexR
			}
			indexL = i + 1 // update indexL to ensure non-repeating
		}
		lookup[currByte] = indexR // update last appeared location
	}

	if length-indexL > maxLeng {
		maxLeng = length - indexL
		substrL = indexL
		substrR = length
	}

	return maxLeng, string(input[substrL:substrR])
}

// GetLongestSubstringUTF8 solves the following problem:
// Given a string, find the longest substring without repeating characters.
// Note: This is to support UTF-8
func GetLongestSubstringUTF8(input string) (int, string, int, int) {
	var currLen, maxiLen, substrL, substrR int
	var hashmap = make(map[rune]int)
	var nextmap = make(map[rune]int)
	var preItem rune
	var posLeft int

	for posCurr, item := range input {
		if posCurr > 0 {
			nextmap[preItem] = posCurr
		}
		preItem = item

		nextIdx := nextmap[item]
		_, okay := hashmap[item]
		if okay && nextIdx >= posLeft {
			u.Debug("current pos %v : %v (next: %v)\n", posLeft, posCurr, nextIdx)
			u.Debug("current input[%v : %v] = %v [%v]\n",
				posLeft, posCurr, input[posLeft:posCurr], currLen)
			u.Debug("    max input[%v : %v] = %v [%v]\n",
				substrL, substrR, input[substrL:substrR], maxiLen)
			if currLen > maxiLen {
				maxiLen = currLen
				substrL = posLeft
				substrR = posCurr
			}
			for range input[posLeft:nextIdx] {
				currLen--
			}
			posLeft = nextIdx
		}

		hashmap[item] = posCurr
		currLen++
	}

	if currLen > maxiLen {
		substrR = len(input)
		substrL = posLeft
		maxiLen = currLen
	}

	var substr = string(input[substrL:substrR])
	u.Debug("Longest: input[%v : %v] = '%v' [%v] in '%v'\n",
		substrL, substrR, substr, maxiLen, input)

	return maxiLen, substr, substrL, substrR
}

// GetLongestUniqueSubstring solves the following problem:
// Given a string, find the longest substring without repeating characters.
// For example:
//   "abc" from "aaabcabcbb", and the length is 3
//   "o" from "oooooo", and the length of 1
func GetLongestUniqueSubstring(input string) string {
	var longest, current []rune
	var hashmap = make(map[rune]int)
	var szInput = len(input)
	var szRunes = len([]rune(input))

	u.Debug("input string: %+v (size= %v, len= )\n", input, szInput, szRunes)
	for stIndex := range input {
		slice := input[stIndex:szInput]
		u.Debug("%v: input slice: %+v\n", stIndex, slice)
		for index, item := range slice {
			u.Debug("start:index[item]= %v:%v[%v], current= %+v, max= %+v, hash= %+v\n",
				stIndex, index, item, string(current), string(longest), hashmap)
			_, okay := hashmap[item]
			if !okay {
				current = append(current, item)
				hashmap[item] = stIndex + index
				continue
			}
			break
		}

		if len(current) > len(longest) {
			longest = current
		}
		hashmap = make(map[rune]int)
		current = []rune{}

		if stIndex >= szInput-len(string(longest)) {
			u.Debug("* break @ start= %v,　size= %v, len= %v, max= %v (%v)\n",
				stIndex, szInput, szRunes, len(longest), len(string(longest)))
			break
		}
	}
	u.Debug("\n")

	if len(current) > len(longest) {
		return string(current)
	}

	return string(longest)
}

// GetMostFrequentRune returns the rune and count of the most appearance
func GetMostFrequentRune(input string) (rune, int) {
	var runeHash = make(map[rune]int)
	var mostFreq = '\x00'
	var maxCount = 0
	for _, r := range input {
		n, okay := runeHash[r]
		if okay {
			runeHash[r] = 1 + n
		} else {
			runeHash[r] = 1
		}
		if maxCount < runeHash[r] {
			maxCount = runeHash[r]
			mostFreq = r
		}
		// u.Debug("%c: %c, %v\n", r, mostFreq, maxCount)
	}
	return mostFreq, maxCount
}
