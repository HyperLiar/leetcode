/*
 * @lc app=leetcode id=894 lang=golang
 *
 * [894] All Possible Full Binary Trees
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
func allPossibleFBT(n int) []*TreeNode {
	allRes := make(map[int][]*TreeNode)
	res := make([]*TreeNode, 1)
	res[0] = &TreeNode{}
	allRes[0] = res
	allRes[1] = []*TreeNode{}

	traversal(n, &allRes)
	return allRes[n-1]
}

func traversal(n int, allRes *map[int][]*TreeNode) []*TreeNode {
	if _, ok := (*allRes)[n-1]; ok {
		return (*allRes)[n-1]
	}

	res := make([]*TreeNode, 0)
	for i := 1; i < n; i += 2 {
		l, r := traversal(i, allRes), traversal(n-i-1, allRes)

		for a := 0; a < len(l); a++ {
			for b := 0; b < len(r); b++ {
				root := &TreeNode{}
				root.Left = l[a]
				root.Right = r[b]
				res = append(res, root)
			}
		}
	}

	(*allRes)[n-1] = res

	return res
}

// @lc code=end

