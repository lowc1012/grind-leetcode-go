/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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
import "strings"
import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root != nil {
		dfs(root, []int{}, &res)
	}
	return res
}

func dfs(node *TreeNode, curr []int, res *[]string) {

	curr = append(curr, node.Val)

	// leaf
	if node.Left == nil && node.Right == nil {
		*res = append(*res, toString(curr))
		return
	} else {
		if node.Left != nil {
			dfs(node.Left, curr, res)
		}

		if node.Right != nil {
			dfs(node.Right, curr, res)
		}
	}

}

func toString(path []int) string {
	var sb strings.Builder
	for _, v := range path {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteString("->")
	}
	return strings.TrimRight(sb.String(), "->")
}

// @lc code=end

