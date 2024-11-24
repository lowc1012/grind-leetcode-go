/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start
// Time Complexity: O(V + E)
// Space Complexity: O(V + E)
func findOrder(numCourses int, prerequisites [][]int) []int {
  return topologicalSort(numCourses, prerequisites)
}

func topologicalSort(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	// TC: O(E), E is the number of prerequisites
	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		graph[prerequisite] = append(graph[prerequisite], course)
		inDegree[course]++
	}

	queue := make([]int, 0)
	// TC: O(V), V is the number of courses
	for course := 0; course < numCourses; course++ {
		if inDegree[course] == 0 {
			queue = append(queue, course)
		}
	}

	coursesVisited := 0
	order := make([]int, 0)
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		coursesVisited++

		order = append(order, course)
		for _, course := range graph[course] {
			inDegree[course]--
			if inDegree[course] == 0 {
				queue = append(queue, course)
			}
		}
	}

	if coursesVisited == numCourses {
		return order
	}
	return []int{}
}
// @lc code=end

