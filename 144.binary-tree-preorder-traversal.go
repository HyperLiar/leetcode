/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) []int {
    r := make([]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		n := q[len(q)-1]
		q = q[0:len(q)-1]
		if n == nil {
			continue
		}

		r = append(r, n.Val)
		q = append(q, n.Right, n.Left)
	}

	//preOrder(root, &r)
	return r
}

func preOrder(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}

	*nums = append(*nums, node.Val)
	preOrder(node.Left, nums)
	preOrder(node.Right, nums)
}
// @lc code=end

