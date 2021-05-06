/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// even
	if fast.Next != nil {
		slow = slow.Next
	}

	// reverse
	var pre, temp *ListNode
	for slow != nil {
		temp = slow.Next
		slow.Next = pre
		pre = slow
		slow = temp
	}

	// compare
	for head != nil && pre != nil {
		if head.Val != pre.Val {
			return false
		}
		head, pre = head.Next, pre.Next
	}

	return true
}
// @lc code=end

