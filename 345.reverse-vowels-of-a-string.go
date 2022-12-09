/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start
func reverseVowels(s string) string {

	r := []byte(s)

	left, right := 0, len(r)-1
	for left < right {

		if r[left] != 'a' && r[left] != 'e' && r[left] != 'i' && r[left] != 'o' && r[left] != 'u' &&
			r[left] != 'A' && r[left] != 'E' && r[left] != 'I' && r[left] != 'O' && r[left] != 'U' {
			left++
		}

		if r[right] != 'a' && r[right] != 'e' && r[right] != 'i' && r[right] != 'o' && r[right] != 'u' &&
			r[right] != 'A' && r[right] != 'E' && r[right] != 'I' && r[right] != 'O' && r[right] != 'U' {
			right--
		}

		if r[left] == 'a' || r[left] == 'e' || r[left] == 'i' || r[left] == 'o' || r[left] == 'u' ||
			r[left] == 'A' || r[left] == 'E' || r[left] == 'I' || r[left] == 'O' || r[left] == 'U' {
			if r[right] == 'a' || r[right] == 'e' || r[right] == 'i' || r[right] == 'o' || r[right] == 'u' ||
				r[right] == 'A' || r[right] == 'E' || r[right] == 'I' || r[right] == 'O' || r[right] == 'U' {
				r[left], r[right] = r[right], r[left]
				left++
				right--
			}
		}

	}

	return string(r)
}

// @lc code=end

