/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
// Time Complexity: O(n)
func lengthOfLongestSubstring(s string) int {
	// initialize the result
	res := 0

	// create a map to store the index of each character
	m := make(map[byte]int)

	// implement "sliding window"
	begin, end := 0, 0
	for end < len(s) {

		// Shrink window from beginning if duplicate found, until no duplicate
		lastIndex, exists := m[s[end]]
		if exists && lastIndex >= begin {
			begin = lastIndex + 1
		}

		// add the character to the map
		m[s[end]] = end
		// update the result
		res = max(res, end-begin+1)
		// move the end pointer
		end++
	}
	return res
}

// @lc code=end

