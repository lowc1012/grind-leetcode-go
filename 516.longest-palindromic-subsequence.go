/*
 * @lc app=leetcode id=516 lang=golang
 *
 * [516] Longest Palindromic Subsequence
 */

// @lc code=start
func longestPalindromeSubseq(s string) int {
	// DP, tabulation
	table := make([][]int, len(s))
	for i, _ := range table {
		table[i] = make([]int, len(s))
	}

	for i := 0; i < len(s); i++ {
		table[i][i] = 1
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				table[i][j] = table[i-1][j+1] + 2
			} else {
				table[i][j] = max(table[i][j+1], table[i-1][j])
			}
		}
	}

	return table[len(s)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

