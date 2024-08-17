/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func sortedSquares(nums []int) []int {
		result := make([]int, len(nums))

    // use two pointers to compare the absolute value of the two ends
		left, right := 0, len(nums)-1

		// fill the result array from the end to the start
		for left <= right {
			if abs(nums[left]) > abs(nums[right]) {
				result[right-left] = nums[left] * nums[left]
				left++
			} else {
				result[right-left] = nums[right] * nums[right]
				right--
			}
		}

		// time complexity: O(n)
		return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

