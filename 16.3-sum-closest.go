/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {

	result := nums[0] + nums[1] + nums[2]

	n := len(nums)

	sort.Ints(nums)
	for i := 0; i < n; i++ {

		left := i+1
		right := n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return target
			} else if sum > target{
				right--
			} else {
				left++
			}
			d1 := math.Abs(float64(sum - target))
			d2 := math.Abs(float64(result - target))
			if d1 < d2 {
				result = sum
			}
		}

	}
    return result
}
// @lc code=end

