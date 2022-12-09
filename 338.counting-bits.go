/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start
func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	num := []int{0, 1}
	for i := 2; i <= n; i++ {
		num = append(num, num[int(i/2)] + num[(i%2)])
	}
	return num
}
// @lc code=end

