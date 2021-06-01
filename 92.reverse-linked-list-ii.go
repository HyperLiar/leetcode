/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseBetween(head *ListNode, left int, right int) *ListNode {
    newHead := &ListNode{Next:head}
    prev := newHead
    for i := 1; i < left; i++ {
        prev = prev.Next
    }
    curr := prev.Next
    var temp *ListNode
    for i := left; i < right; i++ {
        temp = curr.Next
        curr.Next = temp.Next
        temp.Next = prev.Next
        prev.Next = temp
    }

    return newHead.Next
}
// @lc code=end

