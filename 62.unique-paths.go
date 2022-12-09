/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
// https://www.tutorialcup.com/interview/matrix/unique-paths.htm
func uniquePaths(m int, n int) int {
	// We can use Recursion to solve this question.
	// But it's not efficient. O(2^max(m, n))
	// The time complexity can be reduced from exponential order to
	// polynomial order by DP.

	// Dynamic programming (tabulation = bottom-up dynamic programming)
	// Let’s describe a state for our DP problem to be dp[x]
	// with dp[0] as base state and dp[n] as our destination
	// state. So,  we need to find the value of destination
	// state i.e dp[n]. If we start our transition from our
	// base state i.e dp[0] and follow our state transition
	// relation to reach our destination state dp[n], we call
	// it the Bottom-Up approach as it is quite clear that we
	// started our transition from the bottom base state and
	// reached the topmost desired state.

	// table := make([][]int, m)
	// for i, _ := range table {
	// 	table[i] = make([]int, n)
	// }

	// // time complexity: O(mn), space complexity: O(mn)
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if i == 0 || j == 0 {
	// 			// base case
	// 			table[i][j] = 1
	// 		} else {
	// 			// transition from the bottom-most base case
	// 			table[i][j] = table[i-1][j] + table[i][j-1]
	// 		}
	// 	}
	// }
	// return table[m-1][n-1]

	// Optimize the memory usage
	arr := make([]int, m)

	// base case
	for i := 0; i < m; i++ {
		arr[i] = 1
	}

	// recurrence relation
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			arr[j] = arr[j] + arr[j-1]
		}
	}

	// return destination state
	return arr[m-1]
}

// @lc code=end

