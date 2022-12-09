/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(n int) int {
	if n < 2 {
		return n
	}

	f := make([]int, 31)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// @lc code=end

