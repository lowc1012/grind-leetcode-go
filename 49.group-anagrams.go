/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
import (
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	anagrams := make(map[string][]string)

	for _, str := range strs {
		charCount := [26]int{}
		for _, ch := range str {
			charCount[ch-'a']++
		}
		// generate a key as identifier
		var sb strings.Builder
		for _, c := range charCount {
			sb.WriteString(strconv.Itoa(c))
			// Adding a separator for clarity (this avoids ambiguity between counts like "1" and "10")
			sb.WriteString("#")
		}
		key := sb.String()

		anagrams[key] = append(anagrams[key], str)
	}

	for _, entry := range anagrams {
		res = append(res, entry)
	}
	return res
}

// @lc code=end

