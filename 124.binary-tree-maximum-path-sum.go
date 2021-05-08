/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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
func maxPathSum(root *TreeNode) int {
	max := math.MinInt32
	helper(root, &max)
	return max
}

func helper(root *TreeNode, max *int) int {
	if root == nil {
		return math.MinInt32
	}

	passNow, notPassNow, rootNow := root.Val, math.MinInt32, math.MinInt32
	left := helper(root.Left, max)
	right := helper(root.Right, max)

	rootNow = maxInt(passNow, passNow+left, passNow+right, passNow+left+right)
	passNow = maxInt(passNow, passNow+left, passNow+right)
	notPassNow = maxInt(left, right)

	*max = maxInt(*max, rootNow, passNow, notPassNow)
	return passNow
}

func maxInt(nums ...int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
// @lc code=end

