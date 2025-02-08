/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
// detectCycle checks for cycles using DFS

func detectCycle(graph [][]int, state []int, curr int) bool {

	if state[curr] == 1 {
		// detected cycle
		return true
	}

	if state[curr] == 2 {
		// the node is finished checking cycle with all adajacent nodes
		return false
	}

	// Mark as being visited (part of current DFS path)
	state[curr] = 1

	for _, neibor := range graph[curr] {
		if detectCycle(graph, state, neibor) {
			return true
		}

	}

	// mart current node visied
	state[curr] = 2

	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Create adjacency list to build a graph
	graph := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		graph[prerequisite] = append(graph[prerequisite], course)
	}

	state := make([]int, numCourses)

	for idx, _ := range graph {
		if detectCycle(graph, state, idx) {
			return false
		}
	}

	return true
}

// @lc code=end

