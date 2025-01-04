/*
 * @lc app=leetcode id=1716 lang=golang
 *
 * [1716] Calculate Money in Leetcode Bank
 */

// @lc code=start
func totalMoney(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		weekNumber := i/7
		dayInWeek := i%7
		// amount for a day
		res += weekNumber + 1 + dayInWeek
	}
	return res
}

// @lc code=end

