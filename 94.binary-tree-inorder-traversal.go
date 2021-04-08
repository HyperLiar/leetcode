/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
func inorderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	q := make([]*TreeNode, 0)
	n := root

	for n != nil || len(q) > 0 {
		for n != nil {
			q = append(q, n)
			n = n.Left
		}

		top := q[len(q)-1]
		q = q[0:len(q)-1]
		r = append(r, top.Val)
		n = top.Right
	}
	// inOrder(root, &r)

	return r
}

func inOrder(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}

	inOrder(node.Left, nums)
	*nums = append(*nums, node.Val)
	inOrder(node.Right, nums)
}

// @lc code=end

