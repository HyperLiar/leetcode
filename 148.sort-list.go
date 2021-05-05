/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // use merge sort
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next

	for ;fast != nil && fast.Next != nil; slow, fast = slow.Next, fast.Next.Next {}

	slow, slow.Next = slow.Next, nil
	return helper(sortList(head), sortList(slow))
}

func helper(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for ; l1 != nil && l2 != nil; tail = tail.Next {
		if l1.Val <= l2.Val {
			tail.Next, l1 = l1, l1.Next
		} else {
			tail.Next, l2 = l2, l2.Next
		}
	}

	if l1 != nil {
		tail.Next = l1
	}

	if l2 != nil {
		tail.Next = l2
	}

	return head.Next
}
// @lc code=end

