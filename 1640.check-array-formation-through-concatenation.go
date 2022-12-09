/*
 * @lc app=leetcode id=1640 lang=golang
 *
 * [1640] Check Array Formation Through Concatenation
 */

// @lc code=start
func canFormArray(arr []int, pieces [][]int) bool {
	firstPiece := make(map[int][]int)
	for _, p := range pieces {
		firstPiece[p[0]] = p
	}

	concat := make([]int, 0)
	for _, v := range arr {
		if p, exist := firstPiece[v]; exist {
			concat = append(concat, p...)
		}
	}

	if len(arr) != len(concat) {
		return false
	}

	for i, _ := range arr {
		if arr[i] != concat[i] {
			return false
		}
	}

	return true
}

// @lc code=end

