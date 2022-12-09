/*
 * @lc app=leetcode id=318 lang=golang
 *
 * [318] Maximum Product of Word Lengths
 */

import "strings"

// @lc code=start
func maxProduct(words []string) int {
	res := 0

	// walk through all words in the given array
	start, end := 0, len(words)-1
	for start < end {
		lenStart := len(words[start])
		lenEnd := len(words[end])

		// check the word whether is consis of any char in the map
		if !checkHasCommonLetter(words[start], words[end]) {
			l := lenStart * lenEnd
			if l > res {
				res = l
			}
		}

		if lenStart > lenEnd {
			end -= 1
		} else {
			start += 1
		}

	}
	return res
}

func checkHasCommonLetter(a, b string) bool {
	for _, v := range b {
		if strings.IndexRune(a, v) > -1 {
			return true
		}
	}
	return false
}

// @lc code=end

