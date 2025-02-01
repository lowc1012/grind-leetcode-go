/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)

	countP := [26]int{}
	for _, ch := range p {
		countP[ch-'a']++
	}

	slidingWindow := [26]int{}
	for i := 0; i < len(s); i++ {
		slidingWindow[s[i]-'a']++

		// shink the window from left because the size of window should be less than len(p)
		if i >= len(p) {
			// slidingWindow[(first char)-'a']--
			slidingWindow[s[i-len(p)]-'a']--
		}

		if slidingWindow == countP {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

// @lc code=end

