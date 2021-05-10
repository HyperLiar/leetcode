/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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
func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	maxDepth(root, &max)
	return max
}

func maxDepth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left, max)
	right := maxDepth(root.Right, max)

	if left+right > *max {
		*max = left + right
	}

	if left > right {
		return left+1
	}

	return right+1
}

// @lc code=end

