/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
		return false
	}

	p1, p2 := head, head

	for p1 != nil && p2 != nil{
		p1 = p1.Next
		if p2.Next == nil {
			return false
		}
		p2 = p2.Next.Next

		if p1 == p2 {
			return true
		}
	}

	return false
}
// @lc code=end

