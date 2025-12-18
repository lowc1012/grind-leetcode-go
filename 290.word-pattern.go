/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

import "strings"

// @lc code=start
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := words[i]

		if mappedWord, exists := charToWord[char]; exists {
			if mappedWord != word {
				return false
			}
		} else {
			charToWord[char] = word
		}

		if mappedChar, exists := wordToChar[word]; exists {
			if mappedChar != char {
				return false
			}
		} else {
			wordToChar[word] = char
		}
	}

	return true
}

// @lc code=end

