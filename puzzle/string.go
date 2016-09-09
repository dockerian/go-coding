// +build all puzzle string

package puzzle

import (
	u "github.com/dockerian/go-coding/utils"
)

// GetLongestUniqueSubstring solves the following problem:
// Given a string, find the longest non-repeating substring length
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

// GetLongestUniqueSubstring solves the following problem:
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
	for stIndex, _ := range input {
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
