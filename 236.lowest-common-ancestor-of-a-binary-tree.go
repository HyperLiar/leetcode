/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ps, qs := []*TreeNode{}, []*TreeNode{}
	
	dfs(root, p, &ps)
	dfs(root, q, &qs)

	l := len(ps)
	if len(qs) < len(ps) {
		l = len(qs)
	}
	for i := 0; i < l; i++ {
		if ps[i] != qs[i] {
			return ps[i-1]
		}
	}

	return ps[l-1]
}

func dfs(root, p *TreeNode, stack *[]*TreeNode) bool {
	if root == nil {
		return false
	}

	*stack = append(*stack, root)
	if root == p {
		return true
	}

	if dfs(root.Left, p, stack) || dfs(root.Right, p, stack) {
		return true
	}

	*stack = (*stack)[0:len(*stack)-1]
	return false
}

func traversal(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
    l := lowestCommonAncestor(root.Left, p, q)
    r := lowestCommonAncestor(root.Right, p, q)
    if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	}
    return r
}
// @lc code=end

