/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
import "unicode"

func isPalindrome(s string) bool {

	r := []rune(s)
	head, tail := 0, len(r)-1

	for head < tail {
		cHead := r[head]
		cTail := r[tail]

		if !unicode.IsLetter(cHead) && !unicode.IsDigit(cHead) {
			head++
		} else if !unicode.IsLetter(cTail) && !unicode.IsDigit(cTail) {
			tail--
		} else {
			if unicode.ToLower(cHead) != unicode.ToLower(cTail) {
				return false
			}

			head++
			tail--
		}

	}
	return true
}

// @lc code=end

