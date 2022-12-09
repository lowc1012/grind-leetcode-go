/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 */

// @lc code=start

type IntHeap []int

// implement sort interface
func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] =  h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0:len(old)-1]
	return x
}

type KthLargest struct {
    Items *IntHeap
	K int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
	}

	// remove the smaller elements
	for len(*h) > k {
		heap.Pop(h)
	}

	return KthLargest{
		h,
		k,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(*this.Items) < this.K {
		heap.Push(this.Items, val)
	} else if val > (*this.Items)[0] {
		heap.Push(this.Items, val)
		heap.Pop(this.Items)
	}
	return (*this.Items)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end

