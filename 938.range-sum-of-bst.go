/*
 * @lc app=leetcode id=938 lang=golang
 *
 * [938] Range Sum of BST
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
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	return count(root, low, high)
}

func count(root *TreeNode, low int, high int) int {
	sum := 0
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	if root.Left != nil && root.Val > low {
		sum += count(root.Left, low, high)
	}

	if root.Right != nil && root.Val < high {
		sum += count(root.Right, low, high)
	}

	return sum
}
// @lc code=end

