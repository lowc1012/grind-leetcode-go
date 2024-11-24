/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
// detectCycle checks for cycles using DFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	// Create adjacency list and in-degree count
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses) // 有多少個 edge 指向這個 node

	// Build graph and count in-degrees
	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		graph[prerequisite] = append(graph[prerequisite], course)
		inDegree[course]++
	}

	// Initialize queue with all nodes that have no prerequisites
	queue := make([]int, 0)
	for course := 0; course < numCourses; course++ {
		if inDegree[course] == 0 {
			queue = append(queue, course)
		}
	}

	// Process queue
	coursesVisited := 0
	for len(queue) > 0 {
		// Remove a node (pop from queue)
		course := queue[0]
		queue = queue[1:]
		// Mark as visited
		coursesVisited++
		
		// For each neighbor, reduce in-degree and add to queue if in-degree becomes 0
		for _, nextCourse := range graph[course] {
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}
	
	// If we visited all courses, there's no cycle
	return coursesVisited == numCourses
}

// @lc code=end

