/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// initialize the result
	res := 0

	// make a set to store the characters
	set := make(map[byte]struct{})

	// implement "sliding window"
	begin, end := 0, 0
	for end < len(s) {
		// if the character is already in the set, remove the character at the beginning of the string
		_, exists := set[s[end]]
		for exists {
			delete(set, s[begin])
			begin++
			_, exists = set[s[end]]
		}

		// add the character to the set
		set[s[end]] = struct{}{}
		// update the result
		res = max(res, end-begin+1)
		// move the end pointer
		end++
	}
	return res
}

// @lc code=end

