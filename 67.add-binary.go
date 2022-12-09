/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a, b string) string {
	// a and b can not be nil and empty
	// a and b consist only of '0' or '1'

	result := ""
	ptrA := len(a) - 1
	ptrB := len(b) - 1

	carry := 0

	for ptrA >= 0 || ptrB >= 0 {
		sum := carry

		if ptrA >= 0 {
			sum += int(a[ptrA] - '0')
			ptrA--
		}

		if ptrB >= 0 {
			sum += int(b[ptrB] - '0')
			ptrB--
		}

		carry = sum / 2
		result = strconv.Itoa(sum%2) + result
	}

	if carry > 0 {
		result = "1" + result
	}

	return result
}

// @lc code=end

