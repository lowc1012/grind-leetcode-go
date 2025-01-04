/*
 * @lc app=leetcode id=1046 lang=golang
 *
 * [1046] Last Stone Weight
 */

// @lc code=start

import "container/heap"

type IntHeap []int

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// Time Complexity: O(NlogN)
// Space Complexity: O(N)
func lastStoneWeight(stones []int) int {

	// Build a Max-Heap, O(NlogN)
	h := &IntHeap{}
	heap.Init(h) // O(N)
	for _, v := range stones {
		heap.Push(h, v) // O(logN)
	}

	// Find the heavist two stones
	for h.Len() > 1 {
		y := heap.Pop(h).(int)
		x := heap.Pop(h).(int)
		heap.Push(h, y-x)
	}
	return heap.Pop(h).(int)
}

// @lc code=end

