/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
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
func pathSum(root *TreeNode, targetSum int) int {
	return helper(root, 0, targetSum, map[int]int{0: 1})
}

func helper(node *TreeNode, curSum, target int, m map[int]int) int {
	if node == nil {
		return 0
	}
	curSum += node.Val
	summary := m[curSum-target]
	m[curSum]++
	summary += helper(node.Left, curSum, target, m)
	summary += helper(node.Right, curSum, target, m)
	m[curSum]--
	return summary
}

// @lc code=end

