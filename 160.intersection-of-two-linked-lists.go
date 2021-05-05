/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2, diff := 0, 0, 0
	tailA, tailB := headA, headB

	for tailA != nil {
		tailA = tailA.Next
		l1++
	}
	for tailB != nil {
		tailB = tailB.Next
		l2++
	}

	tailA, tailB = headA, headB
	if l1 > l2 {
		tailA, tailB = tailB, tailA
		diff = l1 - l2
	} else {
		diff = l2 - l1
	}

	for i := 0; i < diff; i++ {
		tailB = tailB.Next
	}

	var intersectNode *ListNode
	for tailA != nil {
		if tailA == tailB {
			if intersectNode == nil {
				intersectNode = tailA
			}
		} else {
			intersectNode = nil
		}
		//fmt.Println(tailA, tailB, intersectNode)
		tailA, tailB = tailA.Next, tailB.Next
	}

	return intersectNode
}

// @lc code=end

