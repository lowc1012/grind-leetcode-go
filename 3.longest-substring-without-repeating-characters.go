/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// initialize the result
	res := 0

	// make a set
	set := make(map[string]struct{})

	// use two pointers (begin, end) to implement "sliding window"
	begin := 0
	for end := 0; end < len(s); end++ {
		// if the character is already in the set, remove the character at the beginning of the string
		_, exists := set[string(s[end])]
		for exists {
			delete(set, string(s[begin]))
			begin++
			_, exists = set[string(end)]
		}

		// add the character to the set
		set[string(k)] = struct{}{}
		// update the result
		res = max(res, end-begin+1)
	}
	return res
}

// @lc code=end

