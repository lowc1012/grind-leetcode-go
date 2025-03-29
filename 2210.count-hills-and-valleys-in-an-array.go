/*
 * @lc app=leetcode id=2210 lang=golang
 *
 * [2210] Count Hills and Valleys in an Array
 */

// @lc code=start
func countHillValley(nums []int) int {
    count, l := 0, len(nums)
    for i := 1; i < l - 1; i++ {
        if nums[i] == nums[i-1] {
            continue
        }
        left, right := i-1, i+1
        for right < l && nums[i] == nums[right] {
            right++
        }

        if right < l {
            if nums[i] > nums[left] && nums[i] > nums[right] {
                count++
            }
            if nums[i] < nums[left] && nums[i] < nums[right] {
                count++
            }
        }
    }
    return count
}
// @lc code=end

