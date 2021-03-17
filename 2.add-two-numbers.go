/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := new(ListNode)
    result := sum
	bias := 0
	for l1 != nil || l2 != nil || bias == 1 {
        sum.Next = new(ListNode)
        sum = sum.Next
        if l1 != nil {
            sum.Val += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            sum.Val += l2.Val
            l2 = l2.Next
        }

		sum.Val += bias

		if sum.Val >= 10 {
			sum.Val -= 10
			bias = 1
		} else {
			bias = 0
		}
	}

	return result.Next
}
// @lc code=end

