/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
import "math/rand"

func topKFrequent(nums []int, k int) []int {
	// quick select , TC: O(N)
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v]++
	}

	frequent := make([][]int, len(countMap))
	for i := 0; i < len(countMap); i++ {
		frequent[i] = make([]int, 2)
	}

	i := 0
	for k, v := range countMap {
		frequent[i][0] = k
		frequent[i][1] = v
		i++
	}

	quickSelect(&frequent, len(countMap)-k, 0, len(countMap)-1)

	res := []int{}
	for i := len(countMap) - k; i < len(countMap); i++ {
		res = append(res, frequent[i][0])
	}
	return res
}

func quickSelect(arr *[][]int, kSmall, start, end int) {
	if start == end {
		return
	}

	// select pivot randomly
	pivotIndex := start + rand.Intn(end-start)

	selectedIndex := partition(arr, start, end, pivotIndex)
	if selectedIndex > kSmall {
		quickSelect(arr, kSmall, start, selectedIndex-1)
	} else if selectedIndex < kSmall {
		quickSelect(arr, kSmall, selectedIndex+1, end)
	} else {
		return
	}

}

func partition(arr *[][]int, start, end, pivotIndex int) int {

	pivotValue := (*arr)[pivotIndex][1]

	// move pivot to end position
	(*arr)[pivotIndex], (*arr)[end] = (*arr)[end], (*arr)[pivotIndex]

	nextLeftIndex := start
	for i := start; i < end; i++ {
		if (*arr)[i][1] < pivotValue {
			(*arr)[i], (*arr)[nextLeftIndex] = (*arr)[nextLeftIndex], (*arr)[i]
			nextLeftIndex++
		}
	}

	(*arr)[nextLeftIndex], (*arr)[end] = (*arr)[end], (*arr)[nextLeftIndex]

	return nextLeftIndex
}

// @lc code=end

