/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {

	var res strings.Builder
	carry, p1, p2 := 0, len(num1)-1, len(num2)-1

	for p1 >= 0 || p2 >= 0 {

		x1 := 0
		if p1 >= 0 {
			x1, _ = strconv.Atoi(string(num1[p1]))
		}

		x2 := 0
		if p2 >= 0 {
			x2, _ = strconv.Atoi(string(num2[p2]))
		}

		val := (carry + x1 + x2) % 10
		carry = (carry + x1 + x2) / 10

		res.WriteString(strconv.Itoa(val))
		p1--
		p2--
	}

	if carry != 0 {
		res.WriteString("1")
	}

	return reverseString(res.String())
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = 1+i, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// @lc code=end

