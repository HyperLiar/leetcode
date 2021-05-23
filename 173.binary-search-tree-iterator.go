/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
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
 type BSTIterator struct {
    stack []*TreeNode
}

func (this *BSTIterator) add(root *TreeNode) {
    for root != nil {
        this.stack = append(this.stack, root)
        root = root.Left
    }
}

func Constructor(root *TreeNode) BSTIterator {
    b := BSTIterator{stack:[]*TreeNode{}}
    b.add(root)

    return b
}


func (this *BSTIterator) Next() int {
    top := this.stack[len(this.stack)-1]
    this.stack = this.stack[0:len(this.stack)-1]
    this.add(top.Right)
    return top.Val
}


func (this *BSTIterator) HasNext() bool {
    return len(this.stack) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

