/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// preorder: VLR, inorder: LVR
	// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	// We can know the root is 3, the leftest node is 9

	// Use a stack to store the path we visited while traversing PreOrder array
	stack := make([]TreeNode, len(preorder))

	// Use a set to maintain the node in which the next right subtree is expected
	set := make(map[TreeNode]struct{})

	var root TreeNode
	pre, in := 0, 0
	for pre < len(preorder) {

		var node *TreeNode
		// reach the leftest node
		for {
			node := TreeNode{
				Val: preorder[pre],
			}

			if root == nil {
				root = node
			}

			if len(stack) > 0 {
				if set[stack[len(stack)-1]] != nil {

					delete(set, stack[len(stack)-1])

					peak := stack[len(stack)-1]

					peak.Right = node

					stack = stack[:len(stack)-1]
				} else {
					peak := stack[len(stack)-1]
					peak.left = node
				}
			}
			stack = append(stack, node)

			if pre >= len || preorder[pre] == inorder[in] {
				pre++
				break
			}
			pre++
		}

		node = nil

		for len(stack) > 0 && (in < len(preorder)) && (stack[len(stack)-1].Val == inorder[in]) {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			in++
		}

		if node != nil {
			set[node] = struct{}{}
			stack = append(stack, node)
		}
	}

	return &root
}

// @lc code=end

