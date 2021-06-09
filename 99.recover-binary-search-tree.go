/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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
func recoverTree(root *TreeNode)  {
    var prev, first, second *TreeNode
	dfs(root, &prev, &first, &second)
	(*first).Val, (*second).Val = (*second).Val, (*first).Val
}
func dfs(root *TreeNode, prev, first, second **TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left, prev, first, second)
	//fmt.Println(*prev, *first, *second, root)
	if *prev != nil && (*prev).Val > root.Val {
		if *first == nil {
			*first = *prev
		}
		*second = root
	}
	*prev = root
	dfs(root.Right, prev, first, second)
}
// @lc code=end

