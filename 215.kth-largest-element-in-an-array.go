/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
// https://leetcode.com/discuss/general-discussion/1088565/top-k-problems-sort-heap-and-quickselect
// implement the heap interface
type minHeap []int

func (h minHeap) Len() int { return len(h) }

func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &minHeap{}
	// build a empty min-heap, O(N), N = h.Len() = 0
	heap.Init(h)

	// partial sorting, O(NlogK)
	for _, n := range nums { // O(N)
		if h.Len() < k {
			heap.Push(h, n)
		} else {
			if n > (*h)[0] {
				(*h)[0] = n
				heap.Fix(h, 0) // O(logK)
			}
		}
	}

	// TC: O(NlogK), SC: O(K)
	return (*h)[0]
}

// @lc code=end

