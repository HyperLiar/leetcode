/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
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
func rob(root *TreeNode) int {
	res := dfs(root)
	return max(res[0], res[1])
}

func dfs(root *TreeNode) [2]int {
	res := [2]int{}
	if root == nil {
		return res
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	res[0] = max(left[0], left[1]) + max(right[0], right[1])
	res[1] = left[0] + right[0] + root.Val

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// @lc code=end

