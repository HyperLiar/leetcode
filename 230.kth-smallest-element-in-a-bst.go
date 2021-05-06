/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
func kthSmallest(root *TreeNode, k int) int {

}

func inorder(root *TreeNode, k int, *result) {
	if root == nil {
		return
	}
}

func iterate(root *TreeNode, k int) {
	// inorder traversal
	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}
}

// @lc code=end

