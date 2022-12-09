/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	// create a cache
	cache := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		cache[i] = make([]int, len(text2)+1)
	}

	/*
			text1 = "abcde", text2 = "ace"
			cache:
		  		  "" a b c d e
				0  0 0 0 0 0 0
				a  0 0 0 0 0 0
				c  0 0 0 0 0 0
				e  0 0 0 0 0 0
	*/

	// DP (iteration + caching)
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				cache[i][j] = cache[i-1][j-1] + 1
			} else {
				cache[i][j] = max(cache[i-1][j], cache[i][j-1])
			}
		}
	}

	/*
			text1 = "abcde", text2 = "ace"
			cache:
		  		  "" a b c d e
				0  0 0 0 0 0 0
				a  0 1 1 1 1 1
				c  0 1 1 2 2 2
				e  0 1 1 2 2 3

	*/

	// Using DP (tabulation)
	// if len(text1) = M, len(text2) = N
	// TC: O(M * N), SC: O(M * N)
	return cache[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end=
