/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
func findDuplicate(nums []int) int {
	// example: [1, 3, 4, 2, 2]
	// Binary search
	start, end := 1, len(nums) - 1 // [1, 4]

	for start < end {
		// guess the middle number is the duplicated
		mid := start + (end - start) / 2
		// count the number of elements that are less than or equal to 'mid'
		count := 0
		for _, n := range nums {
			if n <= mid {
				count++
			}
		}
		// (Pigeon-hole principle) if the count is greater than 'mid', the duplicate must be in the range of [start, mid]
		if count > mid {
			end = mid 
		} else { // otherwise, the duplicate must be in the range of [mid+1, end]
			start = mid + 1
		}
	}

	return start
}

// @lc code=end

