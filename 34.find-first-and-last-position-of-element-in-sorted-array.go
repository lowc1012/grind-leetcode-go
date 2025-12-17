/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
// binary search for firt position of element
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right - left) / 2
        if target == nums[mid] {
            l, r := mid, mid
            for l > 0 && nums[l-1] == target {
                l--
            }
            for r < len(nums)-1 && nums[r+1] == target {
                r++
            }
            res = []int{l, r}
            break
        } else if target > nums[mid] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return res
}

// @lc code=end

