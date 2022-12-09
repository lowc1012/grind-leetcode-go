/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		m[v]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

