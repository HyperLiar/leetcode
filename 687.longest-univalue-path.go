/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
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
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxCount := 0
    countPath(root, &maxCount)
	return maxCount
}

func countPath(root *TreeNode, maxCount *int) int {
	if root == nil {
		return 0
	}

	leftLen := countPath(root.Left, maxCount)
	rightLen := countPath(root.Right, maxCount)

	arrowLeft, arrowRight := 0, 0

	if root.Left != nil && root.Left.Val == root.Val {
		arrowLeft = leftLen + 1
	}

	if root.Right != nil && root.Right.Val == root.Val {
		arrowRight = rightLen + 1
	}

	*maxCount = max(*maxCount, arrowLeft + arrowRight)
	return max(arrowLeft, arrowRight)
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
// @lc code=end

