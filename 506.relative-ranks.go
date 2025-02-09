/*
 * @lc app=leetcode id=506 lang=golang
 *
 * [506] Relative Ranks
 */

// @lc code=start

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] > h[j] } // Reverse for descending
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}

// Time complexity: O(NlogN)
// Space complexity: O(N)
func findRelativeRanks(score []int) []string {
    n := len(score)
    // Use minimum heap to store all scores
    // time complexity: O(NlogN)
    h := &MinHeap{}
    for _, s := range score {
        heap.Push(h, s) // O(logN) time
    }
    heap.Init(h) // O(N) space

    // use a HashMap to store the relation between scores(key) and rank(value)
    // time complexity: O(NlogN)
    rankMap := make(map[int]string, len(score)) // O(N) space
    for i := 1; i <= n; i++ {
        s := heap.Pop(h).(int) // O(logN)
        switch i {
        case 1:
            rankMap[s] = "Gold Medal"
        case 2:
            rankMap[s] = "Silver Medal"
        case 3:
            rankMap[s] = "Bronze Medal"
        default:
            rankMap[s] = strconv.Itoa(i)
        }
    }

    // iterate score to retrieve the result
    // O(N)
    result := make([]string, n) // O(N) space
    for i, j := range score {
        result[i] = rankMap[j]
    }

    return result
}

// @lc code=end

