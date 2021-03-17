/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    start := &ListNode{0, head}
	origin := start
	for start.Next != nil {
		if start.Next.Val == val {
			start.Next = start.Next.Next
		} else {
			start = start.Next
		}
	}

	return origin.Next
}
// @lc code=end

