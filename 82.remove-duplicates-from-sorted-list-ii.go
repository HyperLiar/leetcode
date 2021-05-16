/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    m := make(map[int]int)

	temp := head
	for temp != nil {
		m[temp.Val]++
		temp = temp.Next
	}
	prev := &ListNode{Next:head}
	temp = prev
	for prev.Next != nil {
		if m[prev.Next.Val] != 1 {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return temp.Next
}
// @lc code=end

