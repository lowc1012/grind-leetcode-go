/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	// Use a fixed-size array to store the frequency of each character in s1
	count1 := [26]int{}
	for _, ch := range s1 {
		count1[ch-'a']++
	}

	// Use a sliding window to check the frequency of each character in s2
	slidingWindow := [26]int{}
	for i := 0; i < n2; i++ {
		slidingWindow[s2[i]-'a']++
	
		// if the window size is larger than n1, we need to remove the leftmost character
		if i >= n1 {
			slidingWindow[s2[i-n1]-'a']--
		}

		if slidingWindow == count1 {
			return true
		}
	}

	return false
}
// @lc code=end

