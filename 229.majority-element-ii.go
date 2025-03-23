/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 */

// @lc code=start
func majorityElement(nums []int) []int {
	result := make([]int, 0)
	if len(nums) == 0 {
		return result
	} else if len(nums) == 1 {
        return nums
    }

	// boyer-moore voting algorithm
	// Initialize two candidates and two counters
	majority1, majority2 := nums[0], nums[1]
	count1, count2 := 0, 0

	// Find two candidates (majority1, majority2)
	for _, num := range nums {
		if num == majority1 {
			count1++
		} else if num == majority2 {
			count2++
		} else if count1 == 0 {
			majority1 = num
			count1++
		} else if count2 == 0 {
			majority2 = num
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if num == majority1 {
			count1++
		} else if num == majority2 {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		result = append(result, majority1)
	}
	if count2 > len(nums)/3 {
		result = append(result, majority2)
	}
	return result
}
// @lc code=end

