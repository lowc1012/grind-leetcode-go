/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start

// using stack
func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	// stack
	stack := []int{}

	for index, temp := range temperatures {

		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			lastIndex := stack[len(stack)-1]
			days := index - lastIndex
			result[lastIndex] = days

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, index)
	}

	return result
}

// @lc code=end

