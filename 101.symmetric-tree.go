/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root, root)

	for len(queue) > 0 {
		r1 := queue[0]
		r2 := queue[1]
		queue = queue[2:]

		if r1 == nil && r2 == nil {
			continue
		}
		if r1 == nil || r2 == nil {
			return false
		}

		if r1.Val != r2.Val {
			return false
		}

		queue = append(queue, r1.Left)
		queue = append(queue, r2.Right)
		queue = append(queue, r1.Right)
		queue = append(queue, r2.Left)
	}

	return true
}

func isMirror(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}

	return r1.Val == r2.Val &&
		isMirror(r1.Left, r2.Right) && isMirror(r1.Right, r2.Left)
}

// @lc code=end

