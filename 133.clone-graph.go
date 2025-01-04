/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	// 1. The Val is unique, so we can use it to avoid duplicated nodes
	m := make(map[int]*Node)
	return dfsClone(node, m)
}

func dfsClone(node *Node, m map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	// if the node is cloned, return the result
	if n, ok := m[node.Val]; ok {
		return n
	}

	n := &Node{Val: node.Val}
	m[node.Val] = n

	for _, neighbor := range node.Neighbors {
		n.Neighbors = append(n.Neighbors, dfsClone(neighbor, m))
	}
	return n
}

// @lc code=end

