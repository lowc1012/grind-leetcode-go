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
	result := [][]string{}
	m := make(map[string][]string)
	for _, str := range strs {
		intSlice := make([]int, 26)
		for _, r := range str {
			intSlice[int(r-'a')]++
		}
		// convert integer slice to string
		valuesText := []string{}
		for _, i := range intSlice {
			valuesText = append(valuesText, strconv.Itoa(i))
		}

		keyString := strings.Join(valuesText, "")
		if _, exist := m[keyString]; !exist {
			m[keyString] = make([]string, 0)
		}

		m[keyString] = append(m[keyString], str)
	}

	for _, mapEntry := range m {
		result = append(result, mapEntry)
	}

	return result
}

// @lc code=end

