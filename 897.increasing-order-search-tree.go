/*
 * @lc app=leetcode id=897 lang=golang
 *
 * [897] Increasing Order Search Tree
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
func increasingBST(root *TreeNode) *TreeNode {
	queue := make([]*TreeNode, 0)
	res := make([]int, 0)
	n := root
	for n != nil || len(queue) > 0 {
		for n != nil {
			queue = append(queue, n)
			n = n.Left
		}
		top := queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]
		res = append(res, top.Val)
		n = top.Right
	}

	newRoot := &TreeNode{Val: res[0]}
	temp := newRoot
	fmt.Println(newRoot, temp)
	for i := 1; i < len(res); i++ {
		fmt.Println(newRoot, temp)
		temp.Right = &TreeNode{Val: res[i]}
		temp = temp.Right
	}

	return newRoot
}

// @lc code=end

