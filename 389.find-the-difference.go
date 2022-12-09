/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {

	for _, j := range s {
		if strings.Index(s, string(j)) > -1 {
			t = strings.Replace(t, string(j), "", 1)
		}
	}
	b := []byte(t)
    return b[0]
}
// @lc code=end

