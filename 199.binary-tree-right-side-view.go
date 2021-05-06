/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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
func rightSideView(root *TreeNode) []int {
	result := []int{}
	dfs(root, &result, 0)
	return result
}

func dfs(root *TreeNode, result *[]int, level int) {
	if root == nil {
		return
	}
	if len(*result) <= level {
		*result = append(*result, root.Val)
	}
	dfs(root.Right, result, level+1)
	dfs(root.Left, result, level+1)
}

// @lc code=end

