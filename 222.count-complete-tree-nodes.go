/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
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
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := -1, -1
	tempL, tempR := root, root
	for tempL != nil {
		l++
		tempL = tempL.Left
	}
	for tempR != nil {
		r++
		tempR = tempR.Right
	}

	if l == r {
		return 2<<l - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for i := 0; i < len(queue); i++ {
		if queue[i].Left != nil {
			queue = append(queue, queue[i].Left)
		} else {
			break
		}
		if queue[i].Right != nil {
			queue = append(queue, queue[i].Right)
		} else {
			break
		}
	}

	return len(queue)
}

// @lc code=end

