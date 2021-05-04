/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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
func flatten(root *TreeNode) {
	helper(root)
}

func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := helper(root.Left)
	right := helper(root.Right)

	if left == nil && right == nil {
		return root
	} else if left != nil {
		temp := root.Right
		root.Right = root.Left
		root.Left = nil
		left.Right = temp

		fmt.Println(root)

		if right == nil {
			return left
		}
	}

	return right
}
// @lc code=end

