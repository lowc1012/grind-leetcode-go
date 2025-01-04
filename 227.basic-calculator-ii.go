/*
 * @lc app=leetcode id=227 lang=golang
 *
 * [227] Basic Calculator II
 */

// @lc code=start
import "unicode"

func calculate(s string) int {
	stack := make([]int, 0)
	num := 0
	op := '+'

	for i, ch := range s {
		// skip the space character
		// but we need to trigger processing when the last character is ' '
		if ch == ' ' && i != len(s)-1 {
			continue
		}

		// if the character is a digit then convert it to a number
		if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0')
		}

		// if the character is an operator or the last character
		if (!unicode.IsDigit(ch) && ch != ' ') || i == len(s)-1 {
			// current `num` is in front of `op``
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				// pop the top element from the stack
				// and multiply it with the current number
				// then push it back to the stack
				stack[len(stack)-1] *= num
			case '/':
				// pop the top element from the stack
				// and divide it with the current number
				// then push it back to the stack
				stack[len(stack)-1] /= num
			}
			// reset the number and operator
			op = ch
			num = 0
		}
	}

	res := 0
	for _, v := range stack {
		res += v
	}
	return res
}
// @lc code=end

