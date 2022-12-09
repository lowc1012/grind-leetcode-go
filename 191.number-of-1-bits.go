/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	var count = 0

	for i := 0; i < 32; i++ {
		if (num & 1) == 1 {
			count++
		}
		num = num >> 1
	}

	return count
}
// @lc code=end

