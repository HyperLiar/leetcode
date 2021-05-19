/*
 * @lc app=leetcode id=1448 lang=golang
 *
 * [1448] Count Good Nodes in Binary Tree
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
 func goodNodes(root *TreeNode) int {
    max, res := root.Val, 1
    helper(root.Left, max, &res)
    helper(root.Right, max, &res)
    return res
}

func helper(root *TreeNode, max int, res *int) {
    if root == nil {
        return
    }

    if root.Val >= max {
        max = root.Val
        *res++
    }
    helper(root.Left, max, res)
    helper(root.Right, max, res)
}

// once ac
func helper1(root *TreeNode, path []int) int {
    if root == nil {
        return 0
    }

    bigger := 1
    for _, v := range path {
        if v > root.Val {
            bigger = 0
            break
        }
    }

    path = append(path, root.Val)
    return helper1(root.Left, path) + helper1(root.Right, path) + bigger
}
// @lc code=end

