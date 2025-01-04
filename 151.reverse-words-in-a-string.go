/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
import "strings"

func reverseWords(s string) string {
	s = reverse(s)
	// "eulb si yks eht"

	words := strings.Split(s, " ")
	// ["eulb", "si", "yks", "eht"]
	res := ""
	for i := 0; i < len(words); i++ {
		// skip empty string
		if words[i] != "" {
			res += reverse(words[i]) + " "
		}
	}

	return res[:len(res)-1]
}

func reverse(s string) string {
	chars := []rune(s)
	left, right := 0, len(chars)-1
	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
	return string(chars)
}

// @lc code=end

