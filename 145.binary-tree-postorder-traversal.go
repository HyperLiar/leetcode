/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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
func postorderTraversal(root *TreeNode) []int {
    r := []int{}
	q := []*TreeNode{}
	n := root

	for n != nil || len(q) > 0 {
		for n != nil {
			r = append([]int{n.Val}, r...)
			q = append(q, n)
			n = n.Right
		}

		n = q[len(q)-1]
		q = q[0:len(q)-1]
		n = n.Left
	}

	return r
}
// @lc code=end

