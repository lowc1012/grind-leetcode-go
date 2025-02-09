/*
 * @lc app=leetcode id=802 lang=golang
 *
 * [802] Find Eventual Safe States
 */

// @lc code=start

func hasCycle(graph [][]int, state []int, curr int) bool {
	// The state 0: unvisited
	// The state 1: visiting (hasn't found terminal node in the path)
	// The state 2: visied (finished the traversal on the path)

	if state[curr] == 1 {
        return true // Cycle detected
    }
    if state[curr] == 2 {
        return false // Already visited node, no cycle
    }

	state[curr] = 1
	for _, adj := range graph[curr] {
		if hasCycle(graph, state, adj) {
			return true
		}
	}
	state[curr] = 2
	return false
}

func eventualSafeNodes(graph [][]int) []int {
	// Use DFS traverse the graph.
	// If there is no cycle in the path, add the start point to result

	// Use array to record the state of each node
	state := make([]int, len(graph))

	// iterate each node
	res := make([]int, 0)
	for i := 0; i < len(graph); i++ {
		if !hasCycle(graph, state, i) {
			res = append(res, i)
		}
	}
	return res
}
// @lc code=end

