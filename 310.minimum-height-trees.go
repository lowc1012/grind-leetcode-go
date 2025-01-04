/*
 * @lc app=leetcode id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	// Base cases
	if n <= 2 {
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = i
		}
		return result
	}

	// Build the graph
	graph := make([][]int, n)
	inDegree := make([]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
		inDegree[a]++
		inDegree[b]++
	}

	// Find the leaves
	leaves := make([]int, 0)
	for i := 0; i < n; i++ {
		if inDegree[i] == 1 {
			leaves = append(leaves, i)
		}
	}

	// Trim the leaves until reaching the centroids
	remainingNodes := n
	for remainingNodes > 2 {
		remainingNodes -= len(leaves)
		newLeaves := make([]int, 0)
		// Remove the current leaves along with the edges
		for _, leaf := range leaves {
			for _, neighbor := range graph[leaf] {
				inDegree[neighbor]--
				if inDegree[neighbor] == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}
		leaves = newLeaves
	}

	return leaves
}

// @lc code=end

