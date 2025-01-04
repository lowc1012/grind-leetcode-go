/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
    // backtracking
		res := []string{}
		backtrack(&res, "", n, n)
		return res
}

func backtrack(res *[]string, str string, maxOfLeft, maxOfRight int) {
		// return case
		if maxOfLeft == 0 && maxOfRight == 0 {
				*res = append(*res, str)
				return
		}
		// skip invalid combinations
		if maxOfLeft > maxOfRight {
			return
		}

		if maxOfLeft > 0 {
				backtrack(res, str+"(", maxOfLeft-1, maxOfRight)
		}

		if maxOfRight > 0 {
				backtrack(res, str+")", maxOfLeft, maxOfRight-1)
		}
}
// @lc code=end

