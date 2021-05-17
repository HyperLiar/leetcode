/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree  Traversal
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
func Traversal(root *TreeNode) []int {
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
	// (root, &r)

	return r
}

func (node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}

	(node.Left, nums)
	*nums = append(*nums, node.Val)
	(node.Right, nums)
}

// @lc code=end

