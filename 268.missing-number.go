/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
// Time Complexity: O(NlogN)
func missingNumber(nums []int) int {
	// nums = [3, 0, 1], n = 3

	// binary search to find the missing number
	left, right := 0, len(nums)
	for left < right {
		// gusse the missing num is `mid`
		mid := left + (right-left)/2 // mid = 1
		
		// count the number of elements that are less than or equal to mid
		count := 0
		// time complexity: O(N)
		for _, n := range nums {
			if n <= mid {
				count++
			}
		}
		// if count <= mid, the missing number is in the left half
		if count <= mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// @lc code=end

