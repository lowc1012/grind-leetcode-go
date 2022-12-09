/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
import "strconv"

func fizzBuzz(n int) []string {
	s := make([]string, n)

	for i := 1; i <= n; i++ {

		if i%3 == 0 && i%5 == 0 {
			s[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			s[i-1] = "Fizz"
		} else if i%5 == 0 {
			s[i-1] = "Buzz"
		} else {
			s[i-1] = strconv.Itoa(i)
		}
	}
	return s
}

// @lc code=end

