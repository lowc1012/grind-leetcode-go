/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
// binary search for firt position of element
func findFirstElement(target int, nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			// assure first element
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		}

	}
	return -1
}

// binary search for last position of element
func findLastElement(target int, nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			// assure first element
			if mid == len(nums)-1 || (nums[mid+1]) != target {
				return mid
			}
			left = mid + 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	return []int{findFirstElement(target, nums), findLastElement(target, nums)}
}

// @lc code=end

