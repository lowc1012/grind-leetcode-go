/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	// two pointers
	result := [][]int{}
	n := len(nums)
	if n == 0 || n < 3 {
		return result
	}

	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		left := i + 1
		right := n - 1

		for left < right {
			
			sum := nums[left] + nums[right]
			if target == sum {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum > target {
				right--
			} else {
				left++
			}

		}
	}
	return result
}
// @lc code=end

