/*
 * @lc app=leetcode id=1497 lang=golang
 *
 * [1497] Check If Array Pairs Are Divisible by k
 */

// @lc code=start
// Time complexity: O(N)
// Space complexity: O(N)
func canArrange(arr []int, k int) bool {
	// HashMap to store the remainder of each element
	remainderMap := make(map[int]int)

	// Count the remainder of each element in the array
	for _, n := range arr {
		remainder := n % k // -1 % k = -1
		// -109 <= n <= 109
		// If remainder is negative, add k to make it positive
		if remainder < 0 {
			remainder += k
		}
		remainderMap[remainder]++
	}

	// Check if the remainder of each element can be paired
	for i := 1; i < k; i++ {
		if remainderMap[i] != remainderMap[k-i] {
			return false
		}
	}
	return remainderMap[0]%2 == 0
}

// @lc code=end

