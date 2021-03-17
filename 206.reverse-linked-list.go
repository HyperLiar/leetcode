/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var temp, prev *ListNode
    for head != nil {
		temp = head.Next
		head.Next = prev
		prev = head
		head = temp
	}

	return prev
}
// @lc code=end

