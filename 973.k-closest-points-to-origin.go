/*
 * @lc app=leetcode id=973 lang=golang
 *
 * [973] K Closest Points to Origin
 */

// @lc code=start

import (
	"container/heap"
	"math"
)

type Point struct {
	dist float64 // 到原點的距離
	coord []int // 座標
}

type MaxHeap []Point

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].dist > h[j].dist
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeap) Pop() interface{} {
	oldHeap := *h
	l := len(oldHeap)
	item := oldHeap[l-1]
	*h = oldHeap[:l-1]
	return item
}

func kClosest(points [][]int, k int) [][]int {
  // Use a Maximum Heap to maintain the distance between origin and point
	h := &MaxHeap{}
	heap.Init(h)

	// iterate over the points, and 
	for _, point := range points {
		dist := getDistance(point)
		// just need to maintain the heap whose size is equal to k
		if len(*h) < k {
			heap.Push(h, Point{dist: dist, coord: point})
		} else if dist < (*h)[0].dist {
			heap.Pop(h)
			heap.Push(h, Point{dist: dist, coord: point})
		}
	}

	result := make([][]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(Point).coord
	}

	return result
}

func getDistance(point []int) float64 {
	return math.Sqrt(float64(point[0]*point[0] + point[1]*point[1]))
}
// @lc code=end

