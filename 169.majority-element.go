/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
// (method 1) use quick sort (sort.Ints) => TC: O(nlogn), SC: O(1)
// (method 2) use quick selection algorithm => TC: O(n), SC: O(1)
import "math/rand"

func majorityElement(nums []int) int {
	// use quick selection
	kthSmall := 1 + len(nums)/2

	return quickSelect(&nums, kthSmall, 0, len(nums)-1)
}

func quickSelect(nums *[]int, kthSmall, start, end int) int {
	if start == end {
		return (*nums)[start]
	}

	pivotIndex := start + rand.Intn(end-start)

	selectedIndex := partition(nums, start, end, pivotIndex)
	if selectedIndex > kthSmall-1 {
		return quickSelect(nums, kthSmall, start, selectedIndex-1)
	} else if selectedIndex < kthSmall-1 {
		return quickSelect(nums, kthSmall, selectedIndex+1, end)
	} else {
		return (*nums)[selectedIndex]
	}
}

func partition(nums *[]int, start, end, pivotIndex int) int {
	pivotVal := (*nums)[pivotIndex]

	(*nums)[pivotIndex], (*nums)[end] = (*nums)[end], (*nums)[pivotIndex]

	nextLeftIndex := start
	for i := start; i < end; i++ {
		if (*nums)[i] < pivotVal {
			(*nums)[i], (*nums)[nextLeftIndex] = (*nums)[nextLeftIndex], (*nums)[i]
			nextLeftIndex++
		}
	}

	(*nums)[nextLeftIndex], (*nums)[end] = (*nums)[end], (*nums)[nextLeftIndex]

	return nextLeftIndex
}

// @lc code=end

