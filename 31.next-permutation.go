/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
import "sort"

func nextPermutation(nums []int) {

	lastPeakIndex := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			lastPeakIndex = i
			break
		}
	}

	// the last permutation , e.g. [3, 2 , 1]
	if lastPeakIndex == -1 {
		sort.Ints(nums)
		return
	}

	index := lastPeakIndex
	for i := lastPeakIndex; i < len(nums); i++ {
		if nums[i] > nums[lastPeakIndex-1] && nums[i] < nums[index] {
			index = i
		}
	}

	nums[lastPeakIndex-1], nums[index] = nums[index], nums[lastPeakIndex-1]
	sort.Ints(nums[lastPeakIndex:])
}

// @lc code=end

