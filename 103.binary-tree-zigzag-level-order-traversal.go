/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
 func zigzagLevelOrder(root *TreeNode) [][]int {
    var right bool
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    res := make([][]int, 0)
    
    for len(queue) > 0 {
        temp := make([]*TreeNode, 0)
        level := make([]int, 0)
        for i := 0; i < len(queue); i++ {
            if queue[i] == nil {
                continue
            }
            level = append(level, queue[i].Val)
            temp = append(temp, queue[i].Left, queue[i].Right)
        }
        if right {
            for i := 0; i < len(level) / 2; i++ {
                level[i], level[len(level)-1-i] = level[len(level)-1-i], level[i]
            }
        }
        right = !right
        queue = temp
        if len(level) > 0 {
            res = append(res, level)
        }
    }

    return res
}
// @lc code=end

