/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	dfs(root, &cur, &res, targetSum)
	return res
}

func dfs(root *TreeNode, cur *[]int, res *[][]int, targetSum int) {
	if root == nil {
		return
	}
	*cur = append(*cur, root.Val)
	targetSum -= root.Val

	if root.Left == nil && root.Right == nil && targetSum == 0 {
		temp := []int{}
		temp = append(temp, *cur...)
		*res = append(*res, temp)
	}
	dfs(root.Left, cur, res, targetSum)
	dfs(root.Right, cur, res, targetSum)
	*cur = (*cur)[0 : len(*cur)-1]
}

// @lc code=end

